<script setup lang="ts">

import User from "@/components/icons/user.vue";
import Group from "@/components/icons/group.vue";
</script>
<template>
  <div class="card">
    <h1>Light Notes</h1>
    <div class="column user">
      <user />
      <div class="description">
        <span>我想登录个人账号</span>
        <p>便签数据保障，多端同步，更多高级功能。</p>
      </div>
    </div>
    <div class="column anonymous">
      <group />
      <div class="description">
        <span>我只想随便写写</span>
        <p>方便快捷，迅速传发，即用即走。</p>
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
  height: 500px;
  background-color: rgb(35,35,35);
  border-radius: 12px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  padding: 20px;
}

h1 {
  color: white;
  text-align: center;
  margin: 20px;
}

.column {
  display: flex;
  align-items: center;
  flex-direction: row;
  width: calc(100% - 52px);
  margin: 26px;
  padding: 8px 16px;
  border-radius: 8px;
  transition: .25s;
  border: 1px solid rgba(50,50,50);
  user-select: none;
  cursor: pointer;
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
  border: 1px solid #70c000;
}

.column.user:hover {
  background: rgba(88, 166, 255, .1);
  border: 1px solid #58a6ff;
}
</style>
