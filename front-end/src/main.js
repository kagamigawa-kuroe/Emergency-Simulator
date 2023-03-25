import { createApp } from 'vue'
import App from './App.vue'
import '@/assets/global.css'
import {urgences} from "../store";


const app = createApp(App)
app.use(urgences)
app.mount("#app")

