import { View, Text, StyleSheet, StyleProp, ViewStyle } from "react-native";

type EmptyListProps = {
  style?: StyleProp<ViewStyle>;
  text?: string;
};

export default function EmptyList({ text, style }: EmptyListProps) {
  return (
    <View style={[styles.container, style]}>
      <Text style={styles.text}>{text ?? "No items"}</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    width: "100%",
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
  },
  text: {
    fontSize: 18,
    color: "#aaa",
  },
});
