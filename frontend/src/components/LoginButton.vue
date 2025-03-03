<script setup lang="ts">
import { ArrowRightEndOnRectangleIcon } from '@heroicons/vue/24/solid';
import { useAuth } from '@hooks/useAuth';
import CustomButton from './CustomButton.vue';

const auth = useAuth();

function login() {
    auth.login();
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