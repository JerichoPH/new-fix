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
                  <q-input outlined clearable lazy-rules v-model="uniqueCode_search" label="代码" :rules="[]"
                    class="q-mb-md" />
                </div>
                <div class="col">
                  <q-input outlined clearable lazy-rules v-model="name_search" label="名称" :rules="[]" class="q-mb-md" />
                </div>
                <div class="col">
                  <sel-organization-railway_search label-name="所属路局" :ajaxParams="{}" />
                </div>
                <div class="col">
                  <sel-organization-paragraph_search label-name="所属站段" :ajaxParams="{}" />
                </div>
                <div class="col">
                  <sel-organization-workshop_search label-name="所属车间" :ajaxParams="{}" />
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
          <div class="col"><span :style="{ fontSize: '20px' }">中心列表</span></div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="secondary" label="新建中心" icon="add" @click="fnOpenAlertCreateCreateStation" />
              <q-btn color="negative" label="删除中心" icon="deleted" @click="fnDestroyCreateStations" />
            </q-btn-group>
          </div>
        </div>
        <div class="row q-mt-md">
          <div class="col">
            <q-table flat bordered title="中心列表" :rows="rows" row-key="uuid" :pagination="{ rowsPerPage: 200 }"
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
                  <q-td key="operation" :props="props">
                    <q-btn-group>
                      <q-btn @click="fnOpenAlertEditCreateOrganizationCenter(props.row.operation)" color="warning"
                        icon="edit">
                        编辑
                      </q-btn>
                      <q-btn @click="fnDestroyCreateOrganizationCenter(props.row.operation)" color="negative"
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
  <!-- 新建中心弹窗 -->
  <q-dialog v-model="alertCreateOrganizationCenter">
    <q-card style="width: 800px">
      <q-card-section>
        <div class="text-h6">新建中心</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnStoreOrganizationCenter">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="uniqueCode_alertCreateOrganizationCenter" label="代码"
                :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertCreateOrganizationCenter" label="名称"
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
        </q-card-section>
        <q-card-actions align="right">
          <q-btn type="submit" label="确定" icon="check" color="secondary" v-close-popup />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
  <!-- 编辑中心弹窗 -->
  <q-dialog v-model="alertEditOrganizationCenter">
    <q-card style="width: 800px">
      <q-card-section>
        <div class="text-h6">编辑中心</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnUpdateOrganizationCenter">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="uniqueCode_alertEditOrganizationCenter" label="名称"
                :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertEditOrganizationCenter" label="名称" :rules="[]" />
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
        </q-card-section>
        <q-card-actions align="right">
          <q-btn type="submit" label="确定" icon="check" color="secondary" v-close-popup />
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
  ajaxGetOrganizationCenters,
  ajaxGetOrganizationCenter,
  ajaxStoreOrganizationCenter,
  ajaxUpdateOrganizationCenter,
  ajaxDestroyOrganizationCenter,
  ajaxDestroyOrganizationCenters,
} from "src/apis/organization";
import {
  loadingNotify,
  successNotify,
  errorNotify,
  actionNotify,
  destroyActions,
} from "src/utils/notify";
import SelOrganizationRailway_search from "src/components/SelOrganizationRailway_search.vue";
import SelOrganizationParagraph_search from "src/components/SelOrganizationParagraph_search.vue";
import SelOrganizationWorkshop_search from "src/components/SelOrganizationWorkshop_search.vue";
import SelOrganizationRailway_alertCreate from "src/components/SelOrganizationRailway_alertCreate.vue";
import SelOrganizationParagraph_alertCreate from "src/components/SelOrganizationParagraph_alertCreate.vue";
import SelOrganizationWorkshop_alertCreate from "src/components/SelOrganizationWorkshop_alertCreate.vue";
import SelOrganizationRailway_alertEdit from "src/components/SelOrganizationRailway_alertEdit.vue";
import SelOrganizationParagraph_alertEdit from "src/components/SelOrganizationParagraph_alertEdit.vue";
import SelOrganizationWorkshop_alertEdit from "src/components/SelOrganizationWorkshop_alertEdit.vue";

// 搜索栏数据
const uniqueCode_search = ref("");
const name_search = ref("");
const organizationRailwayUuid_search = ref("");
provide("organizationRailwayUuid_search", organizationRailwayUuid_search);
const organizationParagraphUuid_search = ref("");
provide("organizationParagraphUuid_search", organizationParagraphUuid_search);
const organizationWorkshopUuid_search = ref("");
provide("organizationWorkshopUuid_search", organizationWorkshopUuid_search);

// 表格数据
const rows = ref([]);
const selected = ref([]);
const sortBy = ref("");

// 新建中心弹窗数据
const alertCreateOrganizationCenter = ref(false);
const uniqueCode_alertCreateOrganizationCenter = ref("");
const name_alertCreateOrganizationCenter = ref("");
const organizationRailwayUuid_alertCreateOrganizationCenter = ref("");
provide("organizationRailwayUuid_alertCreate", organizationRailwayUuid_alertCreateOrganizationCenter);
const organizationParagraphUuid_alertCreateOrganizationCenter = ref("");
provide("organizationParagraphUuid_alertCreate", organizationParagraphUuid_alertCreateOrganizationCenter);
const organizationWorkshopUuid_alertCreateOrganizationCenter = ref("");
provide("organizationWorkshopUuid_alertCreate", organizationWorkshopUuid_alertCreateOrganizationCenter);

