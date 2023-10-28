<script setup lang="ts">
export interface GridCellProps {
    isSelectable?: boolean
    selected?: boolean
    text?: string | number
}
const { text, isSelectable, selected} = withDefaults(defineProps<GridCellProps>(), {
    isSelectable: false,
    selected: false,
    text: ''
})

defineEmits(['click'])
</script>

<template>
    <div 
        class="flex-1 h-24 flex items-center justify-center"
        :class="{ 
            'cursor-pointer text-gl': isSelectable, 
            'bg-blue-200': selected, 
            'text-3xl text-gray-500': !isSelectable 
        }"
        @click="$emit('click')">
        <div class="flex items-baseline">
            <div>{{ text }}</div>
            <div v-if="!isSelectable && text != ''" class="ml-1 text-lg">oz.</div>
        </div>
    </div>
</template>