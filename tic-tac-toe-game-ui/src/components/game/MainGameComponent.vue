<template>
  <v-col>
    <v-dialog
      v-model="chooseSymbolDialog"
      width="500"
      persistent
    >
      <ChooseSplashScreen
        :send-choose-symbol="sendChooseSymbol"
      />
    </v-dialog>
    <v-dialog
      v-model="waitSymbolChoosing"
      width="500"
      persistent
    >
      <WaitSplashScreen />
    </v-dialog>
    <v-dialog
      v-model="isPrivate"
      width="500"
      persistent
    >
      <PrivatePasswordInput
        :room-id="roomInfo?.id"
        :room-name="roomInfo?.name"
        :connect-to-room="connectToRoom"
      />
    </v-dialog>
    <v-row class="d-flex justify-center">
      <HeaderComponent
        :versus="versus"
        :room-info="roomInfo"
        :exit-from-room="exitFromRoom"
      />
    </v-row>
    <v-row
      v-if="wonFlag === 0 && versus !== null"
      class="d-flex justify-center my-4"
    >
      <span
        v-if="mySymbol === currentPlayer"
        style="color: yellowgreen;"
      >
        Ваш ход
      </span>
      <span
        v-else
        style="color: red;"
      >
        Ходит {{ versus }}
      </span>
    </v-row>
    <v-row
      v-if="wonFlag !== 0 && versus !== null"
      class="d-flex justify-center my-4"
    >
      Игра окончена!
    </v-row>
    <v-row class="d-flex justify-center my-8">
      <GameBoardComponent
        :versus="versus"
        :rows-and-columns="rowsAndColumns"
        :make-step="makeStep"
        :get-cell-style="getCellStyle"
        :get-font-style="getFontStyle"
      />
    </v-row>

    <GameControlsComponent
      v-if="gameStarted === 0 && authStore.user?.id === roomInfo?.creator_id"
      :game-started="gameStarted"
      :rows-and-columns="rowsAndColumns"
      :resize-board="resizeBoard"
      :room-info="roomInfo"
    />

    <GameResultComponent
      :my-symbol="mySymbol"
      :won-flag="wonFlag"
      :do-reset-game="doResetGame"
    />
  </v-col>
</template>

<script setup lang="ts">
import { getCurrentInstance, ref, computed, watch } from 'vue';
import axios from "axios";
import { useAuthStore } from "@/stores/auth";
import { useRouter } from "vue-router"

import HeaderComponent from './HeaderComponent.vue';
import GameBoardComponent from './GameBoardComponent.vue';
import GameControlsComponent from './GameControlsComponent.vue';
import GameResultComponent from './GameResultComponent.vue';
import { toast } from "vue3-toastify";

const { proxy } = getCurrentInstance();
const apiRooms = proxy.$api.rooms;

const authStore = useAuthStore();
const router = useRouter();

const props = defineProps<{
  roomId: string;
}>();
const mySymbol = ref<string>("");
const currentPlayer = ref<string>("");
const roomInfo = ref(null);
const rowsAndColumns = ref<number>(3);
const wonFlag = ref<number>(0);
const gameStarted = ref<number>(0);
const gameEnd = ref<number>(0);
const versusFetchIntervalId = ref<number>(0);
const isPrivate = ref<boolean>(false);
const wssIsSuccess = ref<boolean>(false);
const chooseSymbolDialog = ref<boolean>(false);
const waitSymbolChoosing = ref<boolean>(false);

const xCount = new Array(rowsAndColumns.value).fill(0);
const oCount = new Array(rowsAndColumns.value).fill(0);

let controller:AbortController;
let ws: WebSocket;

