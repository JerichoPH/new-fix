<template>
  <q-card flat bordered>
    <q-card-section>
      <p class="text-body1">{{ labelName }}</p>
      <q-separator class="q-mb-md" />

      <div class="row" v-for="(item, idx) in organizationStations" :key="idx"
        v-show="(organizationWorkshopUuid_alertEdit === item.organizationWorkshop.uuid) || (!organizationWorkshopUuid_alertEdit)">
        <div class="col-12">
          <p class="text-subtitle2 q-mb-none">{{ item.organizationWorkshop.name }}</p>
        </div>
        <template v-for="(stations, idx2) in item.organizationStations" :key="idx2">
          <div class="col-4" v-for="station in stations" :key="station.uuid">
            <q-checkbox v-model="checkedOrganizationStationUuids_alertEdit" :val="station.uuid" :key="station.uuid"
              :label="station.name" />
          </div>
        </template>
      </div>
    </q-card-section>
  </q-card>
</template>
<script setup>
import { ref, onMounted, inject, defineProps } from "vue";
import collect from "collect.js";
import { ajaxGetOrganizationStations } from "src/apis/organization";
import { errorNotify } from "src/utils/notify";

const props = defineProps({
  labelName: {
    type: String,
    default: "",
  },
  ajaxParams: {
    type: Object,
    default: () => {
      return {};
    },
  },
});
const labelName = props.labelName;
const ajaxParams = props.ajaxParams;
const organizationWorkshopUuid_alertEdit = ref("");
const organizationStations = ref([]);
const checkedOrganizationStationUuids_alertEdit = inject("checkedOrganizationStationUuids_alertEdit");

onMounted(() => fnSearch(""));

const fnSearch = () => {
  organizationStations.value = [];

  ajaxGetOrganizationStations({
    ...ajaxParams,
    "@~[]": ["OrganizationWorkshop"],
  })
    .then(res => {
      const tmp = {};
      collect(res.content.organization_stations).each(item => {
        if (!tmp[item.organization_workshop.unique_code]) {
          tmp[item.organization_workshop.unique_code] = { organizationWorkshop: item.organization_workshop, organizationStations: [] };
        }
        tmp[item.organization_workshop.unique_code].organizationStations.push(item);
      });
      organizationStations.value = collect(tmp)
        .map(item => {
          item.organizationStations = collect(item.organizationStations).chunk(3).toArray();
          return item;
        })
        .toArray();
    })
    .catch(e => errorNotify(e.msg));
};
</script>
src/utils/notify
