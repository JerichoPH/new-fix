<template>
  <div class="q-pa-md">
    <q-card>
      <q-card-section>
        <div class="q-mb-md">
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
            <span :style="{ fontSize: '20px' }">器材种类列表</span>
          </div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="secondary" outline label="新建器材种类" icon="add" @click="fnOpenAlertCreateEquipmentKindCategory" />
              <q-btn color="negative" outline label="删除器材种类" icon="deleted" @click="fnDestroyEquipmentKindCategories" />
            </q-btn-group>
          </div>
        </div>
        <div class="row q-mt-md">
          <div class="col">
            <q-table flat bordered title="" :rows="rows" row-key="uuid" :pagination="{ rowsPerPage: 200 }"
              :rows-per-page-options="[50, 100, 200, 0]" rows-per-page-label="分页" selection="multiple"
              v-model:selected="selected">
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
                        fnOpenAlertEditEquipmentKindCategory(
                          props.row.operation
                        )
                        " color="warning" icon="edit" outline>
                        编辑
                      </q-btn>
                      <q-btn @click="
                        fnDestroyEquipmentKindCategory(props.row.operation)
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
  <!-- 新建器材种类弹窗 -->
  <q-dialog v-model="alertCreateEquipmentKindCategory" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">新建器材种类</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnStoreEquipmentKindCategory">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertCreateEquipmentKindCategory" label="名称"
                :rules="[]" />
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
  <!-- 编辑器材种类弹窗 -->
  <q-dialog v-model="alertEditEquipmentKindCategory" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">编辑器材种类</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnUpdateEquipmentKindCategory">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertEditEquipmentKindCategory" label="名称"
                :rules="[]" />
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
  loadingNotify,
  successNotify,
  errorNotify,
  confirmNotify,
  destroyActions,
} from "src/utils/notify";
import {
  ajaxGetEquipmentKindCategories,
  ajaxGetEquipmentKindCategory,
  ajaxStoreEquipmentKindCategory,
  ajaxUpdateEquipmentKindCategory,
  ajaxDestroyEquipmentKindCategory,
  ajaxDestroyEquipmentKindCategories,
} from "src/apis/equipmentKind";

// 搜索栏条件
const name_search = ref("");

// 表格数据
const rows = ref([]);
const selected = ref([]);
const sortBy = ref("");

// 新建器材种类弹窗数据
const alertCreateEquipmentKindCategory = ref(false);
const name_alertCreateEquipmentKindCategory = ref("");

// 编辑器材种类弹窗数据
const alertEditEquipmentKindCategory = ref(false);
const name_alertEditEquipmentKindCategory = ref("");
const currentUuid = ref("");

onMounted(() => fnInit());

/**
 * 初始化
 */
const fnInit = () => fnSearch();;

/**
 * 重置搜索栏条件
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

  ajaxGetEquipmentKindCategories({
    name: name_search.value,
  })
    .then(res => {
      rows.value = res.content.equipment_kind_categories.map(
        (equipmentKindCategory, idx) => {
          return {
            index: idx + 1,
            uuid: equipmentKindCategory.uuid,
            uniqueCode: equipmentKindCategory.unique_code,
            name: equipmentKindCategory.name,
            operation: { uuid: equipmentKindCategory.uuid },
          };
        }
      );
    })
    .catch(e => errorNotify(e.msg));
};

/**
 * 重置新建器材种类弹窗
 */
const fnRestAlertCreateEquipmentCategory = () => {
  name_alertCreateEquipmentKindCategory.value = "";
};

/**
 * 打开新建器材种类弹窗
 */
const fnOpenAlertCreateEquipmentKindCategory = () => {
  alertCreateEquipmentKindCategory.value = true;
};

/**
 * 新建器材种类
 */
const fnStoreEquipmentKindCategory = () => {
  const loading = loadingNotify();

  ajaxStoreEquipmentKindCategory({
    name: name_alertCreateEquipmentKindCategory.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnSearch();
      fnRestAlertCreateEquipmentCategory();

      alertCreateEquipmentKindCategory.value = false;
    })
    .catch(e=>errorNotify(e.msg))
    .finally(loading());
};

/**
 * 打开器材种类编辑弹窗
 * @param {{*}} params
 */
const fnOpenAlertEditEquipmentKindCategory = (params) => {
  if (!params["uuid"]) return;
  currentUuid.value = params["uuid"];

  ajaxGetEquipmentKindCategory(params.uuid)
    .then(res => {
      name_alertEditEquipmentKindCategory.value =
        res.content.equipment_kind_category.name;
      alertEditEquipmentKindCategory.value = true;
    })
    .catch(e=>errorNotify(e.msg))
    .finally(loadingNotify());
};

/**
 * 编辑器材种类
 * @param {{*}} params
 */
const fnUpdateEquipmentKindCategory = (params) => {
  if (!currentUuid.value) return;

  const loading = loadingNotify();

  ajaxUpdateEquipmentKindCategory(currentUuid.value, {
    name: name_alertEditEquipmentKindCategory.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnSearch();

      alertEditEquipmentKindCategory.value = false;
    })
    .catch(e=>errorNotify(e.msg))
    .finally(loading());
};

/**
 * 删除器材种类
 * @param {{*}} params
 */
const fnDestroyEquipmentKindCategory = (params) => {
  if (!params["uuid"]) return;

  confirmNotify(
    destroyActions(() => {
      const loading = loadingNotify();

      ajaxDestroyEquipmentKindCategory(params.uuid)
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
 * 批量删除器材种类
 */
const fnDestroyEquipmentKindCategories = () => {
  confirmNotify(
    destroyActions(() => {
      ajaxDestroyEquipmentKindCategories(
        selected.value.map((item) => item.uuid)
      )
        .then(() => {
          successNotify("删除成功");
          fnSearch();
        })
        .catch(e=>errorNotify(e.msg))
        .finally(loadingNotify());
    })
  );
};
</script>
