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
                  <q-input outlined clearable lazy-rules v-model="name_search" label="名称" :rules="[]" class="q-mb-md" />
                </div>
                <div class="col">
                  <standard-select label-name="所属路局" sechma="search" current-field="organizationRailwayUuid"
                    :data-source="ajaxGetOrganizationRailways" data-source-field="organization_railways"  label-field="short_name" />
                </div>
                <div class="col">
                  <standard-select label-name="所属站段" sechma="search" current-field="organizationParagraphUuid"
                    :data-source="ajaxGetOrganizationParagraphs" data-source-field="organization_paragraphs"
                    parent-field="organizationRailwayUuid" />
                </div>
                <div class="col">
                  <standard-select label-name="所属车间" sechma="search" current-field="organizationWorkshopUuid"
                    :data-source="ajaxGetOrganizationWorkshops" data-source-field="organization_workshops"
                    parent-field="organizationParagraphUuid" />
                </div>
                <div class="col">
                  <standard-select label-name="所属工区" sechma="search" current-field="organizationWorkAreaUuid"
                    :data-source="ajaxGetOrganizationWorkAreas" data-source-field="organization_work_areas"
                    parent-field="organizationWorkshopUuid" />
                </div>
              </div>
              <div class="row q-col-gutter-sm">
                <div class="col">
                  <standard-select label-name="所属仓库-库房" sechma="search" current-field="warehouseStorehouseUuid"
                    :data-source="ajaxGetWarehouseStorehouses" data-source-field="warehouse_storehouses" />
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
              <q-btn color="secondary" outline label="新建仓库-库区" icon="add" @click="fnOpenAlertCreateWarehouseArea" />
              <q-btn color="negative" outline label="删除仓库-库区" icon="deleted" @click="fnDestroyWarehouseAreas" />
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
                  <q-td key="warehouseStorehouse" :props="props">
                    <join-string :values="[
                      props.row.organizationWorkshop.name,
                      props.row.organizationWorkArea.name,
                      props.row.warehouseStorehouse.name,
                    ]" />
                  </q-td>
                  <q-td key="operation" :props="props">
                    <q-btn-group>
                      <q-btn @click="fnOpenAlertEditWarehouseArea(props.row.operation)" color="warning" icon="edit" outline>
                        编辑
                      </q-btn>
                      <q-btn @click="fnDestroyWarehouseArea(props.row.operation)" color="negative" icon="delete" outline>
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
  <!-- 新建库房-库区弹窗 -->
  <q-dialog v-model="alertCreateWarehouseArea" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">新建库房-库区</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnStoreWarehouseArea">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertCreateWarehouseArea" label="名称" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属路局" sechma="alertCreate" current-field="organizationRailwayUuid"
                :data-source="ajaxGetOrganizationRailways" data-source-field="organization_railways"
label-field="short_name" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属站段" sechma="alertCreate" current-field="organizationParagraphUuid"
                :data-source="ajaxGetOrganizationParagraphs" data-source-field="organization_paragraphs"
                parent-field="organizationRailwayUuid" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属车间" sechma="alertCreate" current-field="organizationWorkshopUuid"
                :data-source="ajaxGetOrganizationWorkshops" data-source-field="organization_workshops"
                parent-field="organizationParagraphUuid" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属工区" sechma="alertCreate" current-field="organizationWorkAreaUuid"
                :data-source="ajaxGetOrganizationWorkAreas" data-source-field="organization_work_areas"
                parent-field="organizationWorkshopUuid" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属仓库-库房" sechma="alertCreate" current-field="warehouseStorehouseUuid"
                :data-source="ajaxGetWarehouseStorehouses" data-source-field="warehouse_storehouses" />
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
  <!-- 编辑仓库-库区弹窗-->
  <q-dialog v-model="alertEditWarehouseArea" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">编辑仓库-库区</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnUpdateWarehouseArea">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertEditWarehouseArea" label="名称" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属路局" sechma="alertEdit" current-field="organizationRailwayUuid"
                :data-source="ajaxGetOrganizationRailways" data-source-field="organization_railways"
