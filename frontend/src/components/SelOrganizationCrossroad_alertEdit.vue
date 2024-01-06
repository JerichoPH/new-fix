<template>
  <q-select outlined use-input clearable v-model="organizationCrossroadUuid_alertEdit"
    :options="organizationCrossroads_search" :display-value="organizationCrossroadsMap[organizationCrossroadUuid_alertEdit]"
    :label="labelName" @filter="fnFilter" emit-value map-options :disable="!organizationWorkshopUuid_alertEdit" />
</template>
<script setup>
import { inject, defineProps, onMounted, ref, watch } from "vue";
import { ajaxGetOrganizationCrossroads } from "/src/apis/organization";
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
const organizationCrossroadUuid_alertEdit = inject("organizationCrossroadUuid_alertEdit");
const organizationCrossroads_search = ref([]);
const organizationCrossroads = ref([]);
const organizationCrossroadsMap = ref({});

watch(organizationWorkshopUuid_alertEdit, () => fnSearch());

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      organizationCrossroads_search.value = organizationCrossroads.value;
    });
    return;
  }

  update(() => {
    organizationCrossroads_search.value = organizationCrossroads.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  organizationCrossroads_search.value = [];

  if (!organizationWorkshopUuid_alertEdit.value) {
    organizationCrossroadUuid_alertEdit.value = "";
    return;
  }

  ajaxGetOrganizationCrossroads({
    ...ajaxParams,
    organization_workshop_uuid: organizationWorkshopUuid_alertEdit.value,
  })
    .then((res) => {
      organizationCrossroads.value =
        collect(res.content.organization_crossroads)
          .map(organizationCrossroad => {
            return {
              label: organizationCrossroad.name,
              value: organizationCrossroad.uuid,
            };
          })
          .all();
      organizationCrossroadsMap.value = collect(organizationCrossroads.value).pluck('label', 'value').all();
    })
    .catch(e=>errorNotify(e.msg));
};
</script>
src/utils/notify
