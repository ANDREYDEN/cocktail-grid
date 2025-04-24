import { VmsCocktailVm } from "@/openapi/cocktailGridSchemas";
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

export default function CocktailListItem({
  cocktail,
}: {
  cocktail: VmsCocktailVm;
}) {
  const handleClick = () => {
    router.push(`/cocktails/${cocktail.id}?title=${cocktail.title}`);
  };

  const handleEdit = () => {};
  const handleDelete = () => {};

  return (
    <TouchableOpacity style={styles.container} onPress={handleClick}>
      <View style={styles.leftContent}>
        <Animated.Image
          sharedTransitionTag={`cocktail-image-${cocktail.id}`}
          source={{ uri: "https://picsum.photos/200" }}
          style={styles.image}
        />
        <Text style={styles.title}>{cocktail.title}</Text>
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
