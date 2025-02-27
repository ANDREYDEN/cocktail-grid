import { VmsCocktailVm } from '@/openapi/cocktailGridSchemas'
import { TouchableOpacity, Text, StyleSheet, View } from 'react-native'
import { Image } from 'expo-image'

export default function CocktailListItem({ cocktail }: { cocktail: VmsCocktailVm }) {
    return (
        <TouchableOpacity style={styles.container}>
            <Image source={{ uri: 'https://picsum.photos/200' }} style={styles.image} />
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