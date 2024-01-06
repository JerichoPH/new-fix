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
                  <sel-install-indoor-room-type_search label-name="机房类型" />
                </div>
                <div class="col">
                  <sel-organization-railway_search label-name="所属路局" />
                </div>
                <div class="col">
                  <sel-organization-paragraph_search label-name="所属站段" />
                </div>
                <div class="col">
                  <sel-organization-workshop_search label-name="所属车间" />
                </div>
              </div>
              <div class="row q-col-gutter-sm">
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
          <div class="col"><span :style="{ fontSize: '20px' }">室内上道位置-机房列表</span></div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="secondary" label="新建室内上道位置-机房" icon="add" @click="fnOpenAlertCreateInstallIndoorRoom" />
              <q-btn color="negative" label="删除室内上道位置-机房" icon="deleted" @click="fnDestroyInstallIndoorRooms" />
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
                  <q-th align="left" key="installIndoorRoomType"
                    @click="event => fnColumnReverseSort(event, props, sortBy)">
                    所属机房类型
                  </q-th>
                  <q-th align="left" key="parent" @click="event => fnColumnReverseSort(event, props, sortBy)">
                    所属上级
                  </q-th>
                  <q-th align="right"></q-th>
                </q-tr>
              </template>
              <template v-slot:body="props">
                <q-tr :props="props">
                  <q-td><q-checkbox :key="props.row.uuid" :value="props.row.uuid" v-model="props.selected" /></q-td>
                  <q-td>{{ props.row.index }}</q-td>
                  <q-td key="name" :props="props">{{ props.row.name }}</q-td>
                  <q-td key="installIndoorRoomType" :props="props">{{ props.row.installIndoorRoomType.name }}</q-td>
                  <q-td key="parent" :props="props">
                    <join-string :values="[
                      props.row.organizationRailwayByStation.name,
                      props.row.organizationRailwayByCenter.name,
                      props.row.organizationRailwayByCrossroad.name,
                      props.row.organizationParagraphByStation.name,
                      props.row.organizationParagraphByCenter.name,
                      props.row.organizationParagraphByCrossroad.name,
                      props.row.organizationWorkshopByStation.name,
                      props.row.organizationWorkshopByCenter.name,
                      props.row.organizationWorkshopByCrossroad.name,
                      props.row.organizationStation.name,
                      props.row.organizationCenter.name,
                      props.row.organizationCrossroad.name,
                    ]" />
                  </q-td>
                  <q-td key="operation" :props="props">
                    <q-btn-group>
                      <q-btn @click="fnOpenAlertEditInstallIndoorRoom(props.row.operation)" color="warning" icon="edit">
                        编辑
                      </q-btn>
                      <q-btn @click="fnDestroyInstallIndoorRoom(props.row.operation)" color="negative" icon="delete">
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
  <!-- 新建室内上道位置-机房弹窗 -->
  <q-dialog v-model="alertCreateInstallIndoorRoom" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">室内上道位置-机房</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnStoreInstallIndoorRoom">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertCreateInstallIndoorRoom" label="名称" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-install-indoor-room-type_alertCreate label-name="所属机房类型" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-railway_alert-create label-name="所属路局" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-paragraph_alert-create label-name="所属站段" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-workshop_alert-create label-name="所属车间" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-station_alert-create label-name="所属站场" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-center_alert-create label-name="所属中心" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-crossroad_alert-create label-name="所属道口" />
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
</template>
<script setup>
import { ref, onMounted, provide } from "vue";
import collect from "collect.js";
import { fnColumnReverseSort } from "src/utils/common";
import {
  ajaxGetInstallIndoorRooms,
  ajaxGetInstallIndoorRoom,
  ajaxStoreInstallIndoorRoom,
  ajaxUpdateInstallIndoorRoom,
  ajaxDestroyInstallIndoorRoom,
  ajaxDestroyInstallIndoorRooms,
} from "src/apis/install";
import { notifies, actions } from "src/utils/notify";
import JoinString from "src/components/JoinString.vue";
import SelOrganizationRailway_search from "src/components/SelOrganizationRailway_search.vue";
import SelOrganizationParagraph_search from "src/components/SelOrganizationParagraph_search.vue";
import SelOrganizationWorkshop_search from "src/components/SelOrganizationWorkshop_search.vue";
import SelOrganizationStation_search from "src/components/SelOrganizationStation_search.vue";
import SelOrganizationCenter_search from "src/components/SelOrganizationCenter_search.vue";
import SelOrganizationCrossroad_search from "src/components/SelOrganizationCrossroad_search.vue";
import SelInstallIndoorRoomType_search from "src/components/SelInstallIndoorRoomType_search.vue";
import SelInstallIndoorRoomType_alertCreate from "src/components/SelInstallIndoorRoomType_alertCreate.vue";
import SelOrganizationRailway_alertCreate from "src/components/SelOrganizationRailway_alertCreate.vue";
import SelOrganizationParagraph_alertCreate from "src/components/SelOrganizationParagraph_alertCreate.vue";
import SelOrganizationWorkshop_alertCreate from "src/components/SelOrganizationWorkshop_alertCreate.vue";
import SelOrganizationStation_alertCreate from "src/components/SelOrganizationStation_alertCreate.vue";
import SelOrganizationCenter_alertCreate from "src/components/SelOrganizationCenter_alertCreate.vue";
import SelOrganizationCrossroad_alertCreate from "src/components/SelOrganizationCrossroad_alertCreate.vue";
import SelInstallIndoorRoomType_alertEdit from "src/components/SelInstallIndoorRoomType_alertEdit.vue";
import SelOrganizationRailway_alertEdit from "src/components/SelOrganizationRailway_alertEdit.vue";
import SelOrganizationParagraph_alertEdit from "src/components/SelOrganizationParagraph_alertEdit.vue";
import SelOrganizationWorkshop_alertEdit from "src/components/SelOrganizationWorkshop_alertEdit.vue";
import SelOrganizationStation_alertEdit from "src/components/SelOrganizationStation_alertEdit.vue";
import SelOrganizationCenter_alertEdit from "src/components/SelOrganizationCenter_alertEdit.vue";

