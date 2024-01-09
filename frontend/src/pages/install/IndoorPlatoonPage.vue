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
                  <sel-install-indoor-room-type label-name="机房类型" sechma="search" />
                </div>
                <div class="col">
                  <SelInstallIndoorRoom label-name="所属机房" sechma="search" />
                </div>
                <div class="col">
                  <sel-organization-railway_search label-name="所属路局" />
                </div>
                <div class="col">
                  <sel-organization-paragraph_search label-name="所属站段" />
                </div>
              </div>
              <div class="row q-col-gutter-sm">
                <div class="col">
                  <sel-organization-workshop_search label-name="所属车间" />
                </div>
                <div class="col">
                  <sel-organization-station_search label-name="所属站场" />
                </div>
                <div class="col">
                  <sel-organization-center_search label-name="所属中心" />
                </div>
                <div class="col">
                  <sel-organization-crossroad_search label-name="所属道口" />
                </div>
                <div class="col"></div>
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
import { getVal } from "src/utils/common";
import {
  ajaxGetInstallIndoorPlatoons,
  ajaxGetInstallIndoorPlatoon,
  ajaxStoreInstallIndoorPlatoon,
  ajaxUpdateInstallIndoorPlatoon,
  ajaxDestroyInstallIndoorPlatoon,
  ajaxDestroyInstallIndoorPlatoons,
} from "src/apis/install";
import { notifies, actions } from "src/utils/notify";
import JoinString from "src/components/JoinString.vue";
import SelOrganizationRailway_search from "src/components/SelOrganizationRailway_search.vue";
import SelOrganizationParagraph_search from "src/components/SelOrganizationParagraph_search.vue";
import SelOrganizationWorkshop_search from "src/components/SelOrganizationWorkshop_search.vue";
import SelOrganizationStation_search from "src/components/SelOrganizationStation_search.vue";
import SelOrganizationCenter_search from "src/components/SelOrganizationCenter_search.vue";
import SelOrganizationCrossroad_search from "src/components/SelOrganizationCrossroad_search.vue";
import SelInstallIndoorRoomType from "src/components/SelInstallIndoorRoomType.vue";
import SelInstallIndoorRoom from "src/components/SelInstallIndoorRoom.vue";

import SelOrganizationRailway_alertCreate from "src/components/SelOrganizationRailway_alertCreate.vue";
import SelOrganizationParagraph_alertCreate from "src/components/SelOrganizationParagraph_alertCreate.vue";
import SelOrganizationWorkshop_alertCreate from "src/components/SelOrganizationWorkshop_alertCreate.vue";
import SelOrganizationStation_alertCreate from "src/components/SelOrganizationStation_alertCreate.vue";
import SelOrganizationCenter_alertCreate from "src/components/SelOrganizationCenter_alertCreate.vue";
import SelOrganizationCrossroad_alertCreate from "src/components/SelOrganizationCrossroad_alertCreate.vue";

import SelOrganizationRailway_alertEdit from "src/components/SelOrganizationRailway_alertEdit.vue";
import SelOrganizationParagraph_alertEdit from "src/components/SelOrganizationParagraph_alertEdit.vue";
import SelOrganizationWorkshop_alertEdit from "src/components/SelOrganizationWorkshop_alertEdit.vue";
import SelOrganizationStation_alertEdit from "src/components/SelOrganizationStation_alertEdit.vue";
import SelOrganizationCenter_alertEdit from "src/components/SelOrganizationCenter_alertEdit.vue";
import SelOrganizationCrossroad_alertEdit from "src/components/SelOrganizationCrossroad_alertEdit.vue";
import collect from "collect.js";

// 搜索栏数据
const name_search = ref("");
const installIndoorRoomTypeUuid_search = ref("");
provide("installIndoorRoomTypeUuid_search", installIndoorRoomTypeUuid_search);
const installIndoorRoomUuid_search = ref("");
provide("installIndoorRoomUuid_search", installIndoorRoomUuid_search);
const organizationRailwayUuid_search = ref("");
provide("organizationRailwayUuid_search", organizationRailwayUuid_search);
const organizationParagraphUuid_search = ref("");
provide("organizationParagraphUuid_search", organizationParagraphUuid_search);
const organizationWorkshopUuid_search = ref("");
provide("organizationWorkshopUuid_search", organizationWorkshopUuid_search);
const organizationStationUuid_search = ref("");
provide("organizationStationUuid_search", organizationStationUuid_search);
const organizationCenterUuid_search = ref("");
provide("organizationCenterUuid_search", organizationCenterUuid_search);
const organizationCrossroadUuid_search = ref("");
provide("organizationCrossroadUuid_search", organizationCrossroadUuid_search);

// 表格数据
const rows = ref([]);
const selected = ref([]);
const sortBy = ref("");

onMounted(() => fnInit());

const fnInit = () => fnSearch();

const fnResetSearch = () => {
  name_search.value = "";
  installIndoorRoomTypeUuid_search.value = "";
  installIndoorRoomUuid_search.value = "";
  organizationRailwayUuid_search.value = "";
  organizationParagraphUuid_search.value = "";
  organizationWorkshopUuid_search.value = "";
  organizationStationUuid_search.value = "";
  organizationCenterUuid_search.value = "";
  organizationCrossroadUuid_search.value = "";
};

const fnSearch = () => {
  rows.value = [];
  selected.value = [];

  ajaxGetInstallIndoorPlatoons({
    name: name_search.value,
    install_indoor_room_type_uuid: installIndoorRoomTypeUuid_search.value,
    install_indoor_room_uuid: installIndoorRoomUuid_search.value,
    organization_railway_uuid: organizationRailwayUuid_search.value,
    organization_paragraph_uuid: organizationParagraphUuid_search.value,
    organization_workshop_uuid: organizationWorkshopUuid_search.value,
    organization_station_uuid: organizationStationUuid_search.value,
    organization_center_uuid: organizationCenterUuid_search.value,
    organization_crossroad_uuid: organizationCrossroadUuid_search.value,
  })
    .then(res => {
      collect(getVal(res, 'content.install_indoor_platoons'))
        .map((installIndoorPlatoon, idx) => {
          return {
            index: idx + 1,
            name: getVal(installIndoorPlatoon, 'name'),
            installIndoorRoomType: getVal(installIndoorPlatoon, 'install_indoor_room.install_indoor_room_type'),
            installIndoorRoom: getVal(installIndoorPlatoon, 'install_indoor_room'),
            operation: { uuid: getVal(installIndoorPlatoon, 'uuid') },
          };
        })
        .all();
    })
    .catch(e => notifies.error(e.msg));
};

</script>
