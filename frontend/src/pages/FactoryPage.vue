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
              </div>
            </q-form>
          </div>
        </div>
      </q-card-section>
    </q-card>

    <q-card class="q-mt-md">
      <q-card-section>
        <div class="row">
          <div class="col"><span :style="{ fontSize: '20px' }">厂家列表</span></div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="secondary" label="新建厂家" icon="add" @click="fnOpenAlertCreateFactory" />
              <q-btn color="negative" label="删除厂家" icon="deleted" @click="fnDestroyFactories" />
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
                  <q-th align="left" key="uniqueCode" @click="(event) => fnColumnReverseSort(event, props, sortBy)">
                    代码
                  </q-th>
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
                  <q-td key="uniqueCode" :props="props">{{ props.row.uniqueCode }}</q-td>
                  <q-td key="name" :props="props">{{ props.row.name }}</q-td>
                  <q-td key="operation" :props="props">
                    <q-btn-group>
                      <q-btn @click="fnOpenAlertEditFactory(props.row.operation)" color="warning" icon="edit">
                        编辑
                      </q-btn>
                      <q-btn @click="fnDestroyFactory(props.row.operation)" color="negative" icon="delete">
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
  <!-- 新建厂家弹窗 -->
  <q-dialog v-model="alertCreateFactory" no-backdrop-dismiss>
    <q-card :style="{minWidth: '450px'}">
      <q-card-section>
        <div class="text-h6">新建厂家</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnStoreFactory">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="uniqueCode_alertCreateFactory" label="代码" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertCreateFactory" label="名称" :rules="[]" />
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn type="submit" label="关闭" v-close-popup />
          <q-btn type="submit" label="确定" icon="check" color="secondary" v-close-popup />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
  <!-- 编辑厂家弹窗 -->
  <q-dialog v-model="alertEditFactory" no-backdrop-dismiss>
    <q-card :style="{minWidth: '450px'}">
      <q-card-section>
        <div class="text-h6">编辑厂家</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnUpdateFactory">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="uniqueCode_alertEditFactory" label="代码" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertEditFactory" label="名称" :rules="[]" />
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn type="submit" label="关闭" v-close-popup />
          <q-btn type="submit" label="确定" icon="check" color="secondary" v-close-popup />
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
  ajaxGetFactories,
  ajaxGetFactory,
  ajaxStoreFactory,
  ajaxUpdateFactory,
  ajaxDestroyFactory,
  ajaxDestroyFactories,
} from "src/apis/factory";
import {
  loadingNotify,
  successNotify,
  errorNotify,
  confirmNotify,
  destroyActions,
} from "src/utils/notify";

// 搜索栏数据
const name_search = ref("");

// 表格数据
const rows = ref([]);
const selected = ref([]);
const sortBy = ref("");

// 新建厂家弹窗数据
const alertCreateFactory = ref(false);
const uniqueCode_alertCreateFactory = ref("");
const name_alertCreateFactory = ref("");

// 编辑厂家弹窗数据
const alertEditFactory = ref(false);
const currentUuid = ref("");
const uniqueCode_alertEditFactory = ref("");
const name_alertEditFactory = ref("");

onMounted(() => fnInit());

const fnInit = () => fnSearch();

const fnResetSearch = () => {
  name_search.value = "";
};

const fnSearch = () => {
  rows.value = [];

  ajaxGetFactories({
    name: name_search.value,
  })
    .then(res => {
      rows.value = collect(res.content.factories)
        .map((factory, idx) => {
          return {
            index: idx + 1,
            uuid: factory.uuid,
            uniqueCode: factory.unique_code,
            name: factory.name,
            operation: { uuid: factory.uuid },
          }
        })
        .all();
    })
    .catch(e => errorNotify(e.msg));
};

const fnResetAlertCreateFactory = () => {
  name_search.value = "";
};

const fnOpenAlertCreateFactory = () => {
  alertCreateFactory.value = true;
};

const fnStoreFactory = () => {
  const loading = loadingNotify();
  ajaxStoreFactory({
    unique_code: uniqueCode_alertCreateFactory.value,
    name: name_alertCreateFactory.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnResetAlertCreateFactory();
      fnSearch();
    })
    .catch(e => errorNotify(e.msg))
    .finally(() => loading());
};

const fnOpenAlertEditFactory = params => {
  if (!params["uuid"]) return;
  currentUuid.value = params.uuid;

  ajaxGetFactory(currentUuid.value, params)
    .then(res => {
      console.log(res);
      uniqueCode_alertEditFactory.value = res.content.factory.unique_code;
      name_alertEditFactory.value = res.content.factory.name;
      alertEditFactory.value = true;
    })
    .catch(e => errorNotify(e.msg));
};

const fnUpdateFactory = () => {
  if (!currentUuid.value) return;

  const loading = loadingNotify();

  ajaxUpdateFactory(currentUuid.value, {
    unique_code: uniqueCode_alertEditFactory.value,
    name: name_alertEditFactory.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnSearch();
    })
    .catch(e => errorNotify(e.msg))
    .finally(() => loading());
};

const fnDestroyFactory = params => {
  if (!params["uuid"]) return;

  confirmNotify(
    destroyActions(() => {
      const loading = loadingNotify();

      ajaxDestroyFactory(params.uuid)
        .then(() => {
          successNotify("删除成功");
          fnSearch();
        })
        .catch(e => errorNotify(e.msg))
        .finally(() => loading());
    })
  );
};

const fnDestroyFactories = () => {
  confirmNotify(
    destroyActions(() => {
      const loading = loadingNotify();

      ajaxDestroyFactories(collect(selected.value).pluck("uuid").all())
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