// 搜索栏数据
const name_search = ref("");
const installIndoorRoomTypeUuid_search = ref("");
provide("installIndoorRoomTypeUuid_search", installIndoorRoomTypeUuid_search);
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

// 新建室内上道位置-机房弹窗数据
const alertCreateInstallIndoorRoom = ref(false);
const name_alertCreateInstallIndoorRoom = ref("");
const installIndoorRoomTypeUuid_alertCreateInstallIndoorRoom = ref("");
provide("installIndoorRoomTypeUuid_alertCreate", installIndoorRoomTypeUuid_alertCreateInstallIndoorRoom);
const organizationRailwayUuid_alertCreateInstallIndoorRoom = ref("");
provide("organizationRailwayUuid_alertCreate", organizationRailwayUuid_alertCreateInstallIndoorRoom);
const organizationParagraphUuid_alertCreateInstallIndoorRoom = ref("");
provide("organizationParagraphUuid_alertCreate", organizationParagraphUuid_alertCreateInstallIndoorRoom);
const organizationWorkshopUuid_alertCreateInstallIndoorRoom = ref("");
provide("organizationWorkshopUuid_alertCreate", organizationWorkshopUuid_alertCreateInstallIndoorRoom);
const organizationStationUuid_alertCreateInstallIndoorRoom = ref("");
provide("organizationStationUuid_alertCreate", organizationStationUuid_alertCreateInstallIndoorRoom);
const organizationCenterUuid_alertCreateInstallIndoorRoom = ref("");
provide("organizationCenterUuid_alertCreate", organizationCenterUuid_alertCreateInstallIndoorRoom);
const organizationCrossroadUuid_alertCreateInstallIndoorRoom = ref("");
provide("organizationCrossroadUuid_alertCreate", organizationCrossroadUuid_alertCreateInstallIndoorRoom);

onMounted(() => fnInit());

const fnInit = () => fnSearch();

const fnResetSearch = () => {
  name_search.value = "";
  installIndoorRoomTypeUuid_search.value = "";
  organizationRailwayUuid_search.value = "";
  organizationParagraphUuid_search.value = "";
  organizationWorkshopUuid_search.value = "";
  organizationStationUuid_search.value = "";
  organizationCenterUuid_search.value = "";
  organizationCrossroadUuid_search.value = "";
};