label-field="short_name" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属站段" sechma="alertEdit" current-field="organizationParagraphUuid"
                :data-source="ajaxGetOrganizationParagraphs" data-source-field="organization_paragraphs"
                parent-field="organizationRailwayUuid" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属车间" sechma="alertEdit" current-field="organizationWorkshopUuid"
                :data-source="ajaxGetOrganizationWorkshops" data-source-field="organization_workshops"
                parent-field="organizationParagraphUuid" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属工区" sechma="alertEdit" current-field="organizationWorkAreaUuid"
                :data-source="ajaxGetOrganizationWorkAreas" data-source-field="organization_work_areas"
                parent-field="organizationWorkshopUuid" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属仓库-库房" sechma="alertEdit" current-field="warehouseStorehouseUuid"
                :data-source="ajaxGetWarehouseStorehouses" data-source-field="warehouse_storehouses" />
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
  ajaxGetWarehouseStorehouses,
} from "src/apis/warehouse";
import {
  ajaxGetOrganizationRailways,
  ajaxGetOrganizationParagraphs,
  ajaxGetOrganizationWorkshops,
  ajaxGetOrganizationWorkAreas,
} from "src/apis/organization";
import {
  loadingNotify,
  successNotify,
  errorNotify,
  confirmNotify,
  destroyActions,
} from "src/utils/notify";
import JoinString from "src/components/JoinString.vue";
import StandardSelect from "src/components/StandardSelect.vue";

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

// 新建库房-库区弹窗数据
const alertCreateWarehouseArea = ref(false);
const name_alertCreateWarehouseArea = ref("");
const organizationRailwayUuid_alertCreateWarehouseArea = ref("");
provide("organizationRailwayUuid_alertCreate", organizationRailwayUuid_alertCreateWarehouseArea);
const organizationParagraphUuid_alertCreateWarehouseArea = ref("");
provide("organizationParagraphUuid_alertCreate", organizationParagraphUuid_alertCreateWarehouseArea);
const organizationWorkshopUuid_alertCreateWarehouseArea = ref("");
provide("organizationWorkshopUuid_alertCreate", organizationWorkshopUuid_alertCreateWarehouseArea);
const organizationWorkAreaUuid_alertCreateWarehouseArea = ref("");
provide("organizationWorkAreaUuid_alertCreate", organizationWorkAreaUuid_alertCreateWarehouseArea);
const warehouseStorehouseUuid_alertCreateWarehouseArea = ref("");
provide("warehouseStorehouseUuid_alertCreate", warehouseStorehouseUuid_alertCreateWarehouseArea);

