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
                <div class="col">
                  <q-input outlined clearable lazy-rules v-model="name_search" label="名称" :rules="[]" class="q-mb-md" />
                </div>
                <div class="col">
                  <sel-equipment-kind-category_search labelName="所属种类" v-model="equipmentKindCategoryUuid_search" />
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
          <div class="col"><span :style="{ fontSize: '20px' }">器材类型列表</span></div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="secondary" label="新建器材类型" icon="add" @click="fnOpenAlertCreateEquipmentKindType" />
              <q-btn color="negative" label="删除器材类型" icon="deleted" @click="fnDestroyEquipmentKindTypes" />
            </q-btn-group>
          </div>
        </div>
        <div class="row q-mt-md">
          <div class="col">
            <q-table flat bordered title="器材类型列表" :rows="rows" row-key="uuid" :pagination="{ rowsPerPage: 200 }"
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
                  <q-td key="uniqueCode" :props="props">{{ props.row.uniqueCode }}</q-td>
                  <q-td key="name" :props="props">{{ props.row.name }}</q-td>
                  <q-td key="equipmentKindCategory" :props="props">{{ props.row.equipmentKindCategory.name }}</q-td>
                  <q-td key="operation" :props="props">
                    <q-btn-group>
                      <q-btn @click="
                        fnOpenAlertEditEquipmentKindType(props.row.operation)
                        " color="warning" icon="edit">
                        编辑
                      </q-btn>
                      <q-btn @click="fnDestroyEquipmentKindType(props.row.operation)" color="negative" icon="delete">
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
  <!-- 新建器材类型弹窗 -->
  <q-dialog v-model="alertCreateEquipmentKindType">
    <q-card style="width: 800px">
      <q-card-section>
        <div class="text-h6">新建器材类型</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnStoreEquipmentKindType">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertCreateEquipmentKindType" label="名称" :rules="[]" />
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn type="submit" label="确定" icon="check" color="secondary" v-close-popup />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
  <!-- 编辑器材类型弹窗 -->
</template>

<script setup>
import { ref, onMounted } from "vue";
import { fnColumnReverseSort } from "src/utils/common";
import { ajaxEquipmentKindCategoryList, ajaxEquipmentKindTypeList, ajaxEquipmentKindTypeDetail, ajaxEquipmentKindTypeStore, ajaxEquipmentKindTypeUpdate, ajaxEquipmentKindTypeDestroy, ajaxEquipmentKindTypeDestroyMany, } from "src/apis/equipmentKind";
import { loadingNotify, successNotify, errorNotify, actionNotify, getDestroyActions } from "src/utils/notify";
import {SelEquipmentKindCategory_search} from "src/components/SelEquipmentKindCategory_search.vue";

// 搜索栏条件
const name_search = ref("");
const equipmentKindCategoryUuid_search = ref("");
const equipmentKindCategories = ref([]);

// 表格数据
const rows = ref([]);
const selected = ref([]);
const sortBy = ref("");

// 新建器材类型弹窗数据
const alertCreateEquipmentKindType = ref(false);
const name_alertCreateEquipmentKindType = ref("");
const equipmentKindCategoryUuid_alertCreateEquipmentKindType = ref("");

onMounted(() => {
  fnInit();
});

/**
 * 重置搜索栏条件
 */
const fnResetSearch = () => {
  name_search.value = "";
  equipmentKindCategoryUuid_search.value = "";
};

/**
 * 初始化页面
 */
const fnInit = () => {
  fnGetEquipmentKindCategories();
  fnSearch();
};

/**
 * 获取种类列表
 */
const fnGetEquipmentKindCategories = () => {
  equipmentKindCategories.value = [];
  ajaxEquipmentKindCategoryList()
    .then((res) => {
      equipmentKindCategories.value = res.content.equipment_kind_categories
        .map((item) => {
          return {
            label: item.name,
            value: item.uuid,
          };
        });
    });
};

/**
 * 搜索
 */
const fnSearch = () => {
  rows.value = [];
  ajaxEquipmentKindTypeList({
    name: name_search.value,
    equipment_kind_category_uuid: equipmentKindCategoryUuid_search.value,
  })
    .then((res) => {
      rows.value = res.content.equipment_kind_types
        .map((equipmentKindType, idx) => {
          return {
            index: idx + 1,
            uuid: equipmentKindType.uuid,
            uniqueCode: equipmentKindType.unique_code,
            name: equipmentKindType.name,
            equipmentKindCategory: equipmentKindType.equipment_kind_category,
            operation: { uuid: equipmentKindType.uuid },
          };
        });
    });
};

/**
 * 重置新建器材类型弹窗
 */
const fnResetAlertCreateEquipmentKindType = () => {

};

/**
 * 打开新建器材类型弹窗
 */
const fnOpenAlertCreateEquipmentKindType = () => { };

/**
 * 新建器材类型
 */
const fnStoreEquipmentKindType = () => { };

/**
 * 批量删除器材类型
 */
const fnDestroyEquipmentKindTypes = params => { };


/**
 * 打开编辑器材类型弹窗
 * @param {{*}} params
 */
const fnOpenAlertEditEquipmentKindType = params => { };

/**
 * 编辑器材类型
 * @param {{*}} params
 */
const fnUpdateEquipmentKindType = params => { };

/**
 * 删除器材类型
 */
const fnDestroyEquipmentKindType = params => { };
</script>
