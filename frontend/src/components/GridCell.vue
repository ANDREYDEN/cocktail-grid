<script setup lang="ts">
import { computed, reactive, ref } from 'vue'

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
    'cursor-pointer text-gl': props.isSelectable,
    'bg-blue-200': props.selected,
    'text-3xl text-gray-500': !props.isSelectable
}))

defineEmits(['click'])
</script>

<template>
    <div 
        class="flex-1 h-24 flex items-center justify-center"
        :class="containerClass"
        @click="$emit('click')">
        <div class="flex items-baseline">
            <div>{{ text }}</div>
            <div v-if="!isSelectable && text != ''" class="ml-1 text-lg">oz.</div>
        </div>
    </div>
</template>