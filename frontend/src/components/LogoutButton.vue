<script setup lang="ts">
import { useAuth0 } from '@auth0/auth0-vue';
import { ArrowRightStartOnRectangleIcon } from '@heroicons/vue/24/solid';
import CustomButton from './CustomButton.vue';

export interface LogoutButtonProps {
    loading: boolean
}
withDefaults(defineProps<LogoutButtonProps>(), { loading: false })
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
    <CustomButton :skeleton-loading="loading" @click="logout">
        Log out
        <template v-slot:icon>
            <ArrowRightStartOnRectangleIcon />
        </template>
    </CustomButton>
</template>