<template>
  <q-select outlined use-input clearable v-model="warehousePlatoonUuid_alertCreate" :options="warehousePlatoons_alertCreate"
    :label="labelName" @filter="fnFilter" emit-value map-options
    :display-value="warehousePlatoonsMap[warehousePlatoonUuid_alertCreate]"
    :disable="!warehouseAreaUuid_alertCreate" />
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
const warehouseAreaUuid_alertCreate = inject("warehouseAreaUuid_alertCreate");
const warehousePlatoonUuid_alertCreate = inject("warehousePlatoonUuid_alertCreate");
const warehousePlatoons_alertCreate = ref([]);
const warehousePlatoons = ref([]);
const warehousePlatoonsMap = ref({});

watch(warehouseAreaUuid_alertCreate, () => fnSearch());

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      warehousePlatoons_alertCreate.value = warehousePlatoons.value;
    });
    return;
  }

  update(() => {
    warehousePlatoons_alertCreate.value = warehousePlatoons.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  warehousePlatoons_alertCreate.value = [];

  if (!warehouseAreaUuid_alertCreate.value) {
    warehousePlatoonUuid_alertCreate.value = "";
    return;
  }

  ajaxGetWarehousePlatoons({
    ...ajaxParams,
    warehouse_area_uuid: warehouseAreaUuid_alertCreate.value,
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
    .catch((e) => errorNotify(e.msg));
};
</script>
src/utils/notify
