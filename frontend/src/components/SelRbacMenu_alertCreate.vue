<template>
  <q-select
    outlined
    use-input
    clearable
    v-model="parentUuid_alertCreate"
    :options="options"
    :label="labelName"
    @filter="fnFilter"
    emit-value
    map-options
  />
</template>
<script setup>
import { inject, defineProps, onMounted, ref } from "vue";
import { ajaxRbacMenuList } from "/src/apis/rbac";
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
const parentUuid_alertCreate = inject("parentUuid_alertCreate");
const options = ref([]);
const rbacMenus = ref([]);

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      options.value = rbacMenus.value;
    });
    return;
  }

  update(() => {
    options.value = rbacMenus.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => {
  ajaxRbacMenuList(ajaxParams)
    .then((res) => {
      if (res.content.rbac_menus.length > 0) {
        collect(res.content.rbac_menus).each((rbacMenu) => {
          rbacMenus.value.push({
            label: rbacMenu.name,
            value: rbacMenu.uuid,
          });
        });
      }
    })
    .catch((e) => {
      errorNotify(e.msg);
    });
});
</script>
src/utils/notify
