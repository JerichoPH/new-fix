<template>
  {{ values ? values.filter(val => isEmpty(val)).join(sep) : '' }}
</template>
<script setup>
import { defineProps } from "vue";
import { isEmpty } from "src/utils/common";
const props = defineProps({
  values: {
    type: Array,
    default: null,
    required: true,
  },
  sep: {
    type: String,
    default: " - ",
  },
});

const values = props.values;
const sep = props.sep;
</script>