const fnSearch = () => {
  ajaxGetInstallIndoorRooms({
    "@~[]": [
      "InstallIndoorRoomType",
      "OrganizationStation",
      "OrganizationStation.OrganizationWorkshop",
      "OrganizationStation.OrganizationWorkshop.OrganizationParagraph",
      "OrganizationStation.OrganizationWorkshop.OrganizationParagraph.OrganizationRailway",
      "OrganizationCenter",
      "OrganizationCenter.OrganizationWorkshop",
      "OrganizationCenter.OrganizationWorkshop.OrganizationParagraph",
      "OrganizationCenter.OrganizationWorkshop.OrganizationParagraph.OrganizationRailway",
      "OrganizationCrossroad",
      "OrganizationCrossroad.OrganizationWorkshop",
      "OrganizationCrossroad.OrganizationWorkshop.OrganizationParagraph",
      "OrganizationCrossroad.OrganizationWorkshop.OrganizationParagraph.OrganizationRailway",
    ],
    name: name_search.value,
    installIndoorRoomTypeUuid: installIndoorRoomTypeUuid_search.value,
    organizationRailwayUuid: organizationRailwayUuid_search.value,
    organizationParagraphUuid: organizationParagraphUuid_search.value,
    organizationWorkshopUuid: organizationWorkshopUuid_search.value,
    organizationStationUuid: organizationStationUuid_search.value,
    organizationCenterUuid: organizationCenterUuid_search.value,
    organizationCrossroadUuid: organizationCrossroadUuid_search.value,
  }).then((res) => {
    rows.value = collect(res.content.install_indoor_rooms)
      .map((installIndoorRoom, idx) => {
        return {
          index: idx + 1,
          uuid: installIndoorRoom.uuid,
          name: installIndoorRoom.name,
          installIndoorRoomType: installIndoorRoom.install_indoor_room_type,
          organizationStation: installIndoorRoom.organization_station,
          organizationCenter: installIndoorRoom.organization_center,
          organizationCrossroad: installIndoorRoom.organization_crossroad,
          organizationWorkshopByStation: installIndoorRoom.organization_station.organization_workshop,
          organizationWorkshopByCenter: installIndoorRoom.organization_center.organization_workshop,
          organizationWorkshopByCrossroad: installIndoorRoom.organization_crossroad.organization_workshop,
          organizationParagraphByStation: installIndoorRoom.organization_station.organization_workshop.organization_paragraph,
          organizationParagraphByCenter: installIndoorRoom.organization_center.organization_workshop.organization_paragraph,
          organizationParagraphByCrossroad: installIndoorRoom.organization_crossroad.organization_workshop.organization_paragraph,
          organizationRailwayByStation: installIndoorRoom.organization_station.organization_workshop.organization_paragraph.organization_railway,
          organizationRailwayByCenter: installIndoorRoom.organization_center.organization_workshop.organization_paragraph.organization_railway,
          organizationRailwayByCrossroad: installIndoorRoom.organization_crossroad.organization_workshop.organization_paragraph.organization_railway,
          operation: { uuid: installIndoorRoom.uuid },
        }
      })
      .all();
  });
};

const fnResetAlertCreateInstallIndoorRoom = () => {
  name_alertCreateInstallIndoorRoom.value = "";
  installIndoorRoomTypeUuid_alertCreateInstallIndoorRoom.value = "";
  organizationRailwayUuid_alertCreateInstallIndoorRoom.value = "";
  organizationParagraphUuid_alertCreateInstallIndoorRoom.value = "";
  organizationWorkshopUuid_alertCreateInstallIndoorRoom.value = "";
  organizationStationUuid_alertCreateInstallIndoorRoom.value = "";
  organizationCenterUuid_alertCreateInstallIndoorRoom.value = "";
  organizationCrossroadUuid_alertCreateInstallIndoorRoom.value = "";
};

const fnOpenAlertCreateInstallIndoorRoom = () => { alertCreateInstallIndoorRoom.value = true; };

const fnStoreInstallIndoorRoom = () => {
  const loading = notifies.loading();

  ajaxStoreInstallIndoorRoom({
    name: name_alertCreateInstallIndoorRoom.value,
    install_indoor_room_type_uuid: installIndoorRoomTypeUuid_alertCreateInstallIndoorRoom.value,
    organization_railway_uuid: organizationRailwayUuid_alertCreateInstallIndoorRoom.value,
    organization_paragraph_uuid: organizationParagraphUuid_alertCreateInstallIndoorRoom.value,
    organization_workshop_uuid: organizationWorkshopUuid_alertCreateInstallIndoorRoom.value,
    organization_station_uuid: organizationStationUuid_alertCreateInstallIndoorRoom.value,
    organization_center_uuid: organizationCenterUuid_alertCreateInstallIndoorRoom.value,
    organization_crossroad_uuid: organizationCrossroadUuid_alertCreateInstallIndoorRoom.value,
  })
    .then((res) => {
      notifies.success(res.message);
      fnResetAlertCreateInstallIndoorRoom();
      fnSearch();

      alertCreateInstallIndoorRoom.value = false;
    })
    .catch(e => notifies.error(e.msg))
    .finally(loading());
};

const fnOpenAlertEditInstallIndoorRoom = params => {
};

const fnUpdateInstallIndoorRoom = () => { };

const fnDestroyInstallIndoorRoom = parms => { };

const fnDestroyInstallIndoorRooms = () => { };
</script>
