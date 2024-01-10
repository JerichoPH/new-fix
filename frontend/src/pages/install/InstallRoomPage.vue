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
                  <standard-select label-name="机房类型" sechma="search" current-field="installIndoorRoomTypeUuid"
                    :data-source="ajaxGetInstallIndoorRoomTypes" data-source-field="install_indoor_room_types" />
                </div>
                <div class="col">
                  <standard-select label-name="所属路局" sechma="search" current-field="organizationRailwayUuid"
                    :data-source="ajaxGetOrganizationRailways" data-source-field="organization_railways" />
                </div>
                <div class="col">
                  <standard-select label-name="所属站段" sechma="search" current-field="organizationParagraphUuid"
                    :data-source="ajaxGetOrganizationParagraphs" data-source-field="organization_paragraphs"
                    parent-field="organizationRailwayhUuid" />
                </div>
                <div class="col">
                  <standard-select label-name="所属车间" sechma="search" current-field="organizationWorkshopUuid"
                    :data-source="ajaxGetOrganizationWorkshops" data-source-field="organization_workshops"
                    parent-field="organizationParagraphUuid" />
                </div>
              </div>
              <div class="row q-col-gutter-sm">
                <div class="col">
                  <standard-select label-name="所属站场" sechma="search" current-field="organizationStationUuid"
                    :data-source="ajaxGetOrganizationStations" data-source-field="organization_stations"
                    parent-field="organizationWorkshopUuid" />
                </div>
                <div class="col">
                  <standard-select label-name="所属中心" sechma="search" current-field="organizationCenterUuid"
                    :data-source="ajaxGetOrganizationCenters" data-source-field="organization_centers"
                    parent-field="organizationWorkshopUuid" />
                </div>
                <div class="col">
                  <standard-select label-name="所属道口" sechma="search" current-field="organizationCrossroadUuid"
                    :data-source="ajaxGetOrganizationCrossroads" data-source-field="organization_crossroads"
                    parent-field="organizationWorkshopUuid" />
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
                      getVal(props, 'row.parent.organizationWorkshopByStation.name'),
                      getVal(props, 'row.parent.organizationWorkshopByCenter.name'),
                      getVal(props, 'row.parent.organizationWorkshopByCrossroad.name'),
                      getVal(props, 'row.parent.organizationStation.name'),
                      getVal(props, 'row.parent.organizationCenter.name'),
                      getVal(props, 'row.parent.organizationCrossroad.name'),
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
              <sel-install-indoor-room-type label-name="机房类型" sechma="alertEdit" />
              <standard-select label-name="机房类型" sechma="alertCreate" current-field="installIndoorRoomTypeUuid"
                :data-source="ajaxGetInstallIndoorRoomTypes" data-source-field="install_indoor_room_types" />
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
              <standard-select label-name="所属站场" sechma="alertCreate" current-field="organizationStationUuid"
                :data-source="ajaxGetOrganizationStations" data-source-field="organization_stations"
                parent-field="organizationWorkshopUuid" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属中心" sechma="alertCreate" current-field="organizationCenterUuid"
                :data-source="ajaxGetOrganizationCenters" data-source-field="organization_centers"
                parent-field="organizationWorkshopUuid" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属道口" sechma="alertCreate" current-field="organizationCrossroadUuid"
                :data-source="ajaxGetOrganizationCrossroads" data-source-field="organization_crossroads"
                parent-field="organizationWorkshopUuid" />
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
  <!-- 编辑室内上道位置-机房弹窗 -->
  <q-dialog v-model="alertEditInstallIndoorRoom" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">编辑室内上道位置-机房</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnUpdateInstallIndoorRoom">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertEditInstallIndoorRoom" label="名称" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-install-indoor-room-type label-name="机房类型" sechma="alertEdit" />
              <standard-select label-name="机房类型" sechma="alertEdit" current-field="installIndoorRoomTypeUuid"
                :data-source="ajaxGetInstallIndoorRoomTypes" data-source-field="install_indoor_room_types" />
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
              <standard-select label-name="所属站场" sechma="alertEdit" current-field="organizationStationUuid"
                :data-source="ajaxGetOrganizationStations" data-source-field="organization_stations"
                parent-field="organizationWorkshopUuid" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属中心" sechma="alertEdit" current-field="organizationCenterUuid"
                :data-source="ajaxGetOrganizationCenters" data-source-field="organization_centers"
                parent-field="organizationWorkshopUuid" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属道口" sechma="alertEdit" current-field="organizationCrossroadUuid"
                :data-source="ajaxGetOrganizationCrossroads" data-source-field="organization_crossroads"
                parent-field="organizationWorkshopUuid" />
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
import { fnColumnReverseSort, getVal } from "src/utils/common";
import {
  ajaxGetInstallIndoorRooms,
  ajaxGetInstallIndoorRoom,
  ajaxStoreInstallIndoorRoom,
  ajaxUpdateInstallIndoorRoom,
  ajaxDestroyInstallIndoorRoom,
  ajaxDestroyInstallIndoorRooms,
  ajaxGetOrganizationRailways,
  ajaxGetOrganizationParagraphs,
  ajaxGetOrganizationWorkshops,
  ajaxGetOrganizationStations,
  ajaxGetOrganizationCenters,
  ajaxGetOrganizationCrossroads,
  ajaxGetInstallIndoorRoomTypes,
} from "src/apis/install";
import { notifies, actions } from "src/utils/notify";
import JoinString from "src/components/JoinString.vue";
import StandardSelect from "src/components/StandardSelect.vue";

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

