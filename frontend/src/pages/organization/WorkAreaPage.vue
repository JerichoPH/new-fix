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
                <div class="col-3">
                  <q-input outlined clearable lazy-rules v-model="uniqueCode_search" label="代码" :rules="[]"
                    class="q-mb-md" />
                </div>
                <div class="col-3">
                  <q-input outlined clearable lazy-rules v-model="name_search" label="名称" :rules="[]" class="q-mb-md" />
                </div>
                <div class="col-3">
                  <standard-select label-name="所属路局" sechma="search" current-field="organizationRailwayUuid"
                    :data-source="ajaxGetOrganizationRailways" data-source-field="organization_railways"
                    label-field="short_name" />
                </div>
                <div class="col-3">
                  <standard-select label-name="所属站段" sechma="search" current-field="organizationParagraphUuid"
                    :data-source="ajaxGetOrganizationParagraphs" data-source-field="organization_paragraphs"
                    parent-field="organizationRailwayUuid" />
                </div>
              </div>
              <div class="row q-col-gutter-sm">
                <div class="col-3">
                  <standard-select label-name="所属车间" sechma="search" current-field="organizationWorkshopUuid"
                    :data-source="ajaxGetOrganizationWorkshops" data-source-field="organization_workshops"
                    parent-field="organizationParagraphUuid" />
                </div>
                <div class="col-3">
                  <standard-select label-name="工区类型" sechma="search" current-field="organizationWorkAreaTypeCode"
                    :data-source="ajaxGetOrganizationWorkAreaTypeCodesMap" data-source-field="type_codes_map"
                    label-field="text" value-field="code" />
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
          <div class="col"><span :style="{ fontSize: '20px' }">工区列表</span></div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="secondary" outline label="新建工区" icon="add" @click="fnOpenAlertCreateCreateStation" />
              <q-btn color="negative" outline label="删除工区" icon="deleted" @click="fnDestroyCreateStations" />
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
                  <q-th align="left" key="organizationWorkshop"
                    @click="(event) => fnColumnReverseSort(event, props, sortBy)">
                    所属车间
                  </q-th>
                  <q-th align="left" key="typeText"
                    @click="(event) => fnColumnReverseSort(event, props, sortBy)">工区类型</q-th>
                  <q-th align="right"></q-th>
                </q-tr>
              </template>
              <template v-slot:body="props">
                <q-tr :props="props">
                  <q-td><q-checkbox :key="props.row.uuid" :value="props.row.uuid" v-model="props.selected" /></q-td>
                  <q-td>{{ props.row.index }}</q-td>
                  <q-td key="uniqueCode" :props="props">{{ props.row.uniqueCode }}</q-td>
                  <q-td key="name" :props="props">{{ props.row.name }}</q-td>
                  <q-td key="organizationWorkshop" :props="props">
                    <join-string :values="[
                      props.row.organizationRailway.short_name,
                      props.row.organizationParagraph.name,
                      props.row.organizationWorkshop.name,
                    ]" />
                  </q-td>
                  <q-td key="typeText" :props="props">{{ props.row.typeText }}</q-td>
                  <q-td key="operation" :props="props">
                    <q-btn-group>
                      <q-btn @click="fnOpenAlertEditCreateOrganizationWorkArea(props.row.operation)" color="warning"
                        icon="edit">
                        编辑
                      </q-btn>
                      <q-btn @click="fnDestroyCreateOrganizationWorkArea(props.row.operation)" color="negative"
                        icon="delete">
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
  <!-- 新建工区弹窗 -->
  <q-dialog v-model="alertCreateOrganizationWorkArea" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">新建工区</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnStoreOrganizationWorkArea">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="uniqueCode_alertCreateOrganizationWorkArea" label="代码"
                :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertCreateOrganizationWorkArea" label="名称"
                :rules="[]" />
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
              <standard-select label-name="工区类型" sechma="alertCreate" current-field="organizationWorkAreaTypeCode"
                :data-source="ajaxGetOrganizationWorkAreaTypeCodesMap" data-source-field="type_codes_map"
                label-field="text" value-field="code" />
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
  <!-- 编辑工区弹窗 -->
  <q-dialog v-model="alertEditOrganizationWorkArea" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">编辑工区</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnUpdateOrganizationWorkArea">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="uniqueCode_alertEditOrganizationWorkArea" label="名称"
                :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertEditOrganizationWorkArea" label="名称"
                :rules="[]" />
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
              <standard-select label-name="工区类型" sechma="alertEdit" current-field="organizationWorkAreaTypeCode"
                :data-source="ajaxGetOrganizationWorkAreaTypeCodesMap" data-source-field="type_codes_map"
                label-field="text" value-field="code" />
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
import { ref, onMounted, provide } from "vue";
import collect from "collect.js";
import { fnColumnReverseSort } from "src/utils/common";
import {
  ajaxGetOrganizationWorkAreas,
  ajaxGetOrganizationWorkArea,
  ajaxStoreOrganizationWorkArea,
  ajaxUpdateOrganizationWorkArea,
  ajaxDestroyOrganizationWorkArea,
  ajaxDestroyOrganizationWorkAreas,
  ajaxGetOrganizationWorkAreaTypeCodesMap,
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

// 搜索栏数据
const uniqueCode_search = ref("");
const name_search = ref("");
const organizationRailwayUuid_search = ref("");
provide("organizationRailwayUuid_search", organizationRailwayUuid_search);
const organizationParagraphUuid_search = ref("");
provide("organizationParagraphUuid_search", organizationParagraphUuid_search);
const organizationWorkshopUuid_search = ref("");
provide("organizationWorkshopUuid_search", organizationWorkshopUuid_search);
const organizationWorkAreaTypeCode_search = ref("");
provide("organizationWorkAreaTypeCode_search", organizationWorkAreaTypeCode_search);

// 表格数据
const rows = ref([]);
const selected = ref([]);
const sortBy = ref("");

// 新建工区弹窗数据
const alertCreateOrganizationWorkArea = ref(false);
const uniqueCode_alertCreateOrganizationWorkArea = ref("");
const name_alertCreateOrganizationWorkArea = ref("");
const organizationRailwayUuid_alertCreateOrganizationWorkArea = ref("");
provide("organizationRailwayUuid_alertCreate", organizationRailwayUuid_alertCreateOrganizationWorkArea);
const organizationParagraphUuid_alertCreateOrganizationWorkArea = ref("");
provide("organizationParagraphUuid_alertCreate", organizationParagraphUuid_alertCreateOrganizationWorkArea);
const organizationWorkshopUuid_alertCreateOrganizationWorkArea = ref("");
provide("organizationWorkshopUuid_alertCreate", organizationWorkshopUuid_alertCreateOrganizationWorkArea);
const organizationWorkAreaTypeCode_alertCreate = ref("");
provide("organizationWorkAreaTypeCode_alertCreate", organizationWorkAreaTypeCode_alertCreate);

// 编辑工区弹窗数据
const currentUuid = ref("");
const alertEditOrganizationWorkArea = ref(false);
const uniqueCode_alertEditOrganizationWorkArea = ref("");
const name_alertEditOrganizationWorkArea = ref("");
const organizationRailwayUuid_alertEditOrganizationWorkArea = ref("");
provide("organizationRailwayUuid_alertEdit", organizationRailwayUuid_alertEditOrganizationWorkArea);
const organizationParagraphUuid_alertEditOrganizationWorkArea = ref("");
provide("organizationParagraphUuid_alertEdit", organizationParagraphUuid_alertEditOrganizationWorkArea);
const organizationWorkshopUuid_alertEditOrganizationWorkArea = ref("");
provide("organizationWorkshopUuid_alertEdit", organizationWorkshopUuid_alertEditOrganizationWorkArea);
const organizationWorkAreaTypeCode_alertEdit = ref("");
provide("organizationWorkAreaTypeCode_alertEdit", organizationWorkAreaTypeCode_alertEdit);

onMounted(() => fnInit());

const fnInit = () => fnSearch();

const fnResetSearch = () => {
  uniqueCode_search.value = "";
  name_search.value = "";
  organizationRailwayUuid_search.value = "";
  organizationParagraphUuid_search.value = "";
  organizationWorkshopUuid_search.value = "";
  organizationWorkAreaTypeCode_search.value = "";
};

const fnSearch = () => {
  rows.value = [];
  selected.value = [];

  ajaxGetOrganizationWorkAreas({
    "@~[]": ["OrganizationWorkshop", "OrganizationWorkshop.OrganizationParagraph", "OrganizationWorkshop.OrganizationParagraph.OrganizationRailway"],
    unique_code: uniqueCode_search.value,
    name: name_search.value,
    organization_railway_uuid: organizationRailwayUuid_search.value,
    organization_paragraph_uuid: organizationParagraphUuid_search.value,
    organization_workshop_uuid: organizationWorkshopUuid_search.value,
    type_code: organizationWorkAreaTypeCode_search.value,
  })
    .then(res => {
      rows.value = collect(res.content.organization_work_areas)
        .map((organizationWorkArea, idx) => {
          return {
            index: idx + 1,
            uuid: organizationWorkArea.uuid,
            uniqueCode: organizationWorkArea.unique_code,
            name: organizationWorkArea.name,
            organizationRailway: organizationWorkArea.organization_workshop.organization_paragraph.organization_railway,
            organizationParagraph: organizationWorkArea.organization_workshop.organization_paragraph,
            organizationWorkshop: organizationWorkArea.organization_workshop,
            typeText: organizationWorkArea.type_text,
            operation: { uuid: organizationWorkArea.uuid },
          }
        })
        .all();
    })
    .catch(e => errorNotify(e.msg));
};

const fnResetAlertCreateStation = () => {
  uniqueCode_alertCreateOrganizationWorkArea.value = "";
  name_alertCreateOrganizationWorkArea.value = "";
  organizationRailwayUuid_alertCreateOrganizationWorkArea.value = "";
  organizationParagraphUuid_alertCreateOrganizationWorkArea.value = "";
  organizationWorkshopUuid_alertCreateOrganizationWorkArea.value = "";
  organizationWorkAreaTypeCode_alertCreate.value = "";
};

const fnOpenAlertCreateCreateStation = () => {
  alertCreateOrganizationWorkArea.value = true;
};

const fnStoreOrganizationWorkArea = () => {
  const loading = loadingNotify();

  ajaxStoreOrganizationWorkArea({
    unique_code: uniqueCode_alertCreateOrganizationWorkArea.value,
    name: name_alertCreateOrganizationWorkArea.value,
    organization_railway_uuid: organizationRailwayUuid_alertCreateOrganizationWorkArea.value,
    organization_paragraph_uuid: organizationParagraphUuid_alertCreateOrganizationWorkArea.value,
    organization_workshop_uuid: organizationWorkshopUuid_alertCreateOrganizationWorkArea.value,
    type_code: organizationWorkAreaTypeCode_alertCreate.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnResetAlertCreateStation();
      fnSearch();

      alertCreateOrganizationWorkArea.value = false;
    })
    .catch(e => errorNotify(e.msg))
    .finally(loading());
};

const fnOpenAlertEditCreateOrganizationWorkArea = params => {
  if (!params["uuid"]) return;
  currentUuid.value = params.uuid;

  ajaxGetOrganizationWorkArea(currentUuid.value, {
    "@~[]": ["OrganizationWorkshop", "OrganizationWorkshop.OrganizationParagraph", "OrganizationWorkshop.OrganizationParagraph.OrganizationRailway"],
  })
    .then(res => {
      uniqueCode_alertEditOrganizationWorkArea.value = res.content.organization_work_area.unique_code;
      name_alertEditOrganizationWorkArea.value = res.content.organization_work_area.name;
      organizationRailwayUuid_alertEditOrganizationWorkArea.value = res.content.organization_work_area.organization_workshop.organization_paragraph.organization_railway.uuid;
      organizationParagraphUuid_alertEditOrganizationWorkArea.value = res.content.organization_work_area.organization_workshop.organization_paragraph.uuid;
      organizationWorkshopUuid_alertEditOrganizationWorkArea.value = res.content.organization_work_area.organization_workshop.uuid;
      organizationWorkAreaTypeCode_alertEdit.value = res.content.organization_work_area.type_code;

      alertEditOrganizationWorkArea.value = true;
    })
    .catch(e => errorNotify(e.msg));
};

const fnUpdateOrganizationWorkArea = () => {
  if (!currentUuid.value) return;

  const loading = loadingNotify();
  ajaxUpdateOrganizationWorkArea(currentUuid.value, {
    unique_code: uniqueCode_alertEditOrganizationWorkArea.value,
    name: name_alertEditOrganizationWorkArea.value,
    organization_workshop_uuid: organizationWorkshopUuid_alertEditOrganizationWorkArea.value,
    type_code: organizationWorkAreaTypeCode_alertEdit.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnSearch();

      alertEditOrganizationWorkArea.value = false;
    })
    .catch(e => errorNotify(e.msg))
    .finally(loading());
};

const fnDestroyCreateOrganizationWorkArea = params => {
  if (!params["uuid"]) return;

  confirmNotify(
    destroyActions(() => {
      const loading = loadingNotify();

      ajaxDestroyOrganizationWorkArea(params.uuid)
        .then(() => {
          successNotify("删除成功");
          fnSearch();
        })
        .catch(e => errorNotify(e.msg))
        .finally(loading());
    })
  );
};
const fnDestroyCreateStations = () => {
  confirmNotify(
    destroyActions(() => {
      const loading = loadingNotify();

      ajaxDestroyOrganizationWorkAreas(collect(selected.value).pluck("uuid").all())
        .then(() => {
          successNotify("删除成功");
          fnSearch();
        })
        .catch(e => errorNotify(e.msg))
        .finally(loading());
    })
  );
};
</script>
