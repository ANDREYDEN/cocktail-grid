<script setup lang="ts">
import { provide, ref } from 'vue';
import { ShowInfoParams, Severity } from "./ToastData"
import { toastInjectionKey } from "../../injectionKeys"
import { XMarkIcon } from '@heroicons/vue/16/solid'

const severityToColor: { [key in Severity]: string } = {
    'info': 'bg-blue-300',
    'success': 'bg-green-300',
    'warn': 'bg-yellow-300',
    'error': 'bg-red-300'
}
const title = ref("Title")
const description = ref("Description")
const isOpen = ref(false)
const color = ref(severityToColor['info'])
const currentCloseTimeout = ref<number>()

const showInfoToast = ({ title: newTitle, description: newDescription, severity = 'info' }: ShowInfoParams) => {
    title.value = newTitle
    description.value = newDescription
    color.value = severityToColor[severity]

    open()
}

const open = () => {
    isOpen.value = true

    if (currentCloseTimeout.value) {
        clearTimeout(currentCloseTimeout.value)
    }

    currentCloseTimeout.value = setTimeout(close, 3000);
}

const close = () => {
    isOpen.value = false
}

provide(toastInjectionKey, { showToast: showInfoToast })
</script>

<template>
    <slot />
    <dialog :open="isOpen" class="absolute bottom-2 p-2 w-1/2 rounded-lg" :class="color">
        <button @click="close" class="absolute w-6 h-6 p-1 right-2 top-2 hover:bg-red-500 rounded">
            <XMarkIcon />
        </button>
        <h1 class="text-lg">
            {{ title }}
        </h1>
        <div class="text-sm">
            {{ description }}
        </div>
    </dialog>
</template>