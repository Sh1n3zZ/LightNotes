<script setup lang="ts">
import { onMounted, ref } from "vue";
import axios from "axios";
import { auth, token } from "@/assets/script/auth";
import router from "@/router";
import Error from "@/components/icons/error.vue";

const message = ref("");
onMounted(async () => {
  const url = new URL(location.href);
  const client = url.searchParams.get("token");
  try {
    const res = await axios.post("/login", {
        token: client,
      }),
      data = res.data;
    if (data.status) {
      token.value = data.token;
      auth.value = true;
      message.value = "登录成功！正在跳转中...";
      await router.push("/home");
    } else {
      message.value = "登录失败！请检查您的账号授权是否过期";
    }
  } catch (e) {
    message.value = "登录失败！请检查您的网络连接";
  }
});
</script>

<template>
  <div class="container">
    <template v-if="message">
      <error class="error" />
      <span class="message">{{ message }}</span>
    </template>
    <template v-else>
      <span class="title">正在校验中...</span>
      <div class="wrapper">
        <div class="circle" />
        <div class="circle" />
        <div class="circle" />
        <div class="shadow" />
        <div class="shadow" />
        <div class="shadow" />
      </div>
    </template>
  </div>
</template>

<style scoped>
.container {
  position: absolute;
  display: flex;
  flex-direction: column;
  align-items: center;
  top: calc(50% - 70px);
  left: 50%;
  transform: translate(-50%, -50%);
}

.error {
  fill: var(--text-color-active);
  width: 46px;
  height: 46px;
  animation: FadeInAnimation 0.2s;
  margin: 16px 0;
}

.message {
  color: var(--markdown-color);
  font-size: 18px;
  text-align: center;
  user-select: none;
  animation: FadeInAnimation 0.25s;
}

@keyframes FadeInAnimation {
  0% {
    opacity: 0;
  }

  100% {
    opacity: 1;
  }
}

.title {
  user-select: none;
  padding: 18px 12px;
  margin-bottom: 24px;
  font-size: 24px;
  color: var(--text-color-active);
  text-align: center;
}

.wrapper {
  width: 200px;
  height: 60px;
  position: relative;
  z-index: 1;
}

.circle {
  width: 20px;
  height: 20px;
  position: absolute;
  border-radius: 50%;
  background-color: var(--text-color-full);
  left: 15%;
  transform-origin: 50%;
  animation: CircleAnimation 0.5s alternate infinite ease;
}

@keyframes CircleAnimation {
  0% {
    top: 60px;
    height: 5px;
    border-radius: 50px 50px 25px 25px;
    transform: scaleX(1.7);
  }

  40% {
    height: 20px;
    border-radius: 50%;
    transform: scaleX(1);
  }

  100% {
    top: 0;
  }
}

.circle:nth-child(2) {
  left: 45%;
  animation-delay: 0.2s;
}

.circle:nth-child(3) {
  left: auto;
  right: 15%;
  animation-delay: 0.3s;
}

.shadow {
  width: 20px;
  height: 4px;
  border-radius: 50%;
  background-color: rgba(0, 0, 0, 0.9);
  position: absolute;
  top: 62px;
  transform-origin: 50%;
  z-index: -1;
  left: 15%;
  filter: blur(1px);
  animation: ShadowAnimation 0.5s alternate infinite ease;
}

@keyframes ShadowAnimation {
  0% {
    transform: scaleX(1.5);
  }

  40% {
    transform: scaleX(1);
    opacity: 0.7;
  }

  100% {
    transform: scaleX(0.2);
    opacity: 0.4;
  }
}

.shadow:nth-child(4) {
  left: 45%;
  animation-delay: 0.2s;
}

.shadow:nth-child(5) {
  left: auto;
  right: 15%;
  animation-delay: 0.3s;
}
</style>
