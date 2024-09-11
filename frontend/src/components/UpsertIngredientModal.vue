<script setup lang="ts">
import Modal from "./Modal/Modal.vue"
import CustomButton from "./CustomButton.vue"
import { ModalState } from './Modal/useModal';
import { ref } from "vue";
import { useErrorToast } from "../hooks/useErrorToast";
import { useCreateIngredient, useUpdateIngredient } from "../openapi/cocktailGridHooks";
import { DtosIngredientDto, VmsIngredientVm } from "../openapi/cocktailGridSchemas";

interface UpsertIngredientProps extends ModalState {
  ingredient?: VmsIngredientVm
}

const { onClose, ingredient } = defineProps<UpsertIngredientProps>()
const emit = defineEmits<{
  (e: 'create'): void
}>()
const { mutateAsync: createIngredient, isPending: createIngredientLoading, error: createIngredientError } = useCreateIngredient()
const { mutateAsync: updateIngredient, isPending: updateIngredientLoading, error: updateIngredientError } = useUpdateIngredient()
useErrorToast(createIngredientError || updateIngredientError)
const ingredientName = ref(ingredient?.name)

const isUpdate = !!ingredient

const handleUpsert = async () => {
  const updatedingredient: DtosIngredientDto = {
    ...ingredient,
    name: ingredientName.value
  }

  if (isUpdate) {
    await updateIngredient({
      pathParams: {
        ingredientId: ingredient.id!
      },
      body: updatedingredient
    })
  } else {
    await createIngredient({
      body: updatedingredient
    })
  }

  onClose()
  emit('create')
}

</script>

<template>
  <Modal v-bind="$props" :title="isUpdate ? 'Update Ingredient' : 'Create Ingredient'">
    <template v-slot:body>
      <label for="ingredient-name" class="flex flex-col">
        Name
        <input type="text" name="ingredient-name" id="ingredient-name" v-model="ingredientName" autofocus
          class="border-2 border-black outline-none rounded-md p-1">
      </label>
    </template>
    <template v-slot:actions>
      <CustomButton outlined @click="onClose">Cancel</CustomButton>
      <CustomButton outlined @click="handleUpsert" :loading="createIngredientLoading || updateIngredientLoading">
        {{ isUpdate ? 'Update' : 'Create' }}
      </CustomButton>
    </template>
  </Modal>
</template>