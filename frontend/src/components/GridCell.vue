<script setup lang="ts">
import { computed, ref } from 'vue'
import { CheckCircleIcon, PencilSquareIcon, TrashIcon } from '@heroicons/vue/24/solid'

export interface GridCellProps {
    selectable?: boolean
    selected?: boolean
    editable?: boolean
    text?: string | number
}
const props = withDefaults(defineProps<GridCellProps>(), {
    selectable: false,
    selected: false,
    editable: false,
    text: undefined
})

const containerClass = computed(() => ({
    'cursor-pointer text-lg': props.selectable,
    'hover:bg-blue-100': props.selectable && !props.selected,
    'hover:border-2 hover:border-black': true,
    'bg-blue-300': props.selected,
}))
const textClass = computed(() => ({
    'text-3xl text-gray-500': !props.selectable
}))

const emit = defineEmits<{
    (e: 'click'): void
    (e: 'delete'): void
    (e: 'edit', value: string): void
}>()

const hovering = ref(false)
const inEditMode = ref(false)
const candidate = ref<string>(props.text?.toString() ?? '')
const textInput = ref<HTMLInputElement>()
const hasText = computed(() => !['', undefined, NaN].includes(props.text))
const showActions = computed(() => props.editable && hovering.value)

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
    <div class="relative flex items-center justify-center flex-shrink-0 w-48 h-16 rounded-lg" :class="containerClass"
        @click="$emit('click')" @mouseover="hovering = true" @mouseleave="hovering = false">
        <div v-if="showActions" class="absolute flex gap-1 top-2 right-2">
            <PencilSquareIcon v-if="!inEditMode" class="text-black w-5 h-5 hover:cursor-pointer"
                @click="handleStartEditing" />
            <TrashIcon v-if="!inEditMode" class="text-red-500 w-5 h-5 hover:cursor-pointer" @click="handleDelete" />
            <CheckCircleIcon v-if="inEditMode" class="text-green-500 w-5 h-5 hover:cursor-pointer"
                @click="handleEndEditing" />
        </div>
        <div>
            <div v-if="hasText" :class="textClass" class="flex items-baseline">
                <div class="inline" v-if="!inEditMode">{{ text }}</div>
                <div v-if="!selectable && !inEditMode" class="ml-1 text-lg">oz.</div>
            </div>
            <input v-if="inEditMode" type="number" ref="textInput" v-model="candidate"
                class="outline-none bg-transparent" @keyup.enter="handleEndEditing" />
            <slot v-if="!hasText"></slot>
        </div>
    </div>
</template>