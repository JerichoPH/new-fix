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
                  <sel-organization-railway_search label-name="所属路局" />
                </div>
                <div class="col">
                  <sel-organization-paragraph_search label-name="所属站段" />
                </div>
                <div class="col">
                  <sel-organization-workshop_search label-name="所属车间" />
                </div>
                <div class="col">
                  <sel-organization-work-area_search label-name="所属工区" />
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
          <div class="col"><span :style="{ fontSize: '20px' }">仓库-库房列表</span></div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="secondary" label="新建仓库-库房" icon="add" @click="fnOpenAlertCreateWarehouseStorehouse" />
              <q-btn color="negative" label="删除仓库-库房" icon="deleted" @click="fnDestroyWarehouseStorehouses" />
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
                  <q-th align="left" key="organizationWorkshop"
                    @click="(event) => fnColumnReverseSort(event, props, sortBy)">
                    所属车间
                  </q-th>
                  <q-th align="left" key="organizationWorkArea"
                    @click="(event) => fnColumnReverseSort(event, props, sortBy)">
                    所属工区
                  </q-th>
                  <q-th align="right"></q-th>
                </q-tr>
              </template>
              <template v-slot:body="props">
                <q-tr :props="props">
                  <q-td><q-checkbox :key="props.row.uuid" :value="props.row.uuid" v-model="props.selected" /></q-td>
                  <q-td>{{ props.row.index }}</q-td>
                  <q-td key="name" :props="props">{{ props.row.name }}</q-td>
                  <q-td key="organizationWorkshop" :props="props">
                    <join-string :values="[
                      props.row.organizationRailway.short_name,
                      props.row.organizationParagraph.name,
                      props.row.organizationWorkshop.name,
                    ]" />
                  </q-td>
                  <q-td key="organizationWorkArea" :props="props">{{ props.row.organizationWorkArea ?
                    props.row.organizationWorkArea.name : '' }}</q-td>
                  <q-td key="operation" :props="props">
                    <q-btn-group>
                      <q-btn @click="fnOpenAlertEditWarehouseStorehouse(props.row.operation)" color="warning" icon="edit">
                        编辑
                      </q-btn>
                      <q-btn @click="fnDestroyWarehouseStorehouse(props.row.operation)" color="negative" icon="delete">
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
  <!-- 新建仓库-库房弹窗 -->
  <q-dialog v-model="alertCreateWarehouseStorehouse" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">新建仓库-库房</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnStoreWarehouseStorehouse">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertCreateWarehouseStorehouse" label="名称"
                :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-railway_alert-create label-name="所属路局" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <SelOrganizationParagraph_alertCreate label-name="所属站段" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-workshop_alert-create label-name="所属车间" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <SelOrganizationWorkArea_alertCreate label-name="所属工区" />
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
  <!-- 编辑仓库-库房弹窗 -->
  <q-dialog v-model="alertEditWarehouseStorehouse" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">编辑仓库-库房</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnUpdateWarehouseStorehouse">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertEditWarehouseStorehouse" label="名称" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-railway_alert-edit label-name="所属路局" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-paragraph_alert-edit label-name="所属站段" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-workshop_alert-edit label-name="所属车间" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-work-area_alert-edit label-name="所属工区" />
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
  ajaxGetWarehouseStorehouses,
  ajaxGetWarehouseStorehouse,
  ajaxStoreWarehouseStorehouse,
  ajaxUpdateWarehouseStorehouse,
  ajaxDestroyWarehouseStorehouse,
  ajaxDestroyWarehouseStorehouses,
} from "src/apis/warehouse";
import {
  loadingNotify,
  successNotify,
  errorNotify,
  confirmNotify,
  destroyActions,
} from "src/utils/notify";
import JoinString from "src/components/JoinString.vue";
import SelOrganizationRailway_search from "src/components/SelOrganizationRailway_search.vue";
import SelOrganizationParagraph_search from "src/components/SelOrganizationParagraph_search.vue";
import SelOrganizationWorkshop_search from "src/components/SelOrganizationWorkshop_search.vue";
import SelOrganizationWorkArea_search from "src/components/SelOrganizationWorkArea_search.vue";
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

