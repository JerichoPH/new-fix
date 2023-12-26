<template>
  <q-select outlined use-input clearable v-model="organizationParagraphUuid_alertCreate"
    :options="organizationParagraphs_alertCreate" :label="labelName" @filter="fnFilter" emit-value map-options />
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
const organizationParagraphUuid_alertCreate = inject("organizationParagraphUuid_alertCreate");
const organizationParagraphs_alertCreate = ref([]);
const organizationParagraphs = ref([]);

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

onMounted(() => {
  organizationParagraphs.value = [];

  ajaxGetOrganizationParagraphs(ajaxParams)
    .then(res => {
      if (res.content.organization_paragraphs.length > 0) {
        organizationParagraphs.value = collect(res.content.organization_paragraphs)
          .map(organizationParagraph => {
            return {
              label: organizationParagraph.name,
              value: organizationParagraph.uuid,
            };
          })
          .all();
      }
    })
    .catch((e) => errorNotify(e.msg));
});
</script>
src/utils/notify
