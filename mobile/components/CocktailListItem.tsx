import { VmsCocktailVm } from '@/openapi/cocktailGridSchemas'
import { TouchableOpacity, Text, StyleSheet } from 'react-native'
import { router } from 'expo-router'
import Animated from 'react-native-reanimated'

export default function CocktailListItem({ cocktail }: { cocktail: VmsCocktailVm }) {
    const handleClick = () => {
        router.push(`/cocktails/${cocktail.id}?title=${cocktail.title}`)
    }

    return (
        <TouchableOpacity style={styles.container} onPress={handleClick}>
            <Animated.Image sharedTransitionTag={`cocktail-image-${cocktail.id}`} source={{ uri: 'https://picsum.photos/200' }} style={styles.image} />
            <Text style={styles.title}>{cocktail.title}</Text>
        </TouchableOpacity>
    )
}

const styles = StyleSheet.create({
    container: {
        padding: 10,
        borderBottomWidth: 1,
        borderBottomColor: '#ccc',
        flexDirection: 'row',
        alignItems: 'center'
    },
    title: {
        fontSize: 16
    },
    image: {
        width: 50,
        height: 50,
        borderRadius: 10,
        marginRight: 10
    }
})