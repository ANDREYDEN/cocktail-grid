import Svg, { Path, SvgProps } from "react-native-svg";

export default function OptionsIcon({ stroke = "currentColor" }: SvgProps) {
  return (
    <Svg fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke={stroke}>
      <Path
        stroke-linecap="round"
        stroke-linejoin="round"
        d="M12 6.75a.75.75 0 1 1 0-1.5.75.75 0 0 1 0 1.5ZM12 12.75a.75.75 0 1 1 0-1.5.75.75 0 0 1 0 1.5ZM12 18.75a.75.75 0 1 1 0-1.5.75.75 0 0 1 0 1.5Z"
      />
    </Svg>
  );
}
