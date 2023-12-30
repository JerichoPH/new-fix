import MainLayoutVue from "src/layouts/MainLayout.vue";
import FactoryPage from "src/pages/FactoryPage.vue";

export default [
  {
    path: "/factory",
    component: () => MainLayoutVue,
    children: [{ path: "", name: "factory:index", component: FactoryPage }],
  },
];
