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
              <div class="row q-col-gutter-sm">
                <div class="col">
                  <sel-warehouse-storehouse_search label-name="所属库房" />
                </div>
                <div class="col">
                  <sel-warehouse-area_search label-name="所属库区" />
                </div>
                <div class="col">
                  <sel-warehouse-platoon-type-code_search label-name="排类型" />
                </div>
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
          <div class="col"><span :style="{ fontSize: '20px' }">仓库-排列表</span></div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="secondary" label="新建仓库-排" icon="add" @click="fnOpenAlertCreateWarehousePlatoon" />
              <q-btn color="negative" label="删除仓库-排" icon="deleted" @click="fnDestroyWarehousePlatoons" />
            </q-btn-group>
          </div>
        </div>
        <div class="row q-mt-md">
          <div class="col">
            <q-table flat bordered title="仓库-排列表" :rows="rows" row-key="uuid" :pagination="{ rowsPerPage: 200 }"
              :rows-per-page-options="[50, 100, 200, 0]" rows-per-page-label="分页" :selected-rows-label="() => { }"
              selection="multiple" v-model:selected="selected">
              <template v-slot:header="props">
                <q-tr :props="props">
                  <q-th align="left"><q-checkbox key="allCheck" v-model="props.selected" /></q-th>
                  <q-th align="left">#</q-th>
                  <q-th align="left" key="name" @click="(event) => fnColumnReverseSort(event, props, sortBy)">
                    名称
                  </q-th>
                  <q-th align="left" key="typeText" @click="(event) => fnColumnReverseSort(event, props, sortBy)">
                    排类型
                  </q-th>
                  <q-th align="left" key="warehouseArea" @click="(event) => fnColumnReverseSort(event, props, sortBy)">
                    所属仓库-库区
                  </q-th>
                  <q-th align="right"></q-th>
                </q-tr>
              </template>
              <template v-slot:body="props">
                <q-tr :props="props">
                  <q-td><q-checkbox :key="props.row.uuid" :value="props.row.uuid" v-model="props.selected" /></q-td>
                  <q-td>{{ props.row.index }}</q-td>
                  <q-td key="name" :props="props">{{ props.row.name }}</q-td>
                  <q-td key="typeText" :props="props">{{ props.row.typeText }}</q-td>
                  <q-td key="warehouseArea" :props="props">
                    <join-string :values="[
                      props.row.warehouseStorehouse.name,
                      props.row.warehouseArea.name,
                    ]" />
                  </q-td>
                  <q-td key="operation" :props="props">
                    <q-btn-group>
                      <q-btn @click="fnOpenAlertEditWarehousePlatoon(props.row.operation)" color="warning" icon="edit">
                        编辑
                      </q-btn>
                      <q-btn @click="fnDestroyWarehousePlatoon(props.row.operation)" color="negative" icon="delete">
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
  <!-- 新建仓库-排弹窗 -->
  <q-dialog v-model="alertCreateWarehousePlatoon" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">新建仓库-排</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnStoreWarehousePlatoon">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertCreateWarehousePlatoon" label="名称" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-railway_alert-create label-name="所属路局" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-paragraph_alert-create label-name="所属站段" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-workshop_alert-create label-name="所属车间" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-work-area_alert-create label-name="所属工区" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-warehouse-storehouse_alert-create label-name="所属库房" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-warehouse-area_alert-create label-name="所属库区" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-warehouse-platoon-type-code_alert-create label-name="排类型" />
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
  <!-- 编辑仓库-排弹窗 -->
  <q-dialog v-model="alertWarehousePlatoon" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">编辑仓库-排弹窗</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnUpdateWarehouseArea">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertWarehousePlatoon" label="名称" :rules="[]" />
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
          <div class="row q-mt-md">
            <div class="col">
              <sel-warehouse-storehouse_alert-edit label-name="所属库房" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-warehouse-area_alert-edit label-name="所属库区" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-warehouse-platoon-type-code_alert-edit label-name="排类型" />
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
  ajaxGetWarehousePlatoons,
  ajaxGetWarehousePlatoon,
  ajaxStoreWarehousePlatoon,
  ajaxUpdateWarehousePlatoon,
  ajaxDestroyWarehousePlatoon,
  ajaxDestroyWarehousePlatoons,
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
import SelWarehouseStorehouse_search from "src/components/SelWarehouseStorehouse_search.vue";
import SelWarehouseArea_search from "src/components/SelWarehouseArea_search.vue";
import SelWarehousePlatoonTypeCode_search from "src/components/SelWarehousePlatoonTypeCode_search.vue";
import SelOrganizationRailway_alertCreate from "src/components/SelOrganizationRailway_alertCreate.vue";
import SelOrganizationParagraph_alertCreate from "src/components/SelOrganizationParagraph_alertCreate.vue";
import SelOrganizationWorkshop_alertCreate from "src/components/SelOrganizationWorkshop_alertCreate.vue";
import SelOrganizationWorkArea_alertCreate from "src/components/SelOrganizationWorkArea_alertCreate.vue";
import SelWarehouseStorehouse_alertCreate from "src/components/SelWarehouseStorehouse_alertCreate.vue";
import SelWarehouseArea_alertCreate from "src/components/SelWarehouseArea_alertCreate.vue";
import SelWarehousePlatoonTypeCode_alertCreate from "src/components/SelWarehousePlatoonTypeCode_alertCreate.vue";
import SelOrganizationRailway_alertEdit from "src/components/SelOrganizationRailway_alertEdit.vue";
import SelOrganizationParagraph_alertEdit from "src/components/SelOrganizationParagraph_alertEdit.vue";
import SelOrganizationWorkshop_alertEdit from "src/components/SelOrganizationWorkshop_alertEdit.vue";
import SelOrganizationWorkArea_alertEdit from "src/components/SelOrganizationWorkArea_alertEdit.vue";
import SelWarehouseStorehouse_alertEdit from "src/components/SelWarehouseStorehouse_alertEdit.vue";
import SelWarehouseArea_alertEdit from "src/components/SelWarehouseArea_alertEdit.vue";
import SelWarehousePlatoonTypeCode_alertEdit from "src/components/SelWarehousePlatoonTypeCode_alertEdit.vue";

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
const warehouseAreaUuid_search = ref("");
provide("warehouseAreaUuid_search", warehouseAreaUuid_search);
const warehousePlatoonTypeCode_search = ref("");
provide("warehousePlatoonTypeCode_search", warehousePlatoonTypeCode_search);

