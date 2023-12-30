import ErrorNotFontVue from "src/pages/ErrorNotFound.vue";

export default [
  {
    path: "/:catchAll(.*)*",
    component: () => ErrorNotFontVue,
  },
];
