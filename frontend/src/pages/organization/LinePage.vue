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
                <div class="col-3">
                  <standard-select label-name="所属路局" sechma="search" current-field="organizationRailwayUuid"
                    :data-source="ajaxGetOrganizationRailways" data-source-field="organization_railways" />
                </div>
                <div class="col-3">
                  <standard-select label-name="所属站段" sechma="search" current-field="organizationParagraphUuid"
                    :data-source="ajaxGetOrganizationParagraphs" data-source-field="organization_paragraphs"
                    parent-field="organizationRailwayhUuid" />
                </div>
              </div>
              <div class="row q-col-gutter-sm">
                <div class="col-3">
                  <standard-select label-name="所属车间" sechma="search" current-field="organizationWorkshopUuid"
                    :data-source="ajaxGetOrganizationWorkshops" data-source-field="organization_workshops"
                    parent-field="organizationParagraphUuid" />
                </div>
                <div class="col-3">
                  <sel-organization-station_search label-name="相关站场" :ajax-params="{}" />
                </div>
                <div class="col-3">
                  <sel-organization-crossroad_search label-name="相关道口" :ajax-params="{}" />
                </div>
                <div class="col-3">
                  <sel-organization-center_search label-name="相关中心" :ajax-params="{}" />
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
          <div class="col"><span :style="{ fontSize: '20px' }">线别列表</span></div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="secondary" label="新建线别" icon="add" @click="fnOpenAlertCreateCreateOrganizationLine" />
              <q-btn color="negative" label="删除线别" icon="deleted" @click="fnDestroyCreateOrganizationLines" />
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
                  <q-th align="left" key="uniqueCode"
                    @click="(event) => fnColumnReverseSort(event, props, sortBy)">代码</q-th>
                  <q-th align="left" key="name" @click="(event) => fnColumnReverseSort(event, props, sortBy)">名称</q-th>
                  <q-th align="left" key="organizationParagraph"
                    @click="(event) => fnColumnReverseSort(event, props, sortBy)">所属站段</q-th>
                  <q-th align="left" key="organizationStations"
                    @click="(event) => fnColumnReverseSort(event, props, sortBy)">相关站场</q-th>
                  <q-th align="left" key="organizationStations"
                    @click="(event) => fnColumnReverseSort(event, props, sortBy)">相关道口</q-th>
                  <q-th align="left" key="organizationStations"
                    @click="(event) => fnColumnReverseSort(event, props, sortBy)">相关中心</q-th>
                  <q-th align="right"></q-th>
                </q-tr>
              </template>
              <template v-slot:body="props">
                <q-tr :props="props">
                  <q-td><q-checkbox :key="props.row.uuid" :value="props.row.uuid" v-model="props.selected" /></q-td>
                  <q-td>{{ props.row.index }}</q-td>
                  <q-td key="uniqueCode" :props="props">{{ props.row.uniqueCode }}</q-td>
                  <q-td key="name" :props="props">{{ props.row.name }}</q-td>
                  <q-td key="organizationParagraph" :props="props">
                    <join-string :values="[
                      props.row.organizationRailway.short_name,
                      props.row.organizationParagraph.name
                    ]" />
                  </q-td>
                  <q-td key="organizationStations" :props="props">{{
                    collect(props.row.organizationStations).pluck('name').implode(',') }}</q-td>
                  <q-td key="organizationCrossroads" :props="props">{{
                    collect(props.row.organizationCrossroads).pluck('name').implode(',') }}</q-td>
                  <q-td key="organizationCenters" :props="props">{{
                    collect(props.row.organizationCenters).pluck('name').implode(',') }}</q-td>
                  <q-td key="operation" :props="props">
                    <q-btn-group>
                      <q-btn @click="fnOpenAlertEditCreateOrganizationLine(props.row.operation)" color="warning"
                        icon="edit">
                        编辑
                      </q-btn>
                      <q-btn @click="fnDestroyCreateOrganizationLine(props.row.operation)" color="negative" icon="delete">
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
  <!-- 新建线别弹窗 -->
  <q-dialog v-model="alertCreateOrganizationLine" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">新建线别</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnStoreOrganizationLine">
        <q-card-section class="">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="uniqueCode_alertCreateOrganizationLine" label="代码"
                :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertCreateOrganizationLine" label="名称" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属路局" sechma="alertCreate" current-field="organizationRailwayUuid"
                :data-source="ajaxGetOrganizationRailways" data-source-field="organization_railways" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属站段" sechma="alertCreate" current-field="organizationParagraphUuid"
                :data-source="ajaxGetOrganizationParagraphs" data-source-field="organization_paragraphs"
                parent-field="organizationRailwayhUuid" />
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
              <chk-organization-station_alert-create label-name="相关站场" :ajax-params="{}" />
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
  <!-- 编辑线别弹窗 -->
  <q-dialog v-model="alertEditOrganizationLine" no-backdrop-dismiss>
    <q-card style="width: 1800px">
      <q-card-section>
        <div class="text-h6">编辑线别</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnUpdateOrganizationLine">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="uniqueCode_alertEditOrganizationLine" label="代码"
                :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertEditOrganizationLine" label="名称" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属路局" sechma="alertEdit" current-field="organizationRailwayUuid"
                :data-source="ajaxGetOrganizationRailways" data-source-field="organization_railways" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属站段" sechma="alertEdit" current-field="organizationParagraphUuid"
                :data-source="ajaxGetOrganizationParagraphs" data-source-field="organization_paragraphs"
                parent-field="organizationRailwayhUuid" />
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
              <chk-organization-station_alert-edit label-name="相关站场" :ajax-params="{}" />
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
  ajaxGetOrganizationLines,
  ajaxGetOrganizationLine,
  ajaxStoreOrganizationLine,
  ajaxUpdateOrganizationLine,
  ajaxDestroyOrganizationLine,
  ajaxDestroyOrganizationLines,
  ajaxGetOrganizationRailways,
  ajaxGetOrganizationParagraphs,
  ajaxGetOrganizationWorkshops,
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
import SelOrganizationStation_search from "src/components/SelOrganizationStation_search.vue";
import SelOrganizationCrossroad_search from "src/components/SelOrganizationCrossroad_search.vue";
import SelOrganizationCenter_search from "src/components/SelOrganizationCenter_search.vue";
import ChkOrganizationStation_alertCreate from "src/components/ChkOrganizationStation_alertCreate.vue";
import ChkOrganizationStation_alertEdit from "src/components/ChkOrganizationStation_alertEdit.vue";

