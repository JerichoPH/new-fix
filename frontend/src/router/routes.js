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
import EquipmentKindCategoryPage from "src/pages/equipmentKind/CategoryPage.vue";
import EquipmentKindTypePage from "src/pages/equipmentKind/TypePage.vue";
import EquipmentKindModelPage from "src/pages/equipmentKind/ModelPage.vue";
import OrganizationRailwayPage from "src/pages/organization/RailwayPage.vue";
import OrganizationParagraphPage from "src/pages/organization/ParagraphPage.vue";
import OrganizationWorkshopPage from "src/pages/organization/WorkshopPage.vue";
import OrganizationStationPage from "src/pages/organization/StationPage.vue";
import OrganizationCrossroadPage from "src/pages/organization/CrossroadPage.vue";

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
      {
        path: "category",
        name: "equipmentKindCategory:index",
        component: EquipmentKindCategoryPage,
      },
      {
        path: "type",
        name: "equipmentKindType:index",
        component: EquipmentKindTypePage,
      },
      {
        path: "model",
        name: "equipmentKindModel:index",
        component: () => EquipmentKindModelPage,
      },
    ],
  },
  {
    path: "/organization",
    component: () => MainLayoutVue,
    children: [
      {
        path: "railway",
        name: "organizationRailway:index",
        component: OrganizationRailwayPage,
      },
      {
        path: "paragraph",
        name: "organizationParagraph:index",
        component: OrganizationParagraphPage,
      },
      {
        path: "workshop",
        name: "organizationWorkshop:index",
        component: OrganizationWorkshopPage,
      },
      {
        path: "station",
        name: "organizationStation:index",
        component: OrganizationStationPage,
      },
      {
        path: "crossroad",
        name: "organizationCrossroad:index",
        component: OrganizationCrossroadPage,
      },
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
