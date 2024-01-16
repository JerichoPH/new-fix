<template>
  <div class="row q-mt-md">
    <div class="col">
      <q-select outlined use-input clearable v-model="organizationStations_com" :options="organizationCrossroads"
        label="所属车站" @filter="fnFilterOrganizationStation" emit-value map-options
        :display-value="organizaitonStationsMap[organizationWorkshopUuid]" :disable="!organizationWorkshopUuid"
        :key="`organizationStationUuid_${sechma}`" />
    </div>
  </div>
</template>
<script setup>
import { ref, onMounted, watch, defineProps, inject } from "vue";
import { fnColumnReverseSort, getVal } from "src/utils/common";
import collect from "collect.js";
import { notifies } from "src/utils/notify";
import {
  ajaxGetOrganizationStations,
  ajaxGetOrganizationCrossroads,
  ajaxGetOrganizationCenters,
} from "src/apis/organization";
import StandardSelect from "./StandardSelect.vue";

const props = defineProps({
  sechma: {
    type: String,
    default: "",
  },
});

const sechma = props.sechma;
const organizationStations = ref([]);
const organizationStations_com = ref([]);
const organizaitonStationsMap = ref({});
const organizationStationUuid_com = inject(`organizationStationUuid_${sechma}`);
const organizationCrossroads = ref([]);
const organizationCrossroads_com = ref([]);
const organizationCrossroadsMap = ref([]);
const organizationCrossroadUuid_com = inject(`organizationCrossroadUuid_${sechma}`);
const organizationCenters = ref([]);
const organizationCenters_com = ref([]);
const organizationCentersMap = ref([]);
const organizationCenterUuid_com = inject(`organizationCenterUuid_${sechma}`);
const organizationWorkshopUuid = inject(`organizationWorkshopUuid_${sechma}`);
const strategy_com = inject(`strategy_${sechma}`);

watch(organizationWorkshopUuid, () => { console.log("OK"); })
watch(organizationStationUuid_com, () => fnStrategy("station"));
watch(organizationCrossroadUuid_com, () => fnStrategy("crossroad"));
watch(organizationCenterUuid_com, () => fnStrategy("center"));

const fnFilterOrganizationStation = (val, update) => {
  if (val === "") {
    update(() => {
      organizationStations_com.value = organizationCrossroads.value;
    });
    return;
  }

  update(() => {
    organizationStations_com.value = organizationCrossroads.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

const fnFilterOrganizationCrossroad = (val, update) => {
  if (val === "") {
    update(() => {
      organizationCrossroads_com.value = organizationCrossroads.value;
    });
    return;
  }

  update(() => {
    organizationCrossroads_com.value = organizationCrossroads.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

const fnFilterOrganizationCenter = (val, update) => {
  if (val === "") {
    update(() => {
      organizationCenters_com.value = organizationCenters.value;
    });
    return;
  }

  update(() => {
    organizationCenters_com.value = organizationCenters.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

const fnStrategy = strategy => {
  switch (strategy) {
    case "station":
      organizationCrossroadUuid_com.value = "";
      organizationCenterUuid_com.value = "";
      fnInitOrganizationCrossroads();
      fnInitOrganizationCenters();
      fnSearchOrganizationStations();
      strategy_com.value = strategy;
      break;
    case "crossroad":
      organizationStationUuid_com.value = "";
      organizationCenterUuid_com.value = "";
      fnInitOrganizationStations();
      fnInitOrganizationCenters();
      fnSearchOrganizationCrossroads();
      strategy_com.value = strategy;
      break;
    case "center":
      organizationStationUuid_com.value = "";
      organizationCrossroadUuid_com.value = "";
      fnInitOrganizationStations();
      fnInitOrganizationCrossroads();
      fnSearchOrganizationCenters();
      strategy_com.value = strategy;
      break;
  }
};

const fnInitOrganizationStations = () => {
  organizationCenters.value = [];
  organizationCenters_com.value = [];
  organizationCentersMap.value = {};
};

const fnInitOrganizationCrossroads = () => {
  organizationCrossroads.value = [];
  organizationCrossroads_com.value = [];
  organizationCrossroadsMap.value = {};
};

const fnInitOrganizationCenters = () => {
  organizationCenters.value = [];
  organizationCenters_com.value = [];
  organizationCentersMap.value = {};
};

const fnInit = () => {
  organizationStationUuid_com.value = "";
  organizationCrossroadUuid_com.value = "";
  organizationCenterUuid_com.value = "";
  fnInitOrganizationStations();
  fnInitOrganizationCrossroads();
  fnInitOrganizationCenters();
};

const fnSearchOrganizationStations = () => {
  if (!organizationWorkshopUuid.value) {
    fnInit();
    return;
  }

  ajaxGetOrganizationStations({
    organization_workshop_uuid: organizationWorkshopUuid.value,
  })
    .then(res => {
      organizationCrossroads.value = collect(res.content.organization_stations)
        .each(organizationStation => {
          return {
            label: organizationStation.name,
            value: organizationStation.uuid,
          };
        })
        .all();
      organizaitonStationsMap.value = collect(organizationCrossroads.value).pluck("label", "value").all();
    })
    .catch(e => notifies.error(e.msg));
};

const fnSearchOrganizationCrossroads = () => {
  if (!organizationWorkshopUuid.value) {
    fnInit();
    return;
  }

  ajaxGetOrganizationCrossroads({
    organization_workshop_uuid: organizationWorkshopUuid.value,
  })
    .then(res => {
      organizationCrossroads.value = collect(res.content.organization_crossroads)
        .each(organizationCrossroad => {
          return {
            label: organizationCrossroad.name,
            value: organizationCrossroad.uuid,
          };
        })
        .all();
      organizaitonStationsMap.value = collect(organizationCrossroads.value).pluck("label", "value").all();
    })
    .catch(e => notifies.error(e.msg));
};

const fnSearchOrganizationCenters = () => {
  if (!organizationWorkshopUuid.value) {
    fnInit();
    return;
  }

  ajaxGetOrganizationCenters({
    organization_workshop_uuid: organizationWorkshopUuid.value,
  })
    .then(res => {
      organizationCrossroads.value = collect(res.content.organization_crossroads)
        .each(organizationCrossroad => {
          return {
            label: organizationCrossroad.name,
            value: organizationCrossroad.uuid,
          };
        })
        .all();
      organizaitonStationsMap.value = collect(organizationCrossroads.value).pluck("label", "value").all();
    })
    .catch(e => notifies.error(e.msg));
};

</script>
