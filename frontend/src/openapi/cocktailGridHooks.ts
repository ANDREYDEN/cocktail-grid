import { useAuth } from "@hooks/useAuth";
import { useMutation } from "@tanstack/vue-query";
import { createCocktail, CreateCocktailError, createCocktailIngredient, CreateCocktailIngredientError, CreateCocktailIngredientVariables, CreateCocktailVariables, createIngredient, CreateIngredientError, CreateIngredientVariables, deleteCocktail, DeleteCocktailError, deleteCocktailIngredient, DeleteCocktailIngredientError, DeleteCocktailIngredientVariables, DeleteCocktailVariables, deleteIngredient, DeleteIngredientError, DeleteIngredientVariables, updateCocktail, UpdateCocktailError, updateCocktailIngredient, UpdateCocktailIngredientError, UpdateCocktailIngredientVariables, UpdateCocktailVariables, updateIngredient, UpdateIngredientError, UpdateIngredientVariables } from "./cocktailGridComponents";
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

export function useCreateCocktail() {
    return useMutationWithAuth<Record<string, any>, CreateCocktailError, CreateCocktailVariables>(
        'createCocktail', 
        createCocktail
    )
}

export function useUpdateCocktail() {
    return useMutationWithAuth<Record<string, any>, UpdateCocktailError, UpdateCocktailVariables>(
        'updateCocktail', 
        updateCocktail
    )
}

export function useDeleteCocktail() {
    return useMutationWithAuth<Record<string, any>, DeleteCocktailError, DeleteCocktailVariables>(
        'deleteCocktail', 
        deleteCocktail
    )
}

export function useCreateIngredient() {
    return useMutationWithAuth<Record<string, any>, CreateIngredientError, CreateIngredientVariables>(
        'createIngredient', 
        createIngredient
    )
}

export function useUpdateIngredient() {
    return useMutationWithAuth<Record<string, any>, UpdateIngredientError, UpdateIngredientVariables>(
        'updateIngredient', 
        updateIngredient
    )
}

export function useDeleteIngredient() {
    return useMutationWithAuth<Record<string, any>, DeleteIngredientError, DeleteIngredientVariables>(
        'deleteIngredient', 
        deleteIngredient
    )
}

function useMutationWithAuth<TResult=Record<string, any>,TError=void, TParams={}>(key: string, fetchFn: (params: TParams) => Promise<any>) {
    const auth = useAuth()
    const mutationFnWithAuth = async (params: TParams) => {
        const token = await auth.getToken()
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