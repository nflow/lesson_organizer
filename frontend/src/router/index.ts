import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
import Main from "../views/Main.vue";
import Manager from "../views/Manager.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/:boardId?",
    name: "Main",
    component: Main,
    props: true,
  },
  {
    path: "/manager",
    name: "Manager",
    component: Manager,
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