// 编辑库房-库区弹窗数据
const alertEditWarehouseArea = ref(false);
const currentUuid = ref("");
const name_alertEditWarehouseArea = ref("");
const organizationRailwayUuid_alertEditWarehouseArea = ref("");
provide("organizationRailwayUuid_alertEdit", organizationRailwayUuid_alertEditWarehouseArea);
const organizationParagraphUuid_alertEditWarehouseArea = ref("");
provide("organizationParagraphUuid_alertEdit", organizationParagraphUuid_alertEditWarehouseArea);
const organizationWorkshopUuid_alertEditWarehouseArea = ref("");
provide("organizationWorkshopUuid_alertEdit", organizationWorkshopUuid_alertEditWarehouseArea);
const organizationWorkAreaUuid_alertEditWarehouseArea = ref("");
provide("organizationWorkAreaUuid_alertEdit", organizationWorkAreaUuid_alertEditWarehouseArea);
const warehouseStorehouseUuid_alertEditWarehouseArea = ref("");
provide("warehouseStorehouseUuid_alertEdit", warehouseStorehouseUuid_alertEditWarehouseArea);

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
  rows.value = [];
  selected.value = [];

  ajaxGetWarehouseAreas({
    "@~[]": [
      "WarehouseStorehouse",
      "WarehouseStorehouse.OrganizationWorkshop",
      "WarehouseStorehouse.OrganizationWorkArea",
      "WarehouseStorehouse.OrganizationWorkshop.OrganizationParagraph",
      "WarehouseStorehouse.OrganizationWorkshop.OrganizationParagraph.OrganizationRailway",
    ],
    organization_railway_uuid: organizationRailwayUuid_search.value,
    organization_paragraph_uuid: organizationParagraphUuid_search.value,
    organization_workshop_uuid: organizationWorkshopUuid_search.value,
    organization_work_area_uuid: organizationWorkAreaUuid_search.value,
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
    .catch(e => errorNotify(e.msg));
};

const fnResetAlertCreateWarehouseArea = () => {
  name_alertCreateWarehouseArea.value = "";
  organizationRailwayUuid_alertCreateWarehouseArea.value = "";
  organizationParagraphUuid_alertCreateWarehouseArea.value = "";
  organizationWorkshopUuid_alertCreateWarehouseArea.value = "";
  organizationWorkAreaUuid_alertCreateWarehouseArea.value = "";
  warehouseStorehouseUuid_alertCreateWarehouseArea.value = "";
};

const fnOpenAlertCreateWarehouseArea = () => { alertCreateWarehouseArea.value = true; };

const fnStoreWarehouseArea = () => {
  const loading = loadingNotify();

  ajaxStoreWarehouseArea({
    name: name_alertCreateWarehouseArea.value,
    warehouse_storehouse_uuid: warehouseStorehouseUuid_alertCreateWarehouseArea.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnSearch();
      fnResetAlertCreateWarehouseArea();

      alertCreateWarehouseArea.value = false;
    })
    .catch(e => errorNotify(e.msg))
    .finally(loading());
};

const fnOpenAlertEditWarehouseArea = (params) => {
  if (!params["uuid"]) return;
  currentUuid.value = params.uuid;

  ajaxGetWarehouseArea(currentUuid.value, {
    "@~[]": [
      "WarehouseStorehouse",
      "WarehouseStorehouse.OrganizationWorkshop",
      "WarehouseStorehouse.OrganizationWorkArea",
      "WarehouseStorehouse.OrganizationWorkshop.OrganizationParagraph",
      "WarehouseStorehouse.OrganizationWorkshop.OrganizationParagraph.OrganizationRailway",
    ],
  })
    .then(res => {
      name_alertEditWarehouseArea.value = res.content.warehouse_area.name;
      organizationRailwayUuid_alertEditWarehouseArea.value = res.content.warehouse_area.warehouse_storehouse.organization_workshop.organization_paragraph.organization_railway.uuid;
      organizationParagraphUuid_alertEditWarehouseArea.value = res.content.warehouse_area.warehouse_storehouse.organization_workshop.organization_paragraph.uuid;
      organizationWorkshopUuid_alertEditWarehouseArea.value = res.content.warehouse_area.warehouse_storehouse.organization_workshop.uuid;
      organizationWorkAreaUuid_alertEditWarehouseArea.value = res.content.warehouse_area.warehouse_storehouse.organization_work_area.uuid;
      warehouseStorehouseUuid_alertEditWarehouseArea.value = res.content.warehouse_area.warehouse_storehouse.uuid;

      alertEditWarehouseArea.value = true;
    })
    .catch(e => {
      errorNotify(e.msg)
    });
};

const fnUpdateWarehouseArea = () => {
  if (!currentUuid.value) return;

  const loading = loadingNotify();
  ajaxUpdateWarehouseArea(currentUuid.value, {
    name: name_alertEditWarehouseArea.value,
    warehouse_storehouse_uuid: warehouseStorehouseUuid_alertEditWarehouseArea.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnSearch();
      fnResetAlertCreateWarehouseArea();

      alertEditWarehouseArea.value = false;
    })
    .catch(e => errorNotify(e.msg))
    .finally(loading());
};

const fnDestroyWarehouseArea = (params) => {
  if (!params["uuid"]) return;

  confirmNotify(destroyActions(() => {
    const loading = loadingNotify();

    ajaxDestroyWarehouseArea(params.uuid)
      .then(() => {
        successNotify("删除成功");
        fnSearch();
      })
      .catch(e => errorNotify(e.msg))
      .finally(loading());
  }));
};

const fnDestroyWarehouseAreas = () => {
  confirmNotify(destroyActions(() => {
    const loading = loadingNotify();

    ajaxDestroyWarehouseAreas(collect(selected.value).pluck("uuid").all())
      .then(() => {
        successNotify("删除成功");
        fnSearch();
      })
      .catch(e => errorNotify(e.msg))
      .finally(loading());
  }));
};
</script>
