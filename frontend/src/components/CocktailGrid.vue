<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import axios from 'axios'
import { CocktailDto } from '../types/cocktail';
import { Ingredient } from '../types/ingredient';
import GridCell from './GridCell.vue'
import LoginButton from './LoginButton.vue';
import LogoutButton from './LogoutButton.vue';
import { useAuth0 } from '@auth0/auth0-vue';

const cocktails = ref<CocktailDto[] | null>()
const ingredients = ref<Ingredient[] | null>()
const backendUrl = import.meta.env.VITE_BACKEND_URL
const selectedCocktails = ref<CocktailDto[]>([])
const selectedIngredients = ref<Ingredient[]>([])
const cocktailsAsRows = ref<boolean>(true)

const auth = useAuth0();

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
  const response = await axios.get(`${backendUrl}/cocktails`)

  if (response.status === 200) {
    cocktails.value = response.data
  }
}

const getIngredients = async () => {
  const response = await axios.get(`${backendUrl}/ingredients`)

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

const handleCocktailDelete = (cocktail: CocktailDto) => {

}

const handleIngredientDelete = (ingredient: Ingredient) => {

}

const handleCocktailIngredientDelete = async (cocktail: CocktailDto, ingredient: Ingredient) => {
  const token = await auth.getAccessTokenSilently()
  const response = await axios.delete(`${backendUrl}/cocktails/${cocktail.id}/ingredients/${ingredient.id}`, {
    headers: {
      'Authorization': 'Bearer ' + token
    }
  })
  if (response.status == 204) {
    await getCocktails()
  }
}

const flipAxis = () => {
  cocktailsAsRows.value = !cocktailsAsRows.value
}

const isItemSelected = (row: number, column: number) => {
  if (cocktailsAsRows.value) {
    if (row === 0) {
      const targetedIngredient = ingredientsToRender.value?.[column - 1]
      return selectedIngredients.value?.some(i => i.id === targetedIngredient?.id)
    }
    if (column === 0) {
      const targetedCocktail = cocktailsToRender.value?.[row - 1]
      return selectedCocktails.value?.some(c => c.id === targetedCocktail?.id)
    }
  } else {
    if (row === 0) {
      const targetedCocktail = cocktailsToRender.value?.[column - 1]
      return selectedCocktails.value?.some(c => c.id === targetedCocktail?.id)
    }
    if (column === 0) {
      const targetedIngredient = ingredientsToRender.value?.[row - 1]
      return selectedIngredients.value?.some(i => i.id === targetedIngredient?.id)
    }
  }
  return false
}

const handleItemSelected = (row: number, column: number) => {
  if (cocktailsAsRows.value) {
    if (column === 0) {
      return handleCocktailSelect(cocktails.value?.[row - 1]!)
    }
    if (row === 0) {
      return handleIngredientSelect(ingredients.value?.[column - 1]!)
    }
  } else {
    if (column === 0) {
      return handleIngredientSelect(ingredients.value?.[row - 1]!)
    }
    if (row === 0) {
      return handleCocktailSelect(cocktails.value?.[column - 1]!)
    }
  }
}

const handleItemDelete = (row: number, column: number) => {
  if (cocktailsAsRows.value) {
    if (column === 0) {
      return handleCocktailDelete(cocktails.value?.[row - 1]!)
    }
    if (row === 0) {
      return handleIngredientDelete(ingredients.value?.[column - 1]!)
    }
    return handleCocktailIngredientDelete(cocktails.value?.[row - 1]!, ingredients.value?.[column - 1]!)
  } else {
    if (column === 0) {
      return handleIngredientDelete(ingredients.value?.[row - 1]!)
    }
    if (row === 0) {
      return handleCocktailDelete(cocktails.value?.[column - 1]!)
    }
    return handleCocktailIngredientDelete(cocktails.value?.[column - 1]!, ingredients.value?.[row - 1]!)
  }
}

const itemText = (row: number, column: number) => {
  if (cocktailsAsRows.value) {
    if (row === 0) {
      return ingredientsToRender.value?.[column - 1]?.name ?? ''
    }
    if (column === 0) {
      return cocktailsToRender.value?.[row - 1]?.title ?? ''
    }
    return cocktailsToRender.value?.[row - 1].ingredients.find(ci => ci.ingredientId === ingredientsToRender.value?.[column - 1]?.id)?.quantity ?? ''
  } else {
    if (row === 0) {
      return cocktailsToRender.value?.[column - 1]?.title ?? ''
    }
    if (column === 0) {
      return ingredientsToRender.value?.[row - 1]?.name ?? ''
    }
    return cocktailsToRender.value?.[column - 1].ingredients.find(ci => ci.ingredientId === ingredientsToRender.value?.[row - 1]?.id)?.quantity ?? ''
  }
}

const rowIndexRange = computed(() => {
  const length = (cocktailsAsRows.value ? cocktailsToRender.value?.length : ingredientsToRender.value?.length) ?? 0
  return Array.from({ length: length + 1 }, (_, i) => i)
})

const columnIndexRange = computed(() => {
  const length = (cocktailsAsRows.value ? ingredientsToRender.value?.length : cocktailsToRender.value?.length) ?? 0
  return Array.from({ length: length + 1 }, (_, i) => i)
})

</script>

<template>
  <h1 class="text-center text-5xl mt-8">Cocktail Grid</h1>

  <LoginButton v-if="!auth.isAuthenticated.value" />
  <div v-if="auth.isAuthenticated.value">{{ auth.user.value?.name }}</div>
  <LogoutButton v-if="auth.isAuthenticated.value" />

  <div v-if="!cocktails || !ingredients">
    <p>Loading...</p>
  </div>
  <div v-else class="m-8 p-8 overflow-scroll rounded-lg bg-blue-50">
    <div class="flex gap-2" v-for="row in rowIndexRange">
      <div class="flex gap-2 flex-shrink-0" v-for="column in columnIndexRange">
        <GridCell v-if="row === 0 && column === 0" selectable @click="flipAxis" text="swap" />
        <GridCell v-else :selectable="row === 0 || column === 0" :selected="isItemSelected(row, column)"
          :deletable="auth.isAuthenticated.value" @delete="handleItemDelete(row, column)"
          @click="handleItemSelected(row, column)" :text="itemText(row, column)" />
      </div>
    </div>
  </div>
</template>
