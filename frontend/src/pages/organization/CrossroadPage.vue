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
              <div class="row q-pb-sm q-col-gutter-sm">
                <div class="col">
                  <q-input outlined clearable lazy-rules v-model="uniqueCode_search" label="代码" :rules="[]"
                    class="q-mb-md" />
                </div>
                <div class="col">
                  <q-input outlined clearable lazy-rules v-model="name_search" label="名称" :rules="[]" class="q-mb-md" />
                </div>
                <div class="col">
                  <sel-organization-railway_search labelName="所属路局" :ajaxParams="{}" />
                </div>
                <div class="col">
                  <sel-organization-paragraph_search labelName="所属站段" :ajaxParams="{}" />
                </div>
                <div class="col">
                  <sel-organization-workshop_search labelName="所属车间" :ajaxParams="{}" />
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
          <div class="col"><span :style="{ fontSize: '20px' }">道口列表</span></div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="secondary" label="新建道口" icon="add" @click="fnOpenAlertCreateCreateStation" />
              <q-btn color="negative" label="删除道口" icon="deleted" @click="fnDestroyCreateStations" />
            </q-btn-group>
          </div>
        </div>
        <div class="row q-mt-md">
          <div class="col">
            <q-table flat bordered title="道口列表" :rows="rows" row-key="uuid" :pagination="{ rowsPerPage: 200 }"
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
                      <q-btn @click="fnOpenAlertEditCreateOrganizationCrossroad(props.row.operation)" color="warning"
                        icon="edit">
                        编辑
                      </q-btn>
                      <q-btn @click="fnDestroyCreateOrganizationCrossroad(props.row.operation)" color="negative"
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
  <!-- 新建道口弹窗 -->
  <q-dialog v-model="alertCreateOrganizationCrossroad">
    <q-card style="width: 800px">
      <q-card-section>
        <div class="text-h6">新建道口</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnStoreOrganizationCrossroad">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="uniqueCode_alertCreateOrganizationCrossroad" label="代码"
                :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertCreateOrganizationCrossroad" label="名称"
                :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-railway_alert-create labelName="所属路局" :ajaxParams="{}" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-paragraph_alert-create labelName="所属站段" :ajaxParams="{}" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-workshop_alert-create labelName="所属车间" :ajaxParams="{}" />
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn type="submit" label="确定" icon="check" color="secondary" v-close-popup />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
  <!-- 编辑道口弹窗 -->
  <q-dialog v-model="alertEditOrganizationCrossroad">
    <q-card style="width: 800px">
      <q-card-section>
        <div class="text-h6">编辑道口</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnUpdateOrganizationCrossroad">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="uniqueCode_alertEditOrganizationCrossroad" label="名称"
                :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertEditOrganizationCrossroad" label="名称" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-railway_alert-edit labelName="所属路局" :ajaxParams="{}" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-paragraph_alert-edit labelName="所属站段" :ajaxParams="{}" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-workshop_alert-edit labelName="所属车间" :ajaxParams="{}" />
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
  ajaxGetOrganizationCrossroads,
  ajaxGetOrganizationCrossroad,
  ajaxStoreOrganizationCrossroad,
  ajaxUpdateOrganizationCrossroad,
  ajaxDestroyOrganizationCrossroad,
  ajaxDestroyOrganizationCrossroads,
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

// 新建道口弹窗数据
const alertCreateOrganizationCrossroad = ref(false);
const uniqueCode_alertCreateOrganizationCrossroad = ref("");
const name_alertCreateOrganizationCrossroad = ref("");
const organizationRailwayUuid_alertCreateOrganizationCrossroad = ref("");
provide("organizationRailwayUuid_alertCreate", organizationRailwayUuid_alertCreateOrganizationCrossroad);
const organizationParagraphUuid_alertCreateOrganizationCrossroad = ref("");
provide("organizationParagraphUuid_alertCreate", organizationParagraphUuid_alertCreateOrganizationCrossroad);
const organizationWorkshopUuid_alertCreateOrganizationCrossroad = ref("");
provide("organizationWorkshopUuid_alertCreate", organizationWorkshopUuid_alertCreateOrganizationCrossroad);

