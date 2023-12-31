import MainLayoutVue from "src/layouts/MainLayout.vue";
import SourceTypePage from "src/pages/source/TypePage.vue";
import SourceProjectPage from "src/pages/source/ProjectPage.vue";

export default [
  {
    path: "/source",
    component: () => MainLayoutVue,
    children: [
      { path: "type", name: "sourceType:index", component: SourceTypePage },
      {
        path: "project",
        name: "sourceProject:index",
        component: SourceProjectPage,
      },
    ],
  },
];
