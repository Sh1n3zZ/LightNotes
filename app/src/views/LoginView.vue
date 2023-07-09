<script setup lang="ts">
import { onMounted, ref } from "vue";
import axios from "axios";

const token = ref("");
onMounted(async () => {
  const url = new URL(location.href);
  token.value = url.searchParams.get("token") || "";
  const res = await axios.post("https://deeptrain.vercel.app/api/user", {
    token: token.value,
  });
  console.log(res.data);
});
</script>

<template>

</template>

<style scoped>

</style>
