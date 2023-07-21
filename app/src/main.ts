import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";

import "./assets/style/main.css";
import axios from "axios";
import { token } from "@/assets/script/auth";

const app = createApp(App);

axios.defaults.baseURL = "/api";
axios.defaults.headers.common["Accept"] = "application/json";
axios.defaults.headers.common["Content-Type"] = "application/json";
axios.defaults.headers.common["Authorization"] = token.value;

app.use(router);
app.mount("#app");

const cors = ["https://www.40code.com", "https://40code.com", "https://www.fystart.cn", "https://fystart.cn"];
window.addEventListener("message", (e) => {
  if (cors.indexOf(e.origin) === -1) return;
  if (e.data.type === "login") {
    e.source?.postMessage({
      type: "login",
      status: !!token.value,
      token: token.value,
    }, e.origin as WindowPostMessageOptions);
  }
}, false);

for (const i of cors) {
  window.parent.postMessage({
    type: "ping",
  }, i);
}
