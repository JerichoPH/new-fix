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
                <div class="col">
                  <q-input outlined clearable lazy-rules v-model="name_search" label="名称" :rules="[]" class="q-mb-md" />
                </div>
                <div class="col">
                  <sel-organization-railway_search label-name="所属路局" :ajax-params="{}" />
                </div>
                <div class="col">
                  <sel-oragnization-paragraph_search label-name="所属站段" :ajax-params="{}" />
                </div>
                <div class="col">
                  <sel-organization-workshop_search label-name="所属车间" :ajax-params="{}" />
                </div>
                <div class="col">
                  <sel-organization-work-area_search label-name="所属工区" :ajax-params="{}" />
                </div>
              </div>
              <div class="row q-col-gutter-sm">
                <div class="col">
                  <sel-warehouse-storehouse_search label-name="所属仓库-库房" :ajax-params="{}" />
                </div>
                <div class="col"></div>
                <div class="col"></div>
                <div class="col"></div>
                <div class="col"></div>
              </div>
            </q-form>
          </div>
        </div>
      </q-card-section>
    </q-card>

    <q-card class="q-mt-md">
      <q-card-section>
        <div class="row">
          <div class="col"><span :style="{ fontSize: '20px' }">仓库-库区列表</span></div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="secondary" label="新建仓库-库区" icon="add" @click="fnOpenAlertCreateWarehouseArea" />
              <q-btn color="negative" label="删除仓库-库区" icon="deleted" @click="fnDestroyWarehouseAreas" />
            </q-btn-group>
          </div>
        </div>
        <div class="row q-mt-md">
          <div class="col">
            <q-table flat bordered title="仓库-库区列表" :rows="rows" row-key="uuid" :pagination="{ rowsPerPage: 200 }"
              :rows-per-page-options="[50, 100, 200, 0]" rows-per-page-label="分页" :selected-rows-label="() => { }"
              selection="multiple" v-model:selected="selected">
              <template v-slot:header="props">
                <q-tr :props="props">
                  <q-th align="left"><q-checkbox key="allCheck" v-model="props.selected" /></q-th>
                  <q-th align="left">#</q-th>
                  <q-th align="left" key="name" @click="(event) => fnColumnReverseSort(event, props, sortBy)">
                    名称
                  </q-th>
                  <q-th align="left" key="warehouseStorehouse"
                    @click="(event) => fnColumnReverseSort(event, props, sortBy)">
                    所属仓库-库房
                  </q-th>
                  <q-th align="right"></q-th>
                </q-tr>
              </template>
              <template v-slot:body="props">
                <q-tr :props="props">
                  <q-td><q-checkbox :key="props.row.uuid" :value="props.row.uuid" v-model="props.selected" /></q-td>
                  <q-td>{{ props.row.index }}</q-td>
                  <q-td key="name" :props="props">{{ props.row.name }}</q-td>
                  <q-td key="warehouseStorehouse" :props="props">{{ [
                    props.row.organizationWorkshop.name,
                    props.row.organizationWorkArea.name,
                    props.row.warehouseStorehouse.name,
                  ].filter(val => { return val !== '' && val !== null }).join(' - ') }}</q-td>
                  <q-td key="operation" :props="props">
                    <q-btn-group>
                      <q-btn @click="fnOpenAlertEditWarehouseArea(props.row.operation)" color="warning" icon="edit">
                        编辑
                      </q-btn>
                      <q-btn @click="fnDestroyWarehouseArea(props.row.operation)" color="negative" icon="delete">
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
</template>
<script setup>
import { ref, onMounted, provide } from "vue";
import collect from "collect.js";
import { fnColumnReverseSort } from "src/utils/common";
import {
  ajaxGetWarehouseAreas,
  ajaxGetWarehouseArea,
  ajaxStoreWarehouseArea,
  ajaxUpdateWarehouseArea,
  ajaxDestroyWarehouseArea,
  ajaxDestroyWarehouseAreas,
} from "src/apis/warehouse";
import {
  loadingNotify,
  successNotify,
  errorNotify,
  confirmNotify,
  destroyActions,
} from "src/utils/notify";
import SelOrganizationRailway_search from "src/components/SelOrganizationRailway_search.vue";
import SelOragnizationParagraph_search from "src/components/SelOrganizationParagraph_search.vue";
import SelOrganizationWorkshop_search from "src/components/SelOrganizationWorkshop_search.vue";
import SelOrganizationWorkArea_search from "src/components/SelOrganizationWorkArea_search.vue";
import SelWarehouseStorehouse_search from "src/components/SelWarehouseStorehouse_search.vue";
import SelOrganizationRailway_alertCreate from "src/components/SelOrganizationRailway_alertCreate.vue";
import SelOrganizationParagraph_alertCreate from "src/components/SelOrganizationParagraph_alertCreate.vue";
import SelOrganizationWorkshop_alertCreate from "src/components/SelOrganizationWorkshop_alertCreate.vue";
import SelOrganizationWorkArea_alertCreate from "src/components/SelOrganizationWorkArea_alertCreate.vue";
import SelOrganizationRailway_alertEdit from "src/components/SelOrganizationRailway_alertEdit.vue";
import SelOrganizationParagraph_alertEdit from "src/components/SelOrganizationParagraph_alertEdit.vue";
import SelOrganizationWorkshop_alertEdit from "src/components/SelOrganizationWorkshop_alertEdit.vue";
import SelOrganizationWorkArea_alertEdit from "src/components/SelOrganizationWorkArea_alertEdit.vue";

