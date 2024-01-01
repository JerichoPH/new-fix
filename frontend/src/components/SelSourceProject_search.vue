<template>
  <q-select outlined use-input clearable v-model="sourceProjectUuid_search" :options="sourceProjects_search"
    :label="labelName" @filter="fnFilter" emit-value map-options />
</template>
<script setup>
import { inject, defineProps, onMounted, ref, watch } from "vue";
import { ajaxGetSourceTypes } from "/src/apis/source";
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
const sourceTypeUuid_search = inject("sourceTypeUuid_search");
const sourceProjectUuid_search = inject("sourceTypeUuid_search");
const sourceProjects_search = ref([]);
const sourceProjects = ref([]);

watch(sourceTypeUuid_search, newValue => fnSearch(newValue));

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      sourceProjects_search.value = sourceProjects.value;
    });
    return;
  }

  update(() => {
    sourceProjects_search.value = sourceProjects.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch(""));

const fnSearch = (sourceTypeUuid) => {
  sourceProjects_search.value = [];

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
    })
    .catch((e) => errorNotify(e.msg));
};
</script>
src/utils/notify
