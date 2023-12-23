<template>
  <div class="q-pa-md">
    <q-table
      flat
      bordered
      title="Treats"
      :rows="rows"
      row-key="name"
      selection="multiple"
      v-model:selected="selected"
      :selected-rows-label="getSelectedString"
    >
      <template v-slot:header="props">
        <q-tr :props="props">
          <q-td><q-checkbox key="allCheck" v-model="props.selected" /></q-td>
          <q-td>#</q-td>
          <q-td>名称</q-td>
          <q-td>路由</q-td>
        </q-tr>
      </template>
      <template v-slot:body="props">
        <q-tr :props="props">
          <q-td
            ><q-checkbox
              :key="props.row.uuid"
              :value="props.row.uuid"
              v-model="props.selected"
          /></q-td>
          <q-td>{{ props.row.index }}</q-td>
          <q-td>{{ props.row.name }}</q-td>
          <q-td>{{ props.row.uri }}</q-td>
        </q-tr>
      </template>
    </q-table>

    <div class="q-mt-md">Selected: {{ JSON.stringify(selected) }}</div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { ajaxRbacPermissionList } from "src/apis/rbac";
import collect from "collect.js";

const rows = ref([]);
const selected = ref([]);

onMounted(() => {
  fnInit();
});

const fnInit = () => {
  ajaxRbacPermissionList().then((res) => {
    collect(res.content.rbac_permissions).each((rbacPermission, idx) => {
      rows.value.push({
        index: idx + 1,
        uuid: rbacPermission.uuid,
        name: rbacPermission.name,
        uri: rbacPermission.uri,
      });
    });
  });
};

const getSelectedString = () => {};
</script>
