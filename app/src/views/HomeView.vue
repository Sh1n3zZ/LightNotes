<script setup lang="ts">
import { ref, watch } from "vue";
import axios from "axios";
import { MdEditor } from "md-editor-v3";
import type { ToolbarNames } from "md-editor-v3";
import "md-editor-v3/lib/style.css";

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

const text = ref("");
const title = ref("");
const mobile = ref(document.body.clientWidth < 620);
let tools: ToolbarNames[] = [
  'save',
  'preview',
  'revoke',
  'next',
  '-',
  'bold',
  'underline',
  'italic',
  '-',
  'strikeThrough',
  'title',
  'sub',
  'sup',
  'quote',
  'unorderedList',
  'orderedList',
  'task',
  '-',
  'codeRow',
  'code',
  'link',
]
getNotes();

watch(text, () => {
  const data = text.value.split("\n")[0];
  title.value = data.replace(/^#*|#+$/g, "");
})
</script>

<template>
  <div class="card">
    <div class="header">
      <arrow class="arrow" /><div class="grow" />
      <div class="title">{{ title }}</div><div class="grow" />
      <div class="user">{{ username }}</div>
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
  width: min(720px, 90%);
  max-width: 90%;
  height: max-content;
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

@media (max-width: 520px) {
  .title {
    height: max-content;
    max-width: 40%;
    text-overflow: fade;
  }
}
</style>
