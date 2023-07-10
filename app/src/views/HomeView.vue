<script setup lang="ts">
import { ref } from "vue";
import axios from "axios";
import { MdEditor } from "md-editor-v3";
import "md-editor-v3/lib/style.css";
import Arrow from "@/components/icons/arrow.vue";

const total = ref(0);
const username = ref("");
const page = ref(1);
const data = ref([]);
axios.post('/user/state').then(res => {
  if (res.data.status) username.value = res.data.user;
});

async function getNotes() {
  const resp = await axios.get(`/user/list?page=${page.value}`);
  const res = resp.data;
  total.value = res.total;
}

const text = ref("");
const mobile = ref(document.body.clientWidth < 620);
let tools = [
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
  'task', // ^2.4.0
  '-',
  'codeRow',
  'code',
  'link',
];
getNotes();
</script>

<template>
  <div class="card">
    <div class="header">
      <arrow class="arrow" />
      <div class="user">
        {{ username }}
      </div>
    </div>
    <MdEditor v-model="text" theme="dark" :toolbars="tools" :preview="!mobile" code-theme="github" />
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

.header {
  margin-top: 4px;
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
}

.user {
  border-radius: 4px;
  padding: 4px 8px;
  background: rgb(40,40,40);
  margin: 4px;
  user-select: none;
  float: right;
  transform: translateY(-12px);
}
</style>
