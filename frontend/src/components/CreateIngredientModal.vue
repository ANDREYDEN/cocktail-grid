<script setup lang="ts">
import Modal from "./Modal/Modal.vue"
import CustomButton from "./CustomButton.vue"
import { ModalState } from './Modal/useModal';
import { ref } from "vue";
import { useErrorToast } from "../hooks/useErrorToast";
import { useCreateIngredient } from "../openapi/cocktailGridHooks";

const { onClose } = defineProps<ModalState>()
const emit = defineEmits<{
  (e: 'create'): void
}>()
const { mutateAsync: createIngredient, isPending: createIngredientLoading, error: createIngredientError } = useCreateIngredient()
useErrorToast(createIngredientError)
const ingredientName = ref()

const handleCreate = async (e: Element) => {
  await createIngredient({
    body: {
      name: ingredientName.value
    }
  })

  onClose()
  emit('create')
}

</script>

<template>
  <Modal v-bind="$props" title="Create Ingredient">
    <template v-slot:body>
      <label for="ingredient-name" class="flex flex-col">
        Name
        <input type="text" name="ingredient-name" id="ingredient-name" v-model="ingredientName" autofocus
          class="border-2 border-black outline-none rounded-md p-1">
      </label>
    </template>
    <template v-slot:actions>
      <CustomButton outlined @click="onClose">Cancel</CustomButton>
      <CustomButton outlined @click="handleCreate" :loading="createIngredientLoading">Create</CustomButton>
    </template>
  </Modal>
</template>