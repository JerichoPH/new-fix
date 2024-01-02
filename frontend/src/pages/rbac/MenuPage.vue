<template>
  <div class="q-pa-md">
    <q-card>
      <q-card-section>
        <div class="row q-mb-md">
          <div class="col"><span style="font-size: 20px">搜索</span></div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="primary" label="搜索" icon="search" @click="fnSearch" />
              <q-btn color="primary" label="重置" flat @click="fnResetSearch" />
            </q-btn-group>
          </div>
        </div>
        <div class="row">
          <div class="col">
            <q-form>
              <div class="row q-col-gutter-sm">
                <div class="col-3">
                  <q-input outlined clearable lazy-rules v-model="name_search" label="名称" :rules="[]" />
                </div>
                <div class="col-3">
                  <q-input outlined clearable lazy-rules v-model="uri_search" label="路由" :rules="[]" />
                </div>
                <div class="col-3">
                  <q-input outlined clearable lazy-rules v-model="description_search" label="描述" :rules="[]" />
                </div>
                <div class="col-3">
                  <sel-rbac-menu_search label-name="所属父级" v-if="selRbacMenu_search_enable" />
                </div>
                <div class="col-3">
                  <sel-rbac-role_search label-name="所属角色" />
                </div>
              </div>
            </q-form>
          </div>
        </div>
      </q-card-section>
    </q-card>

    <q-card class="q-mt-md">
      <q-card-section>
        <div class="row">
          <div class="col">
            <span :style="{ fontSize: '20px' }">菜单列表</span>
          </div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="secondary" label="新建菜单" icon="add" @click="fnOpenAlertCreateRbacMenu" />
              <q-btn color="negative" label="删除菜单" icon="delete" @click="fnDestroyRbacMenus" />
            </q-btn-group>
          </div>
        </div>
        <div class="row q-mt-md">
          <div class="col">
            <q-table flat bordered title="" :rows="rows" row-key="index" virtual-scroll :pagination="{ rowsPerPage: 200 }"
              :rows-per-page-options="[50, 100, 200, 0]" rows-per-page-label="分页" selection="multiple"
              v-model:selected="selected">
              <template v-slot:header="props">
                <q-tr :props="props">
                  <q-th align="left"><q-checkbox key="allCheck" v-model="props.selected" /></q-th>
                  <q-th align="left">#</q-th>
                  <q-th align="left" key="name" @click="(event) => fnColumnReverseSort(event, props, sortBy)
                    ">
                    名称
                  </q-th>
                  <q-th align="left" key="uri" @click="(event) => fnColumnReverseSort(event, props, sortBy)
                    ">
                    路由
                  </q-th>
                  <q-th align="left" key="description">描述</q-th>
                  <q-th align="left" key="icon">图标</q-th>
                  <q-th align="left" key="parent">所属父级</q-th>
                  <q-th align="left" key="rbacRoles">所属角色</q-th>
                  <q-th align="right"></q-th>
                </q-tr>
              </template>
              <template v-slot:body="props">
                <q-tr :props="props">
                  <q-td><q-checkbox :key="props.row.uuid" :value="props.row.uuid" v-model="props.selected" /></q-td>
                  <q-td>{{ props.row.index }}</q-td>
                  <q-td key="name" :props="props">{{ props.row.name }}</q-td>
                  <q-td key="uri" :props="props">{{
                    props.row.uri ?? "-"
                  }}</q-td>
                  <q-td key="description" :props="props">
                    {{ props.row.description ?? "-" }}
                  </q-td>
                  <q-td>
                    <q-icon :name="props.row.icon" v-if="props.row.icon" />
                    <span v-else>-</span>
                  </q-td>
                  <q-td key="parent" :props="props">
                    <q-chip color="primary" text-color="white" v-if="props.row.parent">
                      {{ props.row.parent.name }}
                    </q-chip>
                  </q-td>
                  <q-td key="rbacRoles" :props="props">
                    <template v-if="props.row.rbacRoles.length > 0">
                      <q-chip color="primary" text-color="white" v-for="(rbacRole, idx) in props.row.rbacRoles"
                        :key="idx">
                        {{ rbacRole.name }}
                      </q-chip>
                    </template>
                    <template v-else><span>-</span></template>
                  </q-td>
                  <q-td key="operation" :props="props">
                    <q-btn-group>
                      <q-btn @click="fnOpenAlertEditRbacMenu(props.row.operation)" color="warning" icon="edit">
                        编辑
                      </q-btn>
                      <q-btn @click="fnDeconsteRbacMenu(props.row.operation)" color="negative" icon="delete">
                        删除
                      </q-btn>
                    </q-btn-group>
                  </q-td>
                </q-tr>
              </template>
            </q-table>
          </div>
        </div>
      </q-card-section>
    </q-card>
  </div>

  <!-- 对话框 -->
  <!-- 新建菜单对话框 -->
  <q-dialog v-model="alertCreateRbacMenu" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-form class="q-gutter-md" @submit.prevent="fnStoreRbacMenu">
        <q-card-section>
          <div class="text-h6">新建菜单</div>
        </q-card-section>
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertCreateRbacMenu" label="名称" :rules="[]" />
              <q-input outlined clearable lazy-rules v-model="uri_alertCreateRbacMenu" label="路由" :rules="[]"
                class="q-mt-md" />
              <q-input outlined clearable lazy-rules v-model="description_alertCreateRbacMenu" label="描述" :rules="[]"
                class="q-mt-md" />
              <q-input outlined clearable lazy-rules v-model="icon_alertCreateRbacMenu" label="图标" :rules="[]"
                class="q-mt-md" />
              <sel-rbac-menu_alert-create label-name="所属父级" class="q-mt-md" />
              <chk-rbac-role_alert-create label-name="所属角色" class="q-mt-md" />
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn type="button" label="关闭" v-close-popup />
          <q-btn type="button" label="确定" icon="check" color="secondary" />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
  <!-- 编辑菜单对话框 -->
  <q-dialog v-model="alertEditRbacMenu" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">编辑菜单</div>
      </q-card-section>
      <q-card-section class="q-pt-none">
        <q-form class="q-gutter-md" @submit.prevent="fnUpdateRbacMenu">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertEditRbacMenu" label="名称" :rules="[]" />
              <q-input outlined clearable lazy-rules v-model="uri_alertEditRbacMenu" label="路由" :rules="[]"
                class="q-mt-md" />
              <q-input outlined clearable lazy-rules v-model="description_alertEditRbacMenu" label="描述" :rules="[]"
                class="q-mt-md" />
              <q-input outlined clearable lazy-rules v-model="icon_alertEditRbacMenu" label="图标" :rules="[]"
                class="q-mt-md" />
              <sel-rbac-menu_alert-edit label-name="所属父级" :ajaxParams="{
                ':<>*': { uuid: currentUuid },
                not_has_subs: currentUuid,
              }" class="q-mt-md" />
              <chk-rbac-role_alert-edit label-name="所属角色" class="q-mt-md" />
            </div>
          </div>
        </q-form>
      </q-card-section>
      <q-card-actions align="right">
        <q-btn type="button" label="关闭" v-close-popup />
        <q-btn type="button" label="确定" icon="check" color="warning" />
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>
<script setup>
import { ref, onMounted, provide } from "vue";
import { collect } from "collect.js";
import { fnColumnReverseSort } from "src/utils/common.js";
import SelRbacMenu_search from "src/components/SelRbacMenu_search.vue";
import SelRbacMenu_alertCreate from "src/components/SelRbacMenu_alertCreate.vue";
import SelRbacMenu_alertEdit from "src/components/SelRbacMenu_alertEdit.vue";
import SelRbacRole_search from "src/components/SelRbacRole_search.vue";
import ChkRbacRole_alertCreate from "src/components/ChkRbacRole_alertCreate.vue";
import ChkRbacRole_alertEdit from "src/components/ChkRbacRole_alertEdit.vue";
import {
  successNotify,
  errorNotify,
  loadingNotify,
  confirmNotify,
  destroyActions,
} from "src/utils/notify";
import {
  ajaxDestroyRbacMenu,
  ajaxGetRbacMenu,
  ajaxGetRbacMenus,
  ajaxStoreRbacMenu,
  ajaxUpdateRbacMenu,
  ajaxDestroyRbacMenus,
} from "src/apis/rbac";

