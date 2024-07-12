import { createApp } from 'vue'

import './styles.css'
import { createAuth0 } from '@auth0/auth0-vue'
import App from './App.vue'
import { VueQueryPlugin } from '@tanstack/vue-query'

const identityScopes = ["profile", "email"]
var crudScopes =  ["create:cocktail", "delete:cocktail", "create:ingredient", "delete:ingredient", "create:cocktail_ingredient", "delete:cocktail_ingredient", "update:cocktail_ingredient", "delete:cocktail"]
var requestedScopes = [
    ...identityScopes,
    ...crudScopes
]
var scopeString = requestedScopes.join(' ')

var auth0Plugin = createAuth0({
    domain: import.meta.env.VITE_AUTH0_DOMAIN,
    clientId: import.meta.env.VITE_AUTH0_CLIENT_ID,
    authorizationParams: {
        redirect_uri: window.location.origin,
        audience: import.meta.env.VITE_AUTH0_AUDIENCE,
        scope: scopeString
    },
})

createApp(App)
.use(auth0Plugin)
.use(VueQueryPlugin)
.mount('#app')
