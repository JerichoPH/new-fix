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
                  <sel-install-indoor-room label-name="所属机房" sechma="search" />
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
        <div class="row">
          <div class="col"><span :style="{ fontSize: '20px' }">室内上道位置-排列表</span></div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="secondary" label="新建室内上道位置-排" icon="add" @click="fnOpenAlertCreateInstallIndoorPlatoon" />
              <q-btn color="negative" label="删除室内上道位置-排" icon="deleted" @click="fnDestroyInstallIndoorPlatoons" />
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
                  <q-th align="left" key="name" @click="event => fnColumnReverseSort(event, props, sortBy)">
                    名称
                  </q-th>
                  <q-th align="right"></q-th>
                </q-tr>
              </template>
              <template v-slot:body="props">
                <q-tr :props="props">
                  <q-td><q-checkbox :key="props.row.uuid" :value="props.row.uuid" v-model="props.selected" /></q-td>
                  <q-td>{{ props.row.index }}</q-td>
                  <q-td key="name" :props="props">{{ props.row.name }}</q-td>
                  <q-td key="operation" :props="props">
                    <q-btn-group>
                      <q-btn @click="fnOpenAlertEditInstallIndoorPlatoon(props.row.operation)" color="warning"
                        icon="edit">
                        编辑
                      </q-btn>
                      <q-btn @click="fnDestroyInstallIndoorPlatoon(props.row.operation)" color="negative" icon="delete">
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
  <!-- 新建室内上道位置-排弹窗 -->
  <q-dialog v-model="alertCreateInstallIndoorPlatoon" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">新建室内上道位置-排</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnStoreInstallIndoorPlatoon">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertCreateInstallIndoorPlatoon" label="名称"
                :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-install-indoor-room-type label-name="机房类型" sechma="alertCreate" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-install-indoor-room label-name="所属机房" sechma="alertCreate" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-railway_search label-name="所属路局" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-paragraph_search label-name="所属站段" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-workshop_search label-name="所属车间" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-station_search label-name="所属站场" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-railway_search label-name="所属中心" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-crossroad_search label-name="所属道口" />
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
  <!-- 编辑室内上道位置-排弹窗 -->
</template>
<script setup>
import { ref, onMounted, provide } from "vue";
import { fnColumnReverseSort, getVal } from "src/utils/common";
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

// 新建室内上道位置-排弹窗数据
const alertCreateInstallIndoorPlatoon = ref("");
const name_alertCreateInstallIndoorPlatoon = ref("");
const installIndoorRoomTypeUuid_alertCreateInstallIndoorPlatoon = ref("");
provide("installIndoorRoomTypeUuid_alertCreate", installIndoorRoomTypeUuid_alertCreateInstallIndoorPlatoon);
const installIndoorRoomUuid_alertCreateInstallIndoorPlatoon = ref("");
provide("installIndoorRoomUuid_alertCreate", installIndoorRoomUuid_alertCreateInstallIndoorPlatoon);
const organizationRailwayUuid_alertCreateInstallIndoorPlatoon = ref("");
provide("organizationRailwayUuid_alertCreate", organizationRailwayUuid_alertCreateInstallIndoorPlatoon);
const organizationParagraphUuid_alertCreateInstallIndoorPlatoon = ref("");
provide("organizationParagraphUuid_alertCreate", organizationParagraphUuid_alertCreateInstallIndoorPlatoon);
const organizationWorkshopUuid_alertCreateInstallIndoorPlatoon = ref("");
provide("organizationWorkshopUuid_alertCreate", organizationWorkshopUuid_alertCreateInstallIndoorPlatoon);
const organizationStationUuid_alertCreateInstallIndoorPlatoon = ref("");
provide("organizationStationUuid_alertCreate", organizationStationUuid_alertCreateInstallIndoorPlatoon);
const organizationCenterUuid_alertCreateInstallIndoorPlatoon = ref("");
provide("organizationCenterUuid_alertCreate", organizationCenterUuid_alertCreateInstallIndoorPlatoon);
const organizationCrossroadUuid_alertCreateInstallIndoorPlatoon = ref("");
provide("organizationCrossroadUuid_alertCreate", organizationCrossroadUuid_alertCreateInstallIndoorPlatoon);

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
const fnResetAlertCreateInstallIndoorPlatoon = () => { };
const fnOpenAlertCreateInstallIndoorPlatoon = () => {
  alertCreateInstallIndoorPlatoon.value = true;
 };
const fnStoreInstallIndoorPlatoon = params => { };
const fnOpenAlertEditInstallIndoorPlatoon = params => { };
const fnUpdateInstallIndoorPlatoon = params => { };
const fnDestroyInstallIndoorPlatoon = params => { };
const fnDestroyInstallIndoorPlatoons = () => { };
</script>
