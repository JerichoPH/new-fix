<template>
  <q-select outlined use-input clearable v-model="warehouseStorehouseUuid_alertEdit"
    :options="warehouseStorehouses_alertEdit" :label="labelName" @filter="fnFilter" emit-value map-options
    :display-value="warehouseStorehousesMap[warehouseStorehouseUuid_alertEdit]"
    :disable="!organizationWorkshopUuid_alertEdit && !organizationWorkAreaUuid_alertEdit" />
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
const organizationWorkshopUuid_alertEdit = inject("organizationWorkshopUuid_alertEdit");
const organizationWorkAreaUuid_alertEdit = inject("organizationWorkAreaUuid_alertEdit");
const warehouseStorehouseUuid_alertEdit = inject("warehouseStorehouseUuid_alertEdit");
const warehouseStorehouses_alertEdit = ref([]);
const warehouseStorehouses = ref([]);
const warehouseStorehousesMap = ref({});

watch(organizationWorkshopUuid_alertEdit, () => fnSearch());
watch(organizationWorkAreaUuid_alertEdit, () => fnSearch());

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      warehouseStorehouses_alertEdit.value = warehouseStorehouses.value;
    });
    return;
  }

  update(() => {
    warehouseStorehouses_alertEdit.value = warehouseStorehouses.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  warehouseStorehouses_alertEdit.value = [];

  if (!organizationWorkshopUuid_alertEdit.value && !organizationWorkAreaUuid_alertEdit.value) {
    warehouseStorehouseUuid_alertEdit.value = "";
    return;
  }

  ajaxGetWarehouseStorehouses({
    ...ajaxParams,
    organization_workshop_uuid: organizationWorkshopUuid_alertEdit.value,
    organization_work_area_uuid: organizationWorkAreaUuid_alertEdit.value,
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
