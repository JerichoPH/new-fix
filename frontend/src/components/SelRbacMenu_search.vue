<template>
  <q-select outlined use-input clearable v-model="parentUuid_search" :options="rbacMenus_search" :label="labelName"
    @filter="fnFilter" emit-value map-options />
</template>
<script setup>
import { inject, defineProps, onMounted, ref } from "vue";
import { ajaxGetRbacMenus } from "/src/apis/rbac";
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
const parentUuid_search = inject("parentUuid_search");
const rbacMenus_search = ref([]);
const rbacMenus = ref([]);

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      rbacMenus_search.value = rbacMenus.value;
    });
    return;
  }

  update(() => {
    rbacMenus_search.value = rbacMenus.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => {
  rbacMenus_search.value = [];

  ajaxGetRbacMenus(ajaxParams)
    .then((res) => {
      rbacMenus.value = collect(res.content.rbac_menus)
        .map(rbacMenu => {
          return {
            label: rbacMenu.name,
            value: rbacMenu.uuid,
          };
        });
    })
    .catch(e=>errorNotify(e.msg));
});
</script>
src/utils/notify
