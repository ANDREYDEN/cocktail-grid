<script setup lang="ts">
import { provide, ref } from 'vue';
import { ShowInfoParams, Severity } from "./ToastData"
import { toastInjectionKey } from "../../injectionKeys"

const severityToColor: { [key in Severity]: string } = {
    'info': 'bg-blue-300',
    'success': 'bg-green-300',
    'warn': 'bg-yellow-300',
    'error': 'bg-red-300'
}
const title = ref<string>()
const description = ref<string>()
const isOpen = ref(false)
const color = ref(severityToColor['info'])

const showInfoToast = ({ title: newTitle, description: newDescription, severity = 'info' }: ShowInfoParams) => {
    title.value = newTitle
    description.value = newDescription
    isOpen.value = true
    color.value = severityToColor[severity]

    setTimeout(() => {
        isOpen.value = false
    }, 3000);
}

provide(toastInjectionKey, { showToast: showInfoToast })
</script>

<template>
    <slot />
    <dialog :open="isOpen" class="absolute bottom-2 p-2 w-1/2 rounded" :class="color">
        <h1 class="text-lg">
            {{ title }}
        </h1>
        <div class="text-sm">
            {{ description }}
        </div>
    </dialog>
</template>