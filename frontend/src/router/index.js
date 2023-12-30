import { route } from "quasar/wrappers";
import {
  createRouter,
  createMemoryHistory,
  createWebHistory,
  createWebHashHistory,
} from "vue-router";

import errorNotFontRouters from "src/router/errorNotFont";
import testRouters from "src/router/test";
import homeRouters from "src/router/home";
import authRouters from "src/router/auth";
import factoryRouters from "src/router/factory";
import accountRouters from "src/router/account";
import rbacRouters from "src/router/rbac";
import equipmentKindRouters from "src/router/equipmentKind";
import organizationRouters from "src/router/organization";
import breakdownRouters from "src/router/breakdown";

const routes = [
  ...homeRouters,
  ...testRouters,
  ...authRouters,
  ...accountRouters,
  ...factoryRouters,
  ...rbacRouters,
  ...equipmentKindRouters,
  ...organizationRouters,
  ...breakdownRouters,
  // Always leave this as last one,
  // but you can also remove it
  ...errorNotFontRouters,
];

/*
 * If not building with SSR mode, you can
 * directly export the Router instantiation;
 *
 * The function below can be async too; either use
 * async/await or return a Promise which resolves
 * with the Router instance.
 */

export default route(function (/* { store, ssrContext } */) {
  const createHistory = process.env.SERVER
    ? createMemoryHistory
    : process.env.VUE_ROUTER_MODE === "history"
    ? createWebHistory
    : createWebHashHistory;

  const Router = createRouter({
    scrollBehavior: () => ({ left: 0, top: 0 }),
    routes,

    // Leave this as is and make changes in quasar.conf.js instead!
    // quasar.conf.js -> build -> vueRouterMode
    // quasar.conf.js -> build -> publicPath
    history: createHistory(process.env.VUE_ROUTER_BASE),
  });

  return Router;
});
