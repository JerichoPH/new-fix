<template>
  <q-select outlined use-input clearable v-model="warehouseStorehouseUuid_alertCreate"
    :options="warehouseStorehouses_alertCreate" :label="labelName" @filter="fnFilter" emit-value map-options
    :display-value="warehouseStorehousesMap[warehouseStorehouseUuid_alertCreate]"
    :disable="!organizationWorkshopUuid_alertCreate && !organizationWorkAreaUuid_alertCreate" />
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
const organizationWorkshopUuid_alertCreate = inject("organizationWorkshopUuid_alertCreate");
const organizationWorkAreaUuid_alertCreate = inject("organizationWorkAreaUuid_alertCreate");
const warehouseStorehouseUuid_alertCreate = inject("warehouseStorehouseUuid_alertCreate");
const warehouseStorehouses_alertCreate = ref([]);
const warehouseStorehouses = ref([]);
const warehouseStorehousesMap = ref({});

watch(organizationWorkshopUuid_alertCreate, () => fnSearch());
watch(organizationWorkAreaUuid_alertCreate, () => fnSearch());

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      warehouseStorehouses_alertCreate.value = warehouseStorehouses.value;
    });
    return;
  }

  update(() => {
    warehouseStorehouses_alertCreate.value = warehouseStorehouses.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  warehouseStorehouses_alertCreate.value = [];

  if (!organizationWorkshopUuid_alertCreate.value && !organizationWorkAreaUuid_alertCreate.value) {
    warehouseStorehouseUuid_alertCreate.value = "";
    return;
  }

  ajaxGetWarehouseStorehouses({
    ...ajaxParams,
    organization_workshop_uuid: organizationWorkshopUuid_alertCreate.value,
    organization_work_area_uuid: organizationWorkAreaUuid_alertCreate.value,
  })
    .then((res) => {
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
    .catch(e=>errorNotify(e.msg));
};
</script>
src/utils/notify
