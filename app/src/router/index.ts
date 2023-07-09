import { createRouter, createWebHistory } from "vue-router";
import IndexView from "../views/IndexView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: IndexView,
      meta: {
        title: "Light Notes",
      }
    }, {
      path: "/login",
      name: "login",
      component: () => import("../views/LoginView.vue"),
      meta: {
        title: "Login | Light Notes",
      }
    }
  ],
});

export default router;
