<script setup lang="ts">
import { computed, ref } from 'vue'
import { CheckCircleIcon, PencilSquareIcon, TrashIcon, XCircleIcon } from '@heroicons/vue/24/solid'

export interface GridCellProps {
    editable?: boolean
    text?: string | number
}
const props = withDefaults(defineProps<GridCellProps>(), {
    editable: false,
    text: undefined
})

const containerClass = computed(() => ({
    'hover:border-2 hover:border-black': props.editable,
    'border-2 border-black': inEditMode.value,
}))

const emit = defineEmits<{
    (e: 'delete'): void
    (e: 'edit', value: string): void
}>()

const hovering = ref(false)
const inEditMode = ref(false)
const candidate = ref<string>(props.text?.toString() ?? '')
const textInput = ref<HTMLInputElement>()
const hasText = computed(() => !['', undefined, NaN].includes(props.text))
const showActions = computed(() => (props.editable && hovering.value) || inEditMode.value)

function handleDelete(e: MouseEvent) {
    e.stopPropagation()
    emit('delete')
}

function handleStartEditing(e: MouseEvent) {
    e.stopPropagation()
    inEditMode.value = true
    textInput.value?.focus()
}

function handleEndEditing() {
    inEditMode.value = false
    emit('edit', candidate.value)
}
</script>

<template>
    <div class="relative flex items-center justify-center w-28 md:w-48 h-10 md:h-14 rounded-lg text-lg"
        :class="containerClass" @mouseover="hovering = true" @mouseleave="hovering = false">
        <div v-if="showActions" class="absolute flex gap-1 top-2 right-2">
            <PencilSquareIcon v-if="!inEditMode" class="text-black w-5 h-5 hover:cursor-pointer"
                @click="handleStartEditing" />
            <TrashIcon v-if="!inEditMode" class="text-red-500 w-5 h-5 hover:cursor-pointer" @click="handleDelete" />
            <CheckCircleIcon v-if="inEditMode" class="text-green-500 w-5 h-5 hover:cursor-pointer"
                @click="handleEndEditing" />
            <XCircleIcon v-if="inEditMode" class="text-red-500 w-5 h-5 hover:cursor-pointer"
                @click="inEditMode = false" />
        </div>
        <div v-if="hasText && !inEditMode" class="flex items-center text-3xl text-gray-500">
            <div class="flex items-baseline justify-center w-24 md:w-44">
                <div class="inline truncate text-ellipsis capitalize">{{ text }}</div>
                <div class="ml-1 text-lg">oz.</div>
            </div>
        </div>
        <input v-if="inEditMode" type="number" ref="textInput" v-model="candidate"
            class="outline-none bg-transparent w-2/3" @keyup.enter="handleEndEditing" />
    </div>
</template>