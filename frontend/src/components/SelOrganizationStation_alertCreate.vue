<template>
  <q-select outlined use-input clearable v-model="organizationStationUuid_alertCreate" :options="organizationStations_search"
    :display-value="organizationStationsMap[organizationStationUuid_alertCreate]" :label="labelName" @filter="fnFilter"
    emit-value map-options :disable="!organizationWorkshopUuid_alertCreate" />
</template>
<script setup>
import { inject, defineProps, onMounted, ref, watch } from "vue";
import { ajaxGetOrganizationStations } from "/src/apis/organization";
import collect from "collect.js";
import { errorNotify } from "src/utils/notify";

const props = defineProps({
  labelName: {
    type: String,
    default: "",
    required: true,
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
const organizationWorkshopUuid_alertCreate = inject("organizationWorkshopUuid_alertCreate");
const organizationStationUuid_alertCreate = inject("organizationStationUuid_alertCreate");
const organizationStations_search = ref([]);
const organizationStations = ref([]);
const organizationStationsMap = ref({});

watch(organizationWorkshopUuid_alertCreate, () => fnSearch());

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      organizationStations_search.value = organizationStations.value;
    });
    return;
  }

  update(() => {
    organizationStations_search.value = organizationStations.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  organizationStations_search.value = [];

  if (!organizationWorkshopUuid_alertCreate.value) {
    organizationStationUuid_alertCreate.value = "";
    return;
  }

  ajaxGetOrganizationStations({
    ...ajaxParams,
    organization_workshop_uuid: organizationWorkshopUuid_alertCreate.value,
  })
    .then((res) => {
      organizationStations.value =
        collect(res.content.organization_stations)
          .map(organizationStation => {
            return {
              label: organizationStation.name,
              value: organizationStation.uuid,
            };
          })
          .all();
      organizationStationsMap.value = collect(organizationStations.value).pluck('label', 'value').all();
    })
    .catch(e=>errorNotify(e.msg));
};
</script>
src/utils/notify
