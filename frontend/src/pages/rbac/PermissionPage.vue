<template>
  <div class="q-pa-md">
    <q-card>
      <q-card-section>
        <div class="row q-mb-md">
          <div class="col"><span :style="{ fontSize: '20px' }">搜索</span></div>
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
            <span :style="{ fontSize: '20px' }">权限列表</span>
          </div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="secondary" label="新建权限" icon="add" @click="fnOpenAlertCreateRbacPermission" />
              <q-btn color="negative" label="删除权限" icon="delete" @click="fnDestroyRbacPermissions" />
            </q-btn-group>
          </div>
        </div>
        <div class="row q-mt-md">
          <div class="col">
            <q-table flat bordered title="" :rows="rows" row-key="name" :pagination="{ rowsPerPage: 200 }"
              :rows-per-page-options="[50, 100, 200, 0]" rows-per-page-label="分页" selection="multiple"
              v-model:selected="selected">
              <template v-slot:header="props">
                <q-tr :props="props">
                  <q-th align="left"><q-checkbox key="allCheck" v-model="props.selected">全选</q-checkbox></q-th>
                  <q-th align="left">#</q-th>
                  <q-th align="left" key="name" @click="(event) => fnColumnReverseSort(event, props, sortBy)
                    ">
                    名称
                  </q-th>
                  <q-th align="left" key="uri" @click="(event) => fnColumnReverseSort(event, props, sortBy)
                    ">
                    路由
                  </q-th>
                  <q-th align="left" key="description" @click="(event) => fnColumnReverseSort(event, props, sortBy)
                    ">
                    描述
                  </q-th>
                  <q-th align="left" key="rbacRoles">所属角色</q-th>
                  <q-th align="right" key="operation"></q-th>
                </q-tr>
              </template>
              <template v-slot:body="props">
                <q-tr :props="props">
                  <q-td><q-checkbox :key="props.row.uuid" :value="props.row.uuid" v-model="props.selected" /></q-td>
                  <q-td>{{ props.row.index }}</q-td>
                  <q-td key="name" :props="props">{{ props.row.name }}</q-td>
                  <q-td key="uri" :props="props">{{ props.row.uri }}</q-td>
                  <q-td key="description" :props="props">{{ props.row.description }}</q-td>
                  <q-td key="rbacRoles" :props="props">
                    <join-string :values="collect(props.row.rbacRoles).pluck('name').all()" separator=", " />
                  </q-td>
                  <q-td key="operation" :props="props">
                    <q-btn-group>
                      <q-btn @click="fnOpenAlertEditRbacPermission(props.row.operation)" color="warning" icon="edit">
                        编辑
                      </q-btn>
                      <q-btn @click="fnDestroyRbacPermission(props.row.operation)" color="negative" icon="delete">
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
  <!-- 新建权限对话框 -->
  <q-dialog v-model="alertCreateRbacPermission">
    <q-card :style="{ minWidth: '450px' }">
      <q-form class="q-gutter-md" @submit.prevent="fnStoreRbacPermission">
        <q-card-section>
          <div class="text-h6">新建权限</div>
        </q-card-section>
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertCreateRbacPermission" label="名称" :rules="[]" />
              <q-input outlined clearable lazy-rules v-model="uri_alertCreateRbacPermission" label="路由" :rules="[]"
                class="q-mt-md" />
              <q-input outlined clearable lazy-rules v-model="description_alertCreateRbacPermission" label="描述"
                :rules="[]" class="q-mt-md" />
              <chk-rbac-role_alert-create label-name="所属角色" class="q-mt-md" />
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn type="button" label="关闭" v-close-popup />
          <q-btn type="submit" label="确定" icon="check" color="secondary" />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
  <!-- 编辑权限对话框 -->
  <q-dialog v-model="alertEditRbacPermission">
    <q-card :style="{ minWidth: '450px' }">
      <q-form class="q-gutter-md" @submit.prevent="fnUpdateRbacPermission">
        <q-card-section>
          <div class="text-h6">编辑权限</div>
        </q-card-section>
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertEditRbacPermission" label="名称" :rules="[]" />
              <q-input outlined clearable lazy-rules v-model="uri_alertEditRbacPermission" label="路由" :rules="[]"
                class="q-mt-md" />
              <q-input outlined clearable lazy-rules v-model="description_alertEditRbacPermission" label="描述" :rules="[]"
                class="q-mt-md" />
              <chk-rbac-role_alert-edit label-name="所属角色" class="q-mt-md" />
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn type="button" label="关闭" v-close-popup />
          <q-btn type="submit" label="确定" icon="check" color="warning" />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
</template>
<script setup>
import { ref, onMounted, provide } from "vue";
import collect from "collect.js";
import { fnColumnReverseSort } from "src/utils/common";
import SelRbacRole_search from "src/components/SelRbacRole_search.vue";
import ChkRbacRole_alertCreate from "src/components/ChkRbacRole_alertCreate.vue";
import ChkRbacRole_alertEdit from "src/components/ChkRbacRole_alertEdit.vue";
import {
  loadingNotify,
  successNotify,
  errorNotify,
  confirmNotify,
  destroyActions,
} from "src/utils/notify";
import {
  ajaxGetRbacPermissions,
  ajaxGetRbacPermission,
  ajaxStoreRbacPermission,
  ajaxUpdateRbacPermission,
  ajaxDestroyRbacPermission,
  ajaxDestroyRbacPermissions,
} from "src/apis/rbac";
import JoinString from "src/components/JoinString.vue";

// 表格数据
const rows = ref([]);
const selected = ref([]);
const sortBy = ref("");

// 搜索条件
const name_search = ref("");
const uri_search = ref("");
const description_search = ref("");
const rbacRoleUuid_search = ref("");
provide("rbacRoleUuid_search", rbacRoleUuid_search);

