import { VmsCocktailVm } from '@/openapi/cocktailGridSchemas'
import { TouchableOpacity, Text } from 'react-native'

export default function CocktailListItem({ cocktail }: { cocktail: VmsCocktailVm }) {
    return (
        <TouchableOpacity>
            <Text>{cocktail.title}</Text>
        </TouchableOpacity>
    )
}