<template>
  <q-select outlined use-input clearable v-model="organizationWorkshopTypeCode_alertCreate"
    :options="organizationWorkshopTypeCodes_alertCreate" :label="labelName" @filter="fnFilter" emit-value map-options
    :display-value="organizationWorkshopTypeCodesMap[organizationWorkshopTypeCode_alertCreate]"/>
</template>
<script setup>
import { inject, defineProps, onMounted, ref } from "vue";
import collect from "collect.js";
import { ajaxGetOrganizationWorkshopTypeCodesMap } from "/src/apis/organization";
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
const organizationWorkshopTypeCode_alertCreate = inject("organizationWorkshopTypeCode_alertCreate");
const organizationWorkshopTypeCodes_alertCreate = ref([]);
const organizationWorkshopTypeCodes = ref([]);
const organizationWorkshopTypeCodesMap = ref({});


const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      organizationWorkshopTypeCodes_alertCreate.value = organizationWorkshopTypeCodes.value;
    });
    return;
  }

  update(() => {
    organizationWorkshopTypeCodes_alertCreate.value = organizationWorkshopTypeCodes.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => {
  fnSearch();
});

const fnSearch = () => {
  organizationWorkshopTypeCodes.value = [];

  ajaxGetOrganizationWorkshopTypeCodesMap(ajaxParams)
    .then(res => {
      collect(res.content.type_codes_map)
        .map(item => {
          organizationWorkshopTypeCodes.value.push({
            label: item.text,
            value: item.code,
          });
        })
        .all();
      organizationWorkshopTypeCodesMap.value = collect(organizationWorkshopTypeCodes.value).pluck('label', 'value').all();
    })
    .catch((e) => errorNotify(e.msg));
};
</script>
src/utils/notify
