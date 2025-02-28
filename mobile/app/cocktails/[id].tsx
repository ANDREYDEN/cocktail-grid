import { Stack, useLocalSearchParams } from "expo-router";
import { View, Text, StyleSheet } from "react-native";
import { Image } from "expo-image";
import { useGetCocktail, useGetCocktails } from "@/openapi/cocktailGridComponents";
import CocktailIngredientsList from "@/components/CocktailIngredientsList";

type CocktailPageQueryParams = {
    title?: string
}

export default function CocktailPage() {
    const { id, title } = useLocalSearchParams<'/cocktails/[id]', CocktailPageQueryParams>()
    const { data: cocktail, isLoading: cocktailLoading } = useGetCocktail({ pathParams: { cocktailId: +id } })

    return (
        <View style={styles.container}>
            <Stack.Screen name="CocktailPage" options={{ title, headerBackTitle: 'Back' }} />
            <Image source={{ uri: 'https://picsum.photos/200' }} style={styles.image} />
            <Text style={styles.title}>{title}</Text>
            <CocktailIngredientsList style={{ width: '100%' }} ingredients={cocktail?.ingredients ?? []} isLoading={cocktailLoading} />
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        display: 'flex',
        alignItems: 'center',
    },
    image: {
        width: 200,
        height: 200,
        borderRadius: 10,
        marginVertical: 20,
    },
    title: {
        fontSize: 36,
    }
})