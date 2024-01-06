<template>
  <q-card flat bordered>
    <q-card-section>
      <p class="text-body1">{{ labelName }}</p>
      <q-separator class="q-mb-md" />

      <div class="row" v-for="(items, idx) in rbacRoles" :key="idx">
        <div class="col-3" v-for="(rbacRole, rbacRoleIdx) in items" :key="rbacRoleIdx">
          <q-checkbox v-model="checkedRbacRoleUuids_alertEdit" :val="rbacRole.uuid" :key="rbacRole.uuid"
            :label="rbacRole.name" />
        </div>
      </div>
    </q-card-section>
  </q-card>
</template>
<script setup>
import { ref, onMounted, inject, defineProps } from "vue";
import collect from "collect.js";
import { ajaxGetRbacRoles } from "src/apis/rbac";
import { errorNotify } from "src/utils/notify";

const props = defineProps({
  labelName: {
    type: String,
    default: "",
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
const rbacRoles = ref([]);
const checkedRbacRoleUuids_alertEdit = inject("checkedRbacRoleUuids_alertEdit");

onMounted(() => {
  fnGetRbacRoles();
});

const fnGetRbacRoles = () => {
  ajaxGetRbacRoles(ajaxParams)
    .then((res) => {
      if (res.content.rbac_roles.length > 0) {
        rbacRoles.value = collect(res.content.rbac_roles).chunk(4).all();
      }
    })
    .catch(e=>errorNotify(e.msg));
};
</script>
src/utils/notify
