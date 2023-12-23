<template>
  <q-select outlined use-input clearable v-model="equipmentKindCategoryUuid_alertEdit" :options="options"
    :display-value="collect(equipmentKindCategories).where('value', equipmentKindCategoryUuid_alertEdit).first()['label']" :filter="fnFilter"
    :label="labelName" @filter="fnFilter" emit-value map-options />
</template>
<script setup>
import { inject, defineProps, onMounted, ref } from "vue";
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
const equipmentKindCategoryUuid_alertEdit = inject("equipmentKindCategoryUuid_alertEdit");
const options = ref([]);
const equipmentKindCategories = ref([]);

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
  ajaxEquipmentKindCategoryList(ajaxParams)
    .then((res) => {
      if (res.content.equipment_kind_categories.length > 0) {
        equipmentKindCategories.value = res.content.equipment_kind_categories
          .map(equipmentKindCategory => {
            return {
              label: equipmentKindCategory.name,
              value: equipmentKindCategory.uuid,
            };
          });
      }
    })
    .catch((e) => {
      errorNotify(e.msg);
    });
});
</script>
src/utils/notify
