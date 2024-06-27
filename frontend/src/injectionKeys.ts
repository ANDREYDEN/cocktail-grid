import { InjectionKey } from "vue";
import { ToastData } from "./components/Toast/ToastData";

export const toastInjectionKey = Symbol() as InjectionKey<ToastData>