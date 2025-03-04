<template>
  <v-footer height="40" app>
    <div
      class="text-caption text-disabled"
      style="position: absolute; right: 16px;"
    >
      &copy;
      <span class="d-none d-sm-inline-block">
        TicTacToe Game By <a alt="github-link" target="_blank" href="https://github.com/margar-melkonyan" class="social-link">Melkonyan M.A.</a>
      </span>
    </div>
  </v-footer>
</template>

<script setup lang="ts">
  //
</script>

<style scoped lang="sass">
  .social-link
    color: rgba(var(--v-theme-on-background), var(--v-disabled-opacity))
    text-decoration: underline #ff7fea
    transition: .2s ease-in-out

    &:hover
      color: #7fff94
</style>
