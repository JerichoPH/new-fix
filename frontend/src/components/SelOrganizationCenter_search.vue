<template>
  <q-select outlined use-input clearable v-model="organizationCenterUuid_search"
    :options="organizationCenters_search" :display-value="organizationCentersMap[organizationCenterUuid_search]"
    :label="labelName" @filter="fnFilter" emit-value map-options :disable="!organizationWorkshopUuid_search" />
</template>
<script setup>
import { inject, defineProps, onMounted, ref, watch } from "vue";
import { ajaxGetOrganizationCenters } from "/src/apis/organization";
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
const organizationWorkshopUuid_search = inject("organizationWorkshopUuid_search");
const organizationCenterUuid_search = inject("organizationCenterUuid_search");
const organizationCenters_search = ref([]);
const organizationCenters = ref([]);
const organizationCentersMap = ref({});

watch(organizationWorkshopUuid_search, (newValue) => fnSearch(newValue));

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      organizationCenters_search.value = organizationCenters.value;
    });
    return;
  }

  update(() => {
    organizationCenters_search.value = organizationCenters.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch(""));

const fnSearch = (organizationWorkshopUuid) => {
  organizationCenters_search.value = [];

  if (!organizationWorkshopUuid) return;
  ajaxGetOrganizationCenters({
    ...ajaxParams,
    organization_workshop_uuid: organizationWorkshopUuid,
  })
    .then((res) => {
      organizationCenters.value =
        collect(res.content.organization_centers)
          .map(organizationCenter => {
            return {
              label: organizationCenter.name,
              value: organizationCenter.uuid,
            };
          })
          .all();
      organizationCentersMap.value = collect(organizationCenters.value).pluck('label', 'value').all();
    })
    .catch((e) => errorNotify(e.msg));
};
</script>
src/utils/notify
