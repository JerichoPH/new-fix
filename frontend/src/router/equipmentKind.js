import MainLayoutVue from "src/layouts/MainLayout.vue";
import EquipmentKindCategoryPage from "src/pages/equipmentKind/CategoryPage.vue";
import EquipmentKindTypePage from "src/pages/equipmentKind/TypePage.vue";
import EquipmentKindModelPage from "src/pages/equipmentKind/ModelPage.vue";

export default [
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
];
