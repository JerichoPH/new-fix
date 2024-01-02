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
                  <sel-organization-railway_search label-name="所属路局" :ajaxParams="{}" />
                </div>
                <div class="col-3">
                  <sel-organization-paragraph_search label-name="所属站段" :ajaxParams="{}" />
                </div>
              </div>
              <div class="row q-col-gutter-sm">
                <div class="col-3">
                  <sel-organization-workshop_search label-name="所属车间" :ajaxParams="{}" />
                </div>
                <div class="col-3">
                  <sel-organization-work-area-type-code_search label-name="工区类型" :ajaxParams="{}" />
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
              <q-btn color="secondary" label="新建工区" icon="add" @click="fnOpenAlertCreateCreateStation" />
              <q-btn color="negative" label="删除工区" icon="deleted" @click="fnDestroyCreateStations" />
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
                  <q-td key="organizationWorkshop" :props="props">{{ [
                    props.row.organizationRailway.short_name,
                    props.row.organizationParagraph.name,
                    props.row.organizationWorkshop.name,
                  ].join(' - ') }}</q-td>
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
    <q-card :style="{minWidth: '450px'}">
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
              <sel-organization-railway_alert-create label-name="所属路局" :ajaxParams="{}" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-paragraph_alert-create label-name="所属站段" :ajaxParams="{}" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-workshop_alert-create label-name="所属车间" :ajaxParams="{}" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-work-area-type-code_alert-create label-name="工区类型" :ajaxParams="{}" />
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn type="submit" label="关闭" v-close-popup />
          <q-btn type="submit" label="确定" icon="check" color="secondary" />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
  <!-- 编辑工区弹窗 -->
  <q-dialog v-model="alertEditOrganizationWorkArea" no-backdrop-dismiss>
    <q-card :style="{minWidth: '450px'}">
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
              <sel-organization-railway_alert-edit label-name="所属路局" :ajaxParams="{}" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-paragraph_alert-edit label-name="所属站段" :ajaxParams="{}" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-workshop_alert-edit label-name="所属车间" :ajaxParams="{}" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-work-area-type-code_alert-edit label-name="工区类型" :ajaxParams="{}" />
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn type="submit" label="关闭" v-close-popup />
          <q-btn type="submit" label="确定" icon="check" color="secondary" />
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
} from "src/apis/organization";
import {
  loadingNotify,
  successNotify,
  errorNotify,
  confirmNotify,
  destroyActions,
} from "src/utils/notify";
import SelOrganizationRailway_search from "src/components/SelOrganizationRailway_search.vue";
import SelOrganizationParagraph_search from "src/components/SelOrganizationParagraph_search.vue";
import SelOrganizationWorkshop_search from "src/components/SelOrganizationWorkshop_search.vue";
import SelOrganizationWorkAreaTypeCode_search from "src/components/SelOrganizationWorkAreaTypeCode_search.vue";
import SelOrganizationRailway_alertCreate from "src/components/SelOrganizationRailway_alertCreate.vue";
import SelOrganizationParagraph_alertCreate from "src/components/SelOrganizationParagraph_alertCreate.vue";
import SelOrganizationWorkshop_alertCreate from "src/components/SelOrganizationWorkshop_alertCreate.vue";
import SelOrganizationRailway_alertEdit from "src/components/SelOrganizationRailway_alertEdit.vue";
import SelOrganizationParagraph_alertEdit from "src/components/SelOrganizationParagraph_alertEdit.vue";
import SelOrganizationWorkshop_alertEdit from "src/components/SelOrganizationWorkshop_alertEdit.vue";
import SelOrganizationWorkAreaTypeCode_alertCreate from "src/components/SelOrganizationWorkAreaTypeCode_alertCreate.vue";
import SelOrganizationWorkAreaTypeCode_alertEdit from "src/components/SelOrganizationWorkAreaTypeCode_alertEdit.vue";

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