// 表格数据
const rows = ref([]);
const selected = ref([]);
const sortBy = ref("");

// 表格数据
const name_search = ref("");
const uri_search = ref("");
const description_search = ref("");
const parentUuid_search = ref("");
const selRbacMenu_search_enable = ref(true);
const rbacRoleUuid_search = ref("");

// 新建菜单对话框
const alertCreateRbacMenu = ref(false);
const name_alertCreateRbacMenu = ref("");
const uri_alertCreateRbacMenu = ref("");
const description_alertCreateRbacMenu = ref("");
const icon_alertCreateRbacMenu = ref("");
const parentUuid_alertCreateRbacMenu = ref("");
const rbacRoleUuids_alertCreateRbacMenu = ref([]);

// 编辑菜单对话框
const alertEditRbacMenu = ref(false);
const rbacMenus_alertEditRbacMenu = ref([]);
const currentUuid = ref("");
const name_alertEditRbacMenu = ref("");
const uri_alertEditRbacMenu = ref("");
const description_alertEditRbacMenu = ref("");
const icon_alertEditRbacMenu = ref("");
const parentUuid_alertEditRbacMenu = ref("");
const rbacRoleUuids_alertEditRbacMenu = ref([]);

provide("parentUuid_search", parentUuid_search);
provide("parentUuid_alertCreate", parentUuid_alertCreateRbacMenu);
provide("parentUuid_alertEdit", parentUuid_alertEditRbacMenu);
provide("rbacRoleUuid_search", rbacRoleUuid_search);
provide("checkedRbacRoleUuids_alertCreate", rbacRoleUuids_alertCreateRbacMenu);
provide("checkedRbacRoleUuids_alertEdit", rbacRoleUuids_alertEditRbacMenu);

