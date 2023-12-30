import AuthLayoutVue from "src/layouts/AuthLayout.vue";
import LoginPage from "src/pages/auth/LoginPage.vue";
import RegisterPage from "src/pages/auth/RegisterPage.vue";
import LogoutPage from "src/pages/auth/LogoutPage.vue";

export default [
  {
    path: "/auth",
    component: () => AuthLayoutVue,
    children: [
      { path: "login", name: "auth:login", component: LoginPage },
      { path: "register", name: "auth:register", component: RegisterPage },
      { path: "logout", name: "auth:logout", component: LogoutPage },
    ],
  },
];
