<template>
  <q-select outlined use-input clearable v-model="sourceTypeUuid_search" :options="sourceTypes_search"
    :label="labelName" @filter="fnFilter" emit-value map-options />
</template>
<script setup>
import { inject, defineProps, onMounted, ref } from "vue";
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
const sourceTypes_search = ref([]);
const sourceTypes = ref([]);

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      sourceTypes_search.value = sourceTypes.value;
    });
    return;
  }

  update(() => {
    sourceTypes_search.value = sourceTypes.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  sourceTypes_search.value = [];

  ajaxGetSourceTypes(ajaxParams)
    .then(res => {
      sourceTypes.value = collect(res.content.source_types)
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
