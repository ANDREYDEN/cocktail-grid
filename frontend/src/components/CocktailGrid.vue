<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import axios from 'axios'
import GridCell from './GridCell.vue'
import { useAuth0 } from '@auth0/auth0-vue';
import { ArrowsUpDownIcon } from '@heroicons/vue/24/solid';
import Account from './Account.vue'
import { VmsDetailedCocktailVm, VmsIngredientVm } from '../openapi/cocktailGridSchemas'
import { useQuery } from '@tanstack/vue-query';
import { getCocktails } from '../openapi/cocktailGridComponents';

const ingredients = ref<VmsIngredientVm[] | null>()
const backendUrl = import.meta.env.VITE_BACKEND_URL
const selectedCocktails = ref<VmsDetailedCocktailVm[]>([])
const selectedVmsIngredientVms = ref<VmsIngredientVm[]>([])
const cocktailsAsRows = ref<boolean>(true)

const auth = useAuth0();
const { data: cocktails, isLoading: loadingCocktails, refetch: refetchCocktails, isRefetching: refetchingCocktails } = useQuery({
  queryKey: ['getCocktails'],
  queryFn: () => getCocktails({})
})

const cocktailsToRender = computed(() => {
  if (selectedVmsIngredientVms.value.length == 0) {
    return cocktails.value
  }
  return cocktails.value?.filter(c => selectedVmsIngredientVms.value.every(i => c.ingredients?.some(ci => ci.ingredientId === i.id)))
})

const ingredientsToRender = computed(() => {
  if (selectedCocktails.value.length == 0) {
    return ingredients.value
  }
  return ingredients.value?.filter(i => selectedCocktails.value.some(c => c.ingredients?.some(ci => ci.ingredientId === i.id)))
})

onMounted(async () => {
  await getIngredient()
})

const getIngredient = async () => {
  const response = await axios.get(`${backendUrl}/ingredients`)

  if (response.status === 200) {
    ingredients.value = response.data
  }
}

const handleCocktailSelect = (cocktail: VmsDetailedCocktailVm) => {
  if (selectedVmsIngredientVms.value.length > 0) return

  if (selectedCocktails.value.includes(cocktail)) {
    selectedCocktails.value = selectedCocktails.value.filter(c => c.id !== cocktail.id)
  } else {
    selectedCocktails.value.push(cocktail)
  }
}

const handleIngredientSelect = (ingredient: VmsIngredientVm) => {
  if (selectedCocktails.value.length > 0) return

  if (selectedVmsIngredientVms.value.includes(ingredient)) {
    selectedVmsIngredientVms.value = selectedVmsIngredientVms.value.filter(i => i.id !== ingredient.id)
  } else {
    selectedVmsIngredientVms.value.push(ingredient)
  }
}

const handleCocktailDelete = (cocktail: VmsDetailedCocktailVm) => {

}

const handleIngredientDelete = (ingredient: VmsIngredientVm) => {

}

const handleCocktailIngredientDelete = async (cocktail: VmsDetailedCocktailVm, ingredient: VmsIngredientVm) => {
  const token = await auth.getAccessTokenSilently()
  const response = await axios.delete(`${backendUrl}/cocktails/${cocktail.id}/ingredients/${ingredient.id}`, {
    headers: {
      'Authorization': 'Bearer ' + token
    }
  })
  if (response.status == 204) {
    await refetchCocktails()
  }
}

const flipAxis = () => {
  cocktailsAsRows.value = !cocktailsAsRows.value
}

const isItemSelected = (row: number, column: number) => {
  if (cocktailsAsRows.value) {
    if (row === 0) {
      const targetedVmsIngredientVm = ingredientsToRender.value?.[column - 1]
      return selectedVmsIngredientVms.value?.some(i => i.id === targetedVmsIngredientVm?.id)
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
      const targetedVmsIngredientVm = ingredientsToRender.value?.[row - 1]
      return selectedVmsIngredientVms.value?.some(i => i.id === targetedVmsIngredientVm?.id)
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

const handleAddCocktail = () => {

}

const handleAddIngredient = () => {

}

const itemText = (row: number, column: number) => {
  if (cocktailsAsRows.value) {
    if (row === 0) {
      return ingredientsToRender.value?.[column - 1]?.name ?? ''
    }
    if (column === 0) {
      return cocktailsToRender.value?.[row - 1]?.title ?? ''
    }
    return cocktailsToRender.value?.[row - 1].ingredients?.find(ci => ci.ingredientId === ingredientsToRender.value?.[column - 1]?.id)?.quantity ?? ''
  } else {
    if (row === 0) {
      return cocktailsToRender.value?.[column - 1]?.title ?? ''
    }
    if (column === 0) {
      return ingredientsToRender.value?.[row - 1]?.name ?? ''
    }
    return cocktailsToRender.value?.[column - 1].ingredients?.find(ci => ci.ingredientId === ingredientsToRender.value?.[row - 1]?.id)?.quantity ?? ''
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
  <header class="flex items-center">
    <h1 class="text-center text-5xl my-8 grow">Cocktail Grid</h1>

    <Account />
  </header>

  <div v-if="loadingCocktails || refetchingCocktails || !ingredients">
    <p>Loading...</p>
  </div>
  <div v-else>
    <div>
      <button @click="handleAddCocktail">Add Cocktail</button>
      <button @click="handleAddIngredient">Add Ingredient</button>
    </div>
    <div class="m-8 p-8 overflow-scroll rounded-lg bg-blue-50">
      <div class="flex gap-2" v-for="row in rowIndexRange">
        <div class="flex gap-2 flex-shrink-0" v-for="column in columnIndexRange">
          <GridCell v-if="row === 0 && column === 0" selectable @click="flipAxis">
            <ArrowsUpDownIcon class="text-black w-8 h-8 rotate-45" />
          </GridCell>
          <GridCell v-else :selectable="row === 0 || column === 0" :selected="isItemSelected(row, column)"
            :deletable="auth.isAuthenticated.value" @delete="handleItemDelete(row, column)"
            @click="handleItemSelected(row, column)" :text="itemText(row, column)" />
        </div>
      </div>
    </div>
  </div>
</template>
