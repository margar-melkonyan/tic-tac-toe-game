<script setup lang="ts">
import { ref } from 'vue';
import { v4 as uuidv4 } from 'uuid';
/*
 *  X - 0
 *  O - 1
 */

const rowsAndColumns = ref<number>(3);
const currentPlayer = ref<number>(0);
const wonFlag = ref<number>(0)
const gameStarted = ref<number>(0)
const gameEnd = ref<number>(0)
const xCount = new Array(rowsAndColumns.value);
const oCount = new Array(rowsAndColumns.value);

function playerStep(i: number, j: number) {
  gameStarted.value = 1
  const cell = document.querySelector(`.grid-index-${i}-${j}>span`)

  if (wonFlag.value !== 0) {
    return
  }

  if (cell?.textContent === 'O' || cell?.textContent === 'X') {
    return
  }

  if (currentPlayer.value === 0) {
    cell.textContent = 'O'
    currentPlayer.value = 1
  } else {
    cell.textContent = 'X'
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

function resetGame() {
  gameStarted.value = 0
  currentPlayer.value = 0
  wonFlag.value = 0
  resetCounting()
  const grid = document.querySelectorAll('[class^="grid-index-"]>span')

  grid.forEach((cell) => {
    cell.textContent = ''
  })
}
</script>

<template>
  <v-col cols="12">
    <v-row v-if="gameStarted === 0" class="d-flex justify-center">
      <!-- {{ uuidv4() }} -->
      <v-btn color="red" :disabled="rowsAndColumns === 3" @click="--rowsAndColumns">
        <v-icon>
          mdi-minus
        </v-icon>
      </v-btn>
      <v-btn color="green" :disabled="rowsAndColumns === 6" @click="++rowsAndColumns">
        <v-icon>
          mdi-plus
        </v-icon>
      </v-btn>
    </v-row>
    <v-row class="d-flex justify-center" style="max-width: 250px;">
      <div style="display: flex; justify-content: center; max-width: 250px;">
        <div v-for="i in rowsAndColumns" :key="i">
          <div v-for="j in rowsAndColumns" :key="j" :class="`grid-index-${i}-${j}`" @click="playerStep(i, j)"
            style="display: flex; align-items: center; justify-content: center; width: 200px; height: 200px; border: 1px solid black; background-color: white; cursor: pointer;">
            <span class="no-select" style="color: black; font-size: 100px;">
            </span>
          </div>
        </div>
      </div>
    </v-row>
    <v-row class="mt-4 w-100" v-if="wonFlag !== 0">
      <v-col cols="6" v-if="wonFlag === -1">
        Выграл первый игрок (О)
      </v-col>
      <v-col cols="6" v-if="wonFlag === 1">
        Выграл второй игрок (X)
      </v-col>
      <v-col cols="6" v-if="wonFlag === -2">
        Ничья
      </v-col>
      <v-col cols="6">
        <button @click="resetGame">
          Начать заново
        </button>
      </v-col>
    </v-row>
  </v-col>
</template>

<style>
.no-select::-moz-selection {
  background: transparent;
  color: inherit;
}

.no-select::-webkit-selection {
  background: transparent;
  color: inherit;
}
</style>
