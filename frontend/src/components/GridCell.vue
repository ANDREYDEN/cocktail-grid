<script setup lang="ts">
import { computed } from 'vue'

export interface GridCellProps {
    isSelectable?: boolean
    selected?: boolean
    text?: string | number
}
const props = withDefaults(defineProps<GridCellProps>(), {
    isSelectable: false,
    selected: false,
    text: ''
})

const containerClass = computed(() => ({
    'cursor-pointer text-lg': props.isSelectable,
    'hover:bg-blue-100': props.isSelectable && !props.selected,
    'bg-blue-300': props.selected,
}))
const textClass = computed(() => ({
    'text-3xl text-gray-500': !props.isSelectable
}))

defineEmits(['click'])
</script>

<template>
    <div 
        class="flex items-center justify-center flex-shrink-0 w-48 h-16 rounded-lg"
        :class="containerClass"
        @click="$emit('click')">
        <div class="flex items-baseline" :class="textClass">
            <div class="inline">{{ text }}</div>
            <div v-if="!isSelectable && text != ''" class="ml-1 text-lg">oz.</div>
        </div>
    </div>
</template>