// 搜索栏数据
const name_search = ref("");
const organizationRailwayUuid_search = ref("");
provide("organizationRailwayUuid_search", organizationRailwayUuid_search);
const organizationParagraphUuid_search = ref("");
provide("organizationParagraphUuid_search", organizationParagraphUuid_search);
const organizationWorkshopUuid_search = ref("");
provide("organizationWorkshopUuid_search", organizationWorkshopUuid_search);
const organizationWorkAreaUuid_search = ref("");
provide("organizationWorkAreaUuid_search", organizationWorkAreaUuid_search);
const warehouseStorehouseUuid_search = ref("");
provide("warehouseStorehouseUuid_search", warehouseStorehouseUuid_search);

// 表格数据
const rows = ref([]);
const selected = ref([]);
const sortBy = ref("");

onMounted(() => fnInit());

const fnInit = () => fnSearch();

const fnResetSearch = () => {
  name_search.value = "";
  organizationRailwayUuid_search.value = "";
  organizationParagraphUuid_search.value = "";
  organizationWorkshopUuid_search.value = "";
  organizationWorkAreaUuid_search.value = "";
  warehouseStorehouseUuid_search.value = "";
};

const fnSearch = () => {
  ajaxGetWarehouseAreas({
    "@~[]": [
      "WarehouseStorehouse",
      "WarehouseStorehouse.OrganizationWorkshop",
      "WarehouseStorehouse.OrganizationWorkArea",
      "WarehouseStorehouse.OrganizationWorkshop.OrganizationParagraph",
      "WarehouseStorehouse.OrganizationWorkshop.OrganizationParagraph.OrganizationRailway",
    ],
    organization_workshop_uuid: organizationWorkshopUuid_search.value,
    organizaiton_work_area_uuid: organizationWorkAreaUuid_search.value,
    warehouse_storehouse_uuid: warehouseStorehouseUuid_search.value,
  })
    .then(res => {
      rows.value = collect(res.content.warehouse_areas)
        .map((warehouseArea, idx) => {
          return {
            index: idx + 1,
            uuid: warehouseArea.uuid,
            name: warehouseArea.name,
            organizationRailway: warehouseArea.warehouse_storehouse.organization_workshop.organization_paragraph.organization_railway,
            organizationParagraph: warehouseArea.warehouse_storehouse.organization_workshop.organization_paragraph,
            organizationWorkshop: warehouseArea.warehouse_storehouse.organization_workshop,
            organizationWorkArea: warehouseArea.warehouse_storehouse.organization_work_area,
            warehouseStorehouse: warehouseArea.warehouse_storehouse,
            operation: { uuid: warehouseArea.uuid },
          }
        })
        .all();
    })
    .catch(e => {
      console.error('err', e);
      errorNotify(e.msg);
    });
};

const fnResetAlertCreateWarehouseArea = () => { };

const fnOpenAlertCreateWarehouseArea = () => { };

const fnStoreWarehouseArea = () => { };

const fnOpenAlertEditWarehouseArea = (params) => { };

const fnUpdateWarehouseArea = () => { };

const fnDestroyWarehouseArea = (params) => { };

const fnDestroyWarehouseAreas = () => { };
</script>
