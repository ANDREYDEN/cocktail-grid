<script setup lang="ts">
import { useAuth0 } from '@auth0/auth0-vue';

const auth0 = useAuth0();

function login() {
    auth0.loginWithRedirect();
}
</script>

<template>
    <button @click="login">Log in</button>
</template>