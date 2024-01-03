<template>
  <q-select outlined use-input clearable v-model="warehouseAreaUuid_alertEdit" :options="warehouseAreas_alertCreate"
    :label="labelName" @filter="fnFilter" emit-value map-options
    :display-value="warehouseAreasMap[warehouseAreaUuid_alertEdit]"
    :disable="!organizationWorkshopUuid_search && !warehouseStorehouseUuid_alertEdit"/>
</template>
<script setup>
import { inject, defineProps, onMounted, ref } from "vue";
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
const warehouseStorehouseUuid_alertEdit = inject("warehouseStorehouseUuid_alertEdit");
const warehouseAreaUuid_alertEdit = inject("warehouseAreaUuid_alertEdit");
const warehouseAreas_alertCreate = ref([]);
const warehouseAreas = ref([]);
const warehouseAreasMap = ref({});

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      warehouseAreas_alertCreate.value = warehouseAreas.value;
    });
    return;
  }

  update(() => {
    warehouseAreas_alertCreate.value = warehouseAreas.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  warehouseAreas_alertCreate.value = [];

  if(!warehouseStorehouseUuid_alertEdit.value){
    warehouseAreaUuid_alertEdit.value = "";
    return;
  }

  ajaxGetWarehouseAreas({
    ...ajaxParams,
    organization_workshop_uuid: organizationWorkshopUuid_search.value,
    organization_work_area_uuid: warehouseStorehouseUuid_alertEdit.value,
  })
    .then((res) => {
      warehouseAreas.value = collect(res.content.warehouse_areas)
        .map(warehouseArea => {
          return {
            label: warehouseArea.name,
            value: warehouseArea.uuid,
          };
        })
        .all();
        warehouseAreasMap.value = collect(warehouseAreas.value).pluck("label","value").all();
    })
    .catch((e) => {
      errorNotify(e.msg);
    });
};
</script>
src/utils/notify
