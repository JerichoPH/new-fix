<template>
  <q-select outlined use-input clearable v-model="equipmentKindTypeUuid_search"
    :options="equipmentKindTypes_search"
    :display-value="collect(equipmentKindTypes_search).pluck('value','label').all()[equipmentKindTypeUuid_search]"
    :label="labelName" @filter="fnFilter" emit-value map-options />
</template>

<script setup>
import { ref, onMounted, inject, defineProps } from "vue";
import { ajaxEquipmentKindTypeList } from "/src/apis/equipmentKind";
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
const equipmentKindTypeUuid_search = inject("equipmentKindTypeUuid_search");
const equipmentKindTypes_search = ref([]);
const equipmentKindCategories = ref([]);

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      equipmentKindTypes_search.value = equipmentKindCategories.value;
    });
    return;
  }

  update(() => {
    equipmentKindTypes_search.value = equipmentKindCategories.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => {
  ajaxEquipmentKindTypeList(ajaxParams)
    .then((res) => {
      if (res.content.equipment_kind_types.length > 0) {
        equipmentKindTypes_search.value = res.content.equipment_kind_types
        .map(equipmentKindType=>{
          return {
            label: equipmentKindType.name,
            value: equipmentKindType.uuid,
          }
        });
      }
      console.log('equipmentKindTypes_search',equipmentKindTypes_search.value);
    })
    .catch((e) => {
      console.error('err',e);
      errorNotify(e.msg);
    });
});
</script>
