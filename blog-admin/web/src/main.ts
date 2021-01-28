import {createApp} from 'vue'
import App from './App.vue'
import store from './store'
import router from './router'
import ElementPlus from 'element-plus'
import 'normalize.css'

const app = createApp(App)
app.use(ElementPlus)
app.use(store)
app.use(router)

createApp(App).use(store).mount('#app')
