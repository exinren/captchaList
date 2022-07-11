<template>
  <div class="slide">
    <Vue2Verify
      :mode="'fixed'"
      :captcha-type="'blockPuzzle'"
      :captcha-id="'9ca07a9c-c260-50ae-2c13-89cde2f34cb9'"
      :container-id="'#sliderFixed_btn'"
      :img-size="{ width: '330px', height: '155px' }"
    />
  </div>
</template>

<script>
import Vue2Verify from "./verityinfo/Verity.vue"
// https://picsum.photos/id/1084/480/240
export default {
  name: 'MoveCaptcha',
  components: {
    // GoCaptchaBtn,
    Vue2Verify
  },
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
  .slide {
    display: flex;
    justify-content: center;
  }

</style>