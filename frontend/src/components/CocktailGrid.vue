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
const cocktailsToRender = computed(() => selectedIngredients.value.length == 0 ? cocktails.value : cocktails.value?.filter(c => selectedIngredients.value.some(i => c.ingredients.some(ci => ci.ingredientId === i.id))))
const ingredientsToRender = computed(() => selectedCocktails.value.length == 0 ? ingredients.value : ingredients.value?.filter(i => selectedCocktails.value.some(c => c.ingredients.some(ci => ci.ingredientId === i.id))))

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
  <h1 class="text-center text-5xl">Cocktail Grid</h1>

  <div v-if="!cocktails || !ingredients">
    <p>Loading...</p>
  </div>
  <div v-else class="m-10">
    <div class="flex flex-row w-full">
      <grid-cell />
      <grid-cell 
        v-for="ingredient in ingredientsToRender" 
        is-selectable
        :selected="selectedIngredients.includes(ingredient)"
        @click="handleIngredientSelect(ingredient)"
        :text="ingredient.name" />
    </div>
    <div v-for="cocktail in cocktailsToRender" class="flex flex-row w-full justify-between">
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
