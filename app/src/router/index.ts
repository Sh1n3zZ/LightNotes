import { createRouter, createWebHistory } from "vue-router";
import IndexView from "../views/IndexView.vue";
import { auth, awaitUtilSetup } from "@/assets/script/auth";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "index",
      component: IndexView,
      meta: {
        title: "Light Notes",
      },
    },
    {
      path: "/login",
      name: "login",
      component: () => import("../views/LoginView.vue"),
      meta: {
        title: "Login | Light Notes",
      },
    },
    {
      path: "/home",
      name: "home",
      component: () => import("../views/HomeView.vue"),
      meta: {
        title: "Home | Light Notes",
      }
    }
  ],
});

router.beforeEach(async (to, from, next) => {
  document.title = to.meta.title as string;
  await awaitUtilSetup();
  if (to.name === "login" && auth.value) {
    next({ name: "home" });
    return;
  }
  if (to.name === "home" && !auth.value) {
    next({ name: "login" });
    return;
  }
  next();
});
export default router;