// 新建权限对话框
const alertCreateRbacPermission = ref(false);
const name_alertCreateRbacPermission = ref("");
const uri_alertCreateRbacPermission = ref("");
const description_alertCreateRbacPermission = ref("");
const rbacRoleUuids_alertCreateRbacPermission = ref([]);
provide("checkedRbacRoleUuids_alertCreate", rbacRoleUuids_alertCreateRbacPermission);

// 编辑权限对话框
const alertEditRbacPermission = ref(false);
const currentUuid = ref("");
const name_alertEditRbacPermission = ref("");
const uri_alertEditRbacPermission = ref("");
const description_alertEditRbacPermission = ref("");
const rbacRoleUuids_alertEditRbacPermission = ref([]);
provide("checkedRbacRoleUuids_alertEdit", rbacRoleUuids_alertEditRbacPermission);

onMounted(() => fnInit());

/**
 * 初始化
 */
const fnInit = () => fnSearch();

/**
 * 初始化搜索栏
 */
const fnResetSearch = () => {
  name_search.value = "";
  uri_search.value = "";
  description_search.value = "";
  rbacRoleUuid_search.value = "";
};

/**
 * 搜索
 */
const fnSearch = () => {
  rows.value = [];
  selected.value = [];

  ajaxGetRbacPermissions({
    "@~[]": ["RbacRoles"],
    name: name_search.value,
    uri: uri_search.value,
    description: description_search.value,
    rbac_role_uuid: rbacRoleUuid_search.value,
  })
    .then((res) => {
      rows.value = collect(res.content.rbac_permissions)
        .map((rbacPermission, idx) => {
          return {
            ...rbacPermission,
            index: idx + 1,
            rbacRoles: rbacPermission.rbac_roles,
            operation: { uuid: rbacPermission.uuid },
          }
        })
        .all();
      console.log('ok', rows.value);
    })
    .catch((e) => errorNotify(e.msg));
};

/**
 * 重置新建权限对话框
 */
const fnResetAlertCreateRbacPermission = () => {
  name_alertCreateRbacPermission.value = "";
  uri_alertCreateRbacPermission.value = "";
  description_alertCreateRbacPermission.value = "";
  rbacRoleUuids_alertCreateRbacPermission.value = [];
};
/**
 * 打开新建权限对话框
 */
const fnOpenAlertCreateRbacPermission = () => {
  alertCreateRbacPermission.value = true;
};

/**
 * 新建权限
 */
const fnStoreRbacPermission = () => {
  const loading = loadingNotify();

  ajaxStoreRbacPermission({
    name: name_alertCreateRbacPermission.value,
    uri: uri_alertCreateRbacPermission.value,
    description: description_alertCreateRbacPermission.value,
    rbac_role_uuids: rbacRoleUuids_alertCreateRbacPermission.value,
  })
    .then((res) => {
      successNotify(res.msg);
      fnSearch();
      fnResetAlertCreateRbacPermission();

      alertCreateRbacPermission.value = false;
    })
    .catch((e) => errorNotify(e.msg))
    .finally(loading());
};

/**
 * 重置编辑权限对话框
 */
const fnResetAlertEditRbacPermission = () => {
  name_alertEditRbacPermission.value = "";
  uri_alertEditRbacPermission.value = "";
  description_alertEditRbacPermission.value = "";
  rbacRoleUuids_alertEditRbacPermission.value = [];
};

/**
 * 打开编辑权限对话框
 * @param {{*}} params 参数
 */
const fnOpenAlertEditRbacPermission = (params = {}) => {
  if (!params["uuid"]) return;

  currentUuid.value = params.uuid;

  ajaxGetRbacPermission(params.uuid, { "@~[]": ["RbacRoles"] })
    .then((res) => {
      name_alertEditRbacPermission.value = res.content.rbac_permission.name;
      uri_alertEditRbacPermission.value = res.content.rbac_permission.uri;
      description_alertEditRbacPermission.value = res.content.rbac_permission.description;
      rbacRoleUuids_alertEditRbacPermission.value = collect(res.content.rbac_permission.rbac_roles)
        .pluck("uuid")
        .all();

      alertEditRbacPermission.value = true;
    })
    .catch((e) => errorNotify(e.msg));
};

/**
 * 编辑权限
 */
const fnUpdateRbacPermission = () => {
  if (!currentUuid.value) return;

  const loading = loadingNotify();

  ajaxUpdateRbacPermission(currentUuid.value, {
    name: name_alertEditRbacPermission.value,
    uri: uri_alertEditRbacPermission.value,
    description: description_alertEditRbacPermission.value,
    rbac_role_uuids: rbacRoleUuids_alertEditRbacPermission.value,
  })
    .then((res) => {
      successNotify(res.msg);
      fnSearch();
      fnResetAlertEditRbacPermission();

      alertEditRbacPermission.value = false;
    })
    .catch((e) => errorNotify(e.msg))
    .finally(loading());
};

/**
 * 删除权限
 * @param {{*}} params 参数
 */
const fnDestroyRbacPermission = (params = {}) => {
  if (!params["uuid"]) return;

  const loading = loadingNotify();
  confirmNotify(
    destroyActions(() => {
      ajaxDestroyRbacPermission(params.uuid)
        .then((res) => {
          successNotify("删除成功");
          fnSearch();
        })
        .catch((e) => errorNotify(e.msg))
        .finally(loading());
    })
  );
};

/**
 * 批量删除权限
 */
const fnDestroyRbacPermissions = () => {
  confirmNotify(
    destroyActions(() => {
      const loading = loadingNotify();

      ajaxDestroyRbacPermissions(
        collect(selected.value).pluck("uuid").toArray()
      )
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
