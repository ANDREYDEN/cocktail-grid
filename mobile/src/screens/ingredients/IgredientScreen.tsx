import { Stack } from "expo-router";
import { View, Animated, Text, StyleSheet } from "react-native";

type IngredientScreenProps = {
  id: number;
  name?: string;
};

export default function IngredientScreen({ id, name }: IngredientScreenProps) {
  return (
    <View style={styles.container}>
      <Stack.Screen
        name="ingredientPage"
        options={{ title: name, headerBackTitle: "Back" }}
      />
      <Animated.Image
        // sharedTransitionTag={`ingredient-image-${id}`}
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
