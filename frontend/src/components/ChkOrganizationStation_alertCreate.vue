<template>
  <q-card flat bordered>
    <q-card-section>
      <p class="text-body1">{{ labelName }}</p>
      <div class="row" v-for="(items, idx) in organizationStations" :key="idx">
        <div class="col-3" v-for="(organizationStation, idx2) in items" :key="idx2">
          <q-checkbox v-model="checkedOrganizationStationUuids_alertCreate" :val="organizationStation.uuid"
            :key="organizationStation.uuid" :label="organizationStation.name" :true-value="organizationStation.uuid" />
        </div>
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
const organizationStations = ref([]);
const checkedOrganizationStationUuids_alertCreate = inject("checkedOrganizationStationUuids_alertCreate");

onMounted(() => fnSearch(""));

const fnSearch = organizationWorkshopUuid => {
  organizationStations.value = [];

  if (!organizationWorkshopUuid) return;
  ajaxGetOrganizationStations({
    ...ajaxParams,
    "@~[]": ["OrganizationWorkshop"],
  })
    .then(res => {
      organizationStations.value = collect(res.content.organization_stations)
        .groupBy(item => item.organization_workshop.name).
        each(items => items.chunk(4).all())
        .all();
        console.log('ok',organizationStations.value);
    })
    .catch(e => errorNotify(e.msg));
};
</script>
src/utils/notify
