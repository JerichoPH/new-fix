<template>
  <q-select outlined use-input clearable v-model="warehouseAreaUuid_search" :options="warehouseAreas_search"
    :label="labelName" @filter="fnFilter" emit-value map-options
    :display-value="warehouseAreasMap[warehouseAreaUuid_search]"
    :disable="!warehouseStorehouseUuid_search" />
</template>
<script setup>
import { inject, defineProps, onMounted, ref, watch } from "vue";
import { ajaxGetWarehouseAreas } from "/src/apis/warehouse";
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
const warehouseStorehouseUuid_search = inject("warehouseStorehouseUuid_search");
const warehouseAreaUuid_search = inject("warehouseAreaUuid_search");
const warehouseAreas_search = ref([]);
const warehouseAreas = ref([]);
const warehouseAreasMap = ref({});

watch(warehouseStorehouseUuid_search, () => fnSearch());

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      warehouseAreas_search.value = warehouseAreas.value;
    });
    return;
  }

  update(() => {
    warehouseAreas_search.value = warehouseAreas.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  warehouseAreas_search.value = [];

  if (!warehouseStorehouseUuid_search.value) {
    warehouseAreaUuid_search.value = "";
    return;
  }

  ajaxGetWarehouseAreas({
    ...ajaxParams,
    warehouse_storehouse_uuid: warehouseStorehouseUuid_search.value,
  })
    .then((res) => {
      console.log('ok',res);
      warehouseAreas.value = collect(res.content.warehouse_areas)
        .map(warehouseArea => {
          return {
            label: warehouseArea.name,
            value: warehouseArea.uuid,
          };
        })
        .all();
      warehouseAreasMap.value = collect(warehouseAreas.value).pluck("label", "value").all();
    })
    .catch((e) => errorNotify(e.msg));
};
</script>
src/utils/notify