// 表格数据
const rows = ref([]);
const selected = ref([]);
const sortBy = ref("");

// 新建仓库-库房弹窗数据
const alertCreateWarehouseStorehouse = ref(false);
const name_alertCreateWarehouseStorehouse = ref("");
const organizationRailwayUuid_alertCreateWarehouseStorehouse = ref("");
provide("organizationRailwayUuid_alertCreate", organizationRailwayUuid_alertCreateWarehouseStorehouse);
const organizationParagraphUuid_alertCreateWarehouseStorehouse = ref("");
provide("organizationParagraphUuid_alertCreate", organizationParagraphUuid_alertCreateWarehouseStorehouse);
const organizationWorkshopUuid_alertCreateWarehouseStorehouse = ref("");
provide("organizationWorkshopUuid_alertCreate", organizationWorkshopUuid_alertCreateWarehouseStorehouse);
const organizationWorkAreaUuid_alertCreateWarehouseStorehouse = ref("");
provide("organizationWorkAreaUuid_alertCreate", organizationWorkAreaUuid_alertCreateWarehouseStorehouse);

// 编辑仓库-库房弹窗数据
const alertEditWarehouseStorehouse = ref(false);
const currentUuid = ref("");
const name_alertEditWarehouseStorehouse = ref("");
const organizationRailwayUuid_alertEditWarehouseStorehouse = ref("");
provide("organizationRailwayUuid_alertEdit", organizationRailwayUuid_alertEditWarehouseStorehouse);
const organizationParagraphUuid_alertEditWarehouseStorehouse = ref("");
provide("organizationParagraphUuid_alertEdit", organizationParagraphUuid_alertEditWarehouseStorehouse);
const organizationWorkshopUuid_alertEditWarehouseStorehouse = ref("");
provide("organizationWorkshopUuid_alertEdit", organizationWorkshopUuid_alertEditWarehouseStorehouse);
const organizationWorkAreaUuid_alertEditWarehouseStorehouse = ref("");
provide("organizationWorkAreaUuid_alertEdit", organizationWorkAreaUuid_alertEditWarehouseStorehouse);

onMounted(() => fnInit());

const fnInit = () => fnSearch();

const fnResetSearch = () => {
  name_search.value = "";
  organizationRailwayUuid_search.value = "";
  organizationParagraphUuid_search.value = "";
  organizationWorkshopUuid_search.value = "";
};

const fnSearch = () => {
  rows.value = [];
  selected.value = [];

  ajaxGetWarehouseStorehouses({
    "@~[]": [
      "OrganizationWorkshop",
      "OrganizationWorkshop.OrganizationParagraph",
      "OrganizationWorkshop.OrganizationParagraph.OrganizationRailway",
      "OrganizationWorkArea",
    ],
    name: name_search.value,
    organization_railway_uuid: organizationRailwayUuid_search.value,
    organization_paragraph_uuid: organizationParagraphUuid_search.value,
    organization_workshop_uuid: organizationWorkshopUuid_search.value,
    organization_work_area_uuid: organizationWorkAreaUuid_search.value,
  })
    .then(res => {
      rows.value = collect(res.content.warehouse_storehouses)
        .map((warehouseStorehouse, idx) => {
          return {
            index: idx + 1,
            uuid: warehouseStorehouse.uuid,
            name: warehouseStorehouse.name,
            organizationWorkshop: warehouseStorehouse.organization_workshop,
            organizationParagraph: warehouseStorehouse.organization_workshop.organization_paragraph,
            organizationRailway: warehouseStorehouse.organization_workshop.organization_paragraph.organization_railway,
            organizationWorkArea: warehouseStorehouse.organization_work_area,
            operation: { uuid: warehouseStorehouse.uuid },
          }
        })
        .all();
    })
    .catch(e => errorNotify(e.msg));
};