// 编辑中心弹窗数据
const currentUuid = ref("");
const alertEditOrganizationCenter = ref(false);
const uniqueCode_alertEditOrganizationCenter = ref("");
const name_alertEditOrganizationCenter = ref("");
const organizationRailwayUuid_alertEditOrganizationCenter = ref("");
provide("organizationRailwayUuid_alertEdit", organizationRailwayUuid_alertEditOrganizationCenter);
const organizationParagraphUuid_alertEditOrganizationCenter = ref("");
provide("organizationParagraphUuid_alertEdit", organizationParagraphUuid_alertEditOrganizationCenter);
const organizationWorkshopUuid_alertEditOrganizationCenter = ref("");
provide("organizationWorkshopUuid_alertEdit", organizationWorkshopUuid_alertEditOrganizationCenter);

onMounted(() => fnInit());

const fnInit = () => fnSearch();

const fnResetSearch = () => {
  uniqueCode_search.value = "";
  name_search.value = "";
  organizationRailwayUuid_search.value = "";
  organizationParagraphUuid_search.value = "";
  organizationWorkshopUuid_search.value = "";
};

const fnSearch = () => {
  ajaxGetOrganizationCenters({
    "@~[]": ["OrganizationWorkshop", "OrganizationWorkshop.OrganizationParagraph", "OrganizationWorkshop.OrganizationParagraph.OrganizationRailway"],
    unique_code: uniqueCode_search.value,
    name: name_search.value,
    organization_railway_uuid: organizationRailwayUuid_search.value,
    organization_paragraph_uuid: organizationParagraphUuid_search.value,
    organization_workshop_uuid: organizationWorkshopUuid_search.value,
  })
    .then(res => {
      rows.value = collect(res.content.organization_centers)
        .map((organizationCenter, idx) => {
          return {
            index: idx + 1,
            uuid: organizationCenter.uuid,
            uniqueCode: organizationCenter.unique_code,
            name: organizationCenter.name,
            organizationRailway: organizationCenter.organization_workshop.organization_paragraph.organization_railway,
            organizationParagraph: organizationCenter.organization_workshop.organization_paragraph,
            organizationWorkshop: organizationCenter.organization_workshop,
            operation: { uuid: organizationCenter.uuid },
          }
        })
        .all();
    })
    .catch(e => errorNotify(e.msg));
};

const fnResetAlertCreateStation = () => {
  uniqueCode_alertCreateOrganizationCenter.value = "";
  name_alertCreateOrganizationCenter.value = "";
  organizationRailwayUuid_alertCreateOrganizationCenter.value = "";
  organizationParagraphUuid_alertCreateOrganizationCenter.value = "";
  organizationWorkshopUuid_alertCreateOrganizationCenter.value = "";
};

const fnOpenAlertCreateCreateStation = () => {
  alertCreateOrganizationCenter.value = true;
};

const fnStoreOrganizationCenter = () => {
  const loading = loadingNotify();

  ajaxStoreOrganizationCenter({
    unique_code: uniqueCode_alertCreateOrganizationCenter.value,
    name: name_alertCreateOrganizationCenter.value,
    organization_railway_uuid: organizationRailwayUuid_alertCreateOrganizationCenter.value,
    organization_paragraph_uuid: organizationParagraphUuid_alertCreateOrganizationCenter.value,
    organization_workshop_uuid: organizationWorkshopUuid_alertCreateOrganizationCenter.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnResetAlertCreateStation();
      fnSearch();
    })
    .catch(e => errorNotify(e.msg))
    .finally(() => loading());
};

const fnOpenAlertEditCreateOrganizationCenter = params => {
  if (!params["uuid"]) return;
  currentUuid.value = params.uuid;

  ajaxGetOrganizationCenter(currentUuid.value, {
    "@~[]": ["OrganizationWorkshop", "OrganizationWorkshop.OrganizationParagraph", "OrganizationWorkshop.OrganizationParagraph.OrganizationRailway"],
  })
    .then(res => {
      uniqueCode_alertEditOrganizationCenter.value = res.content.organization_center.unique_code;
      name_alertEditOrganizationCenter.value = res.content.organization_center.name;
      organizationRailwayUuid_alertEditOrganizationCenter.value = res.content.organization_center.organization_workshop.organization_paragraph.organization_railway.uuid;
      organizationParagraphUuid_alertEditOrganizationCenter.value = res.content.organization_center.organization_workshop.organization_paragraph.uuid;
      organizationWorkshopUuid_alertEditOrganizationCenter.value = res.content.organization_center.organization_workshop.uuid;

      alertEditOrganizationCenter.value = true;
    })
    .catch(e => errorNotify(e.msg));
};

const fnUpdateOrganizationCenter = () => {
  if (!currentUuid.value) return;

  const loading = loadingNotify();
  ajaxUpdateOrganizationCenter(currentUuid.value, {
    unique_code: uniqueCode_alertEditOrganizationCenter.value,
    name: name_alertEditOrganizationCenter.value,
    organization_workshop_uuid: organizationWorkshopUuid_alertEditOrganizationCenter.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnSearch();
    })
    .catch(e => errorNotify(e.msg))
    .finally(() => loading());
};

const fnDestroyCreateOrganizationCenter = params => {
  if (!params["uuid"]) return;

  actionNotify(
    destroyActions(() => {
      const loading = loadingNotify();

      ajaxDestroyOrganizationCenter(params.uuid)
        .then(() => {
          successNotify("删除成功");
          fnSearch();
        })
        .catch(e => errorNotify(e.msg))
        .finally(() => loading());
    })
  );
};
const fnDestroyCreateStations = () => {
  actionNotify(
    destroyActions(() => {
      const loading = loadingNotify();

      ajaxDestroyOrganizationCenters(collect(selected.value).pluck("uuid").all())
        .then(() => {
          successNotify("删除成功");
          fnSearch();
        })
        .catch(e => errorNotify(e.msg))
        .finally(() => loading());
    })
  );
};
</script>
