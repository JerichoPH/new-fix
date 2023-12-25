<template>
  <q-select outlined use-input clearable v-model="organizationRailwayUuid_search" :options="organizationRailways_search" :label="labelName"
    @filter="fnFilter" emit-value map-options />
</template>
<script setup>
import { inject, defineProps, onMounted, ref } from "vue";
import { ajaxGetRbacMenus } from "/src/apis/rbac";
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

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      organizationRailways_search.value = rbacMenus.value;
    });
    return;
  }

  update(() => {
    organizationRailways_search.value = rbacMenus.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => {
  organizationRailways_search.value = [];

  ajaxGetRbacMenus(ajaxParams)
    .then((res) => {
      if (res.content.organization_railways.length > 0) {
        organizationRailways_search.value = res.content.organization_railways.map(organizationRailway=>{
          return {
            label: organizationRailway.name,
            value: organizationRailway.uuid,
          }
        });
      }
    })
    .catch((e) => {
      errorNotify(e.msg);
    });
});
</script>
src/utils/notify
