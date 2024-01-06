<template>
  <q-select outlined use-input clearable v-model="organizationWorkshopUuid_alertCreate"
    :options="organizationWorkshops_alertCreate" :label="labelName" @filter="fnFilter" emit-value map-options
    :display-value="organizationWorkshopsMap[organizationWorkshopUuid_alertCreate]"
    :disable="!organizationParagraphUuid_alertCreate" />
</template>
<script setup>
import { inject, defineProps, onMounted, ref, watch } from "vue";
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
const organizationParagraphUuid_alertCreate = inject("organizationParagraphUuid_alertCreate");
const organizationWorkshopUuid_alertCreate = inject("organizationWorkshopUuid_alertCreate");
const organizationWorkshops_alertCreate = ref([]);
const organizationWorkshops = ref([]);
const organizationWorkshopsMap = ref({});

watch(organizationParagraphUuid_alertCreate, newValue => fnSearch(newValue));

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      organizationWorkshops_alertCreate.value = organizationWorkshops.value;
    });
    return;
  }

  update(() => {
    organizationWorkshops_alertCreate.value = organizationWorkshops.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  organizationWorkshops.value = [];

  if (!organizationParagraphUuid_alertCreate) {
    organizationWorkshopUuid_alertCreate.value = "";
    return;
  }

  ajaxGetOrganizationWorkshops({
    ...ajaxParams,
    organization_paragarph_uuid: organizationParagraphUuid_alertCreate.value,
  })
    .then(res => {
      organizationWorkshops.value = collect(res.content.organization_workshops)
        .map(organizationWorkshop => {
          return {
            label: organizationWorkshop.name,
            value: organizationWorkshop.uuid,
          };
        })
        .all();
      organizationWorkshopsMap.value = collect(organizationWorkshops.value).pluck('label', 'value').all();
    })
    .catch(e=>errorNotify(e.msg));
};
</script>
src/utils/notify
