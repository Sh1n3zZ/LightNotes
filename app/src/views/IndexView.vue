<script setup lang="ts">
import User from "@/components/icons/user.vue";
import Group from "@/components/icons/group.vue";
import { onMounted, reactive, ref } from "vue";
import PopupWindow from "@/components/PopupWindow.vue";


const active = reactive({
  anonymous: false,
  user: false,
})
const window = reactive({
  send: false,
  receive: false,
})

function toggleAnonymous() {
  if (active.anonymous) return active.anonymous = false;
  setTimeout(() => active.anonymous = true, 100);
  active.user = false;
}

function toggleUser() {
  if (active.user) return active.user = false;
  setTimeout(() => active.user = true, 100);
  active.anonymous = false;
}

function login() {
  location.href = "https://deeptrain.vercel.app/login?app=lightnotes";
}

function register() {
  location.href = "https://deeptrain.vercel.app/register?app=lightnotes";
}
</script>
<template>
  <div class="card">
    <h1>Light Notes</h1>
    <div class="column user" @click="toggleUser">
      <div class="out">
        <user />
        <div class="description">
          <span>我想登录个人账号</span>
          <p>便签数据保障，多端同步，更多高级功能。</p>
        </div>
      </div>
      <div class="embedded" :class="{'active': active.user}">
        <button @click="login" class="button">登录</button>
        <button @click="register" class="button">注册</button>
      </div>
    </div>
    <div class="column anonymous" @click="toggleAnonymous">
      <div class="out">
        <group />
        <div class="description">
          <span>我只想随便写写</span>
          <p>方便快捷，迅速传发，即用即走。</p>
        </div>
      </div>
      <div class="embedded" :class="{'active': active.anonymous}">
        <button @click="window.send = true" class="button">发送</button>
        <button @click="window.receive = true" class="button">接收</button>
      </div>
    </div>
  </div>
  <PopupWindow title="发送" v-model="window.send">

  </PopupWindow>
  <footer>
    <span class="copyright">© 2023 LightXi Cloud</span>
    <a class="icp" href="https://beian.miit.gov.cn/" target="_blank">
      <img src="/gov.webp" alt="icp" />
      <span>粤ICP备2023066011号-1</span>
    </a>
  </footer>
</template>

<style scoped>
.card {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: min(500px, 90%);
  height: max-content;
  overflow: hidden;
  background-color: rgb(30,30,30);
  border-radius: 12px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  padding: 20px;
  transition: .25s, max-height .5s;
  z-index: 1;
}

h1 {
  font-size: 24px;
  user-select: none;
  color: white;
  text-align: center;
  margin: 20px;
}

.column {
  display: flex;
  align-items: center;
  flex-direction: column;
  width: calc(100% - 52px);
  margin: 26px;
  padding: 8px 16px;
  border-radius: 8px;
  transition: .25s;
  border: 1px solid rgba(50,50,50);
  user-select: none;
  cursor: pointer;
}

.column .out {
  display: flex;
  flex-direction: row;
  align-items: center;
  width: 100%;
}

.column .embedded {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  overflow: hidden;
  width: 100%;
  justify-content: center;
  margin: 6px auto;
  gap: 6px;
  height: max-content;
  max-height: 0;
  transition: .5s;
  will-change: height;
}

.button {
  width: 76px;
  height: 38px;
  padding: 6px;
  margin: 4px 12px;
  border: 1px solid rgb(50,50,50);
  border-radius: 6px;
  background: rgb(40,40,40);
  color: #fff;
  cursor: pointer;
}

.column .embedded.active {
  max-height: 100px;
}

.column .description {
  display: flex;
  flex-direction: column;
}

.column .description span {
  font-size: 14px;
  font-weight: 700;
  color: #fff;
}

.column .description p {
  font-size: 14px;
  color: rgba(255, 255, 255, .5);
}

.column svg {
  flex-shrink: 0;
  width: 36px;
  height: 36px;
  margin: 12px 8px;
  fill: #fff;
}

.column.user svg {
  padding: 2px;
}

.column .button {
  cursor: pointer;
  transition: .25s;
}

.column.anonymous:hover {
  background: rgba(112, 192, 0, .1);
  border: 1px solid rgba(112, 192, 0, .6);
}

.column.anonymous .button:hover {
  background: rgba(112, 192, 0, .8);
  border: 1px solid rgba(112, 192, 0, .8);
}

.column.user:hover {
  background: rgba(88, 166, 255, .1);
  border: 1px solid rgba(88, 166, 255, .6);
}

.column.user .button:hover {
  background: rgba(88, 166, 255, .8);
  border: 1px solid rgba(88, 166, 255, .8);
}

footer {
  display: flex;
  flex-direction: column;
  position: absolute;
  bottom: 8px;
  right: 12px;
  z-index: 0;
}

footer .copyright {
  width: max-content;
  color: rgba(255, 255, 255, .6);
  margin-bottom: 4px;
  transition: .25s;
  margin-left: auto;
  user-select: none;
}

.icp {
  color: rgba(255,255,255,.6);
  margin-left: auto;
  text-decoration: none;
  background: none;
  transition: .5s;
  user-select: none;
}

.icp img {
  width: 16px;
  height: 16px;
  margin-right: 4px;
  vertical-align: middle;
  opacity: .8;
  transition: .5s;
}

.icp:hover {
  color: #fff;
}

.icp:hover img {
  opacity: 1;
}

@media (min-width: 500px) {
  footer {
    bottom: 26px;
    right: 32px;
  }
}
</style>
