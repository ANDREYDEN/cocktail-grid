import { Stack, useLocalSearchParams } from "expo-router";
import { StyleSheet, Text, View } from "react-native";
import Animated from "react-native-reanimated";

type IngredientPageQueryParams = {
  name?: string;
};

export default function IngredientPage() {
  const { id, name } = useLocalSearchParams<
    "/ingredients/[id]",
    IngredientPageQueryParams
  >();
  return (
    <View style={styles.container}>
      <Stack.Screen
        name="ingredientPage"
        options={{ title: name, headerBackTitle: "Back" }}
      />
      <Animated.Image
        sharedTransitionTag={`ingredient-image-${id}`}
        source={{ uri: "https://picsum.photos/200" }}
        style={styles.image}
      />
      <Text style={styles.title}>{name}</Text>
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
