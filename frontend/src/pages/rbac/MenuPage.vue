<template>
  <div class="q-pa-md">
    <q-card>
      <q-card-section>
        <div class="row q-mb-md">
          <div class="col"><span style="font-size: 20px">搜索</span></div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="primary" label="搜索" outline icon="search" @click="fnSearch" />
              <q-btn label="重置" outline flat @click="fnResetSearch" />
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
                  <standard-select label-name="所属父级" sechma="search" current-field="parentUuid"
                    :data-source="ajaxGetRbacMenus" data-source-field="rbac_menus" />
                </div>
                <div class="col-3">
                  <standard-select label-name="所属角色" sechma="search" current-field="rbacRoleUuid"
                    :data-source="ajaxGetRbacRoles" data-source-field="rbac_roles" />
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
              <q-btn color="secondary" outline label="新建菜单" icon="add" @click="fnOpenAlertCreateRbacMenu" />
              <q-btn color="negative" outline label="删除菜单" icon="delete" @click="fnDestroyRbacMenus" />
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
                  <q-th align="left" key="name" @click="(event) => fnColumnReverseSort(event, props, sortBy)">
                    名称
                  </q-th>
                  <q-th align="left" key="uri" @click="(event) => fnColumnReverseSort(event, props, sortBy)">
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
                  <q-td key="uri" :props="props">{{ props.row.uri }}</q-td>
                  <q-td key="description" :props="props">{{ props.row.description }}</q-td>
                  <q-td><q-icon :name="props.row.icon" /></q-td>
                  <q-td key="parent" :props="props">
                    {{ props.row.parent ? props.row.parent.name : "" }}
                  </q-td>
                  <q-td key="rbacRoles" :props="props">
                    <join-string :values="collect(props.row.rbacRoles).pluck('name').all()" sep=", " />
                  </q-td>
                  <q-td key="operation" :props="props">
                    <q-btn-group>
                      <q-btn @click="fnOpenAlertEditRbacMenu(props.row.operation)" color="warning" icon="edit" outline>
                        编辑
                      </q-btn>
                      <q-btn @click="fnDeconsteRbacMenu(props.row.operation)" color="negative" icon="delete" outline>
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
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="uri_alertCreateRbacMenu" label="路由" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="description_alertCreateRbacMenu" label="描述" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="icon_alertCreateRbacMenu" label="图标" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属父级" sechma="alertCreate" current-field="parentUuid"
                :data-source="ajaxGetRbacMenus" data-source-field="rbac_menus" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-check label-name="所属角色" current-field="checkedRbacRoleUuids" sechma="alertCreate"
                :ajax-params="{}" :data-source="ajaxGetRbacRoles" data-source-field="rbac_roles" />
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn-group>
            <q-btn type="button" label="关闭" outline v-close-popup />
            <q-btn type="submit" label="确定" outline icon="check" color="secondary" />
          </q-btn-group>
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
  <!-- 编辑菜单对话框 -->
  <q-dialog v-model="alertEditRbacMenu" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-form class="q-gutter-md" @submit.prevent="fnUpdateRbacMenu">
        <q-card-section>
          <div class="text-h6">编辑菜单</div>
        </q-card-section>
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertEditRbacMenu" label="名称" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="uri_alertEditRbacMenu" label="路由" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="description_alertEditRbacMenu" label="描述" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="icon_alertEditRbacMenu" label="图标" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属父级" sechma="alertEdit" current-field="parentUuid"
                :data-source="ajaxGetRbacMenus" data-source-field="rbac_menus"
                :ajax-params="{ '@<>': { uuid: currentUuid }, not_has_subs: currentUuid, }" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-check label-name="所属角色" current-field="checkedRbacRoleUuids" sechma="alertEdit" :ajax-params="{}"
                :data-source="ajaxGetRbacRoles" data-source-field="rbac_roles" />
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn-group>
            <q-btn type="button" label="关闭" outline v-close-popup />
            <q-btn type="submit" label="确定" outline icon="check" color="warning" />
          </q-btn-group>
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
</template>
<script setup>
import { ref, onMounted, provide } from "vue";
import collect from "collect.js";
import { fnColumnReverseSort } from "src/utils/common.js";
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
  ajaxGetRbacRoles,
} from "src/apis/rbac";
import JoinString from "src/components/JoinString.vue";
import StandardSelect from "src/components/StandardSelect.vue";
import StandardCheck from "src/components/StandardCheck.vue";

// 搜索栏数据
const name_search = ref("");
const uri_search = ref("");
const description_search = ref("");
const parentUuid_search = ref("");
provide("parentUuid_search", parentUuid_search);
const selRbacMenuShow_search = ref(true);
const rbacRoleUuid_search = ref("");
provide("rbacRoleUuid_search", parentUuid_search);

// 表格数据
const rows = ref([]);
const selected = ref([]);
const sortBy = ref("");

// 新建菜单对话框
const alertCreateRbacMenu = ref(false);
const name_alertCreateRbacMenu = ref("");
const uri_alertCreateRbacMenu = ref("");
const description_alertCreateRbacMenu = ref("");
const icon_alertCreateRbacMenu = ref("");
const parentUuid_alertCreateRbacMenu = ref("");
provide("parentUuid_alertCreate", parentUuid_alertCreateRbacMenu);
const rbacRoleUuids_alertCreateRbacMenu = ref([]);
provide("checkedRbacRoleUuids_alertCreate", rbacRoleUuids_alertCreateRbacMenu);

// 编辑菜单对话框
const alertEditRbacMenu = ref(false);
const rbacMenus_alertEditRbacMenu = ref([]);
const currentUuid = ref("");
const name_alertEditRbacMenu = ref("");
const uri_alertEditRbacMenu = ref("");
const description_alertEditRbacMenu = ref("");
const icon_alertEditRbacMenu = ref("");
const parentUuid_alertEditRbacMenu = ref("");
provide("parentUuid_alertEdit", parentUuid_alertEditRbacMenu);
const rbacRoleUuids_alertEditRbacMenu = ref([]);
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
    .then(res => {
      rows.value = collect(res.content.rbac_menus)
        .map((rbacMenu, idx) => {
          return {
            ...rbacMenu,
            index: idx + 1,
            parent: rbacMenu.parent,
            rbacRoles: rbacMenu.rbac_roles,
            operation: { uuid: rbacMenu.uuid },
          };
        })
        .all();
    })
    .catch(e => errorNotify(e.msg))
    .finally(() => {
      selRbacMenuShow_search.value = false;
      selRbacMenuShow_search.value = true;
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
    .then(res => {
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
    .then(res => {
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
    .catch(e => errorNotify(e.msg));
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
    .then(res => {
      successNotify(res.msg);
      fnSearch();
      fnResetAlertEditRbacMenu();

      alertEditRbacMenu.value = false;
    })
    .catch(e => errorNotify(e.msg))
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
        .catch(e => errorNotify(e.msg))
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
        .catch(e => errorNotify(e.msg))
        .finally(loading());
    })
  );
};
</script>
src/utils/notify
