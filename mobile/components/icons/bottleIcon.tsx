import Svg, { Path, Rect } from "react-native-svg";

export function BottleIcon() {
    return (
        <Svg fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" className="size-6">
            <Rect x="10" y="0" width="4" height="4" fill="#8B4513" rx="1" />
            <Rect x="8" y="4" width="8" height="16" fill="#b22222" rx="4" />
            <Rect x="9" y="20" width="6" height="2" fill="#800000" rx="1" />
        </Svg>
    )
}