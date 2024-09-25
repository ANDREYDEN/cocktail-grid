<script setup lang="ts">
import { computed, ref } from 'vue';
import { useAuth } from '@hooks/useAuth';

const { isLoading, user } = useAuth();
const hasPicture = computed(() => !!user.value?.picture)
const pictureError = ref()
const firstLetter = computed(() => {
    return user.value?.name?.[0] ?? '?'
})
</script>

<template>
    <div v-if="isLoading" class="bg-slate-300 w-12 h-12 rounded-full animate-pulse"></div>
    <img v-else-if="hasPicture && !pictureError" :src="user?.picture" alt="Profile Picture"
        class="w-12 h-12 rounded-full" @error="pictureError = true" />
    <div v-else class="bg-slate-300 w-12 h-12 rounded-full capitalize flex items-center justify-center">
        {{ firstLetter }}
    </div>
</template>