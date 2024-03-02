<script setup lang="ts">
import { useAuth0 } from '@auth0/auth0-vue';
import { ArrowRightStartOnRectangleIcon } from '@heroicons/vue/24/solid';

const auth0 = useAuth0();

function logout() {
    auth0.logout({
        logoutParams: {
            returnTo: window.location.origin
        }
    });
}
</script>

<template>
    <button class="p-2 hover:bg-blue-100 rounded-lg" @click="logout">
        Log out
        <ArrowRightStartOnRectangleIcon class="w-5 h-5 inline" />
    </button>
</template>