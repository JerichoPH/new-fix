<template>
  <q-select outlined use-input clearable v-model="organizationWorkAreaUuid_alertEdit"
    :options="organizationWorkAreas_alertEdit"
    :display-value="organizationWorkAreasMap[organizationWorkAreaUuid_alertEdit]" :label="labelName" @filter="fnFilter"
    emit-value map-options :disable="!organizationWorkshopUuid_alertEdit" />
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
const organizationWorkshopUuid_alertEdit = inject("organizationWorkshopUuid_alertEdit");
const organizationWorkAreaUuid_alertEdit = inject("organizationWorkAreaUuid_alertEdit");
const organizationWorkAreas_alertEdit = ref([]);
const organizationWorkAreas = ref([]);
const organizationWorkAreasMap = ref({});

watch(organizationWorkshopUuid_alertEdit, () => fnSearch());

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      organizationWorkAreas_alertEdit.value = organizationWorkAreas.value;
    });
    return;
  }

  update(() => {
    organizationWorkAreas_alertEdit.value = organizationWorkAreas.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  organizationWorkAreas_alertEdit.value = [];

  if (!organizationWorkshopUuid_alertEdit.value) {
    organizationWorkAreaUuid_alertEdit.value = "";
    return;
  }

  ajaxGetOrganizationWorkAreas({
    ...ajaxParams,
    organization_workshop_uuid: organizationWorkshopUuid_alertEdit.value,
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
