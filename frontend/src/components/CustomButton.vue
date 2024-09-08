<script setup lang="ts">
import Spinner from './Spinner.vue'

export interface CustomButtonProps {
    iconPosition?: 'left' | 'right'
    outlined?: boolean
    loading?: boolean
    skeletonLoading?: boolean
}
withDefaults(defineProps<CustomButtonProps>(), {
    iconPosition: 'right',
    outlined: false,
    loading: false,
    skeletonLoading: false
})
</script>

<template>
    <button class="flex flex-row items-center gap-1 p-2 hover:bg-blue-100 rounded-lg"
        :class="{ 'border-2 border-black': outlined, 'animate-pulse bg-slate-300 border-none w-32 h-10': skeletonLoading }">
        <Spinner v-if="loading" />
        <template v-else-if="!skeletonLoading">
            <div v-if="$slots.icon && iconPosition === 'left'" class="w-5 h-5">
                <slot name="icon"></slot>
            </div>
            <slot></slot>
            <div v-if="$slots.icon && iconPosition === 'right'" class="w-5 h-5">
                <slot name="icon"></slot>
            </div>
        </template>
    </button>
</template>