import MainLayoutVue from "src/layouts/MainLayout.vue";

import InstallIndoorRoomTypePage from "src/pages/install/IndoorRoomTypePage.vue";

export default [
  {
    path: "/install",
    component: MainLayoutVue,
    children: [
      {
        path: "indoorRoomType",
        name: "installIndoorRoomType:index",
        component: InstallIndoorRoomTypePage,
      },
    ],
  },
];
