<script setup lang="ts">
import { onMounted, ref } from 'vue'
import axios from 'axios'
import { Cocktail } from '../types/cocktail';

const cocktails = ref<Cocktail[] | null>()

onMounted(async () => {
  const response = await axios.get('http://localhost:8080/cocktails')

  if (response.status === 200) {
    console.log(response.data);

    cocktails.value = response.data
  }
})

</script>

<template>
  <h1>Cocktails</h1>

  <div v-if="cocktails">
    <div v-for="cocktail in cocktails" :key="cocktail.id">
      <h2>{{ cocktail.title }}</h2>
    </div>
  </div>
  <div v-else>
    <p>Loading...</p>
  </div>
</template>
