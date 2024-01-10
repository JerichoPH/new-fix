<template>
  <div class="q-pa-md">
    <q-card>
      <q-card-section>
        <div class="row">
          <div class="col"><span :style="{ fontSize: '20px' }">搜索</span></div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="primary" label="搜索" outline icon="search" @click="fnSearch" />
              <q-btn label="重置" outline flat @click="fnResetSearch" />
            </q-btn-group>
          </div>
        </div>
        <div class="row q-mt-md">
          <div class="col">
            <q-form>
              <div class="row q-col-gutter-sm">
                <div class="col">
                  <q-input outlined clearable lazy-rules v-model="uniqueCode_search" label="代码" :rules="[]"
                    class="q-mb-md" />
                </div>
                <div class="col">
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
          <div class="col"><span :style="{ fontSize: '20px' }">来源类型列表</span></div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="secondary" outline label="新建来源类型" icon="add" @click="fnOpenAlertCreateCreateSourceType" />
              <q-btn color="negative" outline label="删除来源类型" icon="deleted" @click="fnDestroyCreateSourceTypes" />
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
                      <q-btn @click="fnOpenAlertEditCreateSourceType(props.row.operation)" color="warning" icon="edit" outline>
                        编辑
                      </q-btn>
                      <q-btn @click="fnDestroyCreateSourceType(props.row.operation)" color="negative" icon="delete" outline>
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
  <!-- 新建来源类型弹窗 -->
  <q-dialog v-model="alertCreateSourceType" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">新建来源类型</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnStoreSourceType">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="uniqueCode_alertCreateSourceType" label="代码" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertCreateSourceType" label="名称" :rules="[]" />
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
  <!-- 编辑来源类型 -->
  <q-dialog v-model="alertEditSourceType" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">编辑来源类型</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnUpdateSourceType">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="uniqueCode_alertEditSourceType" label="代码" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertEditSourceType" label="名称" :rules="[]" />
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
</template>
<script setup>
import { ref, onMounted } from "vue";
import collect from "collect.js";
import { fnColumnReverseSort } from "src/utils/common";
import {
  ajaxGetSourceTypes,
  ajaxGetSourceType,
  ajaxStoreSourceType,
  ajaxUpdateSoruceType,
  ajaxDestroySourceType,
  ajaxDestroySourceTypes,
} from "src/apis/source";
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

// 新建来源类型弹窗数据
const alertCreateSourceType = ref(false);
const uniqueCode_alertCreateSourceType = ref("");
const name_alertCreateSourceType = ref("");

// 编辑来源类型弹窗数据
const alertEditSourceType = ref(false);
const currentUuid = ref("");
const uniqueCode_alertEditSourceType = ref("");
const name_alertEditSourceType = ref("");

onMounted(() => fnInit());

const fnInit = () => fnSearch();

const fnResetSearch = () => {
  uniqueCode_search.value = "";
  name_search.value = "";
};

const fnSearch = () => {
  rows.value = [];
  selected.value = [];

  ajaxGetSourceTypes({
    "@~[]": ["SourceProjects", "Equipments"],
    unique_code: uniqueCode_search.value,
    name: name_search.value,
  })
    .then(res => {
      rows.value = collect(res.content.source_types).
        map((sourceType, idx) => {
          return {
            index: idx + 1,
            uuid: sourceType.uuid,
            uniqueCode: sourceType.unique_code,
            name: sourceType.name,
            operation: { uuid: sourceType.uuid },
          };
        })
        .all();
    })
    .catch(e => errorNotify(e.msg));
};

const fnResetAlertCreateSourceType = () => {
  uniqueCode_alertCreateSourceType.value = "";
  name_alertCreateSourceType.value = "";
};

const fnOpenAlertCreateCreateSourceType = () => { alertCreateSourceType.value = true; };

const fnStoreSourceType = () => {
  const loading = loadingNotify();

  ajaxStoreSourceType({
    unique_code: uniqueCode_alertCreateSourceType.value,
    name: name_alertCreateSourceType.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnSearch();
      fnResetAlertCreateSourceType();

      alertCreateSourceType.value = false;
    })
    .catch(e => errorNotify(e.msg))
    .finally(loading());
};

const fnOpenAlertEditCreateSourceType = params => {
  if (!params["uuid"]) return;
  currentUuid.value = params.uuid;

  ajaxGetSourceType(currentUuid.value)
    .then(res => {
      uniqueCode_alertEditSourceType.value = res.content.source_type.unique_code;
      name_alertEditSourceType.value = res.content.source_type.name;

      alertEditSourceType.value = true;
    })
    .catch(e => {
      errorNotify(e.msg)
    });
};

const fnUpdateSourceType = () => {
  if (!currentUuid.value) return;

  const loading = loadingNotify();

  ajaxUpdateSoruceType(currentUuid.value, {
    unique_code: uniqueCode_alertEditSourceType.value,
    name: name_alertEditSourceType.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnSearch();

      alertEditSourceType.value = false;
    })
    .catch(e => errorNotify(e.msg))
    .finally(loading());
};

const fnDestroyCreateSourceType = params => {
  if (!params["uuid"]) return;

  confirmNotify(destroyActions(() => {
    const loading = loadingNotify();

    ajaxDestroySourceType(params.uuids)
      .then(() => {
        successNotify("删除成功");
        fnSearch();
      })
      .catch(e => errorNotify(e.msg))
      .finally(loading());
  }));
};

const fnDestroyCreateSourceTypes = () => {
  confirmNotify(destroyActions(() => {
    const loading = loadingNotify();

    ajaxDestroySourceTypes(collect(selected.value).pluck("uuid").all())
      .then(() => {
        successNotify("删除成功");
        fnSearch();
      })
      .catch(e => errorNotify(e.msg))
      .finally(loading());
  }));
};
</script>
