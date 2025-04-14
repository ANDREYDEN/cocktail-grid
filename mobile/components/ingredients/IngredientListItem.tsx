import { TouchableOpacity, Text, StyleSheet, View } from "react-native";
import { router } from "expo-router";
import Animated from "react-native-reanimated";
import {
  Menu,
  MenuOption,
  MenuOptions,
  MenuTrigger,
} from "react-native-popup-menu";
import OptionsIcon from "@/components/icons/OptionsIcon";
import { VmsIngredientVm } from "@/openapi/cocktailGridSchemas";

export default function IngredientListItem({
  ingredient,
}: {
  ingredient: VmsIngredientVm;
}) {
  const handleClick = () => {
    router.push(`/ingredients/${ingredient.id}?name=${ingredient.name}`);
  };

  const handleEdit = () => {};
  const handleDelete = () => {};

  return (
    <TouchableOpacity style={styles.container} onPress={handleClick}>
      <View style={styles.leftContent}>
        <Animated.Image
          sharedTransitionTag={`ingredient-image-${ingredient.id}`}
          source={{ uri: "https://picsum.photos/200" }}
          style={styles.image}
        />
        <Text style={styles.title}>{ingredient.name}</Text>
      </View>
      <Menu>
        <MenuTrigger>
          <View style={styles.optionsButton}>
            <OptionsIcon stroke="black" />
          </View>
        </MenuTrigger>
        <MenuOptions>
          <MenuOption onSelect={handleEdit}>
            <Text>Edit</Text>
          </MenuOption>
          <MenuOption onSelect={handleDelete}>
            <Text>Delete</Text>
          </MenuOption>
        </MenuOptions>
      </Menu>
    </TouchableOpacity>
  );
}

const styles = StyleSheet.create({
  container: {
    padding: 10,
    borderBottomWidth: 1,
    borderBottomColor: "#ccc",
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "space-between",
  },
  leftContent: {
    flexDirection: "row",
    alignItems: "center",
  },
  title: {
    fontSize: 16,
  },
  image: {
    width: 50,
    height: 50,
    borderRadius: 10,
    marginRight: 10,
  },
  optionsButton: {
    width: 30,
    height: 30,
  },
});
