
import { fasPlugCircleCheck } from '@quasar/extras/fontawesome-v6';

import collect from 'collect.js';

import { ajaxDestroyBreakdownType } from 'src/apis/breakdown';

import { destroyActions } from 'src/utils/notify';

import { actionNotify } from 'src/utils/notify';

import { errorNotify } from 'src/utils/notify';

import { bexContent } from 'quasar/wrappers';

import collect from 'collect.js';

import { ajaxGetBreakdownTypes } from 'src/apis/breakdown';

import { ajaxGetBreakdownLogs } from 'src/apis/breakdown';

import { onMounted } from 'vue';

import { ref } from 'vue';

import SelEquipmentKindCategory_searchVue from 'src/components/SelEquipmentKindCategory_search.vue';

import { destroyActions } from 'src/utils/notify';

import { actionNotify } from 'src/utils/notify';

<template>
  <div class="q-pa-md">
    <q-card>
      <q-card-section>
        <div class="row">
          <div class="col"><span :style="{ fontSize: '20px' }">搜索</span></div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="primary" label="搜索" icon="search" @click="fnSearch" />
              <q-btn color="primary" label="重置" flat @click="fnResetSearch" />
            </q-btn-group>
          </div>
        </div>
        <div class="row q-mt-md">
          <div class="col">
            <q-form>
              <div class="row q-col-gutter-sm">
                <div class="col-3">
                  <q-input outlined clearable lazy-rules v-model="name_search" label="名称" :rules="[]" class="q-mb-md" />
                </div>
                <div class="col-3">
                  <sel-equipment-kind-category-search v-model="equipmentKindCategoryUuid_search" label="所属设备种类"
                    :rules="[]" class="q-mb-md" />
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
          <div class="col"><span :style="{ fontSize: '20px' }">故障类型列表</span></div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="secondary" label="新建故障类型" icon="add" @click="fnOpenAlertCreateBreakdownType" />
              <q-btn color="negative" label="删除故障类型" icon="deleted" @click="fnDestroyBreakdownTypes" />
            </q-btn-group>
          </div>
        </div>
        <div class="row q-mt-md">
          <div class="col">
            <q-table flat bordered title="故障类型列表" :rows="rows" row-key="uuid" :pagination="{ rowsPerPage: 200 }"
              :rows-per-page-options="[50, 100, 200, 0]" rows-per-page-label="分页" :selected-rows-label="() => { }"
              selection="multiple" v-model:selected="selected">
              <template v-slot:header="props">
                <q-tr :props="props">
                  <q-th align="left"><q-checkbox key="allCheck" v-model="props.selected" /></q-th>
                  <q-th align="left">#</q-th>
                  <q-th align="left" key="name" @click="(event) => fnColumnReverseSort(event, props, sortBy)">
                    名称
                  </q-th>
                  <q-th align="right"></q-th>
                </q-tr>
              </template>
              <template v-slot:body="props">
                <q-tr :props="props">
                  <q-td><q-checkbox :key="props.row.uuid" :value="props.row.uuid" v-model="props.selected" /></q-td>
                  <q-td>{{ props.row.index }}</q-td>
                  <q-td key="name" :props="props">{{ props.row.name }}</q-td>
                  <q-td key="operation" :props="props">
                    <q-btn-group>
                      <q-btn @click="fnOpenAlertEditBreakdownType(props.row.operation)" color="warning" icon="edit">
                        编辑
                      </q-btn>
                      <q-btn @click="fnDestroyBreakdownType(props.row.operation)" color="negative" icon="delete">
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
</template>

