import { createApp } from 'vue'

import './styles.css'
import { createAuth0 } from '@auth0/auth0-vue'
import App from './App.vue'

const auth0Plugin = createAuth0({
    domain: import.meta.env.VITE_AUTH0_DOMAIN,
    clientId: import.meta.env.VITE_AUTH0_CLIENT_ID,
    authorizationParams: {
        redirect_uri: window.location.origin
    },
})

createApp(App).use(auth0Plugin).mount('#app')
