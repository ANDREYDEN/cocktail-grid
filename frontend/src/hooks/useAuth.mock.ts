import { fn } from '@storybook/test'
import * as actual from './useAuth.js'
import { ref } from 'vue'

export * from './useAuth.js'
export const useAuth = fn(actual.useAuth).mockName('useAuth')
export const defaultAuthState: actual.AuthState = {
    isAuthenticated: ref(true),
    isLoading: ref(false),
    user: ref({
        name: 'John Doe'
    })
}