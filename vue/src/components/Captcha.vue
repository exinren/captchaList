<template>
  <div>
    <div class="go-captcha-wrap">
      <!-- <GoCaptchaBtn
              class="go-captcha-btn"
              v-model="captStatus"
              width="100%"
              height="50px"
              :image-base64="captBase64"
              :thumb-base64="captThumbBase64"
              @confirm="handleConfirm"
              @refresh="handleRequestCaptCode"
      /> -->
<!--   弹窗方式   -->
     <GoCaptchaBtnDialog
      class="go-captcha-btn"
      :value="captStatus"
      width="100%"
      height="50px"
      :image-base64="captBase64"
      :thumb-base64="captThumbBase64"
      @confirm="handleConfirm"
      @refresh="handleRequestCaptCode"  />

      <!-- tip -->
      <div class="wg-cap-tip">
        <a target="_blank" href="https://github.com/wenlng/go-captcha/releases"><img src="https://img.shields.io/github/v/release/wenlng/go-captcha.svg" alt="version"></a>
        <!-- Github Button -->
        <!-- Place this tag where you want the button to render. -->
        <a class="github-button" href="https://github.com/wenlng/go-captcha" data-size="large"
           data-show-count="true" aria-label="Star wenlng/go-captcha on GitHub">Star</a>
        <!-- Place this tag where you want the button to render. -->
        <a class="github-button" href="https://github.com/wenlng/go-captcha/fork" data-size="large"
           data-show-count="true" aria-label="Fork wenlng/go-captcha on GitHub">Fork</a>
      </div>
    </div>
  </div>
</template>

<script>
  // import GoCaptchaBtn from './GoCaptchaBtn.vue';
  import GoCaptchaBtnDialog from "./GoCaptchaBtnDialog.vue"
  import githubBtn from '../assets/github-btn.js';
  import Qs from "qs"

  export default {
    name: 'Captcha',
    components: {
      // GoCaptchaBtn,
      GoCaptchaBtnDialog
    },
    data(){
      return {
        // 验证码数据
        needCapt: false,
        popoverVisible: true,
        captBase64: '',
        captThumbBase64: '',
        captKey: '',
        captStatus: 'default',
        captExpires: 0,
        captAutoRefreshCount: 0,
      }
    },
    created() {
      githubBtn()
    },
    methods: {
      /**
       * 处理请求验证码
       */
      handleRequestCaptCode() {
        this.captBase64 = ''
        this.captThumbBase64 = ''
        this.captKey = ''

        this.$axios({
          method: 'get',
          url: '/api/go_captcha_data',
        }).then((response)=>{
          const {data = {}} = response;
          if ((data['code'] || 0) === 0) {
            if (this.$lodash.isEmpty(data)) {
              return
            }
            this.captBase64 = data['image_base64'] || ''
            this.captThumbBase64 = data['thumb_base64'] || ''
            this.captKey = data['captcha_key'] || ''
          } else {
            this.$message({
              message: `获取人机验证数据失败`,
              type: 'warning'
            })
          }
        })
      },
      /**
       * 处理验证码校验请求
       */
      handleConfirm(dots) {
        if (this.$lodash.size(dots) <= 0) {
          this.$message({
            message: `请进行人机验证再操作`,
            type: 'warning'
          })
          return
        }

        let dotArr = []
        this.$lodash.forEach(dots, (dot) => {
          dotArr.push(dot.x, dot.y)
        })

        this.$axios({
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
          },
          method: 'post',
          url: '/api/go_captcha_check_data',
          data: Qs.stringify({
            dots: dotArr.join(','),
            key: this.captKey
          })
        }).then((response)=>{
          const {data = {}} = response;

          if ((data['code'] || 0) === 0) {
            this.$message({
              message: `人机验证成功`,
              type: 'success'
            })
            this.captStatus = 'success'
            this.captAutoRefreshCount = 0
          } else {
            this.$message({
              message: `人机验证失败`,
              type: 'warning'
            })
            if (this.captAutoRefreshCount > 5) {
              this.captAutoRefreshCount = 0
              this.captStatus = 'over'
              return
            }

            this.handleRequestCaptCode()
            this.captAutoRefreshCount += 1
            this.captStatus = 'error'
          }
        })
      }
    }
  }
</script>

<style>
  html{
    width: 100%;
    height: 100%;
    background-color: #16202f;
  }
  body{
    margin:0;
    position: relative;
    padding-bottom: 1200px;
    font-family: "Arial", "Microsoft YaHei", "黑体", "宋体", sans-serif;
  }

  .go-captcha-wrap{
    position: absolute;
    top: 450px;
    left: 50%;
    margin-left: -200px;
    width: 400px;
  }

  .go-captcha-btn{
    width: 300px !important;
    margin: 0 auto !important;
  }

  .wg-cap-tip{
    padding: 50px 0 100px;
    font-size: 13px;
    color: #76839b;
    text-align: center;
    line-height: 180%;
    width: 100%;
    max-width: 680px;
  }

  .wg-cap-tip a{
    display: inline-block;
    margin: 0 5px;
  }

  .wg-cap-tip a img{
    height: 28px;
  }

  .wg-cap-tip > span{
    margin: 0 5px;
  }
</style>