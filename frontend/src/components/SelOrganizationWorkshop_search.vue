<template>
  <q-select outlined use-input clearable v-model="organizationWorkshopUuid_search"
    :options="organizationWorkshops_search" :display-value="organizationWorkshopsMap[organizationWorkshopUuid_search]"
    :label="labelName" @filter="fnFilter" emit-value map-options :disable="!organizationParagraphUuid_search" />
</template>
<script setup>
import { inject, defineProps, onMounted, ref, watch } from "vue";
import { ajaxGetOrganizationWorkshops } from "/src/apis/organization";
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
const organizationParagraphUuid_search = inject("organizationParagraphUuid_search");
const organizationWorkshopUuid_search = inject("organizationWorkshopUuid_search");
const organizationWorkshops_search = ref([]);
const organizationWorkshops = ref([]);
const organizationWorkshopsMap = ref({});

watch(organizationParagraphUuid_search, (newValue) => fnSearch(newValue));

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      organizationWorkshops_search.value = organizationWorkshops.value;
    });
    return;
  }

  update(() => {
    organizationWorkshops_search.value = organizationWorkshops.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  organizationWorkshops_search.value = [];

  if (!organizationParagraphUuid_search.value) {
    organizationWorkshopUuid_search.value = "";
    return;
  }

  ajaxGetOrganizationWorkshops({
    ...ajaxParams,
    organization_paragraph_uuid: organizationParagraphUuid_search.value,
  })
    .then((res) => {
      organizationWorkshops.value =
        collect(res.content.organization_workshops)
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
