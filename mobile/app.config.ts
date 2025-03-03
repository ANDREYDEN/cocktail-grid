import { ConfigContext } from '@expo/config';

export default ({ config }: ConfigContext) => {
    return {
        ...config,
        plugins: [
            ...config.plugins ?? [],
            [
                "react-native-auth0",
                {
                    "domain": process.env.EXPO_PUBLIC_AUTH0_DOMAIN,
                }
            ]
        ]
    }
}