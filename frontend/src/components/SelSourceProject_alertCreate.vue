<template>
  <q-select outlined use-input clearable v-model="sourceProjectUuid_alertCreate" :options="sourceProjects_alertCreate"
    :label="labelName" @filter="fnFilter" emit-value map-options
    :display-value="sourceProjectsMap[sourceProjectUuid_alertCreate]" />
</template>
<script setup>
import { inject, defineProps, onMounted, ref, watch } from "vue";
import collect from "collect.js";
import { ajaxGetSourceTypes } from "/src/apis/source";
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
const sourceTypeUuid_alertCreate = inject("sourceTypeUuid_alertCreate");
const sourceProjectUuid_alertCreate = inject("sourceTypeUuid_alertCreate");
const sourceProjects_alertCreate = ref([]);
const sourceProjects = ref([]);
const sourceProjectsMap = ref({});

watch(sourceTypeUuid_alertCreate, newValue => fnSearch(newValue));

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      sourceProjects_alertCreate.value = sourceProjects.value;
    });
    return;
  }

  update(() => {
    sourceProjects_alertCreate.value = sourceProjects.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch(""));

const fnSearch = (sourceTypeUuid) =>{
  sourceProjects.value = [];

  if(!sourceTypeUuid) return;

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
