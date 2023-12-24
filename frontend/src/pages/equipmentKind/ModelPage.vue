<template>
  <div class="q-pa-md">
    <q-card>
      <q-card-section>
        <div class="q-mb-md">
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
              <div class="row q-pb-sm q-col-gutter-sm">
                <div class="col-3">
                  <q-input outlined clearable lazy-rules v-model="name_search" label="名称" :rules="[]" class="q-mb-md" />
                </div>
                <div class="col-3">
                  <sel-equipment-kind-category_search label-name="器材种类" :ajax-params="{}" />
                </div>
                <div class="col-3">
                  <sel-equipment-kind-type_search label-name="器材类型" :ajax-params="{}" />
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
          <div class="col"><span :style="{ fontSize: '20px' }">器材型号列表</span></div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="secondary" label="新建器材型号" icon="add" @click="fnOpenAlertCreateEquipmentKindModel" />
              <q-btn color="negative" label="删除器材型号" icon="deleted" @click="fnDestroyEquipmentKindModels" />
            </q-btn-group>
          </div>
        </div>
        <div class="row q-mt-md">
          <div class="col">
            <q-table flat bordered title="器材型号列表" :rows="rows" row-key="uuid" :pagination="{ rowsPerPage: 200 }"
              :rows-per-page-options="[50, 100, 200, 0]" rows-per-page-label="分页" :selected-rows-label="getSelectedString"
              selection="multiple" v-model:selected="selected">
              <template v-slot:header="props">
                <q-tr :props="props">
                  <q-th align="left"><q-checkbox key="allCheck" v-model="props.selected" /></q-th>
                  <q-th align="left">#</q-th>
                  <q-th align="left" key="uniqueCode" @click="(event) => fnColumnReverseSort(event, props, sortBy)">
                    代码
                  </q-th>
                  <q-th align="left" key="name" @click="(event) => fnColumnReverseSort(event, props, sortBy)">
                    名称
                  </q-th>
                  <q-th align="left" key="parent" @click="(event) => fnColumnReverseSort(event, props, sortBy)">
                    器材种类类型
                  </q-th>
                  <q-th align="right"></q-th>
                </q-tr>
              </template>
              <template v-slot:body="props">
                <q-tr :props="props">
                  <q-td><q-checkbox :key="props.row.uuid" :value="props.row.uuid" v-model="props.selected" /></q-td>
                  <q-td>{{ props.row.index }}</q-td>
                  <q-td key="uniqueCode" :props="props">{{ props.row.uniqueCode }}</q-td>
                  <q-td key="name" :props="props">{{ props.row.name }}</q-td>
                  <q-td key="parent" :props="props">
                    {{ props.row.parent.equipmentKindType.equipmentKindCategory.name }}
                    {{ props.row.parent.equipmentKindType.name }}
                  </q-td>
                  <q-td key="operation" :props="props">
                    <q-btn-group>
                      <q-btn @click="
                        fnOpenAlertEditEquipmentKindModel(props.row.operation)
                        " color="warning" icon="edit">
                        编辑
                      </q-btn>
                      <q-btn @click="fnDestroyEquipmentKindModel(props.row.operation)" color="negative" icon="delete">
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
  <!-- 新建器材型号弹窗 -->
  <q-dialog v-model="alertCreateEquipmentKindModel">
    <q-card style="width: 800px">
      <q-card-section>
        <div class="text-h6">新建器材型号</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnStoreEquipomentKindModel">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertCreateEquipmentKindModel" label="名称"
                :rules="[]" />
              <sel-equipment-kind-category_alert-create label-name="器材种类" :ajax-params="{}" class="q-mt-md"/>
              <sel-equipment-kind-type_alert-create label-name="器材类型" :ajax-params="{}" class="q-mt-md"/>
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn type="submit" label="确定" icon="check" color="secondary" v-close-popup />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
  <!-- 编辑器材型号弹窗 -->
</template>

