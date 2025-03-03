import { useAuth } from "@/hooks/useAuth";
import { StyleSheet, TouchableOpacity, useColorScheme, View } from "react-native";
import MoonIcon from "./icons/MoonIcon";
import SignInIcon from "./icons/SignInIcon";
import SignOutIcon from "./icons/SignOutIcon";
import SunIcon from "./icons/SunIcon";

export default function HeaderActions() {
    const colorScheme = useColorScheme()
    const { user, login } = useAuth()

    return (
        <View style={styles.container}>
            <TouchableOpacity style={styles.iconButton} onPress={login}>
                {!user ? <SignInIcon /> : <SignOutIcon />}
            </TouchableOpacity>
            <TouchableOpacity style={styles.iconButton}>
                {colorScheme === 'dark' ? <MoonIcon /> : <SunIcon />}
            </TouchableOpacity>
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        flexDirection: 'row',
        alignItems: 'center',
    },
    iconButton: {
        width: 24,
        height: 24,
        marginRight: 10,
    }
})