// 搜索栏数据
const uniqueCode_search = ref("");
const name_search = ref("");
const organizationRailwayUuid_search = ref("");
provide("organizationRailwayUuid_search", organizationRailwayUuid_search);
const organizationParagraphUuid_search = ref("");
provide("organizationParagraphUuid_search", organizationParagraphUuid_search);
const organizationWorkshopUuid_search = ref("");
provide("organizationWorkshopUuid_search", organizationWorkshopUuid_search);
const organizationStationUuid_search = ref("");
provide("organizationStationUuid_search", organizationStationUuid_search);
const organizationCrossroadUuid_search = ref("");
provide("organizationCrossroadUuid_search", organizationCrossroadUuid_search);
const organizationCenterUuid_search = ref("");
provide("organizationCenterUuid_search", organizationCenterUuid_search);

// 表格数据
const rows = ref([]);
const selected = ref([]);
const sortBy = ref("");

// 新建线别数据
const alertCreateOrganizationLine = ref(false);
const uniqueCode_alertCreateOrganizationLine = ref("");
const name_alertCreateOrganizationLine = ref("");
const organizationRailwayUuid_alertCreateOrganizationLine = ref("");
provide("organizationRailwayUuid_alertCreate", organizationRailwayUuid_alertCreateOrganizationLine);
const organizationParagraphUuid_alertCreateOrganizationLine = ref("");
provide("organizationParagraphUuid_alertCreate", organizationParagraphUuid_alertCreateOrganizationLine);
const organizationWorkshopUuid_alertCreateOrganizationLine = ref("");
provide("organizationWorkshopUuid_alertCreate", organizationWorkshopUuid_alertCreateOrganizationLine);
const checkedOrganizationStationUuids_alertCreateOrganizationLine = ref([]);
provide("checkedOrganizationStationUuids_alertCreate", checkedOrganizationStationUuids_alertCreateOrganizationLine);

// 编辑线别数据
const currentUuid = ref("");
const alertEditOrganizationLine = ref(false);
const uniqueCode_alertEditOrganizationLine = ref("");
const name_alertEditOrganizationLine = ref("");
const organizationRailwayUuid_alertEditOrganizationLine = ref("");
provide("organizationRailwayUuid_alertEdit", organizationRailwayUuid_alertEditOrganizationLine);
const organizationParagraphUuid_alertEditOrganizationLine = ref("");
provide("organizationParagraphUuid_alertEdit", organizationParagraphUuid_alertEditOrganizationLine);
const organizationWorkshopUuid_alertEditOrganizationLine = ref("");
provide("organizationWorkshopUuid_alertEdit", organizationWorkshopUuid_alertEditOrganizationLine);
const checkedOrganizationStationUuids_alertEditOrganizationLine = ref([]);
provide("checkedOrganizationStationUuids_alertEdit", checkedOrganizationStationUuids_alertEditOrganizationLine);

onMounted(() => fnInit());

const fnInit = () => fnSearch();

const fnResetSearch = () => {
  uniqueCode_search.value = "";
  name_search.value = "";
  organizationParagraphUuid_search.value = "";
  organizationWorkshopUuid_search.value = "";
  organizationStationUuid_search.value = "";
  organizationCrossroadUuid_search.value = "";
  organizationCenterUuid_search.value = "";
};

