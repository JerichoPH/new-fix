<template>
  {{ values ? values.filter(val => { return val !== "" || val !== null; }).join(separator) : '' }}
</template>
<script setup>
import { defineProps } from "vue";
const props = defineProps({
  values: {
    type: Array,
    default: null,
    required: true,
  },
  separator: {
    type: String,
    default: " - ",
  },
});

const values = props.values;
const separator = props.separator;
</script>
