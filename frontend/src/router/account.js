import MainLayoutVue from "src/layouts/MainLayout.vue";
import AccountPage from "src/pages/AccountPage.vue";

export default [
  {
    path: "/account",
    component: () => MainLayoutVue,
    children: [{ path: "", name: "account:index", component: AccountPage }],
  },
];
