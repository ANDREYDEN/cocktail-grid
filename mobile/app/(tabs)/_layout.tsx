import HeaderActions from "@/components/HeaderActions";
import { BottleIcon } from "@/components/icons/bottleIcon";
import { UnderlineIcon } from "@/components/icons/underlineIcon";
import { Tabs } from "expo-router";

export default function TabLayout() {
    return (
        <Tabs>
            <Tabs.Screen name="index" options={{ title: 'Cocktails', tabBarIcon: () => <UnderlineIcon />, headerRight: HeaderActions }} />
            <Tabs.Screen name="ingredients" options={{ title: 'Ingredients', tabBarIcon: () => <BottleIcon /> }} />
        </Tabs>
    )
}