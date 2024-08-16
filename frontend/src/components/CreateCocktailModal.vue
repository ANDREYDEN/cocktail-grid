<script setup lang="ts">
import Modal from "./Modal/Modal.vue"
import CustomButton from "./CustomButton.vue"
import { ModalState } from './Modal/useModal';
import { ref } from "vue";
import { useCreateCocktail } from "../openapi/cocktailGridHooks";
import { useErrorToast } from "../hooks/useErrorToast";

const { onClose } = defineProps<ModalState>()
const emit = defineEmits<{
  (e: 'create'): void
}>()
const { mutateAsync: createCocktail, isPending: createCocktailLoading, error: createCocktailError } = useCreateCocktail()
useErrorToast(createCocktailError)
const cocktailName = ref()

const handleCreate = async (e: Element) => {
  await createCocktail({
    body: {
      title: cocktailName.value
    }
  })

  onClose()
  emit('create')
}

</script>

<template>
  <Modal v-bind="$props" title="Create Cocktail">
    <template v-slot:body>
      <label for="cocktail-name" class="flex flex-col">
        Title
        <input type="text" name="cocktail-name" id="cocktail-name" v-model="cocktailName" autofocus
          class="border-2 border-black outline-none rounded-md p-1">
      </label>
    </template>
    <template v-slot:actions>
      <CustomButton outlined @click="onClose">Cancel</CustomButton>
      <CustomButton outlined @click="handleCreate" :loading="createCocktailLoading">Create</CustomButton>
    </template>
  </Modal>
</template>