<script setup lang="ts">
import { RouterView } from "vue-router";
</script>

<template>
  <router-view v-slot="{ Component }">
    <transition name="fade">
      <component :is="Component" />
    </transition>
  </router-view>
  <footer>
    <span class="copyright"
    >© 2023
      <a href="https://lightxi.com" target="_blank">LightXi Cloud</a></span
    >
    <a class="icp" href="https://beian.miit.gov.cn/" target="_blank">
      <img src="/gov.webp" alt="icp" />
      <span>粤ICP备2023066011号-1</span>
    </a>
  </footer>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: 0.5s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  filter: blur(10px);
}

.fade-leave-to {
  display: none;
}

footer {
  display: flex;
  flex-direction: column;
  position: absolute;
  bottom: 18px;
  right: 12px;
  z-index: 0;
}

footer .copyright {
  width: max-content;
  color: rgba(255, 255, 255, 0.6);
  margin-bottom: 4px;
  transition: 0.25s;
  margin-left: auto;
  user-select: none;
}

footer .copyright a {
  color: rgba(255, 255, 255, 0.6);
  background: none;
}

footer .copyright a:hover {
  color: rgba(255, 255, 255, 0.8);
}

.icp {
  color: rgba(255, 255, 255, 0.6);
  margin-left: auto;
  text-decoration: none;
  background: none;
  transition: 0.5s;
  user-select: none;
}

.icp img {
  width: 16px;
  height: 16px;
  margin-right: 4px;
  vertical-align: middle;
  opacity: 0.8;
  transition: 0.5s;
}

.icp:hover {
  color: #fff;
}

.icp:hover img {
  opacity: 1;
}

@media (min-width: 500px) {
  footer {
    bottom: 26px;
    right: 32px;
  }
}
</style>
