<template>
  <q-select outlined use-input clearable v-model="sourceTypeUuid_alertEdit" :options="sourceTypes_alertEdit"
    :label="labelName" @filter="fnFilter" emit-value map-options
    :display-value="sourceTypesMap[sourceTypeUuid_alertEdit]" />
</template>
<script setup>
import { inject, defineProps, onMounted, ref } from "vue";
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
const sourceTypeUuid_alertEdit = inject("sourceTypeUuid_alertEdit");
const sourceTypes_alertEdit = ref([]);
const sourceTypes = ref([]);
const sourceTypesMap = ref({});

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      sourceTypes_alertEdit.value = sourceTypes.value;
    });
    return;
  }

  update(() => {
    sourceTypes_alertEdit.value = sourceTypes.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => {
  sourceTypes.value = [];

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
      sourceTypesMap.value = collect(sourceTypes.value).pluck('label', 'value').all();
    })
    .catch(e=>errorNotify(e.msg));
});
</script>
src/utils/notify
