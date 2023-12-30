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
                <div class="col-3">
                  <sel-organization-railway_search label-name="所属路局" :ajax-params="{}" />
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
            <span :style="{ fontSize: '20px' }">站段列表</span>
          </div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="secondary" label="新建站段" icon="add" @click="fnOpenAlertCreateOrganizationParagraph" />
              <q-btn color="negative" label="删除站段" icon="deleted" @click="fnDestroyOrganizationParagraphs" />
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
                  <q-th align="left" key="organizationRailway"
                    @click="(event) => fnColumnReverseSort(event, props, sortBy)">
                    所属路局
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
                  <q-td key="organizationRailway" :props="props">{{ props.row.organizationRailway.short_name }}</q-td>
                  <q-td key="operation" :props="props">
                    <q-btn-group>
                      <q-btn @click="
                        fnOpenAlertEditOrganizationParagraph(
                          props.row.operation
                        )
                        " color="warning" icon="edit">
                        编辑
                      </q-btn>
                      <q-btn @click="
                        fnDestroyOrganizationParagraph(props.row.operation)
                        " color="negative" icon="delete">
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
  <!-- 新建站段弹窗 -->
  <q-dialog v-model="alertCreateOrganizationParagraph">
    <q-card style="width: 800px">
      <q-card-section>
        <div class="text-h6">新建站段</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnStoreOrganizationParagraph">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="uniqueCode_alertCreateOrganizationParagraph" label="代码"
                :rules="[]" />
              <q-input outlined clearable lazy-rules v-model="name_alertCreateOrganizationParagraph" label="名称"
                :rules="[]" class="q-mt-md" />
              <sel-organization-railway_alert-create label-name="所属路局" :ajax-params="{}" class="q-mt-md" />
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn type="submit" label="确定" icon="check" color="secondary" v-close-popup />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
  <!-- 编辑站段弹窗 -->
  <q-dialog v-model="alertEditOrganizationParagraph">
    <q-card style="width: 800px">
      <q-card-section>
        <div class="text-h6">编辑站段</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnUpdateOrganizationParagraph">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertEditOrganizationParagraph" label="名称"
                :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-railway_alert-edit label-name="所属路局" :ajax-params="{}" />
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
import {
  ajaxGetOrganizationParagraphs,
  ajaxGetOrganizationParagraph,
  ajaxStoreOrganizationParagraph,
  ajaxUpdateOrganizationParagraph,
  ajaxDestroyOrganizationParagraph,
  ajaxDestroyOrganizationParagraphs,
} from "src/apis/organization";
import {
  loadingNotify,
  successNotify,
  errorNotify,
  actionNotify,
  destroyActions,
} from "src/utils/notify";
import { fnColumnReverseSort } from "src/utils/common";
import SelOrganizationRailway_search from "src/components/SelOrganizationRailway_search.vue";
import SelOrganizationRailway_alertCreate from "src/components/SelOrganizationRailway_alertCreate.vue";
import SelOrganizationRailway_alertEdit from "src/components/SelOrganizationRailway_alertEdit.vue";

// 搜索栏数据
const uniqueCode_search = ref("");
const name_search = ref("");
const organizationRailwayUuid_search = ref("");
provide("organizationRailwayUuid_search", organizationRailwayUuid_search);
const organizationParagraphUuid_search = ref("");
provide("organizationParagraphUuid_search", organizationParagraphUuid_search);

// 表格数据
const rows = ref([]);
const selected = ref([]);
const sortBy = ref('');

// 新建站段弹窗数据
const alertCreateOrganizationParagraph = ref(false);
const uniqueCode_alertCreateOrganizationParagraph = ref("");
const name_alertCreateOrganizationParagraph = ref("");
const organizationRailwayUuid_alertCreateOrganizationParagraph = ref("");
provide("organizationRailwayUuid_alertCreate", organizationRailwayUuid_alertCreateOrganizationParagraph);

