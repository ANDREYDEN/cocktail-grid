import { Tabs } from "expo-router";

export default function TabLayout() {
    return (
        <Tabs>
            <Tabs.Screen name="index" options={{ title: 'Cocktails' }} />
            <Tabs.Screen name="ingredients" options={{ title: 'Ingredients' }} />
        </Tabs>
    )
}