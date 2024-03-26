import { createApp } from 'vue'

import './styles.css'
import { createAuth0 } from '@auth0/auth0-vue'
import App from './App.vue'
import { VueQueryPlugin } from '@tanstack/vue-query'

const auth0Plugin = createAuth0({
    domain: import.meta.env.VITE_AUTH0_DOMAIN,
    clientId: import.meta.env.VITE_AUTH0_CLIENT_ID,
    authorizationParams: {
        redirect_uri: window.location.origin,
        audience: import.meta.env.VITE_AUTH0_AUDIENCE,
        scope: "delete:cocktail_ingredient"
    },
})

createApp(App)
.use(auth0Plugin)
.use(VueQueryPlugin)
.mount('#app')