const fnResetAlertCreateWarehouseStorehouse = () => {
  name_alertCreateWarehouseStorehouse.value = "";
  organizationWorkshopUuid_alertCreateWarehouseStorehouse.value = "";
  organizationWorkAreaUuid_alertCreateWarehouseStorehouse.value = "";
};

const fnOpenAlertCreateWarehouseStorehouse = () => { alertCreateWarehouseStorehouse.value = true; };

const fnStoreWarehouseStorehouse = () => {
  const loading = loadingNotify();

  ajaxStoreWarehouseStorehouse({
    name: name_alertCreateWarehouseStorehouse.value,
    organization_workshop_uuid: organizationWorkshopUuid_alertCreateWarehouseStorehouse.value,
    organization_work_area_uuid: organizationWorkAreaUuid_alertCreateWarehouseStorehouse.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnSearch();
      fnResetAlertCreateWarehouseStorehouse();

      alertCreateWarehouseStorehouse.value = false;
    })
    .catch(e => errorNotify(e.msg))
    .finally(loading());
};

const fnOpenAlertEditWarehouseStorehouse = (params) => {
  if (!params["uuid"]) return;
  currentUuid.value = params.uuid;

  ajaxGetWarehouseStorehouse(currentUuid.value, {
    "@~[]": [
      "OrganizationWorkshop",
      "OrganizationWorkshop.OrganizationParagraph",
      "OrganizationWorkshop.OrganizationParagraph.OrganizationRailway",
      "OrganizationWorkArea",
    ],
  })
    .then(res => {
      name_alertEditWarehouseStorehouse.value = res.content.warehouse_storehouse.name;
      organizationRailwayUuid_alertEditWarehouseStorehouse.value = res.content.warehouse_storehouse.organization_workshop.organization_paragraph.organization_railway.uuid;
      organizationParagraphUuid_alertEditWarehouseStorehouse.value = res.content.warehouse_storehouse.organization_workshop.organization_paragraph.uuid;
      organizationWorkshopUuid_alertEditWarehouseStorehouse.value = res.content.warehouse_storehouse.organization_workshop.uuid;
      organizationWorkAreaUuid_alertEditWarehouseStorehouse.value = res.content.warehouse_storehouse.organization_work_area ? res.content.warehouse_storehouse.organization_work_area.uuid : '';

      alertEditWarehouseStorehouse.value = true;
    })
    .catch(e => errorNotify(e.msg));
};

const fnUpdateWarehouseStorehouse = () => {
  if (!currentUuid.value) return;

  const loading = loadingNotify();
  ajaxUpdateWarehouseStorehouse(currentUuid.value, {
    name: name_alertEditWarehouseStorehouse.value,
    organization_workshop_uuid: organizationWorkshopUuid_alertEditWarehouseStorehouse.value,
    organization_work_area_uuid: organizationWorkAreaUuid_alertEditWarehouseStorehouse.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnSearch();

      alertEditWarehouseStorehouse.value = false;
    })
    .catch(e => errorNotify(e.msg))
    .finally(loading());
};

const fnDestroyWarehouseStorehouse = (params) => {
  if (!params["uuid"]) return;

  confirmNotify(destroyActions(() => {
    const loading = loadingNotify();

    ajaxDestroyWarehouseStorehouse(params["uuid"])
      .then(() => {
        successNotify("删除成功");
        fnSearch();
      })
      .catch(e => errorNotify(e.msg))
      .finally(loading());
  }));
};

const fnDestroyWarehouseStorehouses = () => {
  confirmNotify(destroyActions(() => {
    const loading = loadingNotify();

    ajaxDestroyWarehouseStorehouses(collect(selected.value).pluck("uuid").all())
      .then(() => {
        successNotify("删除成功");
        fnSearch();
      })
      .catch(e => errorNotify(e.msg))
      .finally(loading());
  }));
};
</script>
