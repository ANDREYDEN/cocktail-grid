import IngredientScreen from "@/screens/ingredients/IgredientScreen";
import { useLocalSearchParams } from "expo-router";

type IngredientPageQueryParams = {
  name?: string;
};

export default function IngredientPage() {
  const { id, name } = useLocalSearchParams<
    "/ingredients/[id]",
    IngredientPageQueryParams
  >();

  return <IngredientScreen id={+id} name={name} />;
}
