import CocktailListItem from "@/components/CocktailListItem";
import Loader from "@/components/Loader";
import { useGetCocktails } from "@/openapi/cocktailGridComponents";
import { FlatList } from "react-native";

export default function Index() {
  const { data: cocktails, isLoading: cocktailsLoading } = useGetCocktails({
    queryParams: { compact: "true" },
  });

  if (cocktailsLoading) return <Loader />;

  return (
    <FlatList
      data={cocktails}
      renderItem={(itemInfo) => <CocktailListItem cocktail={itemInfo.item} />}
    />
  );
}
