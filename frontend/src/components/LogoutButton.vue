<script setup lang="ts">
import { ArrowRightStartOnRectangleIcon } from '@heroicons/vue/24/solid';
import { useAuth } from '@hooks/useAuth';
import CustomButton from './CustomButton.vue';

const auth = useAuth();

function logout() {
    auth.logout({ returnTo: window.location.origin });
}
</script>

<template>
    <CustomButton @click="logout">
        Log out
        <template v-slot:icon>
            <ArrowRightStartOnRectangleIcon />
        </template>
    </CustomButton>
</template>