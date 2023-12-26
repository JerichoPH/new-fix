<template>
  <q-select outlined use-input clearable v-model="equipmentKindCategoryUuid_alertEdit" :options="options"
    :display-value="equipmentKindCategoriesMap[equipmentKindCategoryUuid_alertEdit]" :label="labelName" @filter="fnFilter"
    emit-value map-options />
</template>
<script setup>
import { inject, defineProps, onMounted, ref } from "vue";
import { ajaxGetEquipmentKindCategories } from "/src/apis/equipmentKind";
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
const equipmentKindCategoryUuid_alertEdit = inject("equipmentKindCategoryUuid_alertEdit");
const options = ref([]);
const equipmentKindCategories = ref([]);
const equipmentKindCategoriesMap = ref({});

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      options.value = equipmentKindCategories.value;
    });
    return;
  }

  update(() => {
    options.value = equipmentKindCategories.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => {
  ajaxGetEquipmentKindCategories(ajaxParams)
    .then((res) => {
      equipmentKindCategories.value =
        collect(res.content.equipment_kind_categories)
          .map((equipmentKindCategory) => {
            return {
              label: equipmentKindCategory.name,
              value: equipmentKindCategory.uuid,
            };
          })
          .all();
      equipmentKindCategoriesMap.value = collect(equipmentKindCategories.value).pluck('label', 'value').all();
    })
    .catch((e) => errorNotify(e.msg));
});
</script>
src/utils/notify
