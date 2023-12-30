import MainLayoutVue from "src/layouts/MainLayout.vue";
import HomePage from "src/pages/HomePage.vue";

export default [
  {
    path: "/",
    component: () => MainLayoutVue,
    children: [{ path: "", name: "home:index", component: HomePage }],
  },
];
