<script setup lang="ts">
import { onMounted, ref, watch } from "vue";
import { contain } from "@/assets/script/utils";

const props = defineProps<{
  modelValue: boolean;
}>();
const emit = defineEmits(["update:modelValue", "check"]);
const dialog = ref<HTMLElement>();
let prevent = false;

watch(props, () => {
  if (props.modelValue === true) {
    prevent = false;
  }
});

onMounted(() => {
  window.addEventListener("click", (e) => {
    if (props.modelValue) {
      if (!prevent) {
        return prevent = true;
      }
      if (!contain(dialog.value, e.target as HTMLElement)) {
        e.preventDefault();
        close();
      }
    }
  })
})

function close() {
  emit("update:modelValue", false);
}

function check() {
  emit("update:modelValue", false);
  emit("check");
}
</script>

<template>
  <div class="cover" :class="{ active: props.modelValue }" />
  <div class="dialog" :class="{ active: props.modelValue }" ref="dialog">
      <slot></slot>
      <div class="operation-list">
        <div class="cancel" @click="close">手滑了</div>
        <div class="click" @click="check">确定</div>
      </div>
  </div>
</template>

<style scoped>
.dialog {
  display: flex;
  flex-direction: column;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  z-index: -64;
  opacity: 0;
  width: 260px;
  height: max-content;
  padding: 24px 24px;
  box-shadow: var(--card-shadow);
  background: var(--card-background);
  color: var(--text-color-full);
  backdrop-filter: blur(2px);
  justify-content: center;
  align-items: center;
  transition: .5s;
  user-select: none;
  border-radius: 12px;
}

.cover {
  position: absolute;
  top: 0;
  left: 0;
  z-index: -64;
  width: 100vw;
  height: 100vh;
  background: rgba(0, 0, 0, 0);
  backdrop-filter: blur(2px);
  transition: .5s;
  border-radius: 12px;
}

.cover.active {
  z-index: 32;
  opacity: 1;
  background: var(--dialog-cover);
}

@keyframes PopupAnimation {
  0% {
    transform: translate(-50%, -50%) scale(0.5);
  }
  50% {
    transform: translate(-50%, -50%) scale(1.05);
  }
  100% {
    transform: translate(-50%, -50%) scale(1);
  }
}

.dialog.active {
  z-index: 64;
  opacity: 1;
  animation: PopupAnimation 0.5s ease-in-out;
}

.operation-list {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: space-between;
  margin: 24px 8px 2px;
}

.operation-list div {
  width: max-content;
  height: min-content;
  padding: 4px 12px;
  margin: 4px 8px;
  border-radius: 4px;
  background: var(--card-element);
  border: 1px solid var(--card-element);
  color: var(--text-color-full);
  transition: .25s;
  cursor: pointer;
}

.cancel {
  color: var(--vt-c-blue) !important;
}

.cancel:hover {
  border: 1px solid var(--vt-c-blue) !important;
}

.click {
  background: var(--vt-c-blue) !important;
  color: #fff !important;
}

.click:hover {
  background: var(--vt-c-blue-hover) !important;
}
</style>