// 表格数据
const rows = ref([]);
const selected = ref([]);
const sortBy = ref("");

// 新建仓库-排弹窗数据
const alertCreateWarehousePlatoon = ref(false);
const name_alertCreateWarehousePlatoon = ref("");
const organizationRailwayUuid_alertCreateWarehousePlatoon = ref("");
provide("organizationRailwayUuid_alertCreate", organizationRailwayUuid_alertCreateWarehousePlatoon);
const organizationParagraphUuid_alertCreateWarehousePlatoon = ref("");
provide("organizationParagraphUuid_alertCreate", organizationParagraphUuid_alertCreateWarehousePlatoon);
const organizationWorkshopUuid_alertCreateWarehousePlatoon = ref("");
provide("organizationWorkshopUuid_alertCreate", organizationWorkshopUuid_alertCreateWarehousePlatoon);
const organizationWorkAreaUuid_alertCreateWarehousePlatoon = ref("");
provide("organizationWorkAreaUuid_alertCreate", organizationWorkAreaUuid_alertCreateWarehousePlatoon);
const warehouseStorehouseUuid_alertCreateWarehousePlatoon = ref("");
provide("warehouseStorehouseUuid_alertCreate", warehouseStorehouseUuid_alertCreateWarehousePlatoon);
const warehouseAreaUuid_alertCreateWarehousePlatoon = ref("");
provide("warehouseAreaUuid_alertCreate", warehouseAreaUuid_alertCreateWarehousePlatoon);
const warehousePlatoonTypeCode_alertCreateWarehousePlatoon = ref("");
provide("warehousePlatoonTypeCode_alertCreate", warehousePlatoonTypeCode_alertCreateWarehousePlatoon);

