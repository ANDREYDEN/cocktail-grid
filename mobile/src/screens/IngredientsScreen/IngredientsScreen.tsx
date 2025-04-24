import Loader from "@/components/Loader";
import { FlatList, RefreshControl, Text, StyleSheet } from "react-native";
import IngredientListItem from "./components/IngredientListItem";
import { useGetIngredients } from "@/openapi/cocktailGridComponents";

export default function IngredientsScreen() {
  const {
    data: cocktails,
    isPending: cocktailsLoading,
    refetch: refetchCocktails,
  } = useGetIngredients({});

  if (cocktailsLoading) return <Loader />;

  if (cocktails?.length === 0) {
    return <Text style={styles.emptyStateText}>No ingredients found</Text>;
  }

  return (
    <FlatList
      data={cocktails}
      renderItem={(itemInfo) => (
        <IngredientListItem ingredient={itemInfo.item} />
      )}
      refreshControl={
        <RefreshControl
          refreshing={cocktailsLoading}
          onRefresh={refetchCocktails}
        />
      }
    />
  );
}

const styles = StyleSheet.create({
  emptyStateText: {
    textAlign: "center",
    marginTop: 20,
    fontSize: 18,
    color: "#888",
  },
});
