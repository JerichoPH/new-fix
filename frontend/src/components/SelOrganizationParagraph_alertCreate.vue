<template>
  <q-select outlined use-input clearable v-model="organizationParagraphUuid_alertCreate"
    :options="organizationParagraphs_alertCreate" :label="labelName" @filter="fnFilter" emit-value map-options
    :display-value="organizationParagraphsMap[organizationParagraphUuid_alertCreate]"
    :disable="!organizationRailwayUuid_alertCreate" />
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
const organizationRailwayUuid_alertCreate = inject("organizationRailwayUuid_alertCreate");
const organizationParagraphUuid_alertCreate = inject("organizationParagraphUuid_alertCreate");
const organizationParagraphs_alertCreate = ref([]);
const organizationParagraphs = ref([]);
const organizationParagraphsMap = ref({});

watch(organizationRailwayUuid_alertCreate, () => { fnSearch(); });

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      organizationParagraphs_alertCreate.value = organizationParagraphs.value;
    });
    return;
  }

  update(() => {
    organizationParagraphs_alertCreate.value = organizationParagraphs.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  organizationParagraphs.value = [];

  if (!organizationRailwayUuid_alertCreate.value) {
    organizationParagraphUuid_alertCreate.value = "";
    return;
  }

  if (organizationRailwayUuid) {
    ajaxGetOrganizationParagraphs({
      ...ajaxParams,
      organization_railway_uuid: organizationRailwayUuid_alertCreate.value,
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
      .catch((e) => errorNotify(e.msg));
  }
};
</script>
src/utils/notify
