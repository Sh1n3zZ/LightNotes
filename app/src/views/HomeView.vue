<script setup lang="ts">
import { onMounted, ref, watch } from "vue";
import { MdEditor } from "md-editor-v3";
import type { Themes } from "md-editor-v3";
import "md-editor-v3/lib/style.css";

import { tools } from "@/assets/script/config";
import { username } from "@/assets/script/shared";
import Arrow from "@/components/icons/arrow.vue";
import { api } from "@/assets/script/note";
import Plus from "@/components/icons/plus.vue";
import { formatDate } from "@/assets/script/utils";
import Loading from "@/components/icons/loading.vue";
import Trash from "@/components/icons/trash.vue";
import Dialog from "@/components/Dialog.vue";

const pagination = new api.NotePagination();
const data = pagination.getRef();
const theme = (window.matchMedia('(prefers-color-scheme: dark)').matches ? "dark" : "light") as Themes;


pagination.update();

const deleter = ref(false);
const loader = ref(false);
const toggle = ref(false);
const editor = ref(false);
const id = ref(0);
const text = ref("");
const title = ref("");
const mobile = ref(document.body.clientWidth < 620);
const sync = ref(true);
const syncTimer = ref<Date>(new Date());
const syncText = ref("刚刚");

const list = ref<HTMLElement>();

let timer: number;
watch(text, () => {
  sync.value = false;
  save();
})
watch(title, () => {
  sync.value = false;
  save();
})

onMounted(() => {
  if (!list.value) return;
  const element: HTMLElement = list.value;
  element.addEventListener("scroll", () => {
    const height = element.scrollTop + element.clientHeight;
    const offset = element.scrollHeight - height;
    if (offset < 20) pagination.update();
  });
});

async function create() {
  const id = await api.newNote();
  if (id !== undefined) {
    await activeEditor(id);
    pagination.new(id);
  }
}

function save() {
  clearTimeout(timer);
  timer = Number(setTimeout(async () => {
    await api.saveNoteById(id.value, title.value, text.value);
    syncTimer.value = new Date();
    sync.value = true;
  }, 2000));
}

async function activeEditor(_id: number) {
  editor.value = true;
  loader.value = true;
  id.value = _id;
  const note = await api.getNoteById(_id);
  loader.value = false;
  if (note) {
    syncTimer.value = new Date();
    text.value = note.body;
    title.value = note.title;
  }
}

async function closeEditor() {
  editor.value = false;
  toggle.value = false;
  sync.value = true;
  pagination.save(id.value, title.value, text.value);
  id.value = 0;
  text.value = "";
  title.value = "";
}

function onDelete() {
  editor.value = false;
  toggle.value = false;
  sync.value = true;
  pagination.delete(id.value);
  id.value = 0;
  text.value = "";
  title.value = "";
}

setInterval(() => {
  syncText.value = formatDate(syncTimer.value, false);
}, 1000);
</script>

<template>
  <Dialog v-model="deleter" @check="onDelete">确定删除该便签？</Dialog>
  <div class="card editor" v-if="editor">
    <div class="header">
      <arrow class="arrow" @click="closeEditor" /><div class="grow" />
      <div class="title">
        <loading class="loading" v-if="loader" />
        <input v-model="title" v-if="toggle">
        <span v-else @click="toggle = true">{{ title }}</span>
      </div>
      <div class="grow" />
      <trash class="delete" @click="deleter = true" />
      <div class="user">
        <div class="status" :class="{'sync': sync, 'error': loader}" />
        <span class="name">{{ syncText }}</span>
      </div>
    </div>
    <MdEditor v-model="text" :theme="theme" :toolbars="tools" :preview="!mobile" />
  </div>
  <div class="card" v-else>
    <div class="header">
      <div class="title">Notes</div>
      <div class="grow" />
      <plus class="new" @click="create" />
      <div class="user">
        <div class="status sync" />
        <span class="name">{{ username }}</span>
      </div>
    </div>
    <div class="list" ref="list">
      <div class="item" v-for="item in data" @click="activeEditor(item.id)">
        <div class="header">
          <div class="title">{{ item.title }}</div><div class="grow" />
          <div class="time">{{ formatDate(item.updated_at) }}</div>
        </div>
        <div class="description">{{ item.body }}</div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.card {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: min(620px, 90%);
  max-width: 90%;
  height: max-content;
  max-height: 80vh;
  overflow: hidden;
  background-color: var(--card-background);
  border-radius: 12px;
  box-shadow: var(--card-shadow);
  padding: 20px;
  transition: 0.25s, max-height 0.5s;
  z-index: 1;
}

