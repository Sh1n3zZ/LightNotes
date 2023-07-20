<script setup lang="ts">
import User from "@/components/icons/user.vue";
import Group from "@/components/icons/group.vue";
import { reactive } from "vue";
import { _window } from "@/assets/script/shared";
import SendWindow from "@/components/SendWindow.vue";
import RecvWindow from "@/components/RecvWindow.vue";
import { auth } from "@/assets/script/auth";
import router from "@/router";

const active = reactive({
  anonymous: false,
  user: false,
});

function toggleAnonymous() {
  if (active.anonymous) return (active.anonymous = false);
  setTimeout(() => (active.anonymous = true), 100);
  active.user = false;
}

function toggleUser() {
  if (auth.value) return router.push("/home");
  if (active.user) return (active.user = false);
  setTimeout(() => (active.user = true), 100);
  active.anonymous = false;
}

function login() {
  location.href = "https://deeptrain.net/login?app=lightnotes";
}

function register() {
  location.href = "https://deeptrain.net/register?app=lightnotes";
}
</script>
<template>
  <div class="card">
    <div class="title">
      <img src="/favicon.ico" alt="" />
      <h1>Light Notes</h1>
    </div>
    <div class="column user" @click="toggleUser">
      <div class="out">
        <user />
        <div class="description">
          <span>我想登录个人账号</span>
          <p>便签数据保障，多端同步，更多高级功能。</p>
        </div>
      </div>
      <div class="embedded" :class="{ active: active.user }">
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
      <div class="embedded" :class="{ active: active.anonymous }">
        <button @click="_window.send = true" class="button">发送</button>
        <button @click="_window.receive = true" class="button">接收</button>
      </div>
    </div>
  </div>
  <SendWindow v-model="_window.send"> </SendWindow>
  <RecvWindow v-model="_window.receive"> </RecvWindow>
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
  background-color: var(--card-background);
  border-radius: 12px;
  box-shadow: var(--card-shadow);
  padding: 20px;
  transition: 0.25s, max-height 0.5s;
  z-index: 1;
}

.title {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  margin: 20px auto;
}

.title img {
  width: 42px;
  height: 42px;
  margin-right: 12px;
}

h1 {
  font-size: 24px;
  user-select: none;
  color: var(--text-color-active);
  text-align: center;
}

.column {
  display: flex;
  align-items: center;
  flex-direction: column;
  width: calc(100% - 52px);
  margin: 26px;
  padding: 8px 16px;
  border-radius: 8px;
  transition: 0.25s;
  border: 1px solid var(--card-border);
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
  transition: 0.5s;
  will-change: height;
}

.button {
  width: 76px;
  height: 38px;
  padding: 6px;
  margin: 4px 12px;
  border: 1px solid var(--card-border);
  border-radius: 6px;
  background: var(--card-element);
  color: var(--text-color-full);
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
  color: var(--text-color-active);
}

.column .description p {
  font-size: 14px;
  color: var(--text-color-leave);
}

.column svg {
  flex-shrink: 0;
  width: 36px;
  height: 36px;
  margin: 12px 8px;
  fill: var(--text-color-full);
}

.column.user svg {
  padding: 2px;
}

.column .button {
  cursor: pointer;
  transition: 0.25s;
}

.column.anonymous:hover {
  background: rgba(112, 192, 0, 0.1);
  border: 1px solid rgba(112, 192, 0, 0.6);
}

.column.anonymous .button:hover {
  background: rgba(112, 192, 0, 0.8);
  border: 1px solid rgba(112, 192, 0, 0.8);
}

.column.user:hover {
  background: rgba(88, 166, 255, 0.1);
  border: 1px solid rgba(88, 166, 255, 0.6);
}

.column.user .button:hover {
  background: rgba(88, 166, 255, 0.8);
  border: 1px solid rgba(88, 166, 255, 0.8);
}
</style>
