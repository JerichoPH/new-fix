import MainLayoutVue from "src/layouts/MainLayout.vue";
import BreakdownTypePage from "src/pages/breakdown/TypePage.vue";

export default [
  {
    path: "/breakdown",
    component: () => MainLayoutVue,
    children: [
      {
        path: "type",
        name: "breakdownType:index",
        component: BreakdownTypePage,
      },
    ],
  },
];
