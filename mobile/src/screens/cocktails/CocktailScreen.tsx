import CocktailIngredientsList from "@/screens/cocktails/components/CocktailIngredientsList";
import { useGetCocktail } from "@/openapi/cocktailGridComponents";
import { Stack } from "expo-router";
import { View, Animated, Text, StyleSheet } from "react-native";

type CoctailScreenProps = {
  id: number;
  title?: string;
};

export default function CocktailScreen({ id, title }: CoctailScreenProps) {
  const { data: cocktail, isLoading: cocktailLoading } = useGetCocktail({
    pathParams: { cocktailId: id },
  });
  return (
    <View style={styles.container}>
      <Stack.Screen
        name="CocktailPage"
        options={{ title, headerBackTitle: "Back" }}
      />
      <Animated.Image
        // sharedTransitionTag={`cocktail-image-${id}`}
        source={{ uri: "https://picsum.photos/200" }}
        style={styles.image}
      />
      <Text style={styles.title}>{title}</Text>
      <CocktailIngredientsList
        style={{ width: "100%" }}
        ingredients={cocktail?.ingredients ?? []}
        isLoading={cocktailLoading}
      />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    display: "flex",
    alignItems: "center",
    height: "100%",
  },
  image: {
    width: 200,
    height: 200,
    borderRadius: 10,
    marginVertical: 20,
  },
  title: {
    fontSize: 36,
  },
});
