import { ref, watch } from "vue";

export const token = ref(localStorage.getItem("token") || "");
watch(token, () => {
  localStorage.setItem("token", token.value);
});
