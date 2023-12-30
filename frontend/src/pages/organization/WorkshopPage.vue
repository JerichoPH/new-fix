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
                  <sel-organization-railway_search label-name="所属路局" />
                </div>
                <div class="col-3">
                  <sel-organization-paragraph_search label-name="所属站段" />
                </div>
              </div>
              <div class="row q-col-gutter-sm">
                <div class="col-3">
                  <sel-organization-workshop-type-code_alertCreate label-name="车间类型" />
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
                  <q-th align="left" key="typeText" @click="(event) => fnColumnReverseSort(event, props, sortBy)">
                    类型
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
                  <q-td key="typeText" :props="props">{{ props.row.typeText }}</q-td>
                  <q-td key="organizationParagraph" :props="props">
                    {{ [
                      props.row.organizationRailway.short_name,
                      props.row.organizationParagraph.name,
                    ].join(' - ') }}
                  </q-td>
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
  <q-dialog v-model="alertCreateOrganizationWorkshop">
    <q-card style="width: 800px">
      <q-card-section>
        <div class="text-h6">新建车间</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnStoreOrganizationWorkshop">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="uniqueCode_alertCreateOrganizationWorkshop" label="代码"
                :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertCreateOrganizationWorkshop" label="名称"
                :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-railway_alertCreate label-name="所属路局" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-paragraph_alertCreate label-name="所属站段" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-workshop-type-code_search label-name="车间类型" />
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn type="submit" label="确定" icon="check" color="secondary" v-close-popup />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
  <!-- 编辑车间弹窗 -->
  <q-dialog v-model="alertEditOrganizationWorkshop">
    <q-card style="width: 800px">
      <q-card-section>
        <div class="text-h6">编辑车间</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnUpdateOrganizationWorkshop">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="uniqueCode_alertEditOrganizationWorkshop" label="代码"
                :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertEditOrganizationWorkshop" label="名称"
                :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <SelOrganizationRailway_alertEdit label-name="所属路局" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <SelOrganizationParagraph_alertEdit label-name="所属站段" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <organizationWorkshopTypeCode_alertEdit label-name="车间类型" />
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
import { fnColumnReverseSort } from "src/utils/common";
import SelOrganizationRailway_search from "src/components/SelOrganizationRailway_search.vue";
import SelOrganizationParagraph_search from "src/components/SelOrganizationParagraph_search.vue";
import SelOrganizationWorkshopTypeCode_search from "src/components/SelOrganizationWorkshopTypeCode_search.vue";
import SelOrganizationRailway_alertCreate from "src/components/SelOrganizationRailway_alertCreate.vue";
import SelOrganizationParagraph_alertCreate from "src/components/SelOrganizationParagraph_alertCreate.vue";
import SelOrganizationWorkshopTypeCode_alertCreate from "src/components/SelOrganizationWorkshopTypeCode_alertCreate.vue";
import SelOrganizationRailway_alertEdit from "src/components/SelOrganizationRailway_alertEdit.vue";
import SelOrganizationParagraph_alertEdit from "src/components/SelOrganizationParagraph_alertEdit.vue";
import organizationWorkshopTypeCode_alertEdit from "src/components/SelOrganizationWorkshopTypeCode_alertEdit.vue";

// 搜索栏数据
const uniqueCode_search = ref("");
const name_search = ref("");
const organizationRailwayUuid_search = ref("");
provide("organizationRailwayUuid_search", organizationRailwayUuid_search);
const organizationParagraphUuid_search = ref("");
provide("organizationParagraphUuid_search", organizationParagraphUuid_search);
const organizationWorkshopTypeCode_search = ref("");
provide("organizationWorkshopTypeCode_search", organizationWorkshopTypeCode_search);

// 表格数据
const rows = ref([]);
const selected = ref([]);
const sortBy = ref("");

// 新建车间弹窗数据
const alertCreateOrganizationWorkshop = ref(false);
const uniqueCode_alertCreateOrganizationWorkshop = ref("");
const name_alertCreateOrganizationWorkshop = ref("");
const organizationRailwayUuid_alertCreateOrganizationWorkshop = ref("");
provide("organizationRailwayUuid_alertCreate", organizationRailwayUuid_alertCreateOrganizationWorkshop);
const organizationParagraphUuid_alertCreateOrganizationWorkshop = ref("");
provide("organizationParagraphUuid_alertCreate", organizationParagraphUuid_alertCreateOrganizationWorkshop);
const organizationWorkshopTypeCode_alertCreateOrganizationWorkshop = ref("");
provide("organizationWorkshopTypeCode_alertCreate", organizationWorkshopTypeCode_alertCreateOrganizationWorkshop);

// 编辑车间弹窗数据
const currentUuid = ref("");
const alertEditOrganizationWorkshop = ref(false);
const uniqueCode_alertEditOrganizationWorkshop = ref("");
const name_alertEditOrganizationWorkshop = ref("");
const organizationRailwayUuid_alertEditOrganizationWorkshop = ref("");
provide("organizationRailwayUuid_alertEdit", organizationRailwayUuid_alertEditOrganizationWorkshop);
const organizationParagraphUuid_alertEditOrganizationWorkshop = ref("");
provide("organizationParagraphUuid_alertEdit", organizationParagraphUuid_alertEditOrganizationWorkshop);
const organizationWorkshopTypeCode_alertEditOrganizationWorkshop = ref("");
provide("organizationWorkshopTypeCode_alertEdit", organizationWorkshopTypeCode_alertEditOrganizationWorkshop);

