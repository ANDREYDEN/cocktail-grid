import { QueryClient, QueryClientProvider } from "@tanstack/react-query"
import { Auth0Provider } from "react-native-auth0"

export default function ServiceRegistrant({ children }: React.PropsWithChildren) {
    const queryClient = new QueryClient()

    return <Auth0Provider domain={process.env.EXPO_PUBILC_AUTH0_DOMAIN} clientId={process.env.EXPO_PUBILC_AUTH0_CLIENT_ID}>
        <QueryClientProvider client={queryClient}>
            {children}
        </QueryClientProvider>
    </Auth0Provider>
}