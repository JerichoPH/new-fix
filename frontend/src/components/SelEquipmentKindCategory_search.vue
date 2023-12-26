<template>
  <q-select
    outlined
    use-input
    clearable
    v-model="equipmentKindCategoryUuid_search"
    :options="equipmentKindCategories_search"
    :display-value="
      collect(equipmentKindCategories_search).pluck('value', 'label')[
        equipmentKindCategoryUuid_search
      ]
    "
    :label="labelName"
    @filter="fnFilter"
    emit-value
    map-options
  />
</template>

<script setup>
import { ref, onMounted, inject, defineProps } from "vue";
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
const equipmentKindCategoryUuid_search = inject(
  "equipmentKindCategoryUuid_search"
);
const equipmentKindCategories_search = ref([]);
const equipmentKindCategories = ref([]);

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      equipmentKindCategories.value = equipmentKindCategories.value;
    });
    return;
  }

  update(() => {
    equipmentKindCategories.value = equipmentKindCategories.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => {
  equipmentKindCategories_search.value = [];

  ajaxGetEquipmentKindCategories(ajaxParams)
    .then((res) => {
      if (res.content.equipment_kind_categories.length > 0) {
        equipmentKindCategories.value =
          res.content.equipment_kind_categories.map((equipmentKindCategory) => {
            return {
              label: equipmentKindCategory.name,
              value: equipmentKindCategory.uuid,
            };
          });
      }
    })
    .catch((e) => errorNotify(e.msg));
});
</script>
