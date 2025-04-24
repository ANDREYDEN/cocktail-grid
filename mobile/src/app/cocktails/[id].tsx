import CocktailScreen from "@/screens/cocktails/CocktailScreen";
import { useLocalSearchParams } from "expo-router";

type CocktailPageQueryParams = {
  title?: string;
};

export default function CocktailPage() {
  const { id, title } = useLocalSearchParams<
    "/cocktails/[id]",
    CocktailPageQueryParams
  >();

  return <CocktailScreen id={+id} title={title} />;
}
