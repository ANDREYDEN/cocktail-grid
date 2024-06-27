import { Ref, inject, watch } from "vue";
import { ToastData } from "../components/Toast/ToastData";
import { toastInjectionKey } from "../injectionKeys";

export function useErrorToast(error: Ref<any>) {
    const { showToast } = inject(toastInjectionKey) as ToastData

    watch(error, (newError) => {
        if (!newError) return

        showToast({
            ...handleError(error.value),
            severity: 'error',
        })
    })
}

function handleError(error: any): { title: string, description: string } {
    console.log(error)
    let title = error.message ?? 'Error'
    let description = error.stack?.error ?? 'Unknown error occured'

    return { title, description }
}