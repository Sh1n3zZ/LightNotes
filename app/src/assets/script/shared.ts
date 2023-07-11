import { watch, ref, reactive } from "vue";
import axios from "axios";
import { auth } from "@/assets/script/auth";

export const _window = reactive({
  send: false,
  receive: false,
});
export const username = ref("");

watch(auth, () => {
  if (!auth.value) return;
  axios.post('/user/state').then(res => {
    if (res.data.status) username.value = res.data.user;
  });
})