.card.editor {
  width: min(1020px, 90%);
}

.grow {
  flex-grow: 1;
}

.list {
  display: flex;
  flex-direction: column;
  gap: 2px;
  width: 100%;
  height: 100%;
  overflow-x: hidden;
  overflow-y: auto;
  touch-action: pan-y;
  scrollbar-width: thin;
  max-height: 540px;
}

.item {
  display: flex;
  flex-direction: column;
  width: calc(100% - 12px);
  height: max-content;
  padding: 8px 32px;
  margin: 6px 4px;
  border-radius: 6px;
  border: 1px solid rgba(0, 0, 0, 0.1);
  background: var(--card-element);
  cursor: pointer;
  transition: .5s;
}

.item .title {
  font-size: 20px;
  font-weight: 600;
  color: var(--card-title);
  margin: 16px 4px 0;
  user-select: none;
  max-width: 80%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.item .description {
  font-size: 16px;
  font-weight: 400;
  color: var(--card-text);
  margin: 0 4px 8px;
  user-select: none;
  max-width: 90%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.header {
  display: flex;
  flex-direction: row;
  margin-top: 4px;
}

.title {
  color: var(--card-title);
  margin: 4px;
  user-select: none;
  transform: translateY(-12px);
  max-width: 60%;
  overflow: hidden;
}

.title input {
  background: var(--card-input);
  border: 1px solid rgb(0,0,0,0);
  color: var(--text-color-active);
  margin: 2px 4px;
  padding: 16px;
  width: calc(100% - 4px);
  height: 24px;
  border-radius: 16px;
  outline: none;
  text-align: center;
  transition: .25s;
}

.title input:focus,
.title input:active {
  border: 1px solid var(--card-input-border);
}

.title span {
  font-size: 24px;
  font-weight: 600;
  color: var(--card-title);
  text-overflow: ellipsis;
  white-space: nowrap;
}

.delete {
  fill: var(--card-text);
  padding: 6px;
  margin: 2px;
  width: 32px;
  height: 32px;
  cursor: pointer;
  background: var(--card-element);
  border-radius: 4px;
  transform: translateY(-6px);
  flex-shrink: 0;
}

.time {
  font-size: 16px;
  font-weight: 400;
  color: var(--card-text);
  margin: 8px 4px 4px;
  user-select: none;
  max-width: 40%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.new {
  fill: var(--card-text);
  padding: 6px;
  margin: 2px;
  width: 32px;
  height: 32px;
  cursor: pointer;
  background: var(--card-element);
  border-radius: 4px;
  transform: translateY(-6px);
  flex-shrink: 0;
}

.arrow {
  fill: var(--card-text);
  padding: 6px;
  margin: 2px;
  width: 32px;
  height: 32px;
  cursor: pointer;
  background: var(--card-element);
  border-radius: 50%;
  transform: translateY(-6px);
  flex-shrink: 0;
}

.user {
  border-radius: 4px;
  padding: 4px 8px;
  background: var(--card-element);
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

.status.error {
  background: #f00;
}

.loading {
  fill: var(--card-text);
  width: 24px;
  height: 24px;
  margin: 0 4px;
  transform: translateY(6px);
  animation: RotateAnimation 2s linear infinite;
}

.user .name {
  color: var(--card-text);
}

@media (max-width: 520px) {
  .title {
    height: max-content;
    max-width: 40%;
    text-overflow: fade;
  }

  .user .name {
    display: none;
  }
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
