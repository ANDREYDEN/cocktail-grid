import CocktailListItem from "@/components/CocktailListItem";
import { useGetCocktails } from "@/openapi/cocktailGridComponents";
import { ActivityIndicator, FlatList, Text, View } from "react-native";

export default function Index() {
  const { data: cocktails, isLoading: cocktailsLoading, error: cocktailsError } = useGetCocktails({ queryParams: { compact: 'true' } })

  if (cocktailsLoading) {
    return (
      <View
        style={{
          flex: 1,
          justifyContent: "center",
          alignItems: "center",
        }}
      >
        <ActivityIndicator size="large" />
      </View>
    );
  }

  return (
    <View
      style={{
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <FlatList data={cocktails} renderItem={(itemInfo) => <CocktailListItem cocktail={itemInfo.item} />} />
    </View>
  );
}
