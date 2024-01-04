<template>
  <q-select outlined use-input clearable v-model="warehouseStorehouseUuid_search" :options="warehouseStorehouses_search"
    :label="labelName" @filter="fnFilter" emit-value map-options
    :display-value="warehouseStorehousesMap[warehouseStorehouseUuid_search]"
    :disable="!organizationWorkshopUuid_search && !organizationWorkAreaUuid_search" />
</template>
<script setup>
import { inject, defineProps, onMounted, ref, watch } from "vue";
import { ajaxGetWarehouseStorehouses } from "/src/apis/warehouse";
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
const organizationWorkshopUuid_search = inject("organizationWorkshopUuid_search");
const organizationWorkAreaUuid_search = inject("organizationWorkAreaUuid_search");
const warehouseStorehouseUuid_search = inject("warehouseStorehouseUuid_search");
const warehouseStorehouses_search = ref([]);
const warehouseStorehouses = ref([]);
const warehouseStorehousesMap = ref({});

watch(organizationWorkshopUuid_search, () => fnSearch());
watch(organizationWorkAreaUuid_search, () => fnSearch());

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      warehouseStorehouses_search.value = warehouseStorehouses.value;
    });
    return;
  }

  update(() => {
    warehouseStorehouses_search.value = warehouseStorehouses.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  warehouseStorehouses_search.value = [];

  if (!organizationWorkshopUuid_search.value && !organizationWorkAreaUuid_search.value) {
    warehouseStorehouseUuid_search.value = "";
    return;
  }

  ajaxGetWarehouseStorehouses({
    ...ajaxParams,
    organization_workshop_uuid: organizationWorkshopUuid_search.value,
    organization_work_area_uuid: organizationWorkAreaUuid_search.value,
  })
    .then(res => {
      warehouseStorehouses.value = collect(res.content.warehouse_storehouses)
        .map(warehouseStorehouse => {
          return {
            label: warehouseStorehouse.name,
            value: warehouseStorehouse.uuid,
          };
        })
        .all();
      warehouseStorehousesMap.value = collect(warehouseStorehouses.value).pluck("label", "value").all();
    })
    .catch((e) => errorNotify(e.msg));
};
</script>
src/utils/notify
