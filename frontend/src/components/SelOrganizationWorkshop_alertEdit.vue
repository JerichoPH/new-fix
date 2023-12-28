<template>
  <q-select outlined use-input clearable v-model="organizationWorkshopUuid_alertEdit"
    :options="organizationWorkshops_alertEdit"
    :display-value="organizationWorkshopsMap[organizationWorkshopUuid_alertEdit]" :label="labelName" @filter="fnFilter"
    emit-value map-options />
</template>
<script setup>
import { inject, defineProps, onMounted, ref } from "vue";
import collect from "collect.js";
import { ajaxGetOrganizationWorkshops } from "/src/apis/organization";
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
const organizationWorkshops_alertEdit = ref([]);
const organizationWorkshops = ref([]);
const organizationWorkshopsMap = ref({});

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      organizationWorkshops_alertEdit.value = organizationWorkshops.value;
    });
    return;
  }

  update(() => {
    organizationWorkshops_alertEdit.value = organizationWorkshops.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch(""));

const fnSearch = organizationParagraphUuid => {
  organizationWorkshops.value = [];

  if (organizationParagraphUuid) ajaxParams["organization_paragraph_uuid"] = organizationParagraphUuid;

  ajaxGetOrganizationWorkshops(ajaxParams)
    .then((res) => {
      organizationWorkshops.value = collect(res.content.organization_workshops)
        .map(organizationWorkshop => {
          return {
            label: organizationWorkshop.name,
            value: organizationWorkshop.uuid,
          };
        })
        .all();
        organizationWorkshopsMap.value = collect(organizationWorkshops.value).pluck('label','value').all();
    })
    .catch((e) => {
      console.error('err', e);
      errorNotify(e);
    });
};
</script>
src/utils/notify