<script setup>
import { ref, onMounted, provide } from "vue";
import collect from "collect.js";
import { fnColumnReverseSort } from "src/utils/common";
import {
  ajaxGetBreakdownTypes,
  ajaxGetBreakdownType,
  ajaxStoreBreakdownType,
  ajaxUpdateBreakdownType,
  ajaxDestroyBreakdownType,
  ajaxDestroyBreakdownTypes,
} from "src/apis/breakdown";
import {
  loadingNotify,
  successNotify,
  errorNotify,
  actionNotify,
  destroyActions,
} from "src/utils/notify";
import SelEquipmentKindCategory_search from "src/components/SelEquipmentKindCategory_search.vue";

// 搜索栏数据
const name_search = ref("");
const equipmentKindCategoryUuid_search = ref("");

// 表格数据
const rows = ref([]);
const selected = ref([]);
const sortyBy = ref("");

// 新建故障类型弹窗
const alertCreateBreakdownType = ref(false);
const name_alertCreate = ref("");
const equipmentKindCategoryUuid_alertCreate = ref("");

// 编辑故障类型弹窗
const alertEditBreakdownType = ref(false);
const currentUuid = ref("");
const name_alertEdit = ref("");
const equipmentKindCategoryUuid_alertEdit = ref("");

onMounted(() => fnInit());

const fnInit = () => fnSearch();

const fnResetSearch = () => {
  name_search.value = "";
  equipmentKindCategoryUuid_search.value = "";
};

const fnSearch = () => {
  ajaxGetBreakdownTypes({
    name: name_search.value,
    equipmentKindCategoryUuid: equipmentKindCategoryUuid_search.value,
  })
    .then((res) => {
      rows.value = collect(res.content.breakdown_types)
        .map(breakdownType => {
          return {
            uuid: breakdownType.uuid,
            name: breakdownType.name,
            equipmentKindCategory: breakdownType.equipment_kind_category,
            operation: { uuid: breakdownType.uuid },
          }
        });
    })
    .catch(e => errorNotify(e.msg));
};
const fnResetAlertCreateBreakdownType = () => {
  name_alertCreate.value = "";
  equipmentKindCategoryUuid_alertCreate.value = "";
};

const fnOpenAlertCreateBreakdownType = () => {
  alertcreateBreakdownType.value = true;
};

const fnStoreBreakdownType = () => {
  const loading = loadingNotify();

  ajaxStoreBreakdownType({
    name: name_alertCreate.value,
    equipmentKindCategoryUuid: equipmentKindCategoryUuid_alertCreate.value,
  })
    .then((res) => {
      successNotify(res.msg);
      fnSearch();
      fnResetAlertCreateBreakdownType();
    })
    .catch(e => errorNotify(e.msg))
    .finally(() => loading());
};

const fnOpenAlertEditBreakdownType = params => {
  if (!params["uuid"]) return;
  currentUuid.value = params.uuid;

  alerteditBreakdownType.value = true;
};

const fnUpdateBreakdownType = () => {
  const loading = loadingNotify();

  ajaxUpdateBreakdownType(currentUuid.value, {
    name: name_alertEdit.value,
    equipmentKindCategoryUuid: equipmentKindCategoryUuid_alertEdit.value,
  })
    .then((res) => {
      successNotify(res.msg);
      fnSearch();
    })
    .catch(e => errorNotify(e.msg))
    .finally(() => loading());
};

const fnDestroyBreakdownType = params => {
  if (!params["uuid"]) return;

  actionNotify(
    destroyActions(() => {
      const loading = loadingNotify();

      ajaxDestroyBreakdownType(params.uuid)
        .then(() => {
          successNotify("删除成功");
          fnSearch();
        })
        .catch(e => errorNotify(e.msg))
        .finally(() => loading());
    })
  );
};
const fnDestroyBreakdownTypes = () => {
  actionNotify(
    destroyActions(() => {
      const loading = loadingNotify();

      ajaxDestroyBreakdownTypes(collect(selected.value).pluck("uuid").toArray())
        .then(() => {
          successNotify("删除成功");
          fnSearch();
        })
        .catch(e => errorNotify(e.msg))
        .finally(() => loading());
    })
  );
};

</script>
