<script setup lang="ts">
import PopupWindow from "@/components/PopupWindow.vue";
import { reactive, ref, watch } from "vue";
import { window } from "@/assets/script/shared";
import Loading from "@/components/icons/loading.vue";
import axios from "axios";
import Notification from "@/components/Notification.vue";

const loading = ref(false);
const form = reactive({
  title: "",
  content: "",
});
const message = ref("");
watch(message, (val) => {
  if (!val) return;
  setTimeout(() => (message.value = ""), 5000);
});

async function send() {
  if (loading.value) return;
  loading.value = true;
  if (!form.title || !form.content) {
    message.value = "标题和内容不能为空！";
    loading.value = false;
    return;
  }
  try {
    const data = await axios.post("/anonymous/send", form);
  } catch (e) {
    message.value = "发送失败！请检查您的网络环境";
  }

  loading.value = false;
}
</script>

<template>
  <Notification v-if="message">
    {{ message }}
  </Notification>
  <PopupWindow title="发送" v-model="window.send">
    <div class="form">
      <div class="column">
        <div class="row">
          <input type="text" placeholder="请输入标题" maxlength="120" v-model="form.title" />
        </div>
        <div class="divider" style="background: rgb(50,50,50)" />
        <div class="row textarea">
          <textarea placeholder="请输入便签内容" maxlength="10240" v-model="form.content"></textarea>
        </div>
      </div>
      <button class="button" @click="send">
        <loading class="loading" v-if="loading" />
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

.form .row {
  width: 100%;
}

.loading {
  fill: #fff;
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
</style>
