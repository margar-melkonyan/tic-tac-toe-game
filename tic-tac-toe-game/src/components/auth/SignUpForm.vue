<template>
  <v-card-text>
    <v-col>
      <v-text-field
        hide-details="auto"
        :label="$t('enterForm.name')"
        variant="outlined"
      />
    </v-col>
    <v-col>
      <v-text-field
        hide-details="auto"
        :label="$t('enterForm.email')"
        variant="outlined"
      />
    </v-col>
    <v-col>
      <v-text-field
        hide-details="auto"
        :label="$t('enterForm.password')"
        variant="outlined"
        type="password"
      />
    </v-col>
    <v-col>
      <v-text-field
        hide-details="auto"
        :label="$t('enterForm.password_confirmation')"
        variant="outlined"
      />
    </v-col>
  </v-card-text>
  <v-card-actions>
    <v-btn
      variant="tonal"
      block
    >
      {{ $t('enterForm.sign_up') }}
    </v-btn>
  </v-card-actions>
</template>

<script lang="ts" setup>
import {defineEmits} from 'vue';

const emits = defineEmits([
  'close'
])
</script>
