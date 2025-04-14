import CocktailListItem from "@/components/cocktails/CocktailListItem";
import Loader from "@/components/Loader";
import { useGetCocktails } from "@/openapi/cocktailGridComponents";
import { FlatList, Text, StyleSheet, RefreshControl } from "react-native";

export default function Index() {
  const {
    data: cocktails,
    isPending: cocktailsLoading,
    refetch: refetchCocktails,
  } = useGetCocktails({
    queryParams: { compact: "true" },
  });

  if (cocktailsLoading) return <Loader />;

  if (cocktails?.length === 0) {
    return <Text style={styles.emptyStateText}>No cocktails found</Text>;
  }

  return (
    <FlatList
      data={cocktails}
      renderItem={(itemInfo) => <CocktailListItem cocktail={itemInfo.item} />}
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
