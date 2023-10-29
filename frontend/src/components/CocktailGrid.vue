<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import axios from 'axios'
import { CocktailDto } from '../types/cocktail';
import { Ingredient } from '../types/ingredient';
import GridCell from './GridCell.vue'

const cocktails = ref<CocktailDto[] | null>()
const ingredients = ref<Ingredient[] | null>()
const baseUrl = import.meta.env.VITE_BACKEND_URL
const selectedCocktails = ref<CocktailDto[]>([])
const selectedIngredients = ref<Ingredient[]>([])

const cocktailsToRender = computed(() => {
  if (selectedIngredients.value.length == 0) {
    return cocktails.value
  } 
  return cocktails.value?.filter(c => selectedIngredients.value.every(i => c.ingredients.some(ci => ci.ingredientId === i.id)))
})

const ingredientsToRender = computed(() => {
  if (selectedCocktails.value.length == 0) {
    return ingredients.value
  } 
  return ingredients.value?.filter(i => selectedCocktails.value.some(c => c.ingredients.some(ci => ci.ingredientId === i.id)))
})

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

const handleCocktailSelect = (cocktail: CocktailDto) => {
  if (selectedIngredients.value.length > 0) return

  if (selectedCocktails.value.includes(cocktail)) {
    selectedCocktails.value = selectedCocktails.value.filter(c => c.id !== cocktail.id)
  } else {
    selectedCocktails.value.push(cocktail)
  }
}

const handleIngredientSelect = (ingredient: Ingredient) => {
  if (selectedCocktails.value.length > 0) return

  if (selectedIngredients.value.includes(ingredient)) {
    selectedIngredients.value = selectedIngredients.value.filter(i => i.id !== ingredient.id)
  } else {
    selectedIngredients.value.push(ingredient)
  }
}

</script>

<template>
  <h1 class="text-center text-5xl mt-8">Cocktail Grid</h1>

  <div v-if="!cocktails || !ingredients">
    <p>Loading...</p>
  </div>
  <div v-else class="m-8 p-8 overflow-scroll rounded-lg bg-blue-50">
    <div class="flex gap-2">
      <grid-cell />
      <grid-cell 
        v-for="ingredient in ingredientsToRender" 
        is-selectable
        :selected="selectedIngredients.includes(ingredient)"
        @click="handleIngredientSelect(ingredient)"
        :text="ingredient.name" />
    </div>
    <div v-if="cocktailsToRender?.length == 0" class="h-16 flex items-center align-center">
      <p class="text-3xl text-gray-500">No cocktails match your selection</p>
    </div>
    <div v-else v-for="cocktail in cocktailsToRender" class="flex gap-2">
      <grid-cell 
        is-selectable
        :selected="selectedCocktails.includes(cocktail)"
        @click="handleCocktailSelect(cocktail)"
        :text="cocktail.title"
      />
      <grid-cell 
        v-for="ingredient in ingredientsToRender"
        :text="cocktail.ingredients.find(ci => ci.ingredientId === ingredient.id)?.quantity ?? ''"
      />
    </div>
  </div>
</template>
