<script setup lang="ts">
import { useAuth0 } from '@auth0/auth0-vue';
import LoginButton from './LoginButton.vue';
import LogoutButton from './LogoutButton.vue';

const auth = useAuth0();
</script>

<template>
    <div>
        <LoginButton v-if="!auth.isAuthenticated.value" />
        <div class="px-2" v-if="auth.isAuthenticated.value">{{ auth.user.value?.name }}</div>
        <LogoutButton v-if="auth.isAuthenticated.value" />
    </div>
</template>