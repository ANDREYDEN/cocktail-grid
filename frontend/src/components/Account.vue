<script setup lang="ts">
import { useAuth } from '#hooks/useAuth';
import LoginButton from './LoginButton.vue';
import LogoutButton from './LogoutButton.vue';
import ProfileImage from './ProfileImage.vue';
import { computed } from 'vue';

const auth = useAuth();
const loading = computed(() => auth.isLoading.value);
</script>

<template>
    <div class="flex flex-col md:flex-row items-center gap-2">
        <ProfileImage v-if="loading || auth.isAuthenticated.value" />
        <LogoutButton v-if="auth.isAuthenticated.value" :loading="loading" />
        <LoginButton v-else :loading="loading" />
    </div>
</template>