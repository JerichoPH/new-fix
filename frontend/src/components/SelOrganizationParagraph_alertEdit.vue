<template>
  <q-select outlined use-input clearable v-model="organizationParagraphUuid_alertEdit"
    :options="organizationParagraphs_alertEdit" :label="labelName" @filter="fnFilter"
    :display-value="organizationParagraphsMap[organizationParagraphUuid_alertEdit]" emit-value map-options
    :disable="!organizationRailwayUuid_alertEdit"/>
</template>
<script setup>
import { inject, defineProps, onMounted, ref, watch } from "vue";
import collect from "collect.js";
import { ajaxGetOrganizationParagraphs } from "/src/apis/organization";
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
const organizationRailwayUuid_alertEdit = inject("organizationRailwayUuid_alertEdit");
const organizationParagraphUuid_alertEdit = inject("organizationParagraphUuid_alertEdit");
const organizationParagraphs_alertEdit = ref([]);
const organizationParagraphs = ref([]);
const organizationParagraphsMap = ref({});

watch(organizationRailwayUuid_alertEdit, newValue => {fnSearch()});

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      organizationParagraphs_alertEdit.value = organizationParagraphs.value;
    });
    return;
  }

  update(() => {
    organizationParagraphs_alertEdit.value = organizationParagraphs.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  organizationParagraphs.value = [];

  if(!organizationRailwayUuid_alertEdit.value) {
    organizationParagraphUuid_alertEdit.value = "";
    return;
  }

  ajaxGetOrganizationParagraphs({
    ...ajaxParams,
    organization_railway_uuid: organizationRailwayUuid_alertEdit.value,
  })
    .then(res => {
      organizationParagraphs.value = collect(res.content.organization_paragraphs)
        .map(organizationParagraph => {
          return {
            label: organizationParagraph.name,
            value: organizationParagraph.uuid,
          };
        })
        .all();
      organizationParagraphsMap.value = collect(organizationParagraphs.value).pluck('label', 'value').all();
    })
    .catch(e=>errorNotify(e.msg));
};
</script>
src/utils/notify
