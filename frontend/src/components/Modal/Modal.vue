<script setup lang="ts">
import { ModalState } from './useModal';

interface ModalProps extends ModalState {
    title: string
}

const { onClose } = defineProps<ModalProps>()
const handleBgClick = (e: MouseEvent) => {
    if (e.target === e.currentTarget) {
        onClose()
    }
}
</script>

<template>
    <div class="top-0 left-0 fixed flex justify-center items-center w-screen h-screen backdrop-blur-sm z-10"
        :class="isOpen ? 'block' : 'hidden'" @click="handleBgClick">
        <div class="w-1/2 bg-white p-4 border-2 border-black rounded-md z-20">
            <h2 class="text-xl mb-4">{{ title }}</h2>
            <slot name="body"></slot>
            <div class="flex gap-2 justify-end mt-4">
                <slot name="actions"></slot>
            </div>
        </div>
    </div>
</template>