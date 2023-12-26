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
              <div class="row q-pb-sm q-col-gutter-sm">
                <div class="col-3">
                  <q-input outlined clearable lazy-rules v-model="uniqueCode_search" label="代码" :rules="[]"
                    class="q-mb-md" />
                </div>
                <div class="col-3">
                  <q-input outlined clearable lazy-rules v-model="name_search" label="名称" :rules="[]" class="q-mb-md" />
                </div>
                <div class="col-3">
                  <sel-organization-railway_search label-name="所属路局" />
                </div>
                <div class="col-3">
                  <sel-organization-paragraph_search label-name="所属站段" />
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
          <div class="col"><span :style="{ fontSize: '20px' }">车间列表</span></div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="secondary" label="新建车间" icon="add" @click="fnOpenAlertCreateOrganizationWorkshop" />
              <q-btn color="negative" label="删除车间" icon="deleted" @click="fnDestroyOrganizationWorkshops" />
            </q-btn-group>
          </div>
        </div>
        <div class="row q-mt-md">
          <div class="col">
            <q-table flat bordered title="车间列表" :rows="rows" row-key="uuid" :pagination="{ rowsPerPage: 200 }"
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
                  <q-th align="left" key="organizationParagraph"
                    @click="(event) => fnColumnReverseSort(event, props, sortBy)">
                    所属站段
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
                  <q-td key="organizationParagraph" :props="props">{{ props.row.organizationParagraph.name }}</q-td>
                  <q-td key="operation" :props="props">
                    <q-btn-group>
                      <q-btn @click="fnOpenAlertEditOrganizationWorkshop(props.row.operation)" color="warning"
                        icon="edit">
                        编辑
                      </q-btn>
                      <q-btn @click="fnDestroyOrganizationWorkshop(props.row.operation)" color="negative" icon="delete">
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
  <!-- 新建车间弹窗 -->
  
  <!-- 编辑车间弹窗 -->
</template>

<script setup>
import { ref, onMounted, provide } from "vue";
import collect from "collect.js";
import {
  ajaxGetOrganizationWorkshops,
  ajaxGetOrganizationWorkshop,
  ajaxStoreOrganizationWorkshop,
  ajaxUpdateOrganizationWorkshop,
  ajaxDestroyOrganizationWorkshop,
  ajaxDestroyOrganizationWorkshops,
} from "src/apis/organization";
import {
  loadingNotify,
  successNotify,
  errorNotify,
  actionNotify,
  destroyActions,
} from "src/utils/notify";
import fnColumnReverseSort from "src/utils/common";
import SelOrganizationRailway_search from "src/components/SelOrganizationRailway_search.vue";
import SelOrganizationParagraph_search from "src/components/SelOrganizationParagraph_search.vue";

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
const sortBy = ref("");

onMounted(() => fnInit());

const fnInit = () => fnSearch();

const fnResetSearch = () => {
  uniqueCode_search.value = "";
  name_search.value = "";
  organizationRailwayUuid_search.value = "";
  organizationParagraphUuid_search.value = "";
};

const fnSearch = () => {
  ajaxGetOrganizationWorkshops({
    unique_code: uniqueCode_search.value,
    name: name_search.value,
    organization_railway_uuid: organizationRailwayUuid_search.value,
    organization_paragraph_uuid: organizationParagraphUuid_search.value,
  })
    .then(res => {
      rows.value = collect(res.content.organization_workshops)
        .map((organizationWorkshop, idx) => {
          return {
            index: idx + 1,
            uuid: organizationWorkshop.uuid,
            unique_code: organizationWorkshop.unique_code,
            name: organizationWorkshop.name,
            organizationParagraph: organizationWorkshop.organization_paragraph,
            operation: { uuid: organizationWorkshop.uuid },
          }
        })
        .all()
    })
    .catch(e => errorNotify(e.msg));
};

const fnResetAlertCreateOrganizationWorkshop = () => {

};

const fnOpenAlertCreateOrganizationWorkshop = () => { };
const fnOpenAlertEditOrganizationWorkshop = params => { };
const fnDestroyOrganizationWorkshop = params => { };
const fnDestroyOrganizationWorkshops = () => { };
</script>
