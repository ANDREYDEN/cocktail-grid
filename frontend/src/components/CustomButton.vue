<script setup lang="ts">
export interface CustomButtonProps {
    iconPosition?: 'left' | 'right'
    outlined: boolean
}
withDefaults(defineProps<CustomButtonProps>(), {
    iconPosition: 'right',
    outlined: false,
})
</script>

<template>
    <button 
        class="flex flex-row items-center gap-1 p-2 hover:bg-blue-100 rounded-lg" 
        :class="{ 'border-2 border-black': outlined }"
        @click="$emit('click')"
    >
        <div v-if="$slots.icon && iconPosition === 'left'" class="w-5 h-5">
            <slot name="icon"></slot>
        </div>
        <slot></slot>
        <div v-if="$slots.icon && iconPosition === 'right'" class="w-5 h-5">
            <slot name="icon"></slot>
        </div>
    </button>
</template>