<template>
  <q-select outlined use-input clearable v-model="organizationRailwayUuid_search" :options="organizationRailways_search"
    :label="labelName" @filter="fnFilter" emit-value map-options />
</template>
<script setup>
import { inject, defineProps, onMounted, ref } from "vue";
import { ajaxGetOrganizationRailways } from "/src/apis/organization";
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
const organizationRailwayUuid_search = inject("organizationRailwayUuid_search");
const organizationRailways_search = ref([]);
const organizationRailways = ref([]);

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      organizationRailways_search.value = organizationRailways.value;
    });
    return;
  }

  update(() => {
    organizationRailways_search.value = organizationRailways.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  organizationRailways_search.value = [];

  ajaxGetOrganizationRailways(ajaxParams)
    .then((res) => {
      organizationRailways.value = collect(res.content.organization_railways)
        .map(organizationRailway => {
          return {
            label: organizationRailway.short_name,
            value: organizationRailway.uuid,
          };
        })
        .all();
    })
    .catch(e=>errorNotify(e.msg));
};
</script>
src/utils/notify
