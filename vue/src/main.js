import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import App from './App.vue'
import detectEthereumProvider from '@metamask/detect-provider';
import axios from "axios"
import lodash from "lodash"

const provider = await detectEthereumProvider();

if (provider) {
  startApp(provider); // Initialize your app
} else {
  aelrt('Please install MetaMask!');
}

function startApp(provider) {
  if (provider !== window.ethereum) {
    console.error('Do you have multiple wallets installed?');
  } else {
    console.log('MetaMask Success !');
  }
}

const app = createApp(App)
app.config.globalProperties.$axios = axios
app.config.globalProperties.$lodash = lodash
app.use(ElementPlus, { size: 'small', zIndex: 3000 })
app.mount('#app')