// 编辑道口弹窗数据
const currentUuid = ref("");
const alertEditOrganizationCrossroad = ref(false);
const uniqueCode_alertEditOrganizationCrossroad = ref("");
const name_alertEditOrganizationCrossroad = ref("");
const organizationRailwayUuid_alertEditOrganizationCrossroad = ref("");
provide("organizationRailwayUuid_alertEdit", organizationRailwayUuid_alertEditOrganizationCrossroad);
const organizationParagraphUuid_alertEditOrganizationCrossroad = ref("");
provide("organizationParagraphUuid_alertEdit", organizationParagraphUuid_alertEditOrganizationCrossroad);
const organizationWorkshopUuid_alertEditOrganizationCrossroad = ref("");
provide("organizationWorkshopUuid_alertEdit", organizationWorkshopUuid_alertEditOrganizationCrossroad);

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
  ajaxGetOrganizationCrossroads({
    ":~[]": ["OrganizationWorkshop", "OrganizationWorkshop.OrganizationParagraph", "OrganizationWorkshop.OrganizationParagraph.OrganizationRailway"],
    unique_code: uniqueCode_search.value,
    name: name_search.value,
    organization_railway_uuid: organizationRailwayUuid_search.value,
    organization_paragraph_uuid: organizationParagraphUuid_search.value,
    organization_workshop_uuid: organizationWorkshopUuid_search.value,
  })
    .then(res => {
      rows.value = collect(res.content.organization_crossroads)
        .map((organizationCrossroad, idx) => {
          return {
            index: idx + 1,
            uuid: organizationCrossroad.uuid,
            uniqueCode: organizationCrossroad.unique_code,
            name: organizationCrossroad.name,
            organizationRailway: organizationCrossroad.organization_workshop.organization_paragraph.organization_railway,
            organizationParagraph: organizationCrossroad.organization_workshop.organization_paragraph,
            organizationWorkshop: organizationCrossroad.organization_workshop,
            operation: { uuid: organizationCrossroad.uuid },
          }
        })
        .all();
    })
    .catch(e => errorNotify(e.msg));
};

const fnResetAlertCreateStation = () => {
  uniqueCode_alertCreateOrganizationCrossroad.value = "";
  name_alertCreateOrganizationCrossroad.value = "";
  organizationRailwayUuid_alertCreateOrganizationCrossroad.value = "";
  organizationParagraphUuid_alertCreateOrganizationCrossroad.value = ""; n
  organizationWorkshopUuid_alertCreateOrganizationCrossroad.value = "";
};

const fnOpenAlertCreateCreateStation = () => {
  alertCreateOrganizationCrossroad.value = true;
};

const fnStoreOrganizationCrossroad = () => {
  const loading = loadingNotify();

  ajaxStoreOrganizationCrossroad({
    unique_code: uniqueCode_alertCreateOrganizationCrossroad.value,
    name: name_alertCreateOrganizationCrossroad.value,
    organization_railway_uuid: organizationRailwayUuid_alertCreateOrganizationCrossroad.value,
    organization_paragraph_uuid: organizationParagraphUuid_alertCreateOrganizationCrossroad.value,
    organization_workshop_uuid: organizationWorkshopUuid_alertCreateOrganizationCrossroad.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnResetAlertCreateStation();
      fnSearch();
    })
    .catch(e => errorNotify(e.msg))
    .finally(() => loading());
};

const fnOpenAlertEditCreateOrganizationCrossroad = params => {
  if (!params["uuid"]) return;
  currentUuid.value = params.uuid;

  ajaxGetOrganizationCrossroad(currentUuid.value, {
    ":~[]": ["OrganizationWorkshop", "OrganizationWorkshop.OrganizationParagraph", "OrganizationWorkshop.OrganizationParagraph.OrganizationRailway"],
  })
    .then(res => {
      uniqueCode_alertEditOrganizationCrossroad.value = res.content.organization_station.unique_code;
      name_alertEditOrganizationCrossroad.value = res.content.organization_station.name;
      organizationRailwayUuid_alertEditOrganizationCrossroad.value = res.content.organization_station.organization_workshop.organization_paragraph.organization_railway.uuid;
      organizationParagraphUuid_alertEditOrganizationCrossroad.value = res.content.organization_station.organization_workshop.organization_paragraph.uuid;
      organizationWorkshopUuid_alertEditOrganizationCrossroad.value = res.content.organization_station.organization_workshop.uuid;

      alertEditOrganizationCrossroad.value = true;
    })
    .catch(e => errorNotify(e.msg));
};

const fnUpdateOrganizationCrossroad = () => {
  if (!currentUuid.value) return;

  const loading = loadingNotify();
  ajaxUpdateOrganizationCrossroad(currentUuid.value, {
    unique_code: uniqueCode_alertEditOrganizationCrossroad.value,
    name: name_alertEditOrganizationCrossroad.value,
    organization_workshop_uuid: organizationWorkshopUuid_alertEditOrganizationCrossroad.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnSearch();
    })
    .catch(e => errorNotify(e.msg))
    .finally(() => loading());
};

const fnDestroyCreateOrganizationCrossroad = params => {
  if (!params["uuid"]) return;

  actionNotify(
    destroyActions(() => {
      const loading = loadingNotify();

      ajaxDestroyOrganizationCrossroad(params.uuid)
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

      ajaxDestroyOrganizationCrossroads(collect(selected.value).pluck("uuid").all())
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
