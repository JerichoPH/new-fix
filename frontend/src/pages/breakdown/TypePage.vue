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
                  <sel-equipment-kind-category_search label-name="所属器材种类" />
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
            <q-table flat bordered title="" :rows="rows" row-key="uuid" :pagination="{ rowsPerPage: 200 }"
              :rows-per-page-options="[50, 100, 200, 0]" rows-per-page-label="分页" :selected-rows-label="() => { }"
              selection="multiple" v-model:selected="selected">
              <template v-slot:header="props">
                <q-tr :props="props">
                  <q-th align="left"><q-checkbox key="allCheck" v-model="props.selected" /></q-th>
                  <q-th align="left">#</q-th>
                  <q-th align="left" key="name" @click="(event) => fnColumnReverseSort(event, props, sortBy)">
                    名称
                  </q-th>
                  <q-th align="left" key="equipmentKindCategory"
                    @click="(event) => fnColumnReverseSort(event, props, sortBy)">
                    所属种类
                  </q-th>
                  <q-th align="right"></q-th>
                </q-tr>
              </template>
              <template v-slot:body="props">
                <q-tr :props="props">
                  <q-td><q-checkbox :key="props.row.uuid" :value="props.row.uuid" v-model="props.selected" /></q-td>
                  <q-td>{{ props.row.index }}</q-td>
                  <q-td key="name" :props="props">{{ props.row.name }}</q-td>
                  <q-td key="equipmentKindCategory" :props="props">{{ props.row.equipmentKindCategory.name }}</q-td>
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

  <!-- 弹窗 -->
  <!-- 新建故障类型弹窗 -->
  <q-dialog v-model="alertCreateBreakdownType">
    <q-card :style="{minWidth: '40vw'}">
      <q-card-section>
        <div class="text-h6">新建故障类型</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnStoreBreakdownType">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertCreateBreakdownType" label="名称" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-equipment-kind-category_alert-create label-name="所属器材种类" />
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn type="submit" label="确定" icon="check" color="secondary" v-close-popup />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
  <!-- 编辑故障类型 -->
  <q-dialog v-model="alertEditBreakdownType">
    <q-card :style="{minWidth: '40vw'}">
      <q-card-section>
        <div class="text-h6">编辑故障类型</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnUpdateBreakdownType">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertEditBreakdownType" label="名称" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-equipment-kind-category_alert-edit lable-name="所属器材种类" />
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn type="submit" label="确定" icon="check" color="warning" v-close-popup />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
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
  confirmNotify,
  destroyActions,
} from "src/utils/notify";
import SelEquipmentKindCategory_search from "src/components/SelEquipmentKindCategory_search.vue";
import SelEquipmentKindCategory_alertCreate from "src/components/SelEquipmentKindCategory_alertCreate.vue"
import SelEquipmentKindCategory_alertEdit from "src/components/SelEquipmentKindCategory_alertEdit.vue"

// 搜索栏数据
const name_search = ref("");
const equipmentKindCategoryUuid_search = ref("");
provide("equipmentKindCategoryUuid_search", equipmentKindCategoryUuid_search);

// 表格数据
const rows = ref([]);
const selected = ref([]);
const sortBy = ref("");

// 新建故障类型数据
const alertCreateBreakdownType = ref(false);
const name_alertCreateBreakdownType = ref("");
const equipmentKindCategoryUuid_alertCreateBreakdownType = ref("");
provide("equipmentKindCategoryUuid_alertCreate", equipmentKindCategoryUuid_alertCreateBreakdownType);

// 编辑故障类型数据
const alertEditBreakdownType = ref(false);
const currentUuid = ref("");
const name_alertEditBreakdownType = ref("");
const equipmentKindCategoryUuid_alertEditBreakdownType = ref("");
provide("equipmentKindCategoryUuid_alertEdit", equipmentKindCategoryUuid_alertEditBreakdownType);


onMounted(() => fnInit());

const fnInit = () => fnSearch();

const fnResetSearch = () => {
  name_search.value = "";
  equipmentKindCategoryUuid_search.value = "";
};

const fnSearch = () => {
  ajaxGetBreakdownTypes({
    "@~[]": ["EquipmentKindCategory"],
    name: name_search.value,
    equipment_kind_category_uuid: equipmentKindCategoryUuid_search.value,
  })
    .then(res => {
      rows.value = collect(res.content.breakdown_types)
        .map((breakdownType, idx) => {
          return {
            index: idx + 1,
            uuid: breakdownType.uuid,
            name: breakdownType.name,
            equipmentKindCategory: breakdownType.equipment_kind_category,
            operation: { uuid: breakdownType.uuid },
          }
        })
        .all();
    })
    .catch(e => errorNotify(e.msg));
};

const fnResetAlertCreateBreakdownType = () => {
  name_alertCreateBreakdownType.value = "";
  equipmentKindCategoryUuid_alertCreateBreakdownType.value = "";
};

const fnOpenAlertCreateBreakdownType = () => {
  alertCreateBreakdownType.value = true;
};

const fnStoreBreakdownType = () => {
  const loading = loadingNotify();

  ajaxStoreBreakdownType({
    name: name_alertCreateBreakdownType.value,
    equipment_kind_category_uuid: equipmentKindCategoryUuid_alertCreateBreakdownType.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnSearch();
    })
    .catch(e => errorNotify(e.msg))
    .finally(() => loading());
};

const fnOpenAlertEditBreakdownType = params => {
  if (!params["uuid"]) return;
  currentUuid.value = params.uuid;

  ajaxGetBreakdownType(currentUuid.value, {
    "@~[]": ["EquipmentKindCategory"],
  })
    .then(res => {
      name_alertEditBreakdownType.value = res.content.breakdown_type.name;
      equipmentKindCategoryUuid_alertEditBreakdownType.value = res.content.breakdown_type.equipment_kind_category.uuid;

      alertEditBreakdownType.value = true;
    })
    .catch(e => errorNotify(e.msg));
};

const fnUpdateBreakdownType = () => {
  const loading = loadingNotify();

  ajaxUpdateBreakdownType(currentUuid.value, {
    name: name_alertEditBreakdownType.value,
    equipment_kind_category_uuid: equipmentKindCategoryUuid_alertEditBreakdownType.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnSearch();
    })
    .catch(e => errorNotify(e.msg))
    .finally(() => loading());
};

const fnDestroyBreakdownType = params => {
  if (!params["uuid"]) return;

  confirmNotify(
    destroyActions(() => {
      const loading = loadingNotify();

      ajaxDestroyBreakdownType(params.uuid)
        .then(() => {
          successNotify("删除成功");
          fnSearch();
        })
        .catch(e => errorNotify(e.msg))
        .finally(loading());
    })
  );
};

const fnDestroyBreakdownTypes = () => {
  confirmNotify(
    destroyActions(() => {
      const loading = loadingNotify();

      ajaxDestroyBreakdownTypes(collect(selected.value).pluck("uuid").toArray())
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