const sendChooseSymbol = (symbol:string) => {
  mySymbol.value = symbol
  currentPlayer.value = symbol
  chooseSymbolDialog.value = false
  waitSymbolChoosing.value = false
  ws.send(JSON.stringify({
    action: "select symbol",
    symbol: symbol,
  }))
}
function connectToRoom(id: string, password: string) {
  ws = new WebSocket(`${apiRooms.urls.room(id).replace("http", "ws")}?token=${localStorage.getItem("token")}`);
  ws.onopen = () => {
    ws.send(JSON.stringify({
      action: "new connection to room",
      password: password,
    }))
    wssIsSuccess.value = true;
    isPrivate.value = false;
  };
  ws.onerror = (event) => {
    console.error("WebSocket error observed:", event);
    router.push({ name: "index" })
  };
  ws.onclose = (event) => {
    toast.error(event.reason);
    if( event.code === 1013 ||
        event.reason === 'connection is close' ||
        event.reason === 'cannot find room'
    ) {
      clearInterval(versusFetchIntervalId.value)
      router.push({ name: "index" })
    }
    if (event.code === 1008) {
      wssIsSuccess.value = false;
      isPrivate.value = true;
    }
  };
  ws.onmessage = (event) => {
    const data = JSON.parse(event.data);
    switch(data.action) {
      case "reset game":
        resetGame();
        break;
      case "new connection to room":
        fetchRoom();
        break;
      case "get positions":
        const positions = data.data.positions;
        positions.forEach((position) => {
          const pos = position.id.split("-");
          const i = Number(pos[0]);
          const j = Number(pos[1]);
          playerStep(i, j, position.symbol);
        });
        currentPlayer.value = data.symbol
        if (positions.length === 0) {
          resetGame();
          fetchRoom();
          currentPlayer.value = "X"
        }
        break;
      case "resize":
        rowsAndColumns.value = data.size;
        resizeCountingArrays();
        break;
      case "choose symbol":
        mySymbol.value = ""
        if(authStore?.user?.id === data.user_id) {
          chooseSymbolDialog.value = true
        } else {
          waitSymbolChoosing.value = true
        }
        break
      case "selected symbol":
        mySymbol.value = data.symbol
        currentPlayer.value = 'X'
        waitSymbolChoosing.value = false
        break
      case "restart game":
        currentPlayer.value = 'X'
        break
      case "sync symbol":
        mySymbol.value = data.symbol
        waitSymbolChoosing.value = false
        break
    }
  };
}
if (!isPrivate.value) {
  connectToRoom(props.roomId, "");
}
const versus = computed(() => {
  return roomInfo.value?.users?.filter((user) => user.name != authStore.user?.name)[0]?.name || null;
});