onMounted(() => fnInit());

/**
 * 初始化页面
 */
const fnInit = () => fnSearch();;

/**
 * 重置搜索栏
 */
const fnResetSearch = () => {
  name_search.value = "";
  uri_search.value = "";
  description_search.value = "";
  parentUuid_search.value = "";
  rbacRoleUuid_search.value = "";
};

/**
 * 搜索
 */
const fnSearch = () => {
  rows.value = [];
  selected.value = [];

  ajaxGetRbacMenus({
    "@~[]": ["Parent", "RbacRoles"],
    name: name_search.value,
    uri: uri_search.value,
    description: description_search.value,
    parent_uuid: parentUuid_search.value,
    rbac_role_uuid: rbacRoleUuid_search.value,
  })
    .then((res) => {
      if (res.content.rbac_menus.length > 0) {
        collect(res.content.rbac_menus).each((rbacMenu, idx) => {
          rows.value.push({
            index: idx + 1,
            uuid: rbacMenu.uuid,
            name: rbacMenu.name,
            uri: rbacMenu.uri,
            description: rbacMenu.description,
            icon: rbacMenu.icon,
            parent: rbacMenu.parent,
            rbacRoles: rbacMenu.rbac_roles || [],
            operation: { uuid: rbacMenu.uuid },
          });
        });
      }
    })
    .catch((e) => errorNotify(e.msg))
    .finally(() => {
      selRbacMenu_search_enable.value = false;
      selRbacMenu_search_enable.value = true;
    });
};

/**
 * 重置新建菜单对话框
 */
const fnResetAlertCreateRbace = () => {
  name_alertCreateRbacMenu.value = "";
  uri_alertCreateRbacMenu.value = "";
  description_alertCreateRbacMenu.value = "";
  icon_alertCreateRbacMenu.value = "";
  parentUuid_alertCreateRbacMenu.value = "";
  rbacRoleUuids_alertCreateRbacMenu.value = [];
};

/**
 * 打开新建菜单对话框
 */
const fnOpenAlertCreateRbacMenu = () => {
  alertCreateRbacMenu.value = true;
};

/**
 * 新建菜单
 */
