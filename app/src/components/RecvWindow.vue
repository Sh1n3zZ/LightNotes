<script setup lang="ts">
import PopupWindow from "@/components/PopupWindow.vue";
import { reactive, ref, watch } from "vue";
import { window } from "@/assets/script/shared";
import Loading from "@/components/icons/loading.vue";
import axios from "axios";
import Notification from "@/components/Notification.vue";

const loader = ref(false);
const form = reactive({
  title: "",
  body: "",
});

watch(window, val => {
  if (!val.receive) {
    code.value = "";
    form.title = "";
    form.body = "";
  }
});

const get = ref(false);
const code = ref("");
const message = ref("");
watch(message, (val) => {
  if (!val) return;
  setTimeout(() => (message.value = ""), 5000);
});

async function recv() {
  if (loader.value) return;
  loader.value = true;

  const param = code.value.trim();
  if (param.length !== 8) {
    message.value = "请输入正确的接签码";
    loader.value = false;
    return;
  }

  try {
    const data = (await axios.get("/anonymous/get?code=" + param)).data;
    if (data.status) {
      get.value = true;
      message.value = "接收成功！";
      form.title = data.title;
      form.body = data.body;
    } else {
      message.value = "接收失败！请检查您的接签码是否正确，匿名便签是否过期";
    }
  } catch (e) {
    message.value = "接收失败！请检查您的网络环境";
  }

  loader.value = false;
}
</script>

<template>
  <Notification v-if="message">
    {{ message }}
  </Notification>
  <PopupWindow title="接收" v-model="window.receive">
    <div class="form" v-if="!get">
      <span class="message">请输入您的接签码</span>
      <input class="input" v-model="code" maxlength="8" />
      <button class="button" @click="recv" style="transform: translateY(120px)">
        <loading class="loading" v-if="loader" />
        <span v-else>接收</span>
      </button>
    </div>
    <div class="form" v-else>
      <div class="column">
        <div class="row">
          <input type="text" placeholder="标题" v-model="form.title" readonly />
        </div>
        <div class="divider" style="background: rgb(50,50,50)" />
        <div class="row textarea">
          <textarea placeholder="便签内容" v-model="form.body" readonly />
        </div>
      </div>
    </div>
  </PopupWindow>
</template>

<style scoped>
.form {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
  justify-content: center;
}

.form .row {
  width: 100%;
}

.message {
  font-size: 24px;
  margin: 42px !important;
  text-align: center;
  animation: FadeInAnimation 1s ease-in-out;
}

.loading {
  fill: #fff;
  width: 16px;
  height: 16px;
  margin: 4px;
  animation: RotateAnimation 2s linear infinite;
}

.input {
  width: 100%;
  height: 64px !important;
  font-size: 32px !important;
  margin: 0 24px !important;
  letter-spacing: 6px;
  text-align: center;
  border-radius: 12px;
  max-width: 420px;
}

.form .row.textarea {
  flex-direction: column;
}

.form span {
  white-space: nowrap;
  margin: auto 0;
}

.form input {
  height: 38px;
  font-size: 16px;
  margin: 0 8px;
}

.form textarea {
  scrollbar-width: none;
  font-size: 14px;
  margin-bottom: 12px;
}

.form button {
  justify-content: center;
  margin: 20px 12px;
}

.button {
  display: flex;
  flex-direction: row;
  width: 76px;
  height: 38px;
  padding: 6px;
  margin: 4px 12px;
  border: 1px solid rgb(50, 50, 50);
  border-radius: 6px;
  background: rgb(40, 40, 40);
  color: #fff;
  cursor: pointer;
  transition: .25s;
}

.button span {
  margin: 4px;
}

.button:hover {
  background: rgb(50, 50, 50);
}

@keyframes RotateAnimation {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

@keyframes FadeInAnimation {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}
</style>
