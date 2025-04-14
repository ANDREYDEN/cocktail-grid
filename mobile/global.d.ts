declare namespace NodeJS {
  interface ProcessEnv {
    readonly EXPO_PUBLIC_BACKEND_URL: string;
    readonly EXPO_PUBLIC_AUTH0_DOMAIN: string;
    readonly EXPO_PUBLIC_AUTH0_CLIENT_ID: string;
    readonly EXPO_PUBLIC_AUTH0_AUDIENCE: string;
  }
}
