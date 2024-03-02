<script setup lang="ts">
import { computed, ref } from 'vue'
import { TrashIcon } from '@heroicons/vue/24/solid'

export interface GridCellProps {
    selectable?: boolean
    selected?: boolean
    deletable?: boolean
    text?: string | number
}
const props = withDefaults(defineProps<GridCellProps>(), {
    selectable: false,
    selected: false,
    deletable: false,
    text: undefined
})

const containerClass = computed(() => ({
    'cursor-pointer text-lg': props.selectable,
    'hover:bg-blue-100': props.selectable && !props.selected,
    'bg-blue-300': props.selected,
}))
const textClass = computed(() => ({
    'text-3xl text-gray-500': !props.selectable
}))

const emit = defineEmits(['click', 'delete'])

const hovering = ref(false)
const hasText = computed(() => props.text !== '')
const showDeleteButton = computed(() => props.deletable && hovering.value && hasText.value)

function handleDelete(e: MouseEvent) {
    e.stopPropagation()
    emit('delete')
}
</script>

<template>
    <div class="relative flex items-center justify-center flex-shrink-0 w-48 h-16 rounded-lg" :class="containerClass"
        @click="$emit('click')" @mouseover="hovering = true" @mouseleave="hovering = false">
        <TrashIcon v-if="showDeleteButton" class="text-red-500 w-5 h-5 absolute hover:cursor-pointer top-2 right-2"
            @click="handleDelete" />
        <div class="flex items-baseline" :class="textClass">
            <div class="inline" v-if="text !== undefined">{{ text }}</div>
            <slot v-else></slot>
            <div v-if="!selectable && hasText" class="ml-1 text-lg">oz.</div>
        </div>
    </div>
</template>