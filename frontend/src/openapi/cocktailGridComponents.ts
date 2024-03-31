/**
 * Generated by @openapi-codegen
 *
 * @version 1
 */
import type * as Fetcher from "./cocktailGridFetcher";
import { cocktailGridFetch } from "./cocktailGridFetcher";
import type * as Schemas from "./cocktailGridSchemas";

export type GetCocktailsQueryParams = {
  /**
   * exclude ingredients
   */
  compact?: string;
};

export type GetCocktailsError = Fetcher.ErrorWrapper<undefined>;

export type GetCocktailsResponse = Schemas.VmsDetailedCocktailVm[];

export type GetCocktailsVariables = {
  queryParams?: GetCocktailsQueryParams;
};

/**
 * Retrieves all available cocktails
 */
export const getCocktails = (
  variables: GetCocktailsVariables,
  signal?: AbortSignal
) =>
  cocktailGridFetch<
    GetCocktailsResponse,
    GetCocktailsError,
    undefined,
    {},
    GetCocktailsQueryParams,
    {}
  >({ url: "/cocktails", method: "get", ...variables, signal });

export type CreateCocktailError = Fetcher.ErrorWrapper<undefined>;

export type CreateCocktailVariables = {
  body?: Schemas.DtosCreateCocktailDto;
};

/**
 * Creates a new cocktail
 */
export const createCocktail = (
  variables: CreateCocktailVariables,
  signal?: AbortSignal
) =>
  cocktailGridFetch<
    Schemas.VmsCocktailVm,
    CreateCocktailError,
    Schemas.DtosCreateCocktailDto,
    {},
    {},
    {}
  >({ url: "/cocktails", method: "post", ...variables, signal });

export type CreateCocktailIngredientPathParams = {
  /**
   * Cocktail ID
   */
  cocktailId: number;
  /**
   * Ingredient ID
   */
  ingredientId: number;
};

export type CreateCocktailIngredientError = Fetcher.ErrorWrapper<{
  status: 422;
  payload: Record<string, any>;
}>;

export type CreateCocktailIngredientVariables = {
  body?: Schemas.DtosCocktailIngredientDto;
  pathParams: CreateCocktailIngredientPathParams;
};

/**
 * Adds an existing ingredient to the cocktail
 */
export const createCocktailIngredient = (
  variables: CreateCocktailIngredientVariables,
  signal?: AbortSignal
) =>
  cocktailGridFetch<
    Schemas.VmsCocktailIngredientVm,
    CreateCocktailIngredientError,
    Schemas.DtosCocktailIngredientDto,
    {},
    {},
    CreateCocktailIngredientPathParams
  >({
    url: "/cocktails/{cocktailId}/ingredients/{ingredientId}",
    method: "post",
    ...variables,
    signal,
  });

export type DeleteCocktailIngredientPathParams = {
  /**
   * Cocktail ID
   */
  cocktailId: number;
  /**
   * Ingredient ID
   */
  ingredientId: number;
};

export type DeleteCocktailIngredientError = Fetcher.ErrorWrapper<undefined>;

export type DeleteCocktailIngredientVariables = {
  pathParams: DeleteCocktailIngredientPathParams;
};

/**
 * Deletes an ingredient from a cocktail
 */
export const deleteCocktailIngredient = (
  variables: DeleteCocktailIngredientVariables,
  signal?: AbortSignal
) =>
  cocktailGridFetch<
    Record<string, any>,
    DeleteCocktailIngredientError,
    undefined,
    {},
    {},
    DeleteCocktailIngredientPathParams
  >({
    url: "/cocktails/{cocktailId}/ingredients/{ingredientId}",
    method: "delete",
    ...variables,
    signal,
  });

export type GetIngredientsError = Fetcher.ErrorWrapper<undefined>;

export type GetIngredientsResponse = Schemas.VmsIngredientVm[];

/**
 * Retrieves all available ingredients
 */
export const getIngredients = (signal?: AbortSignal) =>
  cocktailGridFetch<
    GetIngredientsResponse,
    GetIngredientsError,
    undefined,
    {},
    {},
    {}
  >({ url: "/ingredients", method: "get", signal });

export type CreateIngredientError = Fetcher.ErrorWrapper<undefined>;

export type CreateIngredientVariables = {
  body?: Schemas.DtosIngredientDto;
};

/**
 * Creates a new ingredient
 */
export const createIngredient = (
  variables: CreateIngredientVariables,
  signal?: AbortSignal
) =>
  cocktailGridFetch<
    Schemas.VmsIngredientVm,
    CreateIngredientError,
    Schemas.DtosIngredientDto,
    {},
    {},
    {}
  >({ url: "/ingredients", method: "post", ...variables, signal });

export const operationsByTag = {
  cocktails: {
    getCocktails,
    createCocktail,
    createCocktailIngredient,
    deleteCocktailIngredient,
  },
  ingredients: { getIngredients, createIngredient },
};
