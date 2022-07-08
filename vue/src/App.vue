<script setup>
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://vuejs.org/api/sfc-script-setup.html#script-setup
// import HelloWorld from './components/HelloWorld.vue'
import Captcha from "./components/Captcha.vue"
import { ref } from 'vue'
import gaming from './abi/gaming.json'
import Web3 from 'web3/dist/web3.min.js';

// 由于vite删除一些多余的包。无法直接引入web3.js包

const SOL_ADDR = "0x69C1F9B128ab25D46535aD4627007D80b48Ce507";

// let currentAccount = null;
// ethereum.request({method: 'eth_requestAccounts'}).then((accounts) => {
//   if (accounts.length === 0) {
//     // MetaMask is locked or the user has not connected any accounts
//     console.log('Please connect to MetaMask.');
//   } else if (accounts[0] !== currentAccount) {
//     currentAccount = accounts[0];
//   }
// }).catch((err) => {
//   if (err.code === 4001) {
//     console.log('Please connect to MetaMask.');
//   } else {
//     console.error(err);
//   }
// });

let web3 = new Web3(Web3.givenProvider || 'https://evmtestnet.confluxrpc.com/0x69C1F9B128ab25D46535aD4627007D80b48Ce507');

let myContract = new web3.eth.Contract(gaming, SOL_ADDR);
myContract.methods.getTotalBet().call().then(res => {
    console.log(res);
})
const input = ref('')

function submit() {
  // selectedAddress  获取当前用户账号
  myContract.methods.bet(2).send({from: ethereum.selectedAddress,value: web3.utils.toWei('2', 'ether')}).then((result) =>{
    console.log(result);
  }).catch(err => {
    console.log(err);
  })
}

myContract.methods.getNumberOfBets().call().then(res => {
  console.log(res)
})
</script>

<template>
  <img alt="Vue logo" src="./assets/logo.png" />
  <Captcha />
  <!-- <HelloWorld msg="Hello Vue 3 + Vite" /> -->
   <!-- <el-input v-model="input" placeholder="Please input" />
   <el-button type="primary" @click="submit" >提交</el-button> -->
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
