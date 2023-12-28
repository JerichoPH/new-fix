<template>
  <q-card>
    <q-card-section>
      <p class="text-body1">{{ labelName }}</p>
      <div class="row" v-for="(items, idx) in organizationStations" :key="idx">
        <div class="col-3" v-for="(organizationStation, idx) in items" :key="idx">
          <q-checkbox v-model="checkedOrganizationStationUuids_alertEdit" :val="organizationStation.uuid"
            :key="organizationStation.uuid" :label="organizationStation.name" />
        </div>
      </div>
    </q-card-section>
  </q-card>
</template>
<script setup>
import { ref, onMounted, inject, defineProps, watch } from "vue";
import collect from "collect.js";
import { ajaxGetOrganizationStations } from "src/apis/rbac";
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
const organizationWorkshopUuid_alertEdit = inject("organizationWorkshopUuid_alertEdit");
const checkedOrganizationStationUuids_alertEdit = inject("checkedRbacRoleUuids_alertEdit");

watch(organizationWorkshopUuid_alertEdit, newValue => fnSearch(newValue));

onMounted(() => fnSearch(""));

const fnSearch = organizationWorkshopUuid => {
  organizationStations.value = [];

  if (!organizationWorkshopUuid) return;
  ajaxGetOrganizationStations(ajaxParams)
    .then((res) => {
      organizationStations.value = collect(res.content.organization_stations).chunk(4).all();
    })
    .catch((e) => errorNotify(e.msg));
};
</script>
src/utils/notify
