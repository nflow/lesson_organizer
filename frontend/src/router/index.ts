import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
import Landing from "../views/Landing.vue";
import Manager from "../views/Manager.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/:boardId?",
    name: "Landing",
    component: Landing,
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
