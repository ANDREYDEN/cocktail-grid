import { Auth0VueClient, useAuth0 } from "@auth0/auth0-vue";
import { useMutation } from "@tanstack/vue-query";
import { CreateCocktailIngredientError, CreateCocktailIngredientVariables, DeleteCocktailIngredientError, DeleteCocktailIngredientVariables, createCocktailIngredient, deleteCocktailIngredient } from "./cocktailGridComponents";
import { VmsCocktailIngredientVm } from "./cocktailGridSchemas";

export function useCreateCocktailIngredient() {
    const auth = useAuth0();
    const { mutateAsync: mutateCreateCocktailIngredient } = useMutation<VmsCocktailIngredientVm, CreateCocktailIngredientError, CreateCocktailIngredientVariables>({
        mutationKey: ['createCocktailIngredient'],
        mutationFn: (params) => mutateWithAuth(params, createCocktailIngredient, auth)
    })
    return { mutateCreateCocktailIngredient }
}

export function useDeleteCocktailIngredient() {
    const auth = useAuth0()
    const { mutateAsync: mutateDeleteCocktailIngredient } = useMutation<Record<string, any>, DeleteCocktailIngredientError, DeleteCocktailIngredientVariables>({
        mutationKey: ['deleteCocktailIngredient'],
        mutationFn: (params) => mutateWithAuth(params, deleteCocktailIngredient, auth)
    })
    return { mutateDeleteCocktailIngredient }
}

async function mutateWithAuth<TParams, TResult>(params: TParams, fetchFn: (params: TParams) => Promise<TResult>, auth: Auth0VueClient): Promise<TResult> {
    const token = await auth.getAccessTokenSilently()
    return await fetchFn({
        ...params,
        headers: {
            Authorization: `Bearer ${token}`
        }
    } as TParams)
}