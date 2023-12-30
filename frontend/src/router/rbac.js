import MainLayoutVue from "src/layouts/MainLayout.vue";
import RbacRolePage from "src/pages/rbac/RolePage.vue";
import RbacPermissionPage from "src/pages/rbac/PermissionPage.vue";
import RbacMenuPage from "src/pages/rbac/MenuPage.vue";

export default [
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
];
