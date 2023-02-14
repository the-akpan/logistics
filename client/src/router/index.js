import { createRouter, createWebHistory } from "vue-router";
import { AuthLayout, DashboardLayout, HomeLayout } from "../layouts";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "",
      component: HomeLayout,
      children: [
        {
          path: "/",
          name: "index",
          component: () => import("../views/site/HomeView.vue"),
        },
        {
          path: "/about",
          name: "about",
          component: () => import("../views/site/AboutView.vue"),
        },
        {
          path: "/services",
          name: "services",
          component: () => import("../views/site/ServicesView.vue"),
        },
        {
          path: "/track",
          name: "tracking",
          component: () => import("../views/site/TrackingView.vue"),
        },
        {
          path: "/contact-us",
          name: "contact",
          component: () => import("../views/site/ContactView.vue"),
        },
      ],
    },
    {
      path: "/auth",
      name: "auth",
      component: AuthLayout,
      children: [
        {
          path: "login",
          name: "login",
          component: () => import("../views/auth/LoginView.vue"),
        },
        {
          path: "reset-password",
          name: "reset-password",
          component: () => import("../views/auth/ResetPasswordView.vue"),
        },
      ],
    },
    {
      path: "/dashboard",
      name: "dashboard",
      component: DashboardLayout,
      children: [
        {
          path: "",
          name: "dashboard",
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
