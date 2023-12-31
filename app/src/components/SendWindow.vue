<script setup lang="ts">
import PopupWindow from "@/components/PopupWindow.vue";
import { reactive, ref, watch } from "vue";
import { _window } from "@/assets/script/shared";
import { copyText } from "@/assets/script/clipboard";
import Loading from "@/components/icons/loading.vue";
import axios from "axios";
import Notification from "@/components/Notification.vue";

const support = ref(true);
const loader = ref(false);
const form = reactive({
  title: "",
  body: "",
});

watch(_window, (val) => {
  if (!val.send) {
    code.value = "";
    form.title = "";
    form.body = "";
  }
});

const code = ref("");
const message = ref("");
watch(message, (val) => {
  if (!val) return;
  setTimeout(() => (message.value = ""), 5000);
});

async function send() {
  if (loader.value) return;
  loader.value = true;
  if (!form.title || !form.body) {
    message.value = "标题和内容不能为空！";
    loader.value = false;
    return;
  }

  try {
    const data = (await axios.post("/anonymous/send", form)).data;
    if (data.status) {
      code.value = data.code;
      message.value = "发送成功！";
      form.title = "";
      form.body = "";
    } else {
      message.value = "发送失败！请稍后重试";
    }
  } catch (e) {
    message.value = "发送失败！请检查您的网络环境";
  }

  loader.value = false;
}

async function copy() {
  if (!code.value) return;
  if (await copyText(code.value)) {
    message.value = "复制成功！";
  } else {
    support.value = false;
    message.value = "复制失败！请手动复制";
  }
}
</script>

<template>
  <Notification v-if="message">
    {{ message }}
  </Notification>
  <PopupWindow title="发送" v-model="_window.send">
    <div class="form" v-if="code">
      <span class="message">您的接签码为：</span>
      <div class="code">
        <div
          v-for="(value, index) in code"
          :key="index"
          :style="{ 'animation-delay': index * 100 + 'ms' }"
        >
          {{ value }}
        </div>
      </div>
      <button class="button copy" @click="copy" v-if="support">
        <span>复制</span>
      </button>
      <input class="copy-input" v-model="code" readonly v-else>
    </div>
    <div class="form" v-else>
      <div class="column">
        <div class="row">
          <input
            type="text"
            placeholder="请输入标题"
            maxlength="120"
            v-model="form.title"
          />
        </div>
        <div class="divider" style="background: var(--card-border)" />
        <div class="row textarea">
          <textarea
            placeholder="请输入便签内容"
            maxlength="10240"
            v-model="form.body"
          />
        </div>
      </div>
      <button class="button" @click="send">
        <loading class="loading" v-if="loader" />
        <span v-else>发送</span>
      </button>
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

.divider {
  margin: 2px 0;
}

.form .row {
  width: 100%;
}

.code {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  justify-content: center;
  width: 100%;
  margin: 12px 0;
}

.code div {
  height: max-content;
  font-size: 42px;
  padding: 2px 12px;
  margin: 4px;
  opacity: 0;
  color: var(--text-color-active);
  background: var(--card-background);
  border-radius: 4px;
  animation: FadeInAnimation 1s ease-in-out forwards;
  user-select: none;
}

.copy-input {
  text-align: center;
  width: 40vh;
  font-size: 18px !important;
}

.message {
  font-size: 24px;
  margin: 42px !important;
  text-align: center;
  animation: FadeInAnimation 1s ease-in-out;
}

.loading {
  fill: var(--text-color-full);
  width: 16px;
  height: 16px;
  margin: 4px;
  animation: RotateAnimation 2s linear infinite;
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
  margin: 10px 8px 0 !important;
}

.form textarea {
  scrollbar-width: none;
  font-size: 14px;
  margin-bottom: 12px;
}

.form button {
  justify-content: center;
  transform: translateY(18px);
}

.button {
  display: flex;
  flex-direction: row;
  width: 76px;
  height: 38px;
  padding: 6px;
  margin: 4px 12px;
  border: 1px solid var(--card-border);
  border-radius: 6px;
  background: var(--card-element);
  color: var(--text-color-full);
  cursor: pointer;
  transition: 0.25s;
}

.button.copy {
  margin-top: 94px;
}

.button span {
  margin: 4px;
}

.button:hover {
  background: var(--card-border);
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

@media (max-width: 620px) {
  .code {
    flex-wrap: wrap;
  }
  .code div {
    font-size: 32px;
    padding: 2px;
  }
}
</style>
