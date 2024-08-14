import { ComponentPublicInstance, computed, ComputedRef, Ref, ref, VNodeRef } from "vue"

export interface ModalState {
    isOpen: boolean
    onClose: () => void
    onOpen: () => void
}

export const useModal = (): ComputedRef<ModalState> => {
    const isOpen = ref(false)

    const onOpen = () => {
        isOpen.value = true
    }

    const onClose = () => {
        isOpen.value = false
    }

    const result = computed(() => ({ isOpen: isOpen.value, onOpen, onClose }))
    return result
}