<template>
  <q-select outlined use-input clearable v-model="equipmentKindCategoryUuid_search"
    :options="equipmentKindCategories_search"
    :display-value="equipmentKindCategories_search.find((val) => val === equipmentKindCategoryUuid_search)"
    :label="labelName" @filter="fnFilter" emit-value map-options />
</template>

<script setup>
import { ref, onMounted, inject, defineProps } from "vue";
import { ajaxEquipmentKindCategoryList } from "/src/apis/equipmentKind";
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
const equipmentKindCategoryUuid_search = inject("rbacRoleUuid_search");
const equipmentKindCategories_search = ref([]);
const equipmentKindCategories = ref([]);

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      equipmentKindCategories_search.value = equipmentKindCategories.value;
    });
    return;
  }

  update(() => {
    equipmentKindCategories_search.value = equipmentKindCategories.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => {
  ajaxEquipmentKindCategoryList(ajaxParams)
    .then((res) => {
      if (res.content.rbac_roles.length > 0) {
        collect(res.content.equipment_kind_categories)
          .each((equipmentKindCategory) => {
            equipmentKindCategories.value.push({
              label: equipmentKindCategory.name,
              value: equipmentKindCategory.uuid,
            });
          });
      }
    })
    .catch((e) => {
      errorNotify(e.msg);
    });
});
</script>
