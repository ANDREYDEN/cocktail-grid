<script setup lang="ts">
import { useAuth } from '#hooks/useAuth';
import LoginButton from './LoginButton.vue';
import LogoutButton from './LogoutButton.vue';
import ProfileImage from './ProfileImage.vue';

const { isLoading, isAuthenticated } = useAuth();
</script>

<template>
    <div class="flex flex-col md:flex-row items-center gap-2">
        <ProfileImage />
        <LoginButton v-if="!isLoading && !isAuthenticated" />
        <LogoutButton v-else-if="!isLoading" />
    </div>
</template>