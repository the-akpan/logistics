import { createRouter, createWebHistory } from "vue-router";
import { DashboardLayout } from "../layouts";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "",
      name: "home",
      component: () => import("../views/site/HomeView.vue"),
    },
    {
      path: "/dashboard",
      name: "dashboard",
      component: DashboardLayout,
      children: [
        {
          path: "",
          name: "index",
          component: () => import("../views/dashboard/HomeView.vue"),
        },
        {
          path: "profile",
          name: "profile",
          component: () => import("../views/dashboard/ProfileView.vue"),
        },
        {
          path: "orders",
          name: "orders",
          component: () => import("../views/dashboard/OrdersView.vue"),
        },
      ],
    },
  ],
});

export default router;