<script setup>
import { ref, onMounted, provide } from "vue";
import {
  ajaxEquipmentKindModelList,
  ajaxEquipmentKindModelDetail,
  ajaxEquipmentKindModelStore,
  ajaxEquipmentKindModelUpdate,
  ajaxEquipmentKindModelDestroy,
  ajaxEquipmentKindModelDestroyMany,
} from "/src/apis/equipmentKind";
import {
  loadingNotify,
  successNotify,
  errorNotify,
  actionNotify,
  destroyActions,
} from "src/utils/notify";
import { fnColumnReverseSort } from "src/utils/common";
import SelEquipmentKindCategory_search from "src/components/SelEquipmentKindCategory_search.vue";
import SelEquipmentKindType_search from "src/components/SelEquipmentKindType_search.vue";
import SelEquipmentKindCategory_alertCreate from "src/components/SelEquipmentKindCategory_alertCreate.vue";
import SelEquipmentKindType_alertCreate from "src/components/SelEquipmentKindType_alertCreate.vue";

// 搜索栏数据
const name_search = ref("");
const equipmentKindCategoryUuid_search = ref("");
provide("equipmentKindCategoryUuid_search", equipmentKindCategoryUuid_search);
const equipmentKindTypeUuid_search = ref("");
provide("equipmentKindTypeUuid_search", equipmentKindTypeUuid_search);

// 表格数据
const rows = ref([]);
const selected = ref([]);
const sortBy = ref("name");

// 新建器材型号弹窗数据
const alertCreateEquipmentKindModel = ref(false);
const name_alertCreateEquipmentKindModel = ref("");
const equipmentKindCategoryUuid_alertCreate = ref("")
provide("equipmentKindCategoryUuid_alertCreate", equipmentKindCategoryUuid_alertCreate);
const equipmentKindTypeUuid_alertCreate = ref("")
provide("equipmentKindTypeUuid_alertCreate", equipmentKindTypeUuid_alertCreate);

onMounted(() => { fnInit(); });

const fnInit = () => {
  fnSearch();
};

/**
 * 重置搜索栏
 */
const fnResetSearch = () => {
  name_search.value = "";
  equipmentKindCategoryUuid_search.value = "";
  equipmentKindTypeUuid_search.value = "";
};

/**
 * 搜索
 */
const fnSearch = () => {
  rows.value = [];
  selected.value = [];

  ajaxEquipmentKindModelList({
    ":~[]": ["EquipmentKindType", "EquipmentKindType.EquipmentKindCategory"],
    name: name_search.value,
    equipment_kind_category_uuid: equipmentKindCategoryUuid_search.value,
    equipment_kind_type_uuid: equipmentKindTypeUuid_search.value,
  })
    .then((res) => {
      if (res.content.equipment_kind_models.length > 0) {
        collect(res.content.equipment_kind_models)
          .each((equipmentKindModel, index) => {
            rows.value.push({
              index: index + 1,
              uuid: equipmentKindModel.uuid,
              uniqueCode: equipmentKindModel.unique_code,
              name: equipmentKindModel.name,
              parent: {
                equipmentKindCategory: equipmentKindModel.equipment_kind_type.equipment_kind_category,
                equipmentKindType: equipment_kind_type,
              },
              operation: { uuid: equipmentKindModel.uuid },
            });
          });
      }
    })
    .catch((e) => {
      errorNotify(e.msg);
    });
};

/**
 * 重置新建器材型号弹窗
 */
const fnResetAlertCreateEquipmentKindModel = () => {
  name_alertCreateEquipmentKindModel.value = "";
  equipmentKindCategoryUuid_alertCreate.value = "";
  equipmentKindTypeUuid_alertCreate.value = "";
};

/**
 * 打开新建型号弹窗
 */
const fnOpenAlertCreateEquipmentKindModel = () => {
  alertCreateEquipmentKindModel.value = true;
};

/**
 * 新建器材型号
 */
const fnStoreEquipomentKindModel = () =>{
  ajaxEquipmentKindModelStore({
    name: name_alertCreateEquipmentKindModel.value,
    equipment_kind_type_uuid: equipmentKindTypeUuid_alertCreate.value,
  })
    .then((res) => {
      successNotify(res.msg);
      fnResetAlertCreateEquipmentKindModel();
      fnSearch();
    })
    .catch((e) => {
      errorNotify(e.msg);
    });
};

const fnDestroyEquipmentKindModels = () => { };
const fnOpenAlertEditEquipmentKindModel = () => { };
const fnDestroyEquipmentKindModel = () => { };
</script>
