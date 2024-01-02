import MainLayoutVue from "src/layouts/MainLayout.vue";
import WarehouseStorehousePage from "src/pages/warehouse/StorehousePage.vue";

export default [
  {
    path: "/warehouse",
    component: () => MainLayoutVue,
    children: [
      {
        path: "storehouse",
        name: "warehouseStorehouse:index",
        component: WarehouseStorehousePage,
      },
    ],
  },
];
