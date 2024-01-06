<template>
  <q-select outlined use-input clearable v-model="sourceProjectUuid_alertEdit" :options="sourceProjects_alertEdit"
    :label="labelName" @filter="fnFilter" emit-value map-options
    :display-value="sourceProjectsMap[sourceProjectUuid_alertEdit]" />
</template>
<script setup>
import { inject, defineProps, onMounted, ref, watch } from "vue";
import collect from "collect.js";
import { ajaxGetSourceTypes } from "/src/apis/source";
import { errorNotify } from "src/utils/notify";

watch(sourceTypeUuid_alertEdit, newValue => fnSearch(newValue));

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
const sourceTypeUuid_alertEdit = inject("sourceTypeUuid_alertEdit");
const sourceProjectUuid_alertEdit = inject("sourceTypeUuid_alertEdit");
const sourceProjects_alertEdit = ref([]);
const sourceProjects = ref([]);
const sourceProjectsMap = ref({});

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      sourceProjects_alertEdit.value = sourceProjects.value;
    });
    return;
  }

  update(() => {
    sourceProjects_alertEdit.value = sourceProjects.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch(""));

const fnSearch = (sourceTypeUuid) => {
  sourceProjects.value = [];

  if (!sourceTypeUuid) return;

  ajaxGetSourceTypes(ajaxParams)
    .then(res => {
      sourceProjects.value = collect(res.content.source_types)
        .map(sourceType => {
          return {
            label: sourceType.name,
            value: sourceType.uuid,
          };
        })
        .all();
      sourceProjectsMap.value = collect(sourceProjects.value).pluck('label', 'value').all();
    })
    .catch(e=>errorNotify(e.msg));
};
</script>
src/utils/notify