// 编辑仓库-排弹窗数据
const alertWarehousePlatoon = ref(false);
const currentUuid = ref("");
const name_alertWarehousePlatoon = ref("");
const organizationRailwayUuid_alertEditWarehousePlatoon = ref("");
provide("organizationRailwayUuid_alertEdit", organizationRailwayUuid_alertEditWarehousePlatoon);
const organizationParagraphUuid_alertEditWarehousePlatoon = ref("");
provide("organizationParagraphUuid_alertEdit", organizationParagraphUuid_alertEditWarehousePlatoon);
const organizationWorkshopUuid_alertEditWarehousePlatoon = ref("");
provide("organizationWorkshopUuid_alertEdit", organizationWorkshopUuid_alertEditWarehousePlatoon);
const organizationWorkAreaUuid_alertEditWarehousePlatoon = ref("");
provide("organizationWorkAreaUuid_alertEdit", organizationWorkAreaUuid_alertEditWarehousePlatoon);
const warehouseStorehouseUuid_alertEditWarehousePlatoon = ref("");
provide("warehouseStorehouseUuid_alertEdit", warehouseStorehouseUuid_alertEditWarehousePlatoon);
const warehouseAreaUuid_alertEditWarehousePlatoon = ref("");
provide("warehouseAreaUuid_alertEdit", warehouseAreaUuid_alertEditWarehousePlatoon);
const warehousePlatoonTypeCode_alertEditWarehousePlatoon = ref("");
provide("warehousePlatoonTypeCode_alertEdit", warehousePlatoonTypeCode_alertEditWarehousePlatoon);


onMounted(() => fnInit());

const fnInit = () => fnSearch();

const fnResetSearch = () => {
  name_search.value = "";
  organizationRailwayUuid_search.value = "";
  organizationParagraphUuid_search.value = "";
  organizationWorkshopUuid_search.value = "";
  organizationWorkAreaUuid_search.value = "";
  warehouseStorehouseUuid_search.value = "";
  warehouseAreaUuid_search.value = "";
};

const fnSearch = () => {
  rows.value = [];
  selected.value = [];

  ajaxGetWarehousePlatoons({
    "@~[]": [
      "WarehouseArea",
      "WarehouseArea.WarehouseStorehouse",
    ],
    name: name_search.value,
    organization_railway_uuid: organizationRailwayUuid_search.value,
    organization_paragraph_uuid: organizationParagraphUuid_search.value,
    organization_workshop_uuid: organizationWorkshopUuid_search.value,
    organization_work_area_uuid: organizationWorkAreaUuid_search.value,
    warehouse_storehouse_uuid: warehouseStorehouseUuid_search.value,
    warehouse_area_uuid: warehouseAreaUuid_search.value,
    type_code: warehousePlatoonTypeCode_search.value,
  })
    .then(resp => {
      rows.value = collect(resp.content.warehouse_platoons)
        .map((warehousePlatoon, idx) => {
          return {
            index: idx + 1,
            uuid: warehousePlatoon.uuid,
            name: warehousePlatoon.name,
            typeText: warehousePlatoon.type_text,
            warehouseArea: warehousePlatoon.warehouse_area,
            warehouseStorehouse: warehousePlatoon.warehouse_area.warehouse_storehouse,
            operation: { uuid: warehousePlatoon.uuid },
          };
        })
        .all();
    })
    .catch(e => errorNotify(e.msg));
};

const fnResetAlertCreateWarehousePlatoon = () => {
  name_alertCreateWarehousePlatoon.value = "";
  organizationRailwayUuid_alertCreateWarehousePlatoon.value = "";
  organizationParagraphUuid_alertCreateWarehousePlatoon.value = "";
  organizationWorkshopUuid_alertCreateWarehousePlatoon.value = "";
  organizationWorkAreaUuid_alertCreateWarehousePlatoon.value = "";
  warehouseStorehouseUuid_alertCreateWarehousePlatoon.value = "";
  warehouseAreaUuid_alertCreateWarehousePlatoon.value = "";
  warehousePlatoonTypeCode_search.value = "";
};

const fnOpenAlertCreateWarehousePlatoon = () => {
  alertCreateWarehousePlatoon.value = true;
};

