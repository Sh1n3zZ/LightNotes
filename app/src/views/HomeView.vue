<script setup lang="ts">
import { ref, watch } from "vue";
import axios from "axios";
import { MdEditor } from "md-editor-v3";
import "md-editor-v3/lib/style.css";

import { tools } from "@/assets/script/config";
import { username } from "@/assets/script/shared";
import Arrow from "@/components/icons/arrow.vue";

const total = ref(0);
const page = ref(1);
const data = ref([]);

async function getNotes() {
  const resp = await axios.get(`/user/list?page=${page.value}`);
  const res = resp.data;
  total.value = res.total;
}
getNotes();

const text = ref("");
const title = ref("");
const mobile = ref(document.body.clientWidth < 620);
const sync = ref(true);

let timer: number;
watch(text, () => {
  const data = text.value.split("\n")[0];
  title.value = data.replace(/^#*|#+$/g, "");
  sync.value = false;
  clearTimeout(timer);
  timer = setTimeout(async () => {

    sync.value = true;
  }, 3000);
})
</script>

<template>
  <div class="card">
    <div class="header">
      <arrow class="arrow" /><div class="grow" />
      <div class="title">{{ title }}</div><div class="grow" />
      <div class="user">
        <div class="status" :class="{'sync': sync}" />
        {{ username }}
      </div>
    </div>
    <MdEditor v-model="text" theme="dark" :toolbars="tools" :preview="!mobile" />
  </div>
</template>

<style scoped>
.card {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: min(1020px, 90%);
  max-width: 90%;
  height: max-content;
  min-height: 80vh;
  overflow: hidden;
  background-color: rgb(30, 30, 30);
  border-radius: 12px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  padding: 20px;
  transition: 0.25s, max-height 0.5s;
  z-index: 1;
}

.grow {
  flex-grow: 1;
}

.header {
  display: flex;
  flex-direction: row;
  margin-top: 4px;
}

.title {
  font-size: 24px;
  font-weight: 600;
  color: #ddd;
  margin: 4px;
  user-select: none;
  transform: translateY(-12px);
  max-width: 60%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.arrow {
  fill: #ddd;
  padding: 6px;
  margin: 2px;
  width: 32px;
  height: 32px;
  cursor: pointer;
  background: rgb(40,40,40);
  border-radius: 50%;
  transform: translateY(-6px);
  flex-shrink: 0;
}

.user {
  border-radius: 4px;
  padding: 4px 8px;
  background: rgb(40,40,40);
  margin: 6px 4px;
  user-select: none;
  float: right;
  height: max-content;
  transform: translateY(-12px);
}

.status {
  width: 8px;
  height: 8px;
  margin: 9px 4px;
  transform: translateX(-2px);
  border-radius: 50%;
  background: #eac121;
  display: inline-block;
  float: left;
  transition: .5s;
}

.status.sync {
  background: #0fab02;
}

@media (max-width: 520px) {
  .title {
    height: max-content;
    max-width: 40%;
    text-overflow: fade;
  }
}
</style>
