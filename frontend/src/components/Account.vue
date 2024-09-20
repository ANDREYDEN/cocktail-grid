<script setup lang="ts">
import { useAuth } from '#hooks/useAuth';
import CustomButton from './CustomButton.vue';
import LoginButton from './LoginButton.vue';
import LogoutButton from './LogoutButton.vue';
import ProfileImage from './ProfileImage.vue';
import { computed } from 'vue';

const auth = useAuth();
const loading = computed(() => auth.isLoading.value);
const isAuthenticated = computed(() => auth.isAuthenticated.value)
</script>

<template>
    <div class="flex flex-col md:flex-row items-center gap-2">
        <ProfileImage v-if="loading || isAuthenticated" />
        <CustomButton v-if="loading" :skeleton-loading="loading" />
        <div v-else>
            <LogoutButton v-if="isAuthenticated" />
            <LoginButton v-if="!isAuthenticated" />
        </div>
    </div>
</template>