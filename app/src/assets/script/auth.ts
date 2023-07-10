import { ref, watch } from "vue";
import axios from "axios";

export const auth = ref<boolean>(false);
export const token = ref(localStorage.getItem("token") || "");

watch(token, () => {
  localStorage.setItem("token", token.value);
  axios.defaults.headers.common["Authorization"] = token.value;
});

if (token.value) {
  window.addEventListener('load', () => {
    axios.post("/user/state")
      .then(resp => {
        if (resp.data.state === "ok")
          auth.value = Boolean(resp.data.status);
      })
  })
}
