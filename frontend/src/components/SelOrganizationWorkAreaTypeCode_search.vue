<template>
  <q-select outlined use-input clearable v-model="organizationWorkAreaTypeCode_search"
    :options="organizationWorkAreaTypeCodes_search" :label="labelName" @filter="fnFilter" emit-value map-options
    :display-value="organizationWorkAreaTypeCodesMap[organizationWorkAreaTypeCode_search]"/>
</template>
<script setup>
import { inject, defineProps, onMounted, ref } from "vue";
import collect from "collect.js";
import { ajaxGetOrganizationWorkAreaTypeCodesMap } from "/src/apis/organization";
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
const organizationWorkAreaTypeCode_search = inject("organizationWorkAreaTypeCode_search");
const organizationWorkAreaTypeCodes_search = ref([]);
const organizationWorkAreaTypeCodes = ref([]);
const organizationWorkAreaTypeCodesMap = ref({});


const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      organizationWorkAreaTypeCodes_search.value = organizationWorkAreaTypeCodes.value;
    });
    return;
  }

  update(() => {
    organizationWorkAreaTypeCodes_search.value = organizationWorkAreaTypeCodes.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  organizationWorkAreaTypeCodes.value = [];

  ajaxGetOrganizationWorkAreaTypeCodesMap(ajaxParams)
    .then(res => {
      collect(res.content.type_codes_map)
        .map(item => {
          organizationWorkAreaTypeCodes.value.push({
            label: item.text,
            value: item.code,
          });
        })
        .all();
      organizationWorkAreaTypeCodesMap.value = collect(organizationWorkAreaTypeCodes.value).pluck('label', 'value').all();
    })
    .catch(e=>errorNotify(e.msg));
};
</script>
src/utils/notify
