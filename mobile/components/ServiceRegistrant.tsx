import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { Auth0Provider } from "react-native-auth0";
import { MenuProvider } from "react-native-popup-menu";

export default function ServiceRegistrant({
  children,
}: React.PropsWithChildren) {
  const queryClient = new QueryClient();

  return (
    <Auth0Provider
      domain={process.env.EXPO_PUBLIC_AUTH0_DOMAIN}
      clientId={process.env.EXPO_PUBLIC_AUTH0_CLIENT_ID}
    >
      <QueryClientProvider client={queryClient}>
        <MenuProvider>{children}</MenuProvider>
      </QueryClientProvider>
    </Auth0Provider>
  );
}
