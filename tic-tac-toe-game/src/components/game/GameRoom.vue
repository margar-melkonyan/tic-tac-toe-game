<template>
  <v-col>
    <v-row>
      <v-col
        class="d-flex justify-space-between"
        cols="12"
      >
        <span
          v-if="versus"
          class="justify-start"
          style="color:red; font-size: 18px;"
        >
          {{ versus }}
        </span>
        <span v-else>
          Ждем противника!
        </span>
        <img
          src="../../assets/vs.png"
          style="height: 45px;"
        >
        <span
          class="justify-end"
          style="color:greenyellow; font-size: 18px;"
        >
          {{ authStore.user?.name }}
        </span>
      </v-col>
      <v-divider class="my-4" />
    </v-row>
    <v-row class="d-flex justify-center">
      <div style="display: flex; flex-direction: column; width: 100%; max-width: 600px;">
        <div
          v-for="i in rowsAndColumns"
          :key="i"
          style="display: flex; width: 100%;"
        >
          <div
            v-for="j in rowsAndColumns"
            :key="j"
            :class="`grid-index-${i}-${j}`"
            :style="getCellStyle()"
            @click="makeStep(i, j)"
          >
            <span
              class="no-select"
              :style="getFontStyle()"
            />
          </div>
        </div>
      </div>
    </v-row>
    <v-row
      v-if="gameStarted === 0 &&
        authStore.user?.id === roomInfo?.creator_id
      "
      class="d-flex justify-center"
    >
      <v-col
        cols="2"
        class="d-flex justify-space-between"
      >
        <v-btn
          color="red"
          block
          :disabled="rowsAndColumns === 3"
          @click="resizeBoard(-1)"
        >
          <v-icon>
            mdi-minus
          </v-icon>
        </v-btn>
      </v-col>
      <v-col cols="2">
        <v-btn
          color="green"
          block
          :disabled="rowsAndColumns === 10"
          @click="resizeBoard(1)"
        >
          <v-icon>
            mdi-plus
          </v-icon>
        </v-btn>
      </v-col>
    </v-row>
    <v-row
      v-if="wonFlag !== 0"
      class="d-flex justify-center mt-6"
    >
      <v-col
        cols="12"
        class="text-center"
        style="max-width: 500px"
      >
        <v-alert
          v-if="wonFlag === -1"
          type="success"
          colored-border
          elevation="2"
        >
          🟢 Победил первый игрок (О)
        </v-alert>

        <v-alert
          v-if="wonFlag === 1"
          type="info"
          colored-border
          elevation="2"
        >
          🔵 Победил второй игрок (X)
        </v-alert>

        <v-alert
          v-if="wonFlag === -2"
          type="warning"
          colored-border
          elevation="2"
        >
          ⚪ Ничья
        </v-alert>
        <v-btn
          class="mt-4"
          color="primary"
          @click="doResetGame"
        >
          🔄 Начать заново
        </v-btn>
      </v-col>
    </v-row>
  </v-col>
</template>

<script setup lang="ts">
import { getCurrentInstance, ref, computed } from 'vue';
import axios from "axios";
import { useAuthStore } from "@/stores/auth";
const { proxy } = getCurrentInstance();
const apiRooms = proxy.$api.rooms;
/*
 *  X - 0
 *  O - 1
 */
const roomInfo = ref(null)
const authStore = useAuthStore()
const versus = computed(() => {
  return roomInfo.value?.users?.filter((user) => user.name != authStore.user?.name)[0]?.name
})

