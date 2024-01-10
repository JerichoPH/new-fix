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
                  <standard-select label-name="所属来源类型" sechma="search" current-field="sourceTypeUuid"
                :data-source="ajaxGetSourceTypes" data-source-field="source_types" />
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
          <div class="col"><span :style="{ fontSize: '20px' }">来源项目列表</span></div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="secondary" label="新建来源项目" icon="add" @click="fnOpenAlertCreateSourceProject" />
              <q-btn color="negative" label="删除来源项目" icon="deleted" @click="fnDestroySourceProjects" />
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
                  <q-th align="left" key="sourceType" @click="(event) => fnColumnReverseSort(event, props, sortBy)">
                    所属来源类型
                  </q-th>
                  <q-th align="right"></q-th>
                </q-tr>
              </template>
              <template v-slot:body="props">
                <q-tr :props="props">
                  <q-td><q-checkbox :key="props.row.uuid" :value="props.row.uuid" v-model="props.selected" /></q-td>
                  <q-td>{{ props.row.index }}</q-td>
                  <q-td key="name" :props="props">{{ props.row.name }}</q-td>
                  <q-td key="sourceType" :props="props">{{ props.row.sourceType.name }}</q-td>
                  <q-td key="operation" :props="props">
                    <q-btn-group>
                      <q-btn @click="fnOpenAlertEditSourceProject(props.row.operation)" color="warning" icon="edit">
                        编辑
                      </q-btn>
                      <q-btn @click="fnDestroySourceProject(props.row.operation)" color="negative" icon="delete">
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
  <!-- 新建来源项目弹窗 -->
  <q-dialog v-model="alertCreateSourceProject" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">新建来源项目</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnStoreSourceProject">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertCreateSourceProject" label="名称" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属来源类型" sechma="alertCreate" current-field="sourceTypeUuid"
                :data-source="ajaxGetSourceTypes" data-source-field="source_types" />
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
  <!-- 编辑来源项目弹窗 -->
  <q-dialog v-model="alertEditSourceProject" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">编辑来源项目</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnUpdateSourceProject">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertEditSourceProject" label="名称" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属来源类型" sechma="alertEdit" current-field="sourceTypeUuid"
                :data-source="ajaxGetSourceTypes" data-source-field="source_types" />
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
import { ref, onMounted, provide } from "vue";
import collect from "collect.js";
import { fnColumnReverseSort } from "src/utils/common";
import {
  ajaxGetSourceProjects,
  ajaxGetSourceProject,
  ajaxStoreSourceProject,
  ajaxUpdateSourceProject,
  ajaxDestroySourceProject,
  ajaxDestroySourceProjects,
  ajaxGetSourceTypes,
} from "src/apis/source";
import {
  loadingNotify,
  successNotify,
  errorNotify,
  confirmNotify,
  destroyActions,
} from "src/utils/notify";
import StandardSelect from "src/components/StandardSelect.vue";

// 搜索栏数据
const name_search = ref("");
const sourceTypeUuid_search = ref("");
provide("sourceTypeUuid_search", sourceTypeUuid_search);

// 表格数据
const rows = ref([]);
const selected = ref([]);
const sortBy = ref("");

// 新建来源项目弹窗数据
const alertCreateSourceProject = ref(false);
const name_alertCreateSourceProject = ref("");
const sourceTypeUuid_alertCreateSourceProject = ref("");
provide("sourceTypeUuid_alertCreate", sourceTypeUuid_alertCreateSourceProject);

// 编辑来源项目弹窗数据
const alertEditSourceProject = ref(false);
const currentUuid = ref("");
const name_alertEditSourceProject = ref("");
const sourceTypeUuid_alertEditSourceProject = ref("");
provide("sourceTypeUuid_alertEdit", sourceTypeUuid_alertEditSourceProject);

onMounted(() => fnInit());

const fnInit = () => fnSearch();

const fnResetSearch = () => {
  name_search.value = "";
  sourceTypeUuid_search.value = "";
};

const fnSearch = () => {
  rows.value = [];
  selected.value = [];

  ajaxGetSourceProjects({
    "@~[]": ["SourceType", "Equipments"],
    name: name_search.value,
    source_type_uuid: sourceTypeUuid_search.value,
  })
    .then(res => {
      rows.value = collect(res.content.source_projects)
        .map((sourceProject, idx) => {
          return {
            index: idx + 1,
            uuid: sourceProject.uuid,
            name: sourceProject.name,
            sourceType: sourceProject.source_type,
            operation: { uuid: sourceProject.uuid },
          }
        })
        .all();
    })
    .catch(e => errorNotify(e.msg));
};

const fnResetAlertCreateSourceProject = () => {
  name_alertCreateSourceProject.value = "";
  sourceTypeUuid_alertCreateSourceProject.value = "";
};

const fnOpenAlertCreateSourceProject = () => { alertCreateSourceProject.value = true; };

const fnStoreSourceProject = () => {
  ajaxStoreSourceProject({
    name: name_alertCreateSourceProject.value,
    source_type_uuid: sourceTypeUuid_alertCreateSourceProject.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnResetAlertCreateSourceProject();
      fnSearch();

      alertCreateSourceProject.value = false;
    })
    .catch(e => errorNotify(e.msg));
};

const fnOpenAlertEditSourceProject = params => {
  if (!params["uuid"]) return;
  currentUuid.value = params.uuid;

  ajaxGetSourceProject(currentUuid.value, {
    "@~[]": ["SourceType", "Equipments"],
  })
    .then(res => {
      name_alertEditSourceProject.value = res.content.source_project.name;
      sourceTypeUuid_alertEditSourceProject.value = res.content.source_project.source_type.uuid;

      alertEditSourceProject.value = true;
    })
    .catch(e => errorNotify(e.msg));
};

const fnUpdateSourceProject = () => {
  if (!currentUuid.value) return;

  ajaxUpdateSourceProject(currentUuid.value, {
    name: name_alertEditSourceProject.value,
    source_type_uuid: sourceTypeUuid_alertEditSourceProject.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnSearch();

      alertEditSourceProject.value = false;
    })
    .catch(e => errorNotify(e.msg));
};

const fnDestroySourceProject = params => {
  if (!params["uuid"]) return;

  confirmNotify(destroyActions(() => {
    const loading = loadingNotify();

    ajaxDestroySourceProject(params.uuid)
      .then(() => {
        successNotify("删除成功");
        fnSearch();
      })
      .catch(e => errorNotify(e.msg))
      .finally(loading());
  }));
};

const fnDestroySourceProjects = () => {
  confirmNotify(destroyActions(() => {
    const loading = loadingNotify();

    ajaxDestroySourceProjects(collect(selected.value).pluck("uuid").all())
      .then(() => {
        successNotify("删除成功");
        fnSearch();
      })
      .catch(e => errorNotify(e.msg))
      .finally(loading());
  }))
};

</script>
