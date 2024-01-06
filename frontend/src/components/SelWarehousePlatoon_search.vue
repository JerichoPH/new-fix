<template>
  <q-select outlined use-input clearable v-model="warehousePlatoonUuid_search" :options="warehousePlatoons_search"
    :label="labelName" @filter="fnFilter" emit-value map-options
    :display-value="warehousePlatoonsMap[warehousePlatoonUuid_search]"
    :disable="!warehouseAreaUuid_search" />
</template>
<script setup>
import { inject, defineProps, onMounted, ref, watch } from "vue";
import { ajaxGetWarehousePlatoons } from "/src/apis/warehouse";
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
const warehouseAreaUuid_search = inject("warehouseAreaUuid_search");
const warehousePlatoonUuid_search = inject("warehousePlatoonUuid_search");
const warehousePlatoons_search = ref([]);
const warehousePlatoons = ref([]);
const warehousePlatoonsMap = ref({});

watch(warehouseAreaUuid_search, () => fnSearch());

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      warehousePlatoons_search.value = warehousePlatoons.value;
    });
    return;
  }

  update(() => {
    warehousePlatoons_search.value = warehousePlatoons.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  warehousePlatoons_search.value = [];

  if (!warehouseAreaUuid_search.value) {
    warehousePlatoonUuid_search.value = "";
    return;
  }

  ajaxGetWarehousePlatoons({
    ...ajaxParams,
    warehouse_area_uuid: warehouseAreaUuid_search.value,
  })
    .then((res) => {
      warehousePlatoons.value = collect(res.content.warehouse_platoons)
        .map(warehousePlatoon => {
          return {
            label: warehousePlatoon.name,
            value: warehousePlatoon.uuid,
          };
        })
        .all();
      warehousePlatoonsMap.value = collect(warehousePlatoons.value).pluck("label", "value").all();
    })
    .catch(e=>errorNotify(e.msg));d
};
</script>
src/utils/notify
