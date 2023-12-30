<template>
  <q-select outlined use-input clearable v-model="factoryUuid_search"
    :options="factories_search"
    :display-value="collect(factories_search).pluck('value', 'label')[factoryUuid_search]"
    :label="labelName" @filter="fnFilter" emit-value map-options />
</template>

<script setup>
import { ref, onMounted, inject, defineProps } from "vue";
import { ajaxGetFactories } from "/src/apis/factory";
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
const factoryUuid_search = inject("factoryUuid_search");
const factories_search = ref([]);
const factories = ref([]);
const factoriesMap = ref({});

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      factories.value = factories.value;
    });
    return;
  }

  update(() => {
    factories.value = factories.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  factories_search.value = [];

  ajaxGetFactories(ajaxParams)
    .then((res) => {
      factories.value = (res.content.factories)
        .map(factory => {
          return {
            label: factory.name,
            value: factory.uuid,
          };
        });
      factoriesMap.value = collect(factories.value).pluck('label', 'value').all();
    })
    .catch((e) => errorNotify(e.msg));
};
</script>