const versusFetchInterval = ref(0)
versusFetchInterval.value = setInterval(() => {
  if (versus.value == null) {
    fetchRoom()
  } else {
    versusFetchInterval.value = 0
  }
}, 5000);
authStore.currentUser()
const props = defineProps<{
  roomId: string;
}>()
let ws: WebSocket;
function connectToRoom(id) {
  ws = new WebSocket(`ws://192.168.1.4:8000/api/v1/rooms/${id}?token=${localStorage.getItem("token")}`);
  ws.onopen = function () {
    console.log("test")
    ws.send(`{"action": "new connection to room"}`)
  };
  ws.onerror = function (event) {
    console.error("WebSocket error observed:", event);
  };
  ws.onmessage = function (event) {
    const data = JSON.parse(event.data)
    if (data.action == "reset game") {
      resetGame()
    }
    if (data.action == "new connection to room")
    {
      fetchRoom()
    }
    if (data.action == "get positions") {
      const positions = data.data.positions
      positions.forEach((position) => {
        const pos = position.id.split("-")
        const i = pos[0]
        const j = pos[1]
        playerStep(i, j, position.symbol)
      })
    }
    if (data.action == "resize") {
      rowsAndColumns.value = data.size
    }
  }
}
connectToRoom(props.roomId)
const rowsAndColumns = ref<number>(3);
const currentPlayer = ref<number>(0);
const wonFlag = ref<number>(0)
const gameStarted = ref<number>(0)
const gameEnd = ref<number>(0)
const xCount = new Array(rowsAndColumns.value);
const oCount = new Array(rowsAndColumns.value);
async function fetchRoom() {
  await axios.get(apiRooms.urls.roomInfo(props.roomId)).then(({ data }) => {
    roomInfo.value = data.data
  })
}
fetchRoom()
function makeStep(i: number, j: number) {
  const cell = document.querySelector(`.grid-index-${i}-${j}>span`)
  if (cell?.textContent === '') {
    ws.send(`{"data":{"id": "${i}-${j}", "symbol": "X"}, "action": "step"}`)
  }
}
function playerStep(i: number, j: number, symbol: string) {
  gameStarted.value = 1
  const cell = document.querySelector(`.grid-index-${i}-${j}>span`)
  if (wonFlag.value !== 0) {
    return
  }
  if (cell?.textContent === 'O' || cell?.textContent === 'X') {
    return
  }
  if (currentPlayer.value === 0) {
    cell.textContent = symbol
    currentPlayer.value = 1
  } else {
    cell.textContent = symbol
    currentPlayer.value = 0
  }
  resetCounting()
  verticalCheck()
  if (wonFlag.value === 0) {
    resetCounting()
    horizontalCheck()
  }
  if (wonFlag.value === 0) {
    resetCounting()
    mainDiagonalCheck()
  }
  if (wonFlag.value === 0) {
    resetCounting()
    sideDiagonalCheck()
  }
  checkDraw()
}
function resetCounting() {
  xCount.fill(0);
  oCount.fill(0);
}
function mainDiagonalCheck() {
  for (let i = 1; i <= rowsAndColumns.value; i++) {
    const cell = document.querySelector(`.grid-index-${i}-${i}`)

    if (cell?.textContent === 'O') {
      oCount[i - 1] += 1
    }

    if (cell?.textContent === 'X') {
      xCount[i - 1] += 1
    }
  }

  diagonalChecker()
}
function sideDiagonalCheck() {
  for (let i = 0; i < rowsAndColumns.value; i++) {
    const cell = document.querySelector(`.grid-index-${i + 1}-${rowsAndColumns.value - i}`)

    if (cell?.textContent === 'O') {
      oCount[i] += 1
    }

    if (cell?.textContent === 'X') {
      xCount[i] += 1
    }
  }

  diagonalChecker()
}
function diagonalChecker() {
  if (xCount.reduce((x, y) => x + y) === rowsAndColumns.value) {
    wonFlag.value = 1
  }

  if (oCount.reduce((x, y) => x + y) === rowsAndColumns.value) {
    wonFlag.value = -1
  }
}
function horizontalCheck() {
  for (let i = 1; i <= rowsAndColumns.value; i++) {
    for (let j = 1; j <= rowsAndColumns.value; ++j) {
      const cell = document.querySelector(`.grid-index-${i}-${j}`)

      if (cell?.textContent === 'O') {
        oCount[j - 1] += 1
      }

      if (cell?.textContent === 'X') {
        xCount[j - 1] += 1
      }
    }
  }

  lineChecker()
}
function verticalCheck() {
  for (let i = 1; i <= rowsAndColumns.value; i++) {
    for (let j = 1; j <= rowsAndColumns.value; ++j) {
      const cell = document.querySelector(`.grid-index-${i}-${j}`)

      if (cell?.textContent === 'O') {
        oCount[i - 1] += 1
      }

      if (cell?.textContent === 'X') {
        xCount[i - 1] += 1
      }
    }
  }

  lineChecker()
}
function lineChecker() {
  if (Math.max.apply(null, xCount) === rowsAndColumns.value) {
    wonFlag.value = 1
  }

  if (Math.max.apply(null, oCount) === rowsAndColumns.value) {
    wonFlag.value = -1
  }
}
function checkDraw() {
  gameEnd.value = 0
  for (let i = 1; i <= rowsAndColumns.value; i++) {
    for (let j = 1; j <= rowsAndColumns.value; ++j) {
      const cell = document.querySelector(`.grid-index-${i}-${j}`)
      if (cell?.textContent !== '') {
        gameEnd.value += 1
      }
    }
  }

  if (gameEnd.value === rowsAndColumns.value * rowsAndColumns.value) {
    wonFlag.value = -2
  }
}
function doResetGame() {
  ws.send(`{"action": "reset game"}`)
}
function resetGame() {
  gameStarted.value = 0
  currentPlayer.value = 0
  wonFlag.value = 0
  resetCounting()
  const grid = document.querySelectorAll('[class^="grid-index-"]>span')

  grid.forEach((cell) => {2000
    cell.textContent = ''
  })
}
function getCellStyle() {
  return {
    flex: '1',
    aspectRatio: '1',
    border: '1px solid black',
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

function resizeBoard(size) {
  rowsAndColumns.value += size
  ws.send(`{"action":"resize", "size": ${rowsAndColumns.value}}`)
}
</script>

<style>
.game-cell>span {
  font-size: 2.5vw;
  color: black;
}

.no-select::-moz-selection {
  background: transparent;
  color: inherit;
}

.no-select::-webkit-selection {
  background: transparent;
  color: inherit;
}
</style>
