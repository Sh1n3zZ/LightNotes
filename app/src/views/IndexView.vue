<script setup lang="ts">
import User from "@/components/icons/user.vue";
import Group from "@/components/icons/group.vue";
import { onMounted, reactive, ref } from "vue";

const card = ref<HTMLElement>();
const height = ref(NaN);
const active = reactive({
  anonymous: false,
  user: false,
})

function toggleAnonymous() {
  if (active.anonymous) return active.anonymous = false;
  active.anonymous = true;
  active.user = false;
}

function toggleUser() {
  if (active.user) return active.user = false;
  active.user = true;
  active.anonymous = false;
}

onMounted(() => {
  const container = card.value;
  if (!container) return;
  const observer = new MutationObserver(() => {
    height.value = container.offsetHeight;
  });
  observer.observe(container, { childList: true});
})
</script>
<template>
  <div class="card" ref="card" v-bind:style="{'max-height': `${height}px`}">
    <h1>Light Notes</h1>
    <div class="column user" @click="toggleUser">
      <div class="out">
        <user />
        <div class="description">
          <span>我想登录个人账号</span>
          <p>便签数据保障，多端同步，更多高级功能。</p>
        </div>
      </div>
      <div class="embedded" v-if="active.user">
        <button>登录</button>
      </div>
    </div>
    <div class="column anonymous" @click="toggleAnonymous">
      <div class="out">
        <group />
        <div class="description">
          <span>我只想随便写写</span>
          <p>方便快捷，迅速传发，即用即走。</p>
        </div>
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
  width: min(500px, 90%);
  height: max-content;
  overflow: hidden;
  background-color: rgb(35,35,35);
  border-radius: 12px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  padding: 20px;
  transition: .25s, max-height .5s;
}

h1 {
  font-size: 24px;
  user-select: none;
  color: white;
  text-align: center;
  margin: 20px;
}

.column {
  display: flex;
  align-items: center;
  flex-direction: column;
  width: calc(100% - 52px);
  margin: 26px;
  padding: 8px 16px;
  border-radius: 8px;
  transition: .25s;
  border: 1px solid rgba(50,50,50);
  user-select: none;
  cursor: pointer;
}

.column .out {
  display: flex;
  flex-direction: row;
  align-items: center;
  width: 100%;
}

.column .description {
  display: flex;
  flex-direction: column;
}

.column .description span {
  font-size: 14px;
  font-weight: 700;
  color: #fff;
}

.column .description p {
  font-size: 14px;
  color: rgba(255, 255, 255, .5);
}

.column svg {
  width: 36px;
  height: 36px;
  margin: 12px 8px;
  fill: #fff;
}

.column.user svg {
  padding: 2px;
}

.column.anonymous:hover {
  background: rgba(112, 192, 0, .1);
  border: 1px solid rgba(112, 192, 0, .6);
}

.column.user:hover {
  background: rgba(88, 166, 255, .1);
  border: 1px solid rgba(88, 166, 255, .6);
}
</style>
