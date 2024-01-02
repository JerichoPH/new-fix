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
                  <q-input outlined clearable lazy-rules v-model="name_search" label="名称" :rules="[]" class="q-mb-md" />
                </div>
                <div class="col-3">
                  <sel-organization-railway_search label-name="所属路局" />
                </div>
                <div class="col-3">
                  <sel-oragnization-paragraph_search label-name="所属站段" />
                </div>
                <div class="col-3">
                  <sel-organization-workshop_search label-name="所属车间" />
                </div>
              </div>
            </q-form>
          </div>
        </div>
      </q-card-section>
    </q-card>

    <q-card class="q-mt-md">
      <q-card-section>

      </q-card-section>
    </q-card>
  </div>
</template>
<script setup>
import { ref, onMounted, provide } from "vue";
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
  actionNotify,
  destroyActions,
} from "src/utils/notify";
import SelOrganizationRailway_search from "src/components/SelOrganizationRailway_search.vue";
import SelOragnizationParagraph_search from "src/components/SelOragnizationParagraph_search.vue";
import SelOrganizationWorkshop_search from "src/components/SelOrganizationWorkshop_search.vue";

// 搜索栏数据
const name_search = ref("");
const organizationRailwayUuid_search = ref("");
provide("organizationRailwayUuid_search", organizationRailwayUuid_search);
const organizationParagraphUuid_search = ref("");
provide("organizationParagraphUuid_search", organizationParagraphUuid_search);
const organizationWorkshopUuid_search = ref("");
provide("organizationWorkshopUuid_search", organizationWorkshopUuid_search);

onMounted(() => fnInit());

const fnInit = () => fnSearch();

const fnResetSearch = () => {
  name_search.value = "";
  organizationRailwayUuid_search.value = "";
  organizationParagraphUuid_search.value = "";
  organizationWorkshopUuid_search.value = "";
};

const fnSearch = () => {
  ajaxGetWarehouseStorehouses({
    name: name_search.value,
    organization_workshop_uuid: organizationWorkshopUuid_search.value,
  })
    .then(res => {
      
    })
    .catch(e => errorNotify(e.msg));
};

</script>
