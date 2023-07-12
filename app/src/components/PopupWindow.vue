<script setup lang="ts">
import Close from "@/components/icons/close.vue";
import { onMounted, ref } from "vue";

const props = defineProps<{
  title: string;
  modelValue: boolean;
}>();
const emit = defineEmits(["update:modelValue"]);
const window = ref<HTMLElement | null>(null);

const start = ref<number>(NaN);

onMounted(() => {
  if (window.value === null) return;
  window.value.addEventListener("touchstart", (e) => {
    start.value = e.touches[0].clientY;
  });
  window.value.addEventListener("touchmove", (e) => {
    e.preventDefault();
    if (window.value === null) return;
    const current = e.touches[0].clientY;
    const height = current - start.value;
    start.value = current;
    window.value.scrollTop = window.value.scrollTop - height;
  });
});
</script>

<template>
  <div class="window" :class="{ active: props.modelValue }">
    <h1 class="title">{{ title }}</h1>
    <close
      class="close"
      @click="emit('update:modelValue', false)"
      viewBox="0 0 512 512"
    />
    <div class="divider" />
    <div class="main" ref="window">
      <slot />
    </div>
  </div>
</template>
<style>
.divider {
  width: 100%;
  height: 1px;
  background: var(--card-element);
  margin: 8px 0;
}

.window * {
  color: var(--text-color-full);
  user-select: none;
}

.window .form {
  display: flex;
  gap: 16px;
  flex-direction: column;
  margin: 6px;
  width: 100%;
  height: max-content;
}

.window .form .column {
  background: var(--card-element);
  margin: 0;
  padding: 8px 12px;
  border-radius: 6px;
  display: flex;
  width: 100%;
  flex-direction: column;
  flex-wrap: nowrap;
  gap: 4px;
}

.window .form .row {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  margin: 2px 0;
}

.window .form .row.desc {
  flex-direction: column;
}

.window .form .row.desc p {
  font-size: 14px;
  font-weight: 400;
  margin: 4px 0 4px 28px;
  color: var(--card-desc);
}

.window label {
  font-size: 15px;
  font-weight: 700;
  width: max-content;
  text-wrap: none;
  transform: translateY(10px);
}

.window input {
  background: var(--card-input);
  border: 1px solid var(--card-input-border);
  margin: 10px 4px;
  padding: 16px;
  width: 100%;
  height: 36px;
  border-radius: 16px;
  outline: none;
  letter-spacing: 0.01cm;
  transition: .25s;
}

.window textarea {
  background: var(--card-input);
  border: 1px solid var(--card-input-border);
  margin: 2px 8px;
  padding: 16px;
  width: calc(100% - 16px);
  height: 240px;
  border-radius: 16px;
  outline: none;
  letter-spacing: 0.01cm;
  resize: none;
  font-family: var(--fonts);
  transition: .25s;
}

.window input:hover,
.window textarea:hover {
  border: 1px solid var(--card-input-border-focus);
}

.window input[type="checkbox"] {
  width: 16px;
  height: 16px;
  margin: 0;
  padding: 0;
  border: none;
  outline: none;
}

.window input[type="checkbox"]:checked {
  background: var(--text-color-full);
}

.window select {
  background: var(--card-background);
  border: none;
  border-radius: 4px;
  padding: 4px 8px 4px 8px;
  outline: none;
}

.window select option {
  background: var(--card-element);
  border: none;
  border-radius: 4px;
}
</style>
<style scoped>
.title {
  transform: translateY(-6px);
  text-align: center;
  font-size: 20px;
}

.window {
  position: absolute;
  padding: 26px 0;
  border: 0;
  top: 50%;
  left: 50%;
  transition: 0.5s;
  transform: translate(-50%, -50%);
  width: calc(100% - 52px);
  height: min(80vh, 540px);
  background: var(--card-background);
  border-radius: 10px;
  max-width: 640px;
  z-index: -64;
  opacity: 0;
  overflow: hidden;
}

@keyframes PopupAnimation {
  0% {
    transform: translate(-50%, -50%) scale(0.95);
  }
  50% {
    transform: translate(-50%, -50%) scale(1.05);
  }
  100% {
    transform: translate(-50%, -50%) scale(1);
  }
}

.window.active {
  z-index: 64;
  opacity: 1;
  animation: PopupAnimation 0.5s ease-in-out;
}

.main {
  width: 100%;
  height: 60vh;
  overflow-x: hidden;
  overflow-y: auto;
  touch-action: pan-x;
  padding: 0 26px 36px;
}

.close {
  position: absolute;
  padding: 2px;
  width: 32px;
  height: 32px;
  right: 18px;
  top: 18px;
  cursor: pointer;
  transition: 0.25s;
  border-radius: 6px;
  stroke: var(--text-color-active);
}

.close:hover {
  background: var(--card-element);
  stroke: var(--text-color-active);
}
</style>
