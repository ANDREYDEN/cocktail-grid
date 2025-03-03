import { useAuth0, User } from "@auth0/auth0-vue"
import { Ref } from "vue"

export type AuthState = {
    isLoading: Ref<boolean>
    user: Ref<User | undefined>
    isAuthenticated: Ref<boolean>
    login: () => Promise<void>
    logout: (options: LogoutParams) => Promise<void>
    getToken: () => Promise<string>
}

type LogoutParams = {
    returnTo: string
}

export const useAuth = (): AuthState => {
    const { loginWithRedirect, getAccessTokenSilently, logout, ...auth } = useAuth0()

    return {
        ...auth,
        login: loginWithRedirect,
        getToken: getAccessTokenSilently,
        logout: async ({ returnTo }) => {
            await logout({ logoutParams: { returnTo } })
        }
    }
}