onMounted(() => fnInit());

const fnInit = () => fnSearch();

const fnResetSearch = () => {
  uniqueCode_search.value = "";
  name_search.value = "";
  organizationRailwayUuid_search.value = "";
  organizationParagraphUuid_search.value = "";
  organizationWorkshopTypeCode_search.value = "";
};

const fnSearch = () => {
  ajaxGetOrganizationWorkshops({
    "@~[]": ["OrganizationParagraph", "OrganizationParagraph.OrganizationRailway"],
    unique_code: uniqueCode_search.value,
    name: name_search.value,
    organization_railway_uuid: organizationRailwayUuid_search.value,
    organization_paragraph_uuid: organizationParagraphUuid_search.value,
    type_code:organizationWorkshopTypeCode_search.value,
  })
    .then(res => {
      rows.value = collect(res.content.organization_workshops)
        .map((organizationWorkshop, idx) => {
          return {
            index: idx + 1,
            uuid: organizationWorkshop.uuid,
            uniqueCode: organizationWorkshop.unique_code,
            name: organizationWorkshop.name,
            typeText: organizationWorkshop.type_text,
            organizationParagraph: organizationWorkshop.organization_paragraph,
            organizationRailway: organizationWorkshop.organization_paragraph.organization_railway,
            operation: { uuid: organizationWorkshop.uuid },
          }
        })
        .all()
    })
    .catch(e => errorNotify(e.msg));
};

const fnResetAlertCreateOrganizationWorkshop = () => {
  uniqueCode_alertCreateOrganizationWorkshop.value = "";
  name_alertCreateOrganizationWorkshop.value = "";
  organizationRailwayUuid_alertCreateOrganizationWorkshop.value = "";
  organizationParagraphUuid_alertCreateOrganizationWorkshop.value = "";
};

const fnOpenAlertCreateOrganizationWorkshop = () => {
  alertCreateOrganizationWorkshop.value = true;
};

const fnStoreOrganizationWorkshop = () => {
  ajaxStoreOrganizationWorkshop({
    unique_code: uniqueCode_alertCreateOrganizationWorkshop.value,
    name: name_alertCreateOrganizationWorkshop.value,
    organization_paragraph_uuid: organizationParagraphUuid_alertCreateOrganizationWorkshop.value,
    type_code: organizationWorkshopTypeCode_alertCreateOrganizationWorkshop.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnResetAlertCreateOrganizationWorkshop();
      fnSearch();
    })
    .catch(e => errorNotify(e.msg));
};

const fnOpenAlertEditOrganizationWorkshop = params => {
  if (!params["uuid"]) return;
  currentUuid.value = params.uuid;

  ajaxGetOrganizationWorkshop(currentUuid.value, {
    "@~[]": ["OrganizationParagraph", "OrganizationParagraph.OrganizationRailway"],
  })
    .then(res => {
      uniqueCode_alertEditOrganizationWorkshop.value = res.content.organization_workshop.unique_code;
      name_alertEditOrganizationWorkshop.value = res.content.organization_workshop.name;
      organizationRailwayUuid_alertEditOrganizationWorkshop.value = res.content.organization_workshop.organization_paragraph.organization_railway.uuid;
      organizationParagraphUuid_alertEditOrganizationWorkshop.value = res.content.organization_workshop.organization_paragraph.uuid;
      organizationWorkshopTypeCode_alertEditOrganizationWorkshop.value = res.content.organization_workshop.type_code;
      alertEditOrganizationWorkshop.value = true;
    })
    .catch(e => errorNotify(e.msg));
};

const fnUpdateOrganizationWorkshop = () => {
  if (!currentUuid.value) return;

  const loading = loadingNotify();
  ajaxUpdateOrganizationWorkshop(currentUuid.value, {
    unique_code: uniqueCode_alertEditOrganizationWorkshop.value,
    name: name_alertEditOrganizationWorkshop.value,
    organization_paragraph_uuid: organizationParagraphUuid_alertEditOrganizationWorkshop.value,
    type_code: organizationWorkshopTypeCode_alertEditOrganizationWorkshop.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnSearch();
    })
    .catch(e => errorNotify(e.msg))
    .finally(() => loading());
};

const fnDestroyOrganizationWorkshop = params => {
  if (!params["uuid"]) return;

  actionNotify(
    destroyActions(() => {
      const loading = loadingNotify();

      ajaxDestroyOrganizationWorkshop(params.uuid)
        .then(() => {
          successNotify("删除成功");
          fnSearch();
        })
        .catch(e => errorNotify(e.msg))
        .finally(() => loading());
    })
  );
};

const fnDestroyOrganizationWorkshops = () => {
  actionNotify(
    destroyActions(() => {
      const loading = loadingNotify();

      ajaxDestroyOrganizationWorkshops(collect(selected.value).pluck("uuid").all())
        .then(() => {
          successNotify("删除成功");
          fnSearch();
        })
        .catch(e => errorNotify(e.msg))
        .finally(() => loading());
    })
  )
};
</script>
