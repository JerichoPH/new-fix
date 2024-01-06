<template>
  <q-select outlined use-input clearable v-model="warehousePlatoonUuid_alertEdit" :options="warehousePlatoons_alertEdit"
    :label="labelName" @filter="fnFilter" emit-value map-options
    :display-value="warehousePlatoonsMap[warehousePlatoonUuid_alertEdit]"
    :disable="!warehouseAreaUuid_alertEdit" />
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
const warehouseAreaUuid_alertEdit = inject("warehouseAreaUuid_alertEdit");
const warehousePlatoonUuid_alertEdit = inject("warehousePlatoonUuid_alertEdit");
const warehousePlatoons_alertEdit = ref([]);
const warehousePlatoons = ref([]);
const warehousePlatoonsMap = ref({});

watch(warehouseAreaUuid_alertEdit, () => fnSearch());

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      warehousePlatoons_alertEdit.value = warehousePlatoons.value;
    });
    return;
  }

  update(() => {
    warehousePlatoons_alertEdit.value = warehousePlatoons.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  warehousePlatoons_alertEdit.value = [];

  if (!warehouseAreaUuid_alertEdit.value) {
    warehousePlatoonUuid_alertEdit.value = "";
    return;
  }

  ajaxGetWarehousePlatoons({
    ...ajaxParams,
    warehouse_area_uuid: warehouseAreaUuid_alertEdit.value,
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
    .catch(e=>errorNotify(e.msg));
};
</script>
src/utils/notify
