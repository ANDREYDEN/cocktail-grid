<script setup lang="ts">
import { onMounted, ref } from 'vue'
import axios from 'axios'
import { Cocktail } from '../types/cocktail';

const cocktails = ref<Cocktail[] | null>()

onMounted(async () => {
  const baseUrl = import.meta.env.VITE_BACKEND_URL

  const response = await axios.get(`${baseUrl}/cocktails`)

  if (response.status === 200) {
    cocktails.value = response.data
  }
})

</script>

<template>
  <h1>Cocktailss</h1>

  <div v-if="cocktails">
    <div v-for="cocktail in cocktails" :key="cocktail.id">
      <h2>{{ cocktail.title }}</h2>
    </div>
  </div>
  <div v-else>
    <p>Loading...</p>
  </div>
</template>
