<template>
  <q-select outlined use-input clearable v-model="equipmentKindTypeUuid_search" :options="equipmentKindTypes_search"
    :display-value="equipmentKindTypes_search.find((val) => val === equipmentKindTypeUuid_search)" :label="labelName"
    @filter="fnFilter" emit-value map-options />
</template>

<script setup>
import { ref, onMounted, inject, defineProps,watch } from "vue";
import { ajaxEquipmentKindCategoryList } from "/src/apis/equipmentKind";
import { errorNotify } from "src/utils/notify";

watch(equipmentKindCategoryUuid_search, (newVal, oldVal) => {
  fnSearch(newVal);
});

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
const equipmentKindCategoryUuid_search = inject("equipmentKindCategoryUuid_search");
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
  fnSearch(equipmentKindCategoryUuid_search.value);
});

/**
 * 根据器材种类代码查询器材类型
 * @param {string} equipmentKindCategoryUuid
 */
const fnSearch = equipmentKindCategoryUuid => {
  equipmentKindTypes_search.value = [];

  if (equipmentKindCategoryUuid) {
    ajaxEquipmentKindCategoryList({
      ...ajaxParams,
      equipment_kind_category_uuid: equipmentKindCategoryUuid,
    })
      .then((res) => {
        if (res.content.equipment_kind_categories.length > 0) {
          equipmentKindTypes_search.value = res.content.equipment_kind_typs
            .map(equipmentKindType => {
              return {
                label: equipmentKindType.name,
                value: equipmentKindType.uuid,
              };
            });
        }
      })
      .catch((e) => {
        errorNotify(e.msg);
      });
  }
}
</script>
