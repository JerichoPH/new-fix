import MainLayoutVue from "src/layouts/MainLayout.vue";
import TestPage from "src/pages/TestPage.vue";

export default [
  {
    path: "/test",
    component: () => MainLayoutVue,
    children: [{ path: "", name: "test:index", component: TestPage }],
  },
];
