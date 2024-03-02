<script setup lang="ts">
import { useAuth0 } from '@auth0/auth0-vue';
import { ArrowRightEndOnRectangleIcon } from '@heroicons/vue/24/solid';

const auth0 = useAuth0();

function login() {
    auth0.loginWithRedirect();
}
</script>

<template>
    <button class="p-2 hover:bg-blue-100 rounded-lg" @click="login">
        Log in
        <ArrowRightEndOnRectangleIcon class="w-5 h-5 inline" />
    </button>
</template>