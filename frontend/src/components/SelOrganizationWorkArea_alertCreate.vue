<template>
  <q-select outlined use-input clearable v-model="organizationWorkAreaUuid_alertCreate"
    :options="organizationWorkAreas_alertCreate"
    :display-value="organizationWorkAreasMap[organizationWorkAreaUuid_alertCreate]" :label="labelName" @filter="fnFilter"
    emit-value map-options :disable="!organizationWorkshopUuid_alertCreate" />
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
const organizationWorkshopUuid_alertCreate = inject("organizationWorkshopUuid_alertCreate");
const organizationWorkAreaUuid_alertCreate = inject("organizationWorkAreaUuid_alertCreate");
const organizationWorkAreas_alertCreate = ref([]);
const organizationWorkAreas = ref([]);
const organizationWorkAreasMap = ref({});

watch(organizationWorkshopUuid_alertCreate, () => fnSearch());

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      organizationWorkAreas_alertCreate.value = organizationWorkAreas.value;
    });
    return;
  }

  update(() => {
    organizationWorkAreas_alertCreate.value = organizationWorkAreas.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  organizationWorkAreas_alertCreate.value = [];

  if (!organizationWorkshopUuid_alertCreate.value) {
    organizationWorkAreaUuid_alertCreate.value = "";
    return;
  }

  ajaxGetOrganizationWorkAreas({
    ...ajaxParams,
    organization_workshop_uuid: organizationWorkshopUuid_alertCreate.value,
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
