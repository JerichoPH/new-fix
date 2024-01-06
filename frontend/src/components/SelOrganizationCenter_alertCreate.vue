<template>
  <q-select outlined use-input clearable v-model="organizationCenterUuid_alertCreate"
    :options="organizationCenters_search" :display-value="organizationCentersMap[organizationCenterUuid_alertCreate]"
    :label="labelName" @filter="fnFilter" emit-value map-options :disable="!organizationWorkshopUuid_alertCreate" />
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
const organizationWorkshopUuid_alertCreate = inject("organizationWorkshopUuid_alertCreate");
const organizationCenterUuid_alertCreate = inject("organizationCenterUuid_alertCreate");
const organizationCenters_search = ref([]);
const organizationCenters = ref([]);
const organizationCentersMap = ref({});

watch(organizationWorkshopUuid_alertCreate, () => fnSearch());

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

onMounted(() => fnSearch());

const fnSearch = () => {
  organizationCenters_search.value = [];

  if (!organizationWorkshopUuid_alertCreate.value) {
    organizationCenterUuid_alertCreate.value = "";
    return;
  }

  ajaxGetOrganizationCenters({
    ...ajaxParams,
    organization_workshop_uuid: organizationWorkshopUuid_alertCreate.value,
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
    .catch(e=>errorNotify(e.msg));
};
</script>
src/utils/notify
