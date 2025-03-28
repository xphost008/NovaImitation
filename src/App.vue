<script setup lang="ts">
  import NavBar from './views/NavBar.vue';
  import Body from './views/Body.vue'
  import { dark_mode, allowCookie, getCookie, setCookie } from './logic/changeBody'
  import { ref, watch, onMounted } from 'vue'
  import MyDialog from './components/card/MyDialog.vue'
  import { messagebox } from './logic/messagebox';
  const dark = ref(dark_mode.value ? '#1a1a1a' : '#e6e6e6')
  watch(dark_mode, value => dark.value = value ? '#1a1a1a' : '#e6e6e6')
  onMounted(async () => {
    if(getCookie("allowCookie") !== "true") {
      let allow = await messagebox("是否允许使用Cookie？", "是否允许本网站存储Cookie在你的本地？如果你不允许的话，可能会影响到本网站的正常使用。<br>Cookie用来存储你的所有设置信息，你可以自行按F12查看本网站所有的Cookie信息。随时随地添加和删除。<br>本网站除了存储你的离线账号信息以外，不存储你的微软账号登录信息。具体你可以自行按下F12查看本网站源码。<br>如果不允许的话，你每次重启本网站，都需要重新设置一遍你的所有设置信息。【比如暗色模式等】<br>该网站的Cookie将在用户的电脑里存储一个月，一个月之后所有Cookie全部失效。（除非你在此期间再次登录本网站）", 1, ["不允许", "允许"])
      setCookie("allowCookie", (allow == 1 ? true : false).toString())
    }
    allowCookie.value = getCookie("allowCookie") === "true" ? true : false
    // 以下开始判断是否初始化所有的控件
    if(allowCookie.value) {
      dark_mode.value = getCookie("dark_mode") === "true" ? true : false
    }
  })
</script>

<template>
  <div id="all">
    <NavBar id="nav-bar" />
    <main id="main">
      <Body id="body"/>
    </main>
  </div>
  <MyDialog />
</template>

<style scoped>
  #all{
    position: absolute;
    display: flex;
    flex-direction: column;
    height: 100%;
    width: 100%;
  }
  #nav-bar {
    position: absolute;
    width: 100%;
    height: 56px;
    background: linear-gradient(to right, rgb(19, 85, 206), cyan);
    z-index: 100;
  }
  #main {
    position: absolute;
    width: 100%;
    height: calc(100% - 56px);
    top: 56px;
    transition: all 0.2s;
    background-color: v-bind(dark);
  }
  #body {
    position: absolute;
    width: 100%;
    height: calc(100%);
  }
</style>