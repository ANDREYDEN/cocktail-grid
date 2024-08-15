<script setup lang="ts">
export interface CustomButtonProps {
    iconPosition?: 'left' | 'right'
    outlined?: boolean
    loading?: boolean
}
withDefaults(defineProps<CustomButtonProps>(), {
    iconPosition: 'right',
    outlined: false,
    loading: false
})
</script>

<template>
    <button class="min-w-32 flex flex-row items-center gap-1 p-2 hover:bg-blue-100 rounded-lg"
        :class="{ 'border-2 border-black': outlined }">
        <span v-if="loading" class="w-10 h-10 text-4xl animate-spin">+</span>
        <template v-else>
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