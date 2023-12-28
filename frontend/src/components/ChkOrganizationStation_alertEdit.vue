<template>
  <q-card flat bordered>
    <q-card-section>
      <p class="text-body1">{{ labelName }}</p>
      <div class="row" v-for="(items, idx) in organizationStations" :key="idx">
        <div class="col-3" v-for="(organizationStation, idx2) in items" :key="idx2">
          <q-checkbox v-model="checkedOrganizationStationUuids_alertEdit" :val="organizationStation.uuid"
            :key="organizationStation.uuid" :label="organizationStation.name" :true-value="defaultValue.find(organizationStation.uuid) > -1" />
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
  defaultValue: {
    type: Array,
    default: () => {
      return [];
    },
  },
});
const labelName = props.labelName;
const ajaxParams = props.ajaxParams;
const defaultValue = props.defaultValue;
const organizationStations = ref([]);
const checkedOrganizationStationUuids_alertEdit = inject("checkedRbacRoleUuids_alertEdit");

onMounted(() => fnSearch());

const fnSearch = () => {
  organizationStations.value = [];

  ajaxGetOrganizationStations(ajaxParams)
    .then((res) => {
      organizationStations.value = collect(res.content.organization_stations).chunk(4).all();
    })
    .catch((e) => errorNotify(e.msg));
};
</script>
src/utils/notify
