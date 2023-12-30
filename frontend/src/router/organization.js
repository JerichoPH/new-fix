import MainLayoutVue from "src/layouts/MainLayout.vue";

import OrganizationRailwayPage from "src/pages/organization/RailwayPage.vue";
import OrganizationParagraphPage from "src/pages/organization/ParagraphPage.vue";
import OrganizationWorkshopPage from "src/pages/organization/WorkshopPage.vue";
import OrganizationStationPage from "src/pages/organization/StationPage.vue";
import OrganizationCrossroadPage from "src/pages/organization/CrossroadPage.vue";
import OrganizationCenterPage from "src/pages/organization/CenterPage.vue";
import OrganizationWorkAreaPage from "src/pages/organization/WorkAreaPage.vue";
import OrganizationLinePage from "src/pages/organization/LinePage.vue";

export default [
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
      {
        path: "center",
        name: "organizationCenter:index",
        component: OrganizationCenterPage,
      },
      {
        path: "workArea",
        name: "organizationWorkArea:index",
        component: OrganizationWorkAreaPage,
      },
      {
        path: "line",
        name: "organizationLine:index",
        component: OrganizationLinePage,
      },
    ],
  },
];