const fnStoreRbacMenu = () => {
  const loading = loadingNotify();

  ajaxStoreRbacMenu({
    name: name_alertCreateRbacMenu.value,
    uri: uri_alertCreateRbacMenu.value,
    description: description_alertCreateRbacMenu.value,
    icon: icon_alertCreateRbacMenu.value,
    parent_uuid: parentUuid_alertCreateRbacMenu.value,
    rbac_role_uuids: rbacRoleUuids_alertCreateRbacMenu.value,
  })
    .then((res) => {
      successNotify(res.msg);
      fnSearch();
      fnResetAlertCreateRbace();

      alertCreateRbacMenu.value = false;
    })
    .catch((err) => errorNotify(err.msg))
    .finally(loading());
};

/**
 * 重置编辑菜单对话框
 */
const fnResetAlertEditRbacMenu = () => {
  rbacMenus_alertEditRbacMenu.value = [];
  currentUuid.value = "";
  name_alertEditRbacMenu.value = "";
  uri_alertEditRbacMenu.value = "";
  description_alertEditRbacMenu.value = "";
  icon_alertEditRbacMenu.value = "";
  parentUuid_alertEditRbacMenu.value = "";
  rbacRoleUuids_alertEditRbacMenu.value = [];
};

/**
 * 打开编辑菜单对话框
 * @param {{uuid:string}} params 参数
 */
const fnOpenAlertEditRbacMenu = (params = {}) => {
  if (!params["uuid"]) return;

  currentUuid.value = params.uuid;

  ajaxGetRbacMenu(currentUuid.value, {
    "@~[]": ["Parent", "RbacRoles"],
  })
    .then((res) => {
      name_alertEditRbacMenu.value = res.content.rbac_menu.name;
      uri_alertEditRbacMenu.value = res.content.rbac_menu.uri;
      description_alertEditRbacMenu.value = res.content.rbac_menu.description;
      icon_alertEditRbacMenu.value = res.content.rbac_menu.icon;
      parentUuid_alertEditRbacMenu.value = res.content.rbac_menu.parent
        ? res.content.rbac_menu.parent.uuid
        : "";
      rbacRoleUuids_alertEditRbacMenu.value = collect(res.content.rbac_menu.rbac_roles)
        .pluck("uuid")
        .all();

      alertEditRbacMenu.value = true;
    })
    .catch((e) => errorNotify(e.msg));
};

/**
 * 编辑菜单
 */
const fnUpdateRbacMenu = () => {
  if (!currentUuid.value) return;
  const loading = loadingNotify();
  ajaxUpdateRbacMenu(currentUuid.value, {
    name: name_alertEditRbacMenu.value,
    uri: uri_alertEditRbacMenu.value,
    description: description_alertEditRbacMenu.value,
    icon: icon_alertEditRbacMenu.value,
    parent_uuid: parentUuid_alertEditRbacMenu.value,
    rbac_role_uuids: rbacRoleUuids_alertEditRbacMenu.value,
  })
    .then((res) => {
      successNotify(res.msg);
      fnSearch();
      fnResetAlertEditRbacMenu();

      alertEditRbacMenu.value = false;
    })
    .catch((e) => errorNotify(e.msg))
    .finally(loading());
};

/**
 * 删除菜单
 * @param {{uuid:string}} params 参数
 */
const fnDeconsteRbacMenu = (params = {}) => {
  if (!params.uuid) return;

  confirmNotify(
    destroyActions(() => {
      const loading = loadingNotify();
      ajaxDestroyRbacMenu(params.uuid)
        .then(() => {
          successNotify("删除成功");
          fnSearch();
        })
        .catch((e) => errorNotify(e.msg))
        .finally(loading());
    })
  );
};

/**
 * 批量删除菜单
 */
const fnDestroyRbacMenus = () => {
  confirmNotify(
    destroyActions(() => {
      const loading = loadingNotify();

      ajaxDestroyRbacMenus(collect(selected.value).pluck("uuid").toArray())
        .then(() => {
          successNotify("删除成功");
          fnSearch();
        })
        .catch((e) => errorNotify(e.msg))
        .finally(loading());
    })
  );
};
</script>
src/utils/notify
