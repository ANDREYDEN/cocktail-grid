import { VmsCocktailIngredientVm } from "@/openapi/cocktailGridSchemas";
import {
  View,
  Text,
  StyleSheet,
  ScrollView,
  ViewStyle,
  StyleProp,
} from "react-native";
import Loader from "./Loader";
import EmptyList from "./EmptyList";

type CocktailIngredientsListProps = {
  ingredients: VmsCocktailIngredientVm[];
  isLoading: boolean;
  style?: StyleProp<ViewStyle>;
};

export default function CocktailIngredientsList({
  style,
  ingredients,
  isLoading,
}: CocktailIngredientsListProps) {
  if (isLoading) return <Loader />;

  if (!ingredients.length) {
    return <EmptyList style={{ flex: 1 }} text="No ingredients" />;
  }

  return (
    <View style={style}>
      <Text style={styles.title}>Ingredients</Text>
      <View style={styles.container}>
        <Text style={styles.tableHeader}>Name</Text>
        <Text style={styles.tableHeader}>Quantity</Text>
      </View>
      <ScrollView style={{ height: "100%" }}>
        {ingredients.map((ingredient) => (
          <View key={ingredient.ingredientId} style={styles.container}>
            <Text>{ingredient.ingredientName}</Text>
            <Text>{ingredient.quantity} oz.</Text>
          </View>
        ))}
      </ScrollView>
    </View>
  );
}

const styles = StyleSheet.create({
  title: {
    fontSize: 18,
    textTransform: "uppercase",
    margin: 10,
    marginTop: 20,
  },
  container: {
    flexDirection: "row",
    justifyContent: "space-between",
    padding: 10,
    borderBottomWidth: 1,
    borderBottomColor: "#ccc",
  },
  tableHeader: {
    fontWeight: "bold",
  },
});
