<template>
  <div class="q-pa-md">
    <q-card>
      <q-card-section>
        <div class="row q-mb-md">
          <div class="col"><span :style="{ fontSize: '20px' }">搜索</span></div>
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
                  <q-input outlined clearable lazy-rules v-model="uniqueCode_search" label="代码" :rules="[]"
                    class="q-mb-md" />
                </div>
                <div class="col-3">
                  <q-input outlined clearable lazy-rules v-model="name_search" label="名称" :rules="[]" class="q-mb-md" />
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
            <span :style="{ fontSize: '20px' }">路局列表</span>
          </div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="secondary" outline label="新建路局" icon="add" @click="fnOpenAlertCreateOrganizationRailway" />
              <q-btn color="negative" outline label="删除路局" icon="deleted" @click="fnDestroyOrganizationRailways" />
            </q-btn-group>
          </div>
        </div>
        <div class="row q-mt-md">
          <div class="col">
            <q-table flat bordered title="" :rows="rows" row-key="uuid" :pagination="{ rowsPerPage: 200 }"
              :rows-per-page-options="[50, 100, 200, 0]" rows-per-page-label="分页" :selected-rows-label="() => { }"
              selection="multiple" v-model:selected="selected">
              <template v-slot:header="props">
                <q-tr :props="props">
                  <q-th align="left"><q-checkbox key="allCheck" v-model="props.selected" /></q-th>
                  <q-th align="left">#</q-th>
                  <q-th align="left" key="uniqueCode" @click="(event) => fnColumnReverseSort(event, props, sortBy)
                    ">
                    代码
                  </q-th>
                  <q-th align="left" key="name" @click="(event) => fnColumnReverseSort(event, props, sortBy)
                    ">
                    名称
                  </q-th>
                  <q-th align="right"></q-th>
                </q-tr>
              </template>
              <template v-slot:body="props">
                <q-tr :props="props">
                  <q-td><q-checkbox :key="props.row.uuid" :value="props.row.uuid" v-model="props.selected" /></q-td>
                  <q-td>{{ props.row.index }}</q-td>
                  <q-td key="uniqueCode" :props="props">{{
                    props.row.uniqueCode
                  }}</q-td>
                  <q-td key="name" :props="props">{{ props.row.name }}</q-td>
                  <q-td key="operation" :props="props">
                    <q-btn-group>
                      <q-btn @click="
                        fnOpenAlertEditOrganizationRailway(
                          props.row.operation
                        )
                        " color="warning" icon="edit" outline>
                        编辑
                      </q-btn>
                      <q-btn @click="
                        fnDestroyOrganizationRailway(props.row.operation)
                        " color="negative" icon="delete" outline>
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

  <!-- 弹窗 -->
  <!-- 新建路局弹窗 -->
  <q-dialog v-model="alertCreateOrganizationRailway" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">新建路局</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnStoreOrganizationRailway">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="uniqueCode_alertCreateOrganizationRailway" label="代码"
                :rules="[]" />
            </div>
          </div>
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertCreateOrganizationRailway" label="名称" :rules="[]"
                class="q-mt-md" />
            </div>
          </div>
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="shortName_alertCreateOrganizationRailway" label="简称"
                :rules="[]" class="q-mt-md" />
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
  <!-- 编辑路局弹窗 -->
  <q-dialog v-model="alertEditOrganizationrailway" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">编辑路局</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnUpdateOrganizationRailway">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="uniqueCode_alertEditOrganizationrailway" label="代码"
                :rules="[]" />
            </div>
          </div>
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertEditOrganizationrailway" label="名称" :rules="[]"
                class="q-mt-md" />
            </div>
          </div>
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="shortName_alertEditOrganizationrailway" label="简称"
                :rules="[]" class="q-mt-md" />
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
import { ref, onMounted } from "vue";
import collect from "collect.js";
import { fnColumnReverseSort } from "src/utils/common";
import {
  ajaxGetOrganizationRailways,
  ajaxGetOrganizationRailway,
  ajaxStoreOrganizationRailway,
  ajaxUpdateOrganizationRailway,
  ajaxDestroyOrganizationRailway,
  ajaxDestroyOrganizationRailways,
} from "src/apis/organization";
import {
  loadingNotify,
  successNotify,
  errorNotify,
  confirmNotify,
  destroyActions,
} from "src/utils/notify";

// 搜索栏数据
const uniqueCode_search = ref("");
const name_search = ref("");

