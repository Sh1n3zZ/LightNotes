import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";

import "./assets/style/main.css";
import axios from "axios";

const app = createApp(App);

axios.defaults.baseURL = "/api";
axios.defaults.headers.common["Accept"] = "application/json";
axios.defaults.headers.common["Content-Type"] = "application/json";

app.use(router);

app.mount("#app");
