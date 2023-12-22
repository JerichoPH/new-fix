<template>
  <q-select
    outlined
    use-input
    clearable
    v-model="rbacRoleUuid_search"
    :options="rbacRoles_search"
    :display-value="rbacRoles.find((val) => val === rbacRoleUuid_search)"
    :label="labelName"
    @filter="fnFilter"
    emit-value
    map-options
  />
</template>
<script setup>
import { inject, defineProps, onMounted, ref } from "vue";
import { ajaxRbacRoleList } from "/src/apis/rbac";
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
const rbacRoleUuid_search = inject("rbacRoleUuid_search");
const rbacRoles_search = ref([]);
const rbacRoles = ref([]);

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      rbacRoles_search.value = rbacRoles.value;
    });
    return;
  }

  update(() => {
    rbacRoles_search.value = rbacRoles.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => {
  ajaxRbacRoleList(ajaxParams)
    .then((res) => {
      if (res.content.rbac_roles.length > 0) {
        collect(res.content.rbac_roles).each((rbacRole) => {
          rbacRoles.value.push({
            label: rbacRole.name,
            value: rbacRole.uuid,
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
