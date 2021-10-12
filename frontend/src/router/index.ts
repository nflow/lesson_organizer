import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
import Landing from "../views/Landing.vue";
import Manager from "../views/Manager.vue";
import Board from "../views/Board.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "Landing",
    component: Landing,
  },
  {
    path: "/:boardId",
    name: "Board",
    component: Board,
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
