import { VmsCocktailIngredientVm } from "@/openapi/cocktailGridSchemas";
import { View, Text, StyleSheet, ScrollView, ViewStyle, StyleProp } from "react-native";
import Loader from "./Loader";

type CocktailIngredientsListProps = {
    ingredients: VmsCocktailIngredientVm[]
    isLoading: boolean
    style?: StyleProp<ViewStyle>
}

export default function CocktailIngredientsList({ style, ingredients, isLoading }: CocktailIngredientsListProps) {
    if (isLoading) return <Loader />

    return (
        <View style={style}>
            <Text style={styles.title}>Ingredients</Text>
            <View style={styles.container}>
                <Text>Name</Text>
                <Text>Quantity</Text>
            </View>
            <ScrollView>
                {ingredients.map((ingredient) => (
                    <View key={ingredient.ingredientId} style={styles.container}>
                        <Text>{ingredient.ingredientName}</Text>
                        <Text>{ingredient.quantity}</Text>
                    </View>
                ))}
            </ScrollView>
        </View>
    )
}

const styles = StyleSheet.create({
    title: {
        fontSize: 18,
        textTransform: 'uppercase',
        margin: 10,
    },
    container: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        padding: 10,
        borderBottomWidth: 1,
        borderBottomColor: '#ccc',
    }
})