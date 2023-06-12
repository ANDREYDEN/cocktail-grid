<script setup lang="ts">
import { onMounted, ref } from 'vue'
import axios from 'axios'
import { CocktailDto } from '../types/cocktail';
import { Ingredient } from '../types/ingredient';

const cocktails = ref<CocktailDto[] | null>()
const ingredients = ref<Ingredient[] | null>()
const baseUrl = import.meta.env.VITE_BACKEND_URL

onMounted(async () => {
  await getCocktails()
  await getIngredients()
})

const getCocktails = async () => {
  const response = await axios.get(`${baseUrl}/cocktails`)

  if (response.status === 200) {
    cocktails.value = response.data
  }
}

const getIngredients = async () => {
  const response = await axios.get(`${baseUrl}/ingredients`)

  if (response.status === 200) {
    ingredients.value = response.data
  }
}

</script>

<template>
  <h1>Cocktail Grid</h1>

  <div v-if="!cocktails || !ingredients">
    <p>Loading...</p>
  </div>
  <div v-else>
    <div class="flex flex-row w-full">
      <h2 class="flex-1"></h2>
      <h2 v-for="ingredient in ingredients" :key="ingredient.id" class="flex-1">
        {{ ingredient.name }}
      </h2>
    </div>
    <div v-for="cocktail in cocktails" :key="cocktail.id" class="flex flex-row w-full justify-between">
      <h2 class="flex-1">{{ cocktail.title }}</h2>
      <h2 v-for="ingredient in ingredients" :key="ingredient.id" class="flex-1">
        {{ cocktail.ingredients.find(ci => ci.ingredientId === ingredient.id)?.quantity ?? "" }}
      </h2>
    </div>
  </div>
</template>
