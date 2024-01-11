<template>
  <q-select outlined use-input clearable v-model="currentValue" :options="l_com" :label="labelName" @filter="fnFilter"
    emit-value map-options :display-value="m[currentValue]" :disable="parentField ? !parentValue : false"
    :key="`${currentField}_${sechma}`" />
</template>
<script setup>
import { inject, defineProps, onMounted, ref, watch } from "vue";
import { getVal } from "src/utils/common";
import collect from "collect.js";
import { errorNotify } from "src/utils/notify";

const props = defineProps({
  labelName: {
    type: String,
    default: "",
    required: true,
  },
  dataSource: {
    type: Function,
    default: null,
    required: true,
  },
  dataSourceField: {
    type: String,
    default: "",
    required: true,
  },
  currentField: {
    type: String,
    default: "",
    required: true,
  },
  labelField: {
    type: String,
    default: "name",
    required: false,
  },
  valueField: {
    type: String,
    default: "uuid",
    required: false,
  },
  sechma: {
    type: String,
    default: "",
    required: true,
  },
  parentField: {
    type: String,
    default: "",
    required: false,
  },
  ajaxParams: {
    type: Object,
    default: () => {
      return {};
    },
    required: false,
  },
});

const labelName = props.labelName;
const ajaxParams = props.ajaxParams;
const sechma = props.sechma;
const dataSource = props.dataSource;
const dataSourceField = props.dataSourceField;
const currentField = props.currentField;
const labelField = props.labelField;
const valueField = props.valueField;
const parentField = props.parentField;

let parentValue = null;
if (parentField) {
  parentValue = inject(`${parentField}_${sechma}`);
  watch(parentValue, () => fnSearch());
};

const currentValue = inject(`${currentField}_${sechma}`);
const l_com = ref([]);
const l = ref([]);
const m = ref({});

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      l_com.value = l.value;
    });
    return;
  }

  update(() => {
    l_com.value = l.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  l_com.value = [];

  // console.log(currentField, parentField, parentValue, currentValue);

  if (parentField) {
    if (!parentValue.value) {
      currentValue.value = "";
      return;
    }
  }

  dataSource(ajaxParams)
    .then(res => {
      l.value = collect(getVal(res, `content.${dataSourceField}`))
        .map(item => {
          return {
            label: getVal(item, labelField),
            value: getVal(item, valueField),
          };
        })
        .all();
      m.value = collect(l.value).pluck("label", "value").all();
    })
    .catch(e => {
      console.error('err', e);
      errorNotify(e.msg)
    });
};
</script>
