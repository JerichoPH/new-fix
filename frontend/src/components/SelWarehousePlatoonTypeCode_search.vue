<template>
  <q-select outlined use-input clearable v-model="warehousePlatoonTypeCode_search"
    :options="warehousePlatoonTypeCodes_search" :label="labelName" @filter="fnFilter" emit-value map-options
    :display-value="warehousePlatoonTypeCodesMap[warehousePlatoonTypeCode_search]"/>
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
const warehousePlatoonTypeCode_search = inject("warehousePlatoonTypeCode_search");
const warehousePlatoonTypeCodes_search = ref([]);
const warehousePlatoonTypeCodes = ref([]);
const warehousePlatoonTypeCodesMap = ref({});


const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      warehousePlatoonTypeCodes_search.value = warehousePlatoonTypeCodes.value;
    });
    return;
  }

  update(() => {
    warehousePlatoonTypeCodes_search.value = warehousePlatoonTypeCodes.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  warehousePlatoonTypeCodes.value = [];
  warehousePlatoonTypeCode_search.value = "";

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
    .catch(e=>errorNotify(e.msg));
};
</script>
src/utils/notify
