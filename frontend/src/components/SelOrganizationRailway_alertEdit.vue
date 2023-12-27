<template>
  <q-select outlined use-input clearable v-model="organizationRailwayUuid_alertEdit"
    :options="organizationRailways_alertEdit" :label="labelName" @filter="fnFilter" emit-value map-options
    :display-value="organizationRailwaysMap[organizationRailwayUuid_alertEdit]" />
</template>
<script setup>
import { inject, defineProps, onMounted, ref } from "vue";
import collect from "collect.js";
import { ajaxGetOrganizationRailways } from "/src/apis/organization";
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
const organizationRailwayUuid_alertEdit = inject("organizationRailwayUuid_alertEdit");
const organizationRailways_alertEdit = ref([]);
const organizationRailways = ref([]);
const organizationRailwaysMap = ref({});

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      organizationRailways_alertEdit.value = organizationRailways.value;
    });
    return;
  }

  update(() => {
    organizationRailways_alertEdit.value = organizationRailways.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => {
  ajaxGetOrganizationRailways(ajaxParams)
    .then((res) => {
      organizationRailways.value = collect(res.content.organization_railways)
        .map((organizationRailway) => {
          return {
            label: organizationRailway.short_name,
            value: organizationRailway.uuid,
          };
        })
        .all();
      organizationRailwaysMap.value = collect(organizationRailways.value).pluck('label', 'value').all();
    })
    .catch((e) => errorNotify(e.msg));
});
</script>
src/utils/notify
