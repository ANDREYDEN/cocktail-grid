import { useState } from 'react'
import { useAuth0, User } from 'react-native-auth0'

export type AuthState = {
    isLoading: boolean
    user: User | null
    isAuthenticated: boolean
    accessToken?: string
    login: () => Promise<void>
    logout: () => Promise<void>
}

export const useAuth = (): AuthState => {
    const { authorize, clearSession, getCredentials, ...auth } = useAuth0()
    const [ accessToken, setAccessToken ] = useState<string>()

    const login = async () => {
        const credentials = await authorize()

        setAccessToken(credentials?.accessToken)
    }

    return {
        ...auth,
        isAuthenticated: !!auth.user,
        accessToken,
        login,
        logout: clearSession,
    }
}