// 表格数据
const rows = ref([]);
const selected = ref([]);
const sortBy = ref("");

// 新建路局弹窗数据
const alertCreateOrganizationRailway = ref(false);
const uniqueCode_alertCreateOrganizationRailway = ref("");
const name_alertCreateOrganizationRailway = ref("");
const shortName_alertCreateOrganizationRailway = ref("");

// 编辑路局弹窗数据
const currentUuid = ref("");
const alertEditOrganizationrailway = ref(false);
const uniqueCode_alertEditOrganizationrailway = ref("");
const name_alertEditOrganizationrailway = ref("");
const shortName_alertEditOrganizationrailway = ref("");

onMounted(() => fnInit());

/**
 * 初始化页面
 */
const fnInit = () => fnSearch();

/**
 * 初始化搜索条件
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

  ajaxGetOrganizationRailways({
    unique_code: uniqueCode_search.value,
    name: name_search.value,
  })
    .then(res => {
      rows.value = collect(res.content.organization_railways)
        .map((organizationRailway, idx) => {
          return {
            index: idx + 1,
            uniqueCode: organizationRailway.unique_code,
            name: organizationRailway.name,
            uuid: organizationRailway.uuid,
            operation: { uuid: organizationRailway.uuid },
          };
        })
        .all();
    });
};

/**
 * 重置新建路局弹窗
 */
const fnResetAlertCreateOrganizationRailway = () => {
  name_alertCreateOrganizationRailway.value = "";
};

/**
 * 打开新建路局弹窗
 */
const fnOpenAlertCreateOrganizationRailway = () => {
  alertCreateOrganizationRailway.value = true;
};

/**
 * 新建路局
 * @param {{*}} params
 */
const fnStoreOrganizationRailway = (params) => {
  const loading = loadingNotify();

  ajaxStoreOrganizationRailway({
    unique_code: uniqueCode_alertCreateOrganizationRailway.value,
    name: name_alertCreateOrganizationRailway.value,
    short_name: shortName_alertCreateOrganizationRailway.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnResetAlertCreateOrganizationRailway();
      fnSearch();

      alertCreateOrganizationRailway.value = false;
    })
    .catch(e=>errorNotify(e.msg))
    .finally(loading());
};

/**
 * 打开编辑
 * @param {{*}} params
 */
const fnOpenAlertEditOrganizationRailway = (params) => {
  if (!params["uuid"]) return;

  currentUuid.value = params.uuid;

  ajaxGetOrganizationRailway(params.uuid)
    .then(res => {
      uniqueCode_alertEditOrganizationrailway.value = res.content.organization_railway.unique_code;
      name_alertEditOrganizationrailway.value = res.content.organization_railway.name;
      shortName_alertEditOrganizationrailway.value = res.content.organization_railway.short_name;
      alertEditOrganizationrailway.value = true;
    })
    .catch(e=>errorNotify(e.msg));
};

/**
 * 编辑路局
 * @param {string} uuid
 * @param {{*}} params
 */
const fnUpdateOrganizationRailway = () => {
  if (!currentUuid.value) return;

  const loading = loadingNotify();

  ajaxUpdateOrganizationRailway(currentUuid.value, {
    unique_code: uniqueCode_alertEditOrganizationrailway.value,
    name: name_alertEditOrganizationrailway.value,
    short_name: shortName_alertEditOrganizationrailway.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnSearch();

      alertEditOrganizationrailway.value = false;
    })
    .catch(e=>errorNotify(e.msg))
    .finally(loading());
};

/**
 * 删除路局
 * @param {{*}} params
 */
const fnDestroyOrganizationRailway = (params) => {
  if (!params["uuid"]) return;

  confirmNotify(
    destroyActions(() => {
      const loading = loadingNotify();

      ajaxDestroyOrganizationRailway(params.uuid)
        .then(() => {
          successNotify("删除成功");
          fnSearch();
        })
        .catch(e=>errorNotify(e.msg))
        .finally(loading());
    })
  );
};

/**
 * 批量删除路局
 * @param {[string]} uuids
 */
const fnDestroyOrganizationRailways = (uuids) => {
  confirmNotify(
    destroyActions(() => {
      const loading = loadingNotify();

      ajaxDestroyOrganizationRailways(selected.value)
        .then(() => {
          successNotify("删除成功");
          fnSearch();
        })
        .catch(e=>errorNotify(e.msg))
        .finally(loading());
    })
  );
};
</script>
src/apis/organization
