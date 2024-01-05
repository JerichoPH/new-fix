<template>
  <q-select outlined use-input clearable v-model="warehousePlatoonTypeCode_alertCreate"
    :options="warehousePlatoonTypeCodes_alertCreate" :label="labelName" @filter="fnFilter" emit-value map-options
    :display-value="warehousePlatoonTypeCodesMap[warehousePlatoonTypeCode_alertCreate]"/>
</template>
<script setup>
import { inject, defineProps, onMounted, ref } from "vue";
import collect from "collect.js";
import { ajaxGetWarehousePlatoonTypeCodesMap } from "/src/apis/warehouse";
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
const warehousePlatoonTypeCode_alertCreate = inject("warehousePlatoonTypeCode_alertCreate");
const warehousePlatoonTypeCodes_alertCreate = ref([]);
const warehousePlatoonTypeCodes = ref([]);
const warehousePlatoonTypeCodesMap = ref({});


const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      warehousePlatoonTypeCodes_alertCreate.value = warehousePlatoonTypeCodes.value;
    });
    return;
  }

  update(() => {
    warehousePlatoonTypeCodes_alertCreate.value = warehousePlatoonTypeCodes.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  warehousePlatoonTypeCodes.value = [];
  warehousePlatoonTypeCode_alertCreate.value = "";

  ajaxGetWarehousePlatoonTypeCodesMap(ajaxParams)
    .then(res => {
      collect(res.content.type_codes_map)
        .map(item => {
          warehousePlatoonTypeCodes.value.push({
            label: item.text,
            value: item.code,
          });
        })
        .all();
      warehousePlatoonTypeCodesMap.value = collect(warehousePlatoonTypeCodes.value).pluck('label', 'value').all();
    })
    .catch((e) => errorNotify(e.msg));
};
</script>
src/utils/notify