// 编辑室内上道位置-机房弹窗数据
const alertEditInstallIndoorRoom = ref(false);
const currentUuid = ref("");
const name_alertEditInstallIndoorRoom = ref("");
const installIndoorRoomTypeUuid_alertEditInstallIndoorRoom = ref("");
provide("installIndoorRoomTypeUuid_alertEdit", installIndoorRoomTypeUuid_alertEditInstallIndoorRoom);
const organizationRailwayUuid_alertEditInstallIndoorRoom = ref("");
provide("organizationRailwayUuid_alertEdit", organizationRailwayUuid_alertEditInstallIndoorRoom);
const organizationParagraphUuid_alertEditInstallIndoorRoom = ref("");
provide("organizationParagraphUuid_alertEdit", organizationParagraphUuid_alertEditInstallIndoorRoom);
const organizationWorkshopUuid_alertEditInstallIndoorRoom = ref("");
provide("organizationWorkshopUuid_alertEdit", organizationWorkshopUuid_alertEditInstallIndoorRoom);
const organizationStationUuid_alertEditInstallIndoorRoom = ref("");
provide("organizationStationUuid_alertEdit", organizationStationUuid_alertEditInstallIndoorRoom);
const organizationCenterUuid_alertEditInstallIndoorRoom = ref("");
provide("organizationCenterUuid_alertEdit", organizationCenterUuid_alertEditInstallIndoorRoom);
const organizationCrossroadUuid_alertEditInstallIndoorRoom = ref("");
provide("organizationCrossroadUuid_alertEdit", organizationCrossroadUuid_alertEditInstallIndoorRoom);

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
  rows.value = [];
  selected.value = [];

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
  }).then(res => {
    rows.value = collect(res.content.install_indoor_rooms)
      .map((installIndoorRoom, idx) => {
        return {
          index: idx + 1,
          uuid: installIndoorRoom.uuid,
          name: installIndoorRoom.name,
          installIndoorRoomType: installIndoorRoom.install_indoor_room_type,
          parent: {
            organizationStation: installIndoorRoom.organization_station,
            organizationCenter: installIndoorRoom.organization_center,
            organizationCrossroad: installIndoorRoom.organization_crossroad,
            organizationWorkshopByStation: getVal(installIndoorRoom, 'organization_station.organization_workshop'),
            organizationWorkshopByCenter: getVal(installIndoorRoom, 'organization_center.organization_workshop'),
            organizationWorkshopByCrossroad: getVal(installIndoorRoom, 'organization_crossroad.organization_workshop'),
            organizationParagraphByStation: getVal(installIndoorRoom, 'organization_station.organization_workshop.organization_paragraph'),
            organizationParagraphByCenter: getVal(installIndoorRoom, 'organization_center.organization_workshop.organization_paragraph'),
            organizationParagraphByCrossroad: getVal(installIndoorRoom, 'organization_crossroad.organization_workshop.organization_paragraph'),
            organizationRailwayByStation: getVal(installIndoorRoom, 'organization_station.organization_workshop.organization_paragraph.organization_railway'),
            organizationRailwayByCenter: getVal(installIndoorRoom, 'organization_center.organization_workshop.organization_paragraph.organization_railway'),
            organizationRailwayByCrossroad: getVal(installIndoorRoom, 'organization_crossroad.organization_workshop.organization_paragraph.organization_railway'),
          },
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
    .then(res => {
      notifies.success(res.msg);
      fnResetAlertCreateInstallIndoorRoom();
      fnSearch();

      alertCreateInstallIndoorRoom.value = false;
    })
    .catch(e => notifies.error(e.msg))
    .finally(loading());
};

const fnOpenAlertEditInstallIndoorRoom = params => {
  if (!getVal(params, "uuid")) return;
  currentUuid.value = getVal(params, "uuid");

  ajaxGetInstallIndoorRoom(currentUuid.value, {
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
  })
    .then(res => {
      name_alertEditInstallIndoorRoom.value = getVal(res, "content.install_indoor_room.name");
      installIndoorRoomTypeUuid_alertEditInstallIndoorRoom.value = getVal(res, "content.install_indoor_room.install_indoor_room_type_uuid");
      if (getVal(res, "content.install_indoor_room.organization_station")) {
        organizationRailwayUuid_alertEditInstallIndoorRoom.value = getVal(res, "content.install_indoor_room.organization_station.organization_workshop.organization_paragraph.organization_railway.uuid");
        organizationParagraphUuid_alertEditInstallIndoorRoom.value = getVal(res, "content.install_indoor_room.organization_station.organization_workshop.organization_paragraph.uuid")
        organizationWorkshopUuid_alertEditInstallIndoorRoom.value = getVal(res, "content.install_indoor_room.organization_station.organization_workshop.uuid")
        organizationStationUuid_alertEditInstallIndoorRoom.value = getVal(res, "content.install_indoor_room.organization_station.uuid")
      } else if (getVal(res, "content.install_indoor_room.organization_crossroad")) {
        organizationRailwayUuid_alertEditInstallIndoorRoom.value = getVal(res, "content.install_indoor_room.organization_crossroad.organization_workshop.organization_paragraph.organization_railway.uuid");
        organizationParagraphUuid_alertEditInstallIndoorRoom.value = getVal(res, "content.install_indoor_room.organization_crossroad.organization_workshop.organization_paragraph.uuid")
        organizationWorkshopUuid_alertEditInstallIndoorRoom.value = getVal(res, "content.install_indoor_room.organization_crossroad.organization_workshop.uuid")
        organizationCrossroadUuid_alertEditInstallIndoorRoom.value = getVal(res, "content.install_indoor_room.organization_crossroad.uuid")
      } else if (getVal(res, "content.install_indoor_room.organization_center")) {
        organizationRailwayUuid_alertEditInstallIndoorRoom.value = getVal(res, "content.install_indoor_room.organization_center.organization_workshop.organization_paragraph.organization_railway.uuid");
        organizationParagraphUuid_alertEditInstallIndoorRoom.value = getVal(res, "content.install_indoor_room.organization_center.organization_workshop.organization_paragraph.uuid")
        organizationWorkshopUuid_alertEditInstallIndoorRoom.value = getVal(res, "content.install_indoor_room.organization_center.organization_workshop.uuid")
        organizationCenterUuid_alertEditInstallIndoorRoom.value = getVal(res, "content.install_indoor_room.organization_center.uuid")
      }

      alertEditInstallIndoorRoom.value = true;
    })
    .catch(e => notifies.error(e.msg));
};

const fnUpdateInstallIndoorRoom = () => {
  if (!currentUuid.value) return;

  const loading = notifies.loading();

  ajaxUpdateInstallIndoorRoom(currentUuid.value, {
    name: name_alertEditInstallIndoorRoom.value,
    install_indoor_room_type_uuid: installIndoorRoomTypeUuid_alertEditInstallIndoorRoom.value,
    organization_station_uuid: organizationStationUuid_alertEditInstallIndoorRoom.value,
    organization_crossroad_uuid: organizationCrossroadUuid_alertEditInstallIndoorRoom.value,
    organization_center_uuid: organizationCenterUuid_alertEditInstallIndoorRoom.value,
  })
    .then(res => {
      notifies.success(res.msg);
      fnSearch();

      alertEditInstallIndoorRoom.value = false;
    })
    .catch(e => notifies.error(e.msg))
    .finally(loading());
};

const fnDestroyInstallIndoorRoom = parms => {
  if (!getVal(parms, "uuid")) return;

  notifies.confirm(actions.destory(() => {
    const loading = notifies.loading();

    ajaxDestroyInstallIndoorRoom(getVal(parms, "uuid"))
      .then(() => {
        notifies.success("删除成功");
        fnSearch();
      })
      .catch(e => notifies.error(e.msg))
      .finally(loading());
  }));
};

const fnDestroyInstallIndoorRooms = () => {
  notifies.confirm(actions.destory(() => {
    const loading = notifies.loading();

    ajaxDestroyInstallIndoorRooms(collect(selected.value).pluck("uuid"))
      .then(() => {
        notifies.success("删除成功");
        fnSearch();
      })
      .catch(e => notifies.error(e.msg))
      .finally(loading());
  }));

};
</script>