const fnSearch = () => {
  rows.value = [];
  selected.value = [];

  ajaxGetOrganizationLines({
    "@~[]": [
      "OrganizationParagraph",
      "OrganizationParagraph.OrganizationRailway",
      "OrganizationStations",
      "OrganizationStations.OrganizationWorkshop",
      "OrganizationCrossroads",
      "OrganizationCrossroads.OrganizationWorkshop",
      "OrganizationCenters",
      "OrganizationCenters.OrganizationWorkshop",
    ],
    unique_code: uniqueCode_search.value,
    name: name_search.value,
    organization_paragraph_uuid: organizationParagraphUuid_search.value,
    organization_station_uuid: organizationStationUuid_search.value,
    organization_crossroad_uuid: organizationCrossroadUuid_search.value,
    organization_center_uuid: organizationCenterUuid_search.value,
  })
    .then(res => {
      rows.value = collect(res.content.organization_lines)
        .map((organizationLine, idx) => {
          return {
            index: idx + 1,
            uuid: organizationLine.uuid,
            uniqueCode: organizationLine.uniqueCode,
            name: organizationLine.name,
            organizationParagraph: organizationLine.organization_paragraph,
            organizationRailway: organizationLine.organization_paragraph.organization_railway,
            organizationStations: organizationLine.organization_stations,
            organizationCrossroads: organizationLine.organization_crossroads,
            organizationCenters: organizationLine.organization_centers,
            operation: { uuid: organizationLine.uuid },
          }
        })
        .all();
    })
    .catch(e => errorNotify(e.msg));
};

const fnResetAlertCreateOrganizationLine = () => {
  uniqueCode_alertCreateOrganizationLine.value = "";
  name_alertCreateOrganizationLine.value = "";
  organizationRailwayUuid_alertCreateOrganizationLine.value = "";
  organizationParagraphUuid_alertCreateOrganizationLine.value = "";
  organizationWorkshopUuid_alertCreateOrganizationLine.value = "";
  checkedOrganizationStationUuids_alertCreateOrganizationLine.value = "";
};

const fnOpenAlertCreateCreateOrganizationLine = () => {
  alertCreateOrganizationLine.value = true;
};

const fnStoreOrganizationLine = () => {
  const loading = loadingNotify();

  ajaxStoreOrganizationLine({
    unique_code: uniqueCode_alertCreateOrganizationLine.value,
    name: name_alertCreateOrganizationLine.value,
    organization_paragraph_uuid: organizationParagraphUuid_alertCreateOrganizationLine.value,
    organization_station_uuids: checkedOrganizationStationUuids_alertCreateOrganizationLine.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnResetAlertCreateOrganizationLine();
      fnSearch();

      alertCreateOrganizationLine.value = false;
    })
    .catch(e => errorNotify(e.msg))
    .finally(loading());
};

const fnOpenAlertEditCreateOrganizationLine = params => {
  if (!params["uuid"]) return;
  currentUuid.value = params["uuid"];

  ajaxGetOrganizationLine(currentUuid.value, {
    "@~[]": [
      "OrganizationParagraph",
      "OrganizationParagraph.OrganizationRailway",
      "OrganizationStations",
      "OrganizationStations.OrganizationWorkshop",
      "OrganizationCrossroads",
      "OrganizationCrossroads.OrganizationWorkshop",
      "OrganizationCenters",
      "OrganizationCenters.OrganizationWorkshop",
    ]
  })
    .then(res => {
      uniqueCode_alertEditOrganizationLine.value = res.content.organization_line.unique_code;
      name_alertEditOrganizationLine.value = res.content.organization_line.name;
      organizationRailwayUuid_alertEditOrganizationLine.value = res.content.organization_line.organization_paragraph.organization_railway.uuid;
      organizationParagraphUuid_alertEditOrganizationLine.value = res.content.organization_line.organization_paragraph.uuid;
      checkedOrganizationStationUuids_alertEditOrganizationLine.value = collect(res.content.organization_line.organization_stations).pluck('uuid').all();

      alertEditOrganizationLine.value = true;
    })
    .catch(e => errorNotify(e.msg));
};

const fnUpdateOrganizationLine = () => {
  if (!currentUuid.value) return;

  const loading = loadingNotify();

  ajaxUpdateOrganizationLine(currentUuid.value, {
    unique_code: uniqueCode_alertEditOrganizationLine.value,
    name: name_alertEditOrganizationLine.value,
    organization_paragraph_uuid: organizationParagraphUuid_alertEditOrganizationLine.value,
    organization_station_uuids: checkedOrganizationStationUuids_alertEditOrganizationLine.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnSearch();

      alertEditOrganizationLine.value = false;
    })
    .catch(e => errorNotify(e.msg))
    .finally(loading());
};

const fnDestroyCreateOrganizationLines = () => {
  confirmNotify(
    destroyActions(() => {
      const loading = loadingNotify();

      ajaxDestroyOrganizationLines(collect(selected.value).pluck("uuid").all())
        .then(() => {
          successNotify("删除成功");
          fnSearch();
        })
        .catch(e => errorNotify(e.msg))
        .finally(loading());
    })
  );
};

const fnDestroyCreateOrganizationLine = params => {
  if (!params["uuid"]) return;

  confirmNotify(
    destroyActions(() => {
      const loading = loadingNotify();

      ajaxDestroyOrganizationLine(params.uuid)
        .then(() => {
          successNotify("删除成功");
          fnSearch();
        })
        .catch(e => errorNotify(e.msg))
        .finally(loading());
    })
  )

};
</script>
