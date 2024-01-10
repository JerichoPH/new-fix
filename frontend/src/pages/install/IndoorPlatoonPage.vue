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
                  <standard-select label-name="机房类型" sechma="search" :data-source="ajaxGetInstallIndoorRoomTypes"
                    data-source-field="install_indoor_room_types" current-field="installIndoorRoomTypeUuid" />
                </div>
                <div class="col">
                  <standard-select label-name="所属机房" sechma="search" :data-source="ajaxGetInstallIndoorRooms"
                    data-source-field="install_indoor_rooms" current-field="installIndoorRoomUuid"
                    parent-field="installIndoorRoomTypeUuid" />
                </div>
                <div class="col">
                  <standard-select label-name="所属路局" sechma="search" :data-source="ajaxGetOrganizationRailways"
                    data-source-field="organization_railways" current-field="organizationRailwayUuid"
                    label-field="short_name" />
                </div>
                <div class="col">
                  <standard-select label-name="所属站段" sechma="search" :data-source="ajaxGetOrganizationParagraphs"
                    data-source-field="organization_paragraphs" current-field="organizationParagraphUuid"
                    parent-field="organizationRailwayUuid" />
                </div>
              </div>
              <div class="row q-col-gutter-sm">
                <div class="col">
                  <standard-select label-name="所属车间" sechma="search" :data-source="ajaxGetOrganizationWorkshops"
                    data-source-field="organization_workshops" current-field="organizationWorkshopUuid"
                    parent-field="organizationParagraphUuid" />
                </div>
                <div class="col">
                  <standard-select label-name="所属站场" sechma="search" :data-source="ajaxGetOrganizationStations"
                    data-source-field="organization_stations" current-field="organizationStationUuid"
                    parent-field="organizationWorkshopUuid" />
                </div>
                <div class="col">
                  <standard-select label-name="所属中心" sechma="search" :data-source="ajaxGetOrganizationCenters"
                    data-source-field="organization_centers" current-field="organizationCenterUuid"
                    parent-field="organizationWorkshopUuid" />
                </div>
                <div class="col">
                  <standard-select label-name="所属道口" sechma="search" :data-source="ajaxGetOrganizationCrossroads"
                    data-source-field="organization_crossroads" current-field="organizationCrossroadUuid"
                    parent-field="organizationWorkshopUuid" />
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
              <standard-select label-name="机房类型" sechma="alertCreate" :data-source="ajaxGetInstallIndoorRoomTypes"
                data-source-field="install_indoor_room_types" current-field="installIndoorRoomTypeUuid" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属机房" sechma="alertCreate" :data-source="ajaxGetInstallIndoorRooms"
                data-source-field="install_indoor_rooms" current-field="installIndoorRoomUuid"
                parent-field="installIndoorRoomTypeUuid" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属路局" sechma="alertCreate" :data-source="ajaxGetOrganizationRailways"
                data-source-field="organization_railways" current-field="organizationRailwayUuid"
                label-field="short_name" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属站段" sechma="alertCreate" :data-source="ajaxGetOrganizationParagraphs"
                data-source-field="organization_paragraphs" current-field="organizationParagraphUuid"
                parent-field="organizationRailwayUuid" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属车间" sechma="alertCreate" :data-source="ajaxGetOrganizationWorkshops"
                data-source-field="organization_workshops" current-field="organizationWorkshopUuid"
                parent-field="organizationParagraphUuid" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属站场" sechma="alertCreate" :data-source="ajaxGetOrganizationStations"
                data-source-field="organization_stations" current-field="organizationStationUuid"
                parent-field="organizationWorkshopUuid" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属中心" sechma="alertCreate" :data-source="ajaxGetOrganizationCenters"
                data-source-field="organization_centers" current-field="organizationCenterUuid"
                parent-field="organizationWorkshopUuid" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <standard-select label-name="所属道口" sechma="alertCreate" :data-source="ajaxGetOrganizationCrossroads"
                data-source-field="organization_crossroads" current-field="organizationCrossroadUuid"
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
  ajaxGetInstallIndoorRoomTypes,
  ajaxGetInstallIndoorRooms,
} from "src/apis/install";
import {
  ajaxGetOrganizationRailways,
  ajaxGetOrganizationParagraphs,
  ajaxGetOrganizationWorkshops,
  ajaxGetOrganizationStations,
  ajaxGetOrganizationCenters,
  ajaxGetOrganizationCrossroads,
} from "src/apis/organization";
import { notifies, actions } from "src/utils/notify";
import JoinString from "src/components/JoinString.vue";
import StandardSelect from "src/components/StandardSelect.vue";
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
const alertCreateInstallIndoorPlatoon = ref(false``);
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
