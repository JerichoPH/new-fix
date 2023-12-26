<template>
  <q-select outlined use-input clearable v-model="parentUuid_alertEdit" :options="options"
    :display-value="collect(rbacMenus).where('value', parentUuid_alertEdit).first()['label']" :filter="fnFilter"
    :label="labelName" @filter="fnFilter" emit-value map-options />
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
const parentUuid_alertEdit = inject("parentUuid_alertEdit");
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
  ajaxGetRbacMenus(ajaxParams)
    .then((res) => {
      if (res.content.rbac_menus.length > 0) {
        rbacMenus.value = res.content.rbac_menus
          .map(rbacMenu => {
            return {
              label: rbacMenu.name,
              value: rbacMenu.uuid,
            };
          });
      }
    })
    .catch((e) => errorNotify(e.msg));
});
</script>
src/utils/notify
