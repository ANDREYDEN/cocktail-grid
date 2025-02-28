import { Stack, useLocalSearchParams, useNavigation } from "expo-router";
import { View, Text } from "react-native";

type CocktailPageQueryParams = {
    title?: string
}

export default function CocktailPage() {
    const { id, title } = useLocalSearchParams<'/cocktails/[id]', CocktailPageQueryParams>()

    return (
        <View>
            <Stack.Screen name="CocktailPage" options={{ title, headerBackTitle: 'Back' }} />
            <Text>Cocktail Page {id}</Text>
        </View>
    )
}