import ServiceRegistrant from "@/components/ServiceRegistrant";
import { Stack } from "expo-router";

export default function RootLayout() {
  return (
    <ServiceRegistrant>
      <Stack>
        <Stack.Screen name="(tabs)" options={{ headerShown: false }} />
      </Stack>
    </ServiceRegistrant>
  )
}
