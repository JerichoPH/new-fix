<template>
  <q-select outlined use-input clearable v-model="organizationRailwayUuid_alertCreate"
    :options="organizationRailways_search" :label="labelName" :option-disable="collect(organizationRailways).pluck('value', 'label').all()[
      organizationRailwayUuid_alertCreate
    ]
      " @filter="fnFilter" emit-value map-options />
</template>
<script setup>
import { inject, defineProps, onMounted, ref, watch } from "vue";
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
const organizationRailwayUuid_alertCreate = inject("organizationRailwayUuid_alertCreate");
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

onMounted(() => {
  organizationRailways.value = [];

  ajaxGetOrganizationRailways(ajaxParams)
    .then(res => {
      if (res.content.organization_railways.length > 0) {
        organizationRailways.value = collect(res.content.organization_railways)
          .map(organizationRailway => {
            return {
              label: organizationRailway.short_name,
              value: organizationRailway.uuid,
            };
          })
          .all();
      }
    })
    .catch((e) => errorNotify(e.msg));
});
</script>
src/utils/notify
