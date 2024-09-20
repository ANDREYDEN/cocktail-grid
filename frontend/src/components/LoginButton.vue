<script setup lang="ts">
import { useAuth0 } from '@auth0/auth0-vue';
import { ArrowRightEndOnRectangleIcon } from '@heroicons/vue/24/solid';
import CustomButton from './CustomButton.vue';

export interface LoginButtonProps {
    loading: boolean
}
const { loading } = withDefaults(defineProps<LoginButtonProps>(), {
    loading: false
})
const auth0 = useAuth0();

function login() {
    auth0.loginWithRedirect();
}
</script>

<template>
    <CustomButton :skeleton-loading="loading" @click="login">
        Log in
        <template v-slot:icon>
            <ArrowRightEndOnRectangleIcon />
        </template>
    </CustomButton>
</template>