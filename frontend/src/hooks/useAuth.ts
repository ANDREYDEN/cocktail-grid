import { useAuth0, User } from "@auth0/auth0-vue"
import { ref, Ref } from "vue"

export type AuthState = {
    isLoading: Ref<boolean>
    user: Ref<User | undefined>
    isAuthenticated: Ref<boolean>
}

export const useAuth = (): AuthState => {
    return useAuth0()
}