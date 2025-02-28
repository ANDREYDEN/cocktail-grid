import { TouchableOpacity, useColorScheme } from "react-native";
import MoonIcon from "./icons/MoonIcon";
import SunIcon from "./icons/SunIcon";

export default function HeaderActions() {
    const colorScheme = useColorScheme()

    return (
        <TouchableOpacity style={{ width: 24, height: 24, marginRight: 10 }}>
            {colorScheme === 'dark' ? <MoonIcon /> : <SunIcon />}
        </TouchableOpacity>
    )
}