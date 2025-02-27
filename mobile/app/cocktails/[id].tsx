import { useLocalSearchParams, useNavigation } from "expo-router";
import { View, Text } from "react-native";

export default function CocktailPage() {
    const { id } = useLocalSearchParams<'/cocktails/[id]'>()
    return (
        <View>
            <Text>Cocktail Page {id}</Text>
        </View>
    )
}