<template>
  <q-select outlined use-input clearable v-model="organizationWorkAreaUuid_search" :options="organizationWorkAreas_search"
    :display-value="organizationWorkAreasMap[organizationWorkAreaUuid_search]" :label="labelName" @filter="fnFilter"
    emit-value map-options :disable="!organizationWorkshopUuid_search" />
</template>
<script setup>
import { inject, defineProps, onMounted, ref, watch } from "vue";
import { ajaxGetOrganizationWorkAreas } from "/src/apis/organization";
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
const organizationWorkAreaUuid_search = inject("organizationWorkAreaUuid_search");
const organizationWorkAreas_search = ref([]);
const organizationWorkAreas = ref([]);
const organizationWorkAreasMap = ref({});

watch(organizationWorkshopUuid_search, () => fnSearch());

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      organizationWorkAreas_search.value = organizationWorkAreas.value;
    });
    return;
  }

  update(() => {
    organizationWorkAreas_search.value = organizationWorkAreas.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  organizationWorkAreas_search.value = [];


  if (!organizationWorkshopUuid_search.value) {
    organizationWorkAreaUuid_search.value = "";
    return;
  }

  ajaxGetOrganizationWorkAreas({
    ...ajaxParams,
    organization_workshop_uuid: organizationWorkshopUuid_search.value,
  })
    .then((res) => {
      organizationWorkAreas.value =
        collect(res.content.organization_work_areas)
          .map(organizationWorkArea => {
            return {
              label: organizationWorkArea.name,
              value: organizationWorkArea.uuid,
            };
          })
          .all();
      organizationWorkAreasMap.value = collect(organizationWorkAreas.value).pluck('label', 'value').all();
    })
    .catch((e) => errorNotify(e.msg));
};
</script>
src/utils/notify
