<script setup lang="ts">
import Modal from "./Modal/Modal.vue"
import CustomButton from "./CustomButton.vue"
import { ModalState } from './Modal/useModal';

export interface CreateCocktailModalProps extends ModalState {
}
const props = defineProps<CreateCocktailModalProps>()
console.log({ ccm: props });

const handleCreate = () => {

}

</script>

<template>
  <Modal v-bind="$props" title="Create Cocktail">
    <template v-slot:body>Hi</template>
    <template v-slot:actions>
      <CustomButton outlined @click="onClose">Cancel</CustomButton>
      <CustomButton outlined @click="handleCreate">Create</CustomButton>
    </template>
  </Modal>
</template>./Modal/Modal.vue