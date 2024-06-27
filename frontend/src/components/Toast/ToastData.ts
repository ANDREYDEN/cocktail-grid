export type Severity = 'info' | 'success' | 'warn' | 'error'

export interface ShowInfoParams {
    title: string
    description: string
    severity?: Severity
}

export interface ToastData {
    showToast: (params: ShowInfoParams) => void
}