// 编辑站段弹窗数据
const currentUuid = ref("");
const alertEditOrganizationParagraph = ref(false);
const name_alertEditOrganizationParagraph = ref("");
const organizationRailwayUuid_alertEditOrganizationParagraph = ref("");
provide("organizationRailwayUuid_alertEdit", organizationRailwayUuid_alertEditOrganizationParagraph);

onMounted(() => fnInit());

const fnInit = () => fnSearch();

const fnResetSearch = () => {
  uniqueCode_search.value = "";
  name_search.value = "";
  organizationRailwayUuid_search.value = "";
  organizationParagraphUuid_search.value = "";
};

const fnSearch = () => {
  rows.value = [];
  selected.value = [];

  ajaxGetOrganizationParagraphs({
    "@~[]": ["OrganizationRailway"],
    unique_code: uniqueCode_search.value,
    name: name_search.value,
    organization_railway_uuid: organizationRailwayUuid_search.value,
    organization_paragraph_uuid: organizationParagraphUuid_search.value,
  })
    .then(res => {
      rows.value = collect(res.content.organization_paragraphs)
        .map((organizationRailway, idx) => {
          return {
            index: idx + 1,
            uuid: organizationRailway.uuid,
            uniqueCode: organizationRailway.unique_code,
            name: organizationRailway.name,
            organizationRailway: organizationRailway.organization_railway,
            operation: { uuid: organizationRailway.uuid },
          }
        })
        .toArray();
    })
    .catch(e => errorNotify(e.msg));
};

const fnResetAlertOrganizationParagraph = () => {
  uniqueCode_alertCreateOrganizationParagraph.value = "";
  name_search.value = "";
  organizationRailwayUuid_alertCreateOrganizationParagraph.value = "";
};

const fnOpenAlertCreateOrganizationParagraph = () => {
  alertCreateOrganizationParagraph.value = true;
};

const fnStoreOrganizationParagraph = () => {
  const loading = loadingNotify();

  ajaxStoreOrganizationParagraph({
    unique_code: uniqueCode_alertCreateOrganizationParagraph.value,
    name: name_alertCreateOrganizationParagraph.value,
    organization_railway_uuid: organizationRailwayUuid_alertCreateOrganizationParagraph.value,
  })
    .then(res => {
      successNotify(res.message);
      fnResetAlertOrganizationParagraph();
      fnSearch();
    })
    .catch(e => errorNotify(e.msg))
    .finally(() => loading());
};

const fnOpenAlertEditOrganizationParagraph = params => {
  if (!params["uuid"]) return;

  currentUuid.value = params["uuid"];

  ajaxGetOrganizationParagraph(currentUuid.value, {
    "@~[]": ["OrganizationRailway"],
  })
    .then(res => {
      name_alertEditOrganizationParagraph.value = res.content.organization_paragraph.name;
      organizationRailwayUuid_alertEditOrganizationParagraph.value = res.content.organization_paragraph.organization_railway.uuid;
      alertEditOrganizationParagraph.value = true;
    })
    .catch(e => errorNotify(e.msg));
};

const fnUpdateOrganizationParagraph = () => {
  if (!currentUuid.value) return;

  const loading = loadingNotify();
  ajaxUpdateOrganizationPragraph(currentUuid.value, {
    name: name_alertEditOrganizationParagraph.value,
    organization_railway_uuid: organizationRailwayUuid_alertEditOrganizationParagraph.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnSearch();
    })
    .catch(e => errorNotify(e.msg))
    .finally(() => loading());
};

const fnDestroyOrganizationParagraph = params => {
  if (!params["uuid"]) return;

  actionNotify(
    destroyActions(() => {
      ajaxDestroyOrganizationParagraph(params["uuid"])
        .then(res => {
          successNotify("删除成功");
          fnSearch();
        })
        .catch(e => errorNotify(e.msg));
    })
  )
};

const fnDestroyOrganizationParagraphs = () => {
  actionNotify(
    destroyActions(() => {
      ajaxDestroyOrganizationParagraphs(selected.value)
        .then(res => {
          successNotify("删除成功");
          fnSearch();
        })
        .catch(e => errorNotify(e.msg));
    })
  )
};
</script>
src/apis/organization
