import MainLayoutVue from "src/layouts/MainLayout.vue";

import InstallIndoorRoomTypePage from "src/pages/install/IndoorRoomTypePage.vue";
import InstallIndoorRoomPage from "src/pages/install/InstallRoomPage.vue";
import InstallIndoorPlatoonPage from "src/pages/install/IndoorPlatoonPage.vue";

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
      {
        path: "indoorRoom",
        name: "installIndoorRoom:index",
        component: InstallIndoorRoomPage,
      },
      {
        path: "indoorPlatoon",
        name: "installIndoorPlatoon:index",
        component: InstallIndoorPlatoonPage,
      },
    ],
  },
];
