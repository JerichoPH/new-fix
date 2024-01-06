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
          <div class="col"><span :style="{ fontSize: '20px' }">室内上道位置-机房类型列表</span></div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="secondary" label="新建室内上道位置-机房类型" icon="add" @click="fnOpenAlertCreateInstallIndoorRoomType" />
              <q-btn color="negative" label="删除室内上道位置-机房类型" icon="deleted" @click="fnDestroyInstallIndoorRoomTypes" />
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
                  <q-th align="left" key="uniqueCode" @click="event => fnColumnReverseSort(event, props, sortBy)">
                    代码
                  </q-th>
                  <q-th align="left" key="name" @click="event => fnColumnReverseSort(event, props, sortBy)">
                    名称
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
                  <q-td key="operation" :props="props">
                    <q-btn-group>
                      <q-btn @click="fnOpenAlertEditInstallIndoorRoomType(props.row.operation)" color="warning"
                        icon="edit">
                        编辑
                      </q-btn>
                      <q-btn @click="fnDestroyInstallIndoorRoomType(props.row.operation)" color="negative" icon="delete">
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
  <!-- 新建室内上道位置-机房类型弹窗 -->
  <q-dialog v-model="alertCreateInstallIndoorRoomType" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">新建室内上道位置-机房类型</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnStoreInstallIndoorRoomType">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="uniqueCode_alertCreateInstallIndoorRoomType" label="代码"
                :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertCreateInstallIndoorRoomType" label="名称"
                :rules="[]" />
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
  <!-- 编辑室内上道位置-机房类型弹窗 -->
  <q-dialog v-model="alertEditInstallIndoorRoomType" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">编辑室内上道位置-机房类型</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnUpdateInstallIndoorRoomType">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="uniqueCode_alertEditInstallIndoorRoomType" label="名称"
                :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertEditInstallIndoorRoomType" label="名称"
                :rules="[]" />
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn-group>
            <q-btn type="button" label="关闭" v-close-popup />
            <q-btn type="submit" label="确定" icon="check" color="warning" />
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
  ajaxGetInstallIndoorRoomTypes,
  ajaxGetInstallIndoorRoomType,
  ajaxStoreInstallIndoorRoomType,
  ajaxUpdateInstallIndoorRoomType,
  ajaxDestroyInstallIndoorRoomType,
  ajaxDestroyInstallIndoorRoomTypes,
} from "src/apis/install";
import { notifies, actions } from "src/utils/notify";

// 搜索栏数据
const uniqueCode_search = ref("");
const name_search = ref("");

// 表格数据
const rows = ref([]);
const selected = ref([]);
const sortBy = ref("");

// 新建室内上道位置-机房类型数据
const alertCreateInstallIndoorRoomType = ref(false);
const uniqueCode_alertCreateInstallIndoorRoomType = ref("");
const name_alertCreateInstallIndoorRoomType = ref("");

// 编辑室内上道位置-机房类型数据
const alertEditInstallIndoorRoomType = ref(false);
const currentUuid = ref("");
const uniqueCode_alertEditInstallIndoorRoomType = ref("");
const name_alertEditInstallIndoorRoomType = ref("");

onMounted(() => fnInit());

const fnInit = () => fnSearch();

const fnResetSearch = () => {
  uniqueCode_search.value = "";
  name_search.value = "";
};

const fnSearch = () => {
  ajaxGetInstallIndoorRoomTypes({
    unique_code: uniqueCode_search.value,
    name: name_search.value,
  })
    .then(res => {
      rows.value = collect(res.content.install_indoor_room_types)
        .map((installIndoorRoomType, idx) => {
          return {
            index: idx + 1,
            uuid: installIndoorRoomType.uuid,
            uniqueCode: installIndoorRoomType.unique_code,
            name: installIndoorRoomType.name,
            operation: { uuid: installIndoorRoomType.uuid },
          };
        })
        .all();
    })
    .catch(e => notifies.error(e.msg));
};

const fnResetAlertCreateInstallIndoorRoomType = () => {
  uniqueCode_alertCreateInstallIndoorRoomType.value = "";
  name_alertCreateInstallIndoorRoomType.value = "";
};

const fnOpenAlertCreateInstallIndoorRoomType = () => { alertCreateInstallIndoorRoomType.value = true; }

const fnStoreInstallIndoorRoomType = () => {
  const loading = notifies.loading();

  ajaxStoreInstallIndoorRoomType({
    unique_code: uniqueCode_alertCreateInstallIndoorRoomType.value,
    name: name_alertCreateInstallIndoorRoomType.value,
  })
    .then(res => {
      notifies.success(res.msg);
      fnSearch();
      fnResetAlertCreateInstallIndoorRoomType();

      alertCreateInstallIndoorRoomType.value = false;
    })
    .catch(e => error(e.msg))
    .finally(loading());
};

const fnOpenAlertEditInstallIndoorRoomType = params => {
  if (!params["uuid"]) return;
  currentUuid.value = params.uuid;

  ajaxGetInstallIndoorRoomType(currentUuid.value)
    .then(res => {
      uniqueCode_alertEditInstallIndoorRoomType.value = res.content.install_indoor_room_type.unique_code;
      name_alertEditInstallIndoorRoomType.value = res.content.install_indoor_room_type.name;

      alertEditInstallIndoorRoomType.value = true;
    })
    .catch(e => notifies.error(e.msg));
};

const fnUpdateInstallIndoorRoomType = () => {
  if (!currentUuid.value) return;

  const loading = notifies.loading();

  ajaxUpdateInstallIndoorRoomType(currentUuid.value, {
    unique_code: uniqueCode_alertEditInstallIndoorRoomType.value,
    name: name_alertEditInstallIndoorRoomType.value,
  })
    .then(res => {
      notifies.success(res.msg);
      fnSearch();

      alertEditInstallIndoorRoomType.value = false;
    })
    .catch(e => notifies.error(e.msg))
    .finally(loading());
};

const fnDestroyInstallIndoorRoomType = params => {
  if (!params["uuid"]) return;

  notifies.confirm(actions.destory(() => {
    const loading = notifies.loading();

    ajaxDestroyInstallIndoorRoomType(params.uuid)
      .then(() => {
        notifies.success("删除成功");
        fnSearch();
      })
      .catch(e => notifies.error(e.msg))
      .finally(loading());
  }));
};

const fnDestroyInstallIndoorRoomTypes = () => {
  notifies.confirm(actions.destory(() => {
    const loading = notifies.loading();

    ajaxDestroyInstallIndoorRoomTypes(collect(selected.value).pluck("uuid"))
      .then(() => {
        notifies.success("删除成功");
        fnSearch();
      })
      .catch(e => notifies.error(e.msg))
      .finally(loading());
  }));
};
</script>
