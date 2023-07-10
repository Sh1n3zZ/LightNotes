<script setup lang="ts">
import { ref } from "vue";
import axios from "axios";

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
  console.log(res);
}

getNotes();
</script>

<template>
  <div class="card">
    <div class="user">
      {{ username }}
    </div>
  </div>
</template>

<style scoped>
.card {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: min(760px, 90%);
  height: max-content;
  overflow: hidden;
  background-color: rgb(30, 30, 30);
  border-radius: 12px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  padding: 20px;
  transition: 0.25s, max-height 0.5s;
  z-index: 1;
}

.user {
  margin-right: 12px;
  user-select: none;
  float: right;
}
</style>
