<template>
  <q-select outlined use-input clearable v-model="organizationParagraphUuid_search"
    :options="organizationParagraphs_search" :display-value="organizationParagraphsMap[organizationParagraphUuid_search]"
    :label="labelName" @filter="fnFilter" emit-value map-options :disable="!organizationRailwayUuid_search" />
</template>
<script setup>
import { inject, defineProps, onMounted, ref, watch } from "vue";
import { ajaxGetOrganizationParagraphs } from "/src/apis/organization";
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
const organizationRailwayUuid_search = inject("organizationRailwayUuid_search");
const organizationParagraphUuid_search = inject("organizationParagraphUuid_search");
const organizationParagraphs_search = ref([]);
const organizationParagraphs = ref([]);
const organizationParagraphsMap = ref({});

watch(organizationRailwayUuid_search, (newValue) => fnSearch(newValue));

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      organizationParagraphs_search.value = organizationParagraphs.value;
    });
    return;
  }

  update(() => {
    organizationParagraphs_search.value = organizationParagraphs.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch(""));

const fnSearch = (organizationRailwayUuid) => {
  organizationParagraphs_search.value = [];

  if (!organizationRailwayUuid) return;
  ajaxGetOrganizationParagraphs({
    ...ajaxParams,
    organization_railway_uuid: organizationRailwayUuid,
  })
    .then((res) => {
      organizationParagraphs.value =
        collect(res.content.organization_paragraphs)
          .map(organizationParagraph => {
            return {
              label: organizationParagraph.name,
              value: organizationParagraph.uuid,
            };
          })
          .all();
      organizationParagraphsMap.value = collect(organizationParagraphs.value).pluck('label', 'value').all();
    })
    .catch((e) => errorNotify(e.msg));
};
</script>
src/utils/notify
