import { BottleIcon } from "@/components/icons/bottleIcon";
import { UnderlineIcon } from "@/components/icons/underlineIcon";
import { Tabs } from "expo-router";

export default function TabLayout() {
    return (
        <Tabs>
            <Tabs.Screen name="cocktails" options={{ title: 'Cocktails', tabBarIcon: () => <UnderlineIcon /> }} />
            <Tabs.Screen name="ingredients" options={{ title: 'Ingredients', tabBarIcon: () => <BottleIcon /> }} />
        </Tabs>
    )
}