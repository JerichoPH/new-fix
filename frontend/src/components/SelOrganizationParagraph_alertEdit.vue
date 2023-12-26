<template>
  <q-select outlined use-input clearable v-model="organizationParagraphUuid_alertEdit"
    :options="organizationParagraphs_alertEdit" :display-value="organizationParagraphsMap[organizationParagraphUuid_alertEdit]"
    :label="labelName" @filter="fnFilter" emit-value map-options />
</template>
<script setup>
import { inject, defineProps, onMounted, ref } from "vue";
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
const organizationParagraphUuid_alertEdit = inject("organizationParagraphUuid_alertEdit");
const organizationParagraphs_alertEdit = ref([]);
const organizationParagraphs = ref([]);
const organizationParagraphsMap = ref({});

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

onMounted(() => {
  ajaxGetOrganizationParagraphs(ajaxParams)
    .then((res) => {
      organizationParagraphs.value = collect(res.content.organization_paragraphs)
        .map((organizationParagraph) => {
          return {
            label: organizationParagraph.name,
            value: organizationParagraph.uuid,
          };
        })
        .all();
      organizationParagraphsMap.value = collect(organizationParagraphs.value).pluck('label', 'value').all();
    })
    .catch((e) => errorNotify(e.msg));
});
</script>
src/utils/notify
