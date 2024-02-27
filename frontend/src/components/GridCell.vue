<script setup lang="ts">
import { computed, ref } from 'vue'

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
    text: ''
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

function handleDelete(e: MouseEvent) {
    e.stopPropagation()
    emit('delete')
}
</script>

<template>
    <div class="relative flex items-center justify-center flex-shrink-0 w-48 h-16 rounded-lg" :class="containerClass"
        @click="$emit('click')" @mouseover="hovering = true" @mouseleave="hovering = false">
        <button v-if="deletable && hovering" class="bg-red-400 text-white w-5 h-5 absolute top-2 right-2"
            @click="handleDelete">x</button>
        <div class="flex items-baseline" :class="textClass">
            <div class="inline">{{ text }}</div>
            <div v-if="!selectable && text != ''" class="ml-1 text-lg">oz.</div>
        </div>
    </div>
</template>