import ErrorNotFontVue from "pages/ErrorNotFound.vue";
import MainLayoutVue from "layouts/MainLayout.vue";
import AuthLayoutVue from "layouts/AuthLayout.vue";
import IndexPage from "pages/IndexPage.vue";
import LoginPage from "pages/auth/LoginPage.vue";
import RegisterPage from "pages/auth/RegisterPage.vue";
import AccountPage from "pages/AccountPage.vue";
import RbacRolePage from "pages/rbac/RolePage.vue";
import RbacPermissionPage from "pages/rbac/PermissionPage.vue";
import RbacMenuPage from "pages/rbac/MenuPage.vue";
import TestPage from "src/pages/TestPage.vue";

const routes = [
  {
    path: "/",
    component: () => MainLayoutVue,
    children: [{ path: "", name: "home:index", component: IndexPage }],
  },
  {
    path: "/test",
    component: () => MainLayoutVue,
    children: [{ path: "", name: "test:index", component: TestPage }],
  },
  {
    path: "/auth",
    component: () => AuthLayoutVue,
    children: [
      { path: "login", name: "auth:login", component: LoginPage },
      { path: "register", name: "auth:register", component: RegisterPage },
    ],
  },
  {
    path: "/account",
    component: () => MainLayoutVue,
    children: [{ path: "", name: "account:index", component: AccountPage }],
  },
  {
    path: "/rbac",
    component: () => MainLayoutVue,
    children: [
      { path: "role", name: "rbacRole:index", component: RbacRolePage },
      {
        path: "permission",
        name: "rbacPermission:index",
        component: RbacPermissionPage,
      },
      { path: "menu", name: "rbacMenu:index", component: RbacMenuPage },
    ],
  },
  {
    path: "/equipmentKind",
    component: () => MainLayoutVue,
    children: [
      { path: "category", name: "equipmentKindCategory:index", component: () => import("src/pages/equipmentKind/CategoryPage.vue") },
    ],
  },
  // Always leave this as last one,
  // but you can also remove it
  {
    path: "/:catchAll(.*)*",
    component: () => ErrorNotFontVue,
  },
];

export default routes;
