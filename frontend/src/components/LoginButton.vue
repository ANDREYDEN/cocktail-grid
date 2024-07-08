<script setup lang="ts">
import { useAuth0 } from '@auth0/auth0-vue';
import { ArrowRightEndOnRectangleIcon } from '@heroicons/vue/24/solid';
import CustomButton from './CustomButton.vue';

const auth0 = useAuth0();

function login() {
    auth0.loginWithRedirect();
}
</script>

<template>
    <CustomButton @click="login">
        Log in
        <template v-slot:icon>
            <ArrowRightEndOnRectangleIcon />
        </template>
    </CustomButton>
</template>