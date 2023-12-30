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

export default routes;
