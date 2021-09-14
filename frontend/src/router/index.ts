import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
import Board from "../views/Board.vue";
import Manager from "../views/Manager.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/:boardId?",
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
