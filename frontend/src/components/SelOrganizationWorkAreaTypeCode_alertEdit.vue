<template>
  <q-select outlined use-input clearable v-model="organizationWorkAreaTypeCode_alertEdit"
    :options="organizationWorkAreaTypeCodes_alertEdit" :label="labelName" @filter="fnFilter" emit-value map-options
    :display-value="organizationWorkAreaTypeCodesMap[organizationWorkAreaTypeCode_alertEdit]"/>
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
const organizationWorkAreaTypeCode_alertEdit = inject("organizationWorkAreaTypeCode_alertEdit");
const organizationWorkAreaTypeCodes_alertEdit = ref([]);
const organizationWorkAreaTypeCodes = ref([]);
const organizationWorkAreaTypeCodesMap = ref({});


const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      organizationWorkAreaTypeCodes_alertEdit.value = organizationWorkAreaTypeCodes.value;
    });
    return;
  }

  update(() => {
    organizationWorkAreaTypeCodes_alertEdit.value = organizationWorkAreaTypeCodes.value.filter(
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