async function fetchRoom() {
  if (controller) controller.abort();
  controller = new AbortController();
  try {
    const { data } = await axios.get(apiRooms.urls.roomInfo(props.roomId));
    roomInfo.value = data.data;
    const currentUser = roomInfo.value.users.filter((user) => user.id === authStore?.user?.id)[0]
    if (currentUser?.symbol !== undefined && currentUser?.symbol !== "") {
      mySymbol.value = currentUser?.symbol
      chooseSymbolDialog.value = false
      waitSymbolChoosing.value = false
    }
    if (!wssIsSuccess.value) {
      isPrivate.value = roomInfo.value.is_private
    }
  } catch (err) {
    console.error("Failed to fetch room info:", err);
  }
}
onMounted(() => {
  versusFetchIntervalId.value = setInterval(() => {
    if (!versus.value) {
      fetchRoom();
    } else {
      clearInterval(versusFetchIntervalId.value);
    }
  }, 2000);
})
onBeforeUnmount(() => {
  clearInterval(versusFetchIntervalId.value)
  if (controller) controller.abort();
})
function makeStep(i: number, j: number) {
  const cell = document.querySelector(`.grid-index-${i}-${j}>span`);
  if (
    cell && cell.textContent === '' && wonFlag.value === 0 && versus.value !== null &&
    currentPlayer.value == mySymbol.value
  ) {
    ws.send(JSON.stringify({
      data: {
        id: `${i}-${j}`,
        symbol: mySymbol.value,
      },
      action: "step",
    }))
  }
}
function playerStep(i: number, j: number, symbol: string) {
  gameStarted.value = 1;
  const cell = document.querySelector(`.grid-index-${i}-${j}>span`);
  if (wonFlag.value !== 0 || !cell) {
    return;
  }
  if (cell.textContent === 'O' || cell.textContent === 'X') {
    return;
  }
  cell.textContent = symbol;
  resetCounting();
  verticalCheck();
  if (wonFlag.value === 0) {
    resetCounting();
    horizontalCheck();
  }
  if (wonFlag.value === 0) {
    resetCounting();
    mainDiagonalCheck();
  }
  if (wonFlag.value === 0) {
    resetCounting();
    sideDiagonalCheck();
  }
  checkDraw();
}
function resetCounting() {
  xCount.fill(0);
  oCount.fill(0);
}
function resizeCountingArrays() {
  xCount.length = rowsAndColumns.value;
  oCount.length = rowsAndColumns.value;
  xCount.fill(0);
  oCount.fill(0);
}
function mainDiagonalCheck() {
  for (let i = 1; i <= rowsAndColumns.value; i++) {
    const cell = document.querySelector(`.grid-index-${i}-${i}>span`);
    if (!cell) continue;
    if (cell.textContent === 'O') oCount[i - 1] += 1;
    if (cell.textContent === 'X') xCount[i - 1] += 1;
  }
  diagonalChecker();
}
function sideDiagonalCheck() {
  for (let i = 0; i < rowsAndColumns.value; i++) {
    const cell = document.querySelector(`.grid-index-${i + 1}-${rowsAndColumns.value - i}>span`);
    if (!cell) continue;
    if (cell.textContent === 'O') oCount[i] += 1;
    if (cell.textContent === 'X') xCount[i] += 1;
  }
  diagonalChecker();
}
function diagonalChecker() {
  if (xCount.reduce((a, b) => a + b, 0) === rowsAndColumns.value) {
    wonFlag.value = 1;
  }
  if (oCount.reduce((a, b) => a + b, 0) === rowsAndColumns.value) {
    wonFlag.value = -1;
  }
}
function horizontalCheck() {
  for (let i = 1; i <= rowsAndColumns.value; i++) {
    for (let j = 1; j <= rowsAndColumns.value; ++j) {
      const cell = document.querySelector(`.grid-index-${i}-${j}>span`);
      if (!cell) continue;
      if (cell.textContent === 'O') oCount[j - 1] += 1;
      if (cell.textContent === 'X') xCount[j - 1] += 1;
    }
  }
  lineChecker();
}
function verticalCheck() {
  for (let i = 1; i <= rowsAndColumns.value; i++) {
    for (let j = 1; j <= rowsAndColumns.value; ++j) {
      const cell = document.querySelector(`.grid-index-${i}-${j}>span`);
      if (!cell) continue;
      if (cell.textContent === 'O') oCount[i - 1] += 1;
      if (cell.textContent === 'X') xCount[i - 1] += 1;
    }
  }
  lineChecker();
}
function lineChecker() {
  if (Math.max(...xCount) === rowsAndColumns.value) {
    wonFlag.value = 1;
  }
  if (Math.max(...oCount) === rowsAndColumns.value) {
    wonFlag.value = -1;
  }
}
function checkDraw() {
  gameEnd.value = 0;
  for (let i = 1; i <= rowsAndColumns.value; i++) {
    for (let j = 1; j <= rowsAndColumns.value; ++j) {
      const cell = document.querySelector(`.grid-index-${i}-${j}>span`);
      if (cell && cell.textContent !== '') {
        gameEnd.value += 1;
      }
    }
  }
  if (gameEnd.value === rowsAndColumns.value * rowsAndColumns.value && wonFlag.value === 0) {
    wonFlag.value = -2;
  }
}
function doResetGame() {
  ws.send(JSON.stringify({
    action: 'reset game'
  }));
}
function resetGame() {
  gameStarted.value = 0;
  wonFlag.value = 0;
  resetCounting();
  resetGameBoardCells();
}
function resetGameBoardCells() {
  const gridSpans = document.querySelectorAll('[class^="grid-index-"] > span');
  gridSpans.forEach(span => {
    span.textContent = '';
  });
}
function getCellStyle() {
  return {
    flex: '1',
    aspectRatio: '1',
    border: '0.25rem solid #ff7fea',
    'border-radius': '0.25rem',
    margin: '0.25rem',
    backgroundColor: 'white',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    cursor: 'pointer',
  };
}
function getFontStyle() {
  const base = 300;
  const size = Math.floor(base / rowsAndColumns.value);
  return {
    color: 'black',
    fontSize: `${size}px`,
  };
}
function resizeBoard(size: number) {
  rowsAndColumns.value += size;
  resizeCountingArrays();
  resetGameBoardCells();
  ws.send(JSON.stringify({
    action: "resize",
    size: rowsAndColumns.value
  }))
}
function exitFromRoom() {
  if (versus.value !== null && authStore?.user.id === roomInfo.value.creator_id) {
    ws.send(JSON.stringify({
      action: "close room"
    }))
    axios.delete(apiRooms.urls.room(props.roomId))
  }
  if (authStore?.user.id !== roomInfo.value.creator_id) {
    ws.send(JSON.stringify({
      action: "exit room",
    }))
  }
  router.push({ name: "index" })
}
watch(wonFlag, function () {
  const wonState = () => {
    if (mySymbol.value === 'X' && wonFlag.value === 1) {
      return 1
    }
    if (mySymbol.value === 'O' && wonFlag.value === -1) {
      return 1
    }
    if (wonFlag.value === -2) {
      return -1
    }
    return 0
  }
  if (wonFlag.value !== 0) {
    ws.send(JSON.stringify({
      "action": "game end",
      "data": {
        "is_won": wonState(),
        "user_id": authStore?.user?.id,
        "versus_player_nickname": versus.value
        }
    }))
  }
})
</script>

<style scoped>
.no-select::-moz-selection {
  background: transparent;
  color: inherit;
}

.no-select::-webkit-selection {
  background: transparent;
  color: inherit;
}
</style>
