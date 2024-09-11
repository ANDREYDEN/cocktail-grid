<script setup lang="ts">
import Modal from "./Modal/Modal.vue"
import CustomButton from "./CustomButton.vue"
import { ModalState } from './Modal/useModal';
import { ref } from "vue";
import { useCreateCocktail, useUpdateCocktail } from "../openapi/cocktailGridHooks";
import { useErrorToast } from "../hooks/useErrorToast";
import { DtosCocktailDto, VmsCocktailVm } from "../openapi/cocktailGridSchemas";

interface UpsertCocktailProps extends ModalState {
  cocktail?: VmsCocktailVm
}

const { onClose, cocktail } = defineProps<UpsertCocktailProps>()
const emit = defineEmits<{
  (e: 'create'): void
}>()
const { mutateAsync: createCocktail, isPending: createCocktailLoading, error: createCocktailError } = useCreateCocktail()
const { mutateAsync: updateCocktail, isPending: updateCocktailLoading, error: updateCocktailError } = useUpdateCocktail()
useErrorToast(createCocktailError || updateCocktailError)
const cocktailName = ref(cocktail?.title)

const isUpdate = !!cocktail

const handleUpsert = async () => {
  const updatedCocktail: DtosCocktailDto = {
    ...cocktail,
    title: cocktailName.value
  }

  if (isUpdate) {
    await updateCocktail({
      pathParams: {
        cocktailId: cocktail.id!
      },
      body: updatedCocktail
    })
  } else {
    await createCocktail({
      body: updatedCocktail
    })
  }

  onClose()
  emit('create')
}
</script>

<template>
  <Modal v-bind="$props" :title="isUpdate ? 'Update Cocktail' : 'Create Cocktail'">
    <template v-slot:body>
      <label for="cocktail-name" class="flex flex-col">
        Title
        <input type="text" name="cocktail-name" id="cocktail-name" v-model="cocktailName" autofocus
          class="border-2 border-black outline-none rounded-md p-1">
      </label>
    </template>
    <template v-slot:actions>
      <CustomButton outlined @click="onClose">Cancel</CustomButton>
      <CustomButton outlined @click="handleUpsert" :loading="createCocktailLoading || updateCocktailLoading">
        {{ isUpdate ? 'Update' : 'Create' }}
      </CustomButton>
    </template>
  </Modal>
</template>