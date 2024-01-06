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
            <span :style="{ fontSize: '20px' }">角色列表</span>
          </div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="secondary" label="添加角色" icon="add" @click="fnOpenAlertCreateRbacRole" />
              <q-btn color="negative" label="删除角色" icon="delete" @click="fnDestroyRbacRoles" />
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
                  <q-th align="left"><q-checkbox key="allCheck" v-model="props.selected" /></q-th>
                  <q-th align="left">#</q-th>
                  <q-th align="left" key="name" @click="(event) => fnColumnReverseSort(event, props, sortBy)
                    ">
                    名称
                  </q-th>
                  <q-th align="right"></q-th>
                </q-tr>
              </template>
              <template v-slot:body="props">
                <q-tr :props="props">
                  <q-td><q-checkbox :key="props.row.uuid" v-model="props.selected" /></q-td>
                  <q-td>{{ props.row.index }}</q-td>
                  <q-td key="name" :props="props">{{ props.row.name }}</q-td>
                  <q-td key="operation" :props="props">
                    <q-btn-group>
                      <q-btn @click="fnOpenAlertEditRbacRole(props.row.operation)" color="warning" icon="edit">
                        编辑
                      </q-btn>
                      <q-btn @click="fnDestroyRbacRole(props.row.operation)" color="negative" icon="delete">
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
  <!-- 新建角色对话框 -->
  <q-dialog v-model="alertCreateRbacRole" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-form class="q-gutter-md" @submit.prevent="fnStoreRbacRole">
        <q-card-section>
          <div class="text-h6">新建角色</div>
        </q-card-section>
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertCreateRbacRole" label="名称" :rules="[]" />
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn-group>
            <q-btn type="button" label="关闭" v-close-popup />
            <q-btn type="submit" label="确定" icon="check" color="secondary" />
          </q-btn-group>
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
  <!-- 编辑角色对话框 -->
  <q-dialog v-model="alertEditRbacRole" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">编辑角色</div>
      </q-card-section>
      <q-card-section class="q-pt-none">
        <q-form class="q-gutter-md" @submit.prevent="fnUpdateRbacRole">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertEditRbacRole" label="名称" :rules="[]" />
            </div>
          </div>
        </q-form>
      </q-card-section>
      <q-card-actions align="right">
        <q-btn-group>
          <q-btn type="button" label="关闭" v-close-popup />
          <q-btn type="submit" label="确定" icon="check" color="warning" />
        </q-btn-group>
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script setup>
import { ref, onMounted } from "vue";
import collect from "collect.js";
import { fnColumnReverseSort } from "src/utils/common";
import {
  ajaxGetRbacRoles,
  ajaxGetRbacRole,
  ajaxStoreRbacRole,
  ajaxUpdateRbacRole,
  ajaxDestroyRbacRole,
  ajaxDestroyRbacRoles,
} from "src/apis/rbac";
import {
  loadingNotify,
  errorNotify,
  successNotify,
  confirmNotify,
  destroyActions,
} from "src/utils/notify";

const name_search = ref("");
const rows = ref([]);
const selected = ref([]);
const sortBy = ref("");
const alertCreateRbacRole = ref(false);
const name_alertCreateRbacRole = ref("");
const alertEditRbacRole = ref(false);
const name_alertEditRbacRole = ref("");
const currentUuid = ref("");

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
};

/**
 * 搜索
 */
const fnSearch = () => {
  rows.value = [];
  selected.value = [];

  ajaxGetRbacRoles({
    name: name_search.value,
  }).then((res) => {
    if (res.content.rbac_roles.length > 0) {
      collect(res.content.rbac_roles).each((rbacRole, idx) => {
        rows.value.push({
          index: idx + 1,
          uuid: rbacRole.uuid,
          name: rbacRole.name,
          operation: { uuid: rbacRole.uuid },
        });
      });
    }
  });
};

/**
 * 重置新建角色对话框
 */
const fnResetAlertCreateRbacRole = () => {
  name_alertCreateRbacRole.value = "";
};

/**
 * 打开新建角色对话框
 */
const fnOpenAlertCreateRbacRole = () => {
  alertCreateRbacRole.value = true;
};

/**
 * 新建角色
 */
const fnStoreRbacRole = () => {
  const loading = loadingNotify();

  ajaxStoreRbacRole({
    name: name_alertCreateRbacRole.value,
  })
    .then((res) => {
      successNotify(res.msg);
      fnSearch();
      fnResetAlertCreateRbacRole();

      alertCreateRbacRole.value = false;
    })
    .catch(e=>errorNotify(e.msg))
    .finally(loading());
};

/**
 * 重置编辑角色对话框
 */
const fnResetAlertEditRbacRole = () => {
  name_alertEditRbacRole.value = "";
};

/**
 * 打开编辑角色对话框
 */
const fnOpenAlertEditRbacRole = (params = {}) => {
  if (!params["uuid"]) return;

  currentUuid.value = params.uuid;

  ajaxGetRbacRole(currentUuid.value)
    .then((res) => {
      if (res.content.rbac_role) {
        name_alertEditRbacRole.value = res.content.rbac_role.name;
        alertEditRbacRole.value = true;
      }
    })
    .catch(e=>errorNotify(e.msg));
};

/**
 * 编辑角色
 */
const fnUpdateRbacRole = () => {
  if (!currentUuid.value) return;

  ajaxUpdateRbacRole(currentUuid.value, {
    name: name_alertEditRbacRole.value,
  })
    .then((res) => {
      successNotify(res.msg);
      fnSearch();
      fnResetAlertEditRbacRole();

      alertEditRbacRole.value = false;
    })
    .catch(e=>errorNotify(e.msg));
};

/**
 * 删除角色
 */
const fnDestroyRbacRole = (params = {}) => {
  if (!params["uuid"]) return;

  confirmNotify(
    destroyActions(() => {
      const loading = loadingNotify();

      ajaxDestroyRbacRole(params.uuid)
        .then((res) => {
          successNotify("删除成功");
          fnSearch();
        })
        .catch(e=>errorNotify(e.msg))
        .finally(loading());
    })
  );
};

/**
 * 批量删除角色
 */
const fnDestroyRbacRoles = () => {
  confirmNotify(
    destroyActions(() => {
      const loading = loadingNotify();

      ajaxDestroyRbacRoles(collect(selected.value).pluck("uuid").toArray())
        .then((res) => {
          successNotify("删除成功");
          fnSearch();
        })
        .catch(e=>errorNotify(e.msg))
        .finally(loading());
    })
  );
};
</script>
src/utils/notify
