import { useAuth0 } from "@auth0/auth0-vue";
import { useMutation } from "@tanstack/vue-query";
import { CreateCocktailIngredientError, CreateCocktailIngredientVariables, DeleteCocktailError, DeleteCocktailIngredientError, DeleteCocktailIngredientVariables, DeleteCocktailVariables, UpdateCocktailIngredientError, UpdateCocktailIngredientVariables, createCocktailIngredient, deleteCocktail, deleteCocktailIngredient, updateCocktailIngredient } from "./cocktailGridComponents";
import { VmsCocktailIngredientVm } from "./cocktailGridSchemas";

export function useCreateCocktailIngredient() {
    return useMutationWithAuth<VmsCocktailIngredientVm, CreateCocktailIngredientError, CreateCocktailIngredientVariables>(
        'createCocktailIngredient', 
        createCocktailIngredient
    )
}

export function useUpdateCocktailIngredient() {
    return useMutationWithAuth<VmsCocktailIngredientVm, UpdateCocktailIngredientError, UpdateCocktailIngredientVariables>(
        'updateCocktailIngredient', 
        updateCocktailIngredient
    )
}

export function useDeleteCocktailIngredient() {
    return useMutationWithAuth<Record<string, any>, DeleteCocktailIngredientError, DeleteCocktailIngredientVariables>(
        'deleteCocktailIngredient', 
        deleteCocktailIngredient
    )
}

export function useDeleteCocktail() {
    return useMutationWithAuth<Record<string, any>, DeleteCocktailError, DeleteCocktailVariables>(
        'deleteCocktail', 
        deleteCocktail
    )
}

function useMutationWithAuth<TResult=Record<string, any>,TError=void, TParams={}>(key: string, fetchFn: (params: TParams) => Promise<any>) {
    const auth = useAuth0()
    const mutationFnWithAuth = async (params: TParams) => {
        const token = await auth.getAccessTokenSilently()
        return await fetchFn({
            ...params,
            headers: {
                Authorization: `Bearer ${token}`
            }
        })
   }
    return useMutation<TResult, TError, TParams>({
        mutationKey: [key],
        mutationFn: mutationFnWithAuth
    })
}