const fnStoreWarehousePlatoon = () => {
  ajaxStoreWarehousePlatoon({
    name: name_alertCreateWarehousePlatoon.value,
    organization_railway_uuid: organizationRailwayUuid_alertCreateWarehousePlatoon.value,
    organization_paragraph_uuid: organizationParagraphUuid_alertCreateWarehousePlatoon.value,
    organization_workshop_uuid: organizationWorkshopUuid_alertCreateWarehousePlatoon.value,
    organization_work_area_uuid: organizationWorkAreaUuid_alertCreateWarehousePlatoon.value,
    warehouse_storehouse_uuid: warehouseStorehouseUuid_alertCreateWarehousePlatoon.value,
    warehouse_area_uuid: warehouseAreaUuid_alertCreateWarehousePlatoon.value,
    type_code: warehousePlatoonTypeCode_alertCreateWarehousePlatoon.value,
  })
    .then(response => {
      successNotify(response.msg);
      fnResetAlertCreateWarehousePlatoon();
      fnSearch();

      alertCreateWarehousePlatoon.value = false;
    })
    .catch(e => errorNotify(e.msg));
};

const fnOpenAlertEditWarehousePlatoon = params => {
  if (!params["uuid"]) return;
  currentUuid.value = params.uuid;

  ajaxGetWarehousePlatoon(params.uuid, {
    "@~[]": [
      "WarehouseArea",
      "WarehouseArea.WarehouseStorehouse",
      "WarehouseArea.WarehouseStorehouse.OrganizationWorkArea",
      "WarehouseArea.WarehouseStorehouse.OrganizationWorkshop",
      "WarehouseArea.WarehouseStorehouse.OrganizationWorkshop.OrganizationParagraph",
      "WarehouseArea.WarehouseStorehouse.OrganizationWorkshop.OrganizationParagraph.OrganizationRailway",
    ],
  })
    .then(res => {
      name_alertWarehousePlatoon.value = res.content.warehouse_platoon.name;
      organizationRailwayUuid_alertEditWarehousePlatoon.value = res.content.warehouse_platoon.warehouse_area.warehouse_storehouse.organization_workshop.organization_paragraph.organization_railway.uuid;
      organizationParagraphUuid_alertEditWarehousePlatoon.value = res.content.warehouse_platoon.warehouse_area.warehouse_storehouse.organization_workshop.organization_paragraph.uuid;
      organizationWorkshopUuid_alertEditWarehousePlatoon.value = res.content.warehouse_platoon.warehouse_area.warehouse_storehouse.organization_workshop.uuid;
      organizationWorkAreaUuid_alertEditWarehousePlatoon.value = res.content.warehouse_platoon.warehouse_area.warehouse_storehouse.organization_work_area.uuid;
      warehouseStorehouseUuid_alertEditWarehousePlatoon.value = res.content.warehouse_platoon.warehouse_area.warehouse_storehouse.uuid;
      warehouseAreaUuid_alertEditWarehousePlatoon.value = res.content.warehouse_platoon.warehouse_area.uuid;
      warehousePlatoonTypeCode_alertEditWarehousePlatoon.value = res.content.warehouse_platoon.type_code;

      alertWarehousePlatoon.value = true;
    })
    .catch(e => errorNotify(e.msg));
};

const fnUpdateWarehouseArea = () => {
  if (!currentUuid.value) return;

  ajaxUpdateWarehousePlatoon(currentUuid.value, {
    name: name_alertWarehousePlatoon.value,
    warehouse_area_uuid: warehouseAreaUuid_alertEditWarehousePlatoon.value,
    type_code: warehousePlatoonTypeCode_alertEditWarehousePlatoon.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnSearch();

      alertWarehousePlatoon.value = false;
    })
    .catch(e => errorNotify(e.msg));
};

const fnDestroyWarehousePlatoon = params => {
  if (!params["uuid"]) return;

  confirmNotify(destroyActions(() => {
    const loading = loading();

    ajaxDestroyWarehousePlatoon(params.uuid)
      .then(() => {
        successNotify("删除成功");
        fnSearch();
      })
      .catch(e => errorNotify(e.msg))
      .finally(loading());
  }));
};

const fnDestroyWarehousePlatoons = () => {
  confirmNotify(destroyActions(() => {
    const loading = loading();

    ajaxDestroyWarehousePlatoons(collect(selected.value).pluck("uuid").all())
      .then(() => {
        successNotify("删除成功");
        fnSearch();
      })
      .catch(e => errorNotify(e.msg))
      .finally(loading());
  }));
};

</script>
