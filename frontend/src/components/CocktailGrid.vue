<script setup lang="ts">
import { useAuth0 } from '@auth0/auth0-vue';
import { ArrowsUpDownIcon, PlusCircleIcon } from '@heroicons/vue/24/solid';
import { useQuery } from '@tanstack/vue-query';
import { computed, inject, ref } from 'vue';
import { getCocktails, getIngredients } from '../openapi/cocktailGridComponents';
import { useCreateCocktailIngredient, useDeleteCocktailIngredient, useUpdateCocktailIngredient } from '../openapi/cocktailGridHooks';
import { VmsDetailedCocktailVm, VmsIngredientVm } from '../openapi/cocktailGridSchemas';
import Account from './Account.vue';
import GridCell from './GridCell.vue';
import { useErrorToast } from '../hooks/useErrorToast'
import CustomButton from './CustomButton.vue';

const selectedCocktails = ref<VmsDetailedCocktailVm[]>([])
const selectedVmsIngredientVms = ref<VmsIngredientVm[]>([])
const cocktailsAsRows = ref<boolean>(true)

const auth = useAuth0();
const { data: cocktails, isLoading: loadingCocktails, refetch: refetchCocktails, isRefetching: refetchingCocktails } = useQuery({
  queryKey: ['getCocktails'],
  queryFn: () => getCocktails({})
})
const { data: ingredients, isLoading: loadingIngredients, refetch: refetchIngredients, isRefetching: refetchingIngredients } = useQuery({
  queryKey: ['getIngredients'],
  queryFn: () => getIngredients()
})

const { mutateAsync: mutateCreateCocktailIngredient, error: createCocktailIngredientError } = useCreateCocktailIngredient()
const { mutateAsync: mutateUpdateCocktailIngredient, error: updateCocktailIngredientError } = useUpdateCocktailIngredient()
const { mutateAsync: mutateDeleteCocktailIngredient, error: deleteCocktailIngredientError } = useDeleteCocktailIngredient()

useErrorToast(createCocktailIngredientError)
useErrorToast(updateCocktailIngredientError)
useErrorToast(deleteCocktailIngredientError)

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

const handleItemDelete = async (row: number, column: number) => {
  if (cocktailsAsRows.value) {
    if (column === 0) {
      return handleCocktailDelete(cocktails.value?.[row - 1]!)
    }
    if (row === 0) {
      return handleIngredientDelete(ingredients.value?.[column - 1]!)
    }
    await mutateDeleteCocktailIngredient({
      pathParams: {
        cocktailId: cocktails.value?.[row - 1]?.id!,
        ingredientId: ingredients.value?.[column - 1]?.id!
      }
    })
    await refetchCocktails()
  } else {
    if (column === 0) {
      return handleIngredientDelete(ingredients.value?.[row - 1]!)
    }
    if (row === 0) {
      return handleCocktailDelete(cocktails.value?.[column - 1]!)
    }
    await mutateDeleteCocktailIngredient({
      pathParams: {
        cocktailId: cocktails.value?.[column - 1]?.id!,
        ingredientId: ingredients.value?.[row - 1]?.id!
      }
    })
    await refetchCocktails()
  }
}

const handleItemEdit = async (row: number, column: number, value: string) => {
  if (cocktailsAsRows.value) {
    if (column === 0) {
      return
    }
    if (row === 0) {
      return
    }

    const ingredient = ingredients.value?.[column - 1]
    const cocktailIngredientExists = cocktails.value?.[row - 1]?.ingredients?.some(i => i.ingredientId === ingredient?.id)

    const upsertCocktailIngredient = (cocktailIngredientExists ? mutateUpdateCocktailIngredient : mutateCreateCocktailIngredient)

    await upsertCocktailIngredient({
      pathParams: {
        cocktailId: cocktails.value?.[row - 1]?.id!,
        ingredientId: ingredient?.id!,
      },
      body: {
        quantity: +value
      }
    })
    await refetchCocktails()
  } else {
    if (column === 0) {
      return
    }
    if (row === 0) {
      return
    }

    const ingredient = ingredients.value?.[row - 1]
    const cocktailIngredientExists = cocktails.value?.[column - 1]?.ingredients?.some(i => i.ingredientId === ingredient?.id)

    const upsertCocktailIngredient = (cocktailIngredientExists ? mutateUpdateCocktailIngredient : mutateCreateCocktailIngredient)

    await upsertCocktailIngredient({
      pathParams: {
        cocktailId: cocktails.value?.[column - 1]?.id!,
        ingredientId: ingredients.value?.[row - 1]?.id!
      },
      body: {
        quantity: +value
      }
    })
    await refetchCocktails()
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
const loading = computed(() => {
  return loadingCocktails.value || refetchingCocktails.value || loadingIngredients.value || refetchingIngredients.value
})

</script>

<template>
  <header class="flex items-center">
    <h1 class="text-center text-5xl my-8 grow">Cocktail Grid</h1>

    <Account />
  </header>

  <div v-if="loading">
    <p>Loading...</p>
  </div>
  <div class="m-8" v-else>
    <div class="flex flex-row gap-4 mb-4">
      <CustomButton outlined icon-position="left" @click="handleAddCocktail">
        Cocktail
        <template v-slot:icon>
          <PlusCircleIcon />
        </template>
      </CustomButton>
      <CustomButton outlined icon-position="left" @click="handleAddIngredient">
        Ingredient
        <template v-slot:icon>
          <PlusCircleIcon />
        </template>
      </CustomButton>
    </div>
    <div class="p-8 overflow-scroll rounded-lg bg-blue-50">
      <div class="flex gap-2" v-for="row in rowIndexRange">
        <div class="flex gap-2 flex-shrink-0" v-for="column in columnIndexRange">
          <GridCell v-if="row === 0 && column === 0" selectable @click="flipAxis">
            <ArrowsUpDownIcon class="text-black w-8 h-8 rotate-45" />
          </GridCell>
          <GridCell v-else :selectable="row === 0 || column === 0" :selected="isItemSelected(row, column)"
            :editable="auth.isAuthenticated.value" @edit="(value) => handleItemEdit(row, column, value)"
            @delete="handleItemDelete(row, column)" @click="handleItemSelected(row, column)"
            :text="itemText(row, column)" />
        </div>
      </div>
    </div>
  </div>
</template>
