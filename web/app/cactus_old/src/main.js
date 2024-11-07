import { createApp } from 'vue'
import './assets/css/main.css';
import App from './App.vue'
import router from './router'
import pinia from "./stores/store";

createApp(App)
    .use(router)
    .use(pinia)
    .mount('#app')
