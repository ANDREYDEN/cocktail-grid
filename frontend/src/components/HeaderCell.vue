<script setup lang="ts">
import { PencilSquareIcon, TrashIcon } from '@heroicons/vue/24/solid'
import { computed, ref } from 'vue'

export interface GridCellProps {
    selected?: boolean
    editable?: boolean
    text?: string | number
}
const props = withDefaults(defineProps<GridCellProps>(), {
    selected: false,
    editable: false,
    text: undefined
})

const containerClass = computed(() => ({
    'hover:bg-blue-100': !props.selected,
    'bg-blue-300': props.selected,
    'hover:border-2 hover:border-black': props.editable,
}))

const emit = defineEmits<{
    (e: 'click'): void
    (e: 'delete'): void
    (e: 'edit'): void
}>()

const hovering = ref(false)
const hasText = computed(() => !['', undefined, NaN].includes(props.text))
const showActions = computed(() => (props.editable && hovering.value))

function handleDelete(e: MouseEvent) {
    e.stopPropagation()
    emit('delete')
}

function handleEdit(e: MouseEvent) {
    e.stopPropagation()
    emit('edit')
}
</script>

<template>
    <div class="cursor-pointer relative flex items-center justify-center w-28 md:w-48 h-10 md:h-14 rounded-lg text-lg"
        :class="containerClass" @click="$emit('click')" @mouseover="hovering = true" @mouseleave="hovering = false">
        <div v-if="showActions" class="absolute flex gap-1 top-2 right-2">
            <PencilSquareIcon class="text-black w-5 h-5 hover:cursor-pointer" @click="handleEdit" />
            <TrashIcon class="text-red-500 w-5 h-5 hover:cursor-pointer" @click="handleDelete" />
        </div>
        <div v-if="hasText" class="flex items-center justify-center w-24 md:w-44">
            <div class="inline truncate text-ellipsis capitalize">{{ text }}</div>
        </div>
        <slot v-else></slot>
    </div>
</template>