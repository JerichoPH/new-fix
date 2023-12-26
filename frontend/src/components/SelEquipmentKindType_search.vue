<template>
  <q-select outlined use-input clearable v-model="equipmentKindTypeUuid_search" :options="equipmentKindTypes_search"
    :display-value="collect(equipmentKindTypes_search).pluck('value', 'label').all()[equipmentKindTypeUuid_search]"
    :label="labelName" @filter="fnFilter" emit-value map-options :disable="!equipmentKindCategoryUuid_search" />
</template>

<script setup>
import { ref, onMounted, inject, defineProps, watch } from "vue";
import { ajaxGetEquipmentKindTypes } from "/src/apis/equipmentKind";
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
const equipmentKindCategoryUuid_search = inject("equipmentKindCategoryUuid_search");
const equipmentKindTypeUuid_search = inject("equipmentKindTypeUuid_search");
const equipmentKindTypes_search = ref([]);
const equipmentKindTypes = ref([]);

watch(equipmentKindCategoryUuid_search, newValue => {
  fnSearch(newValue);
});

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      equipmentKindTypes_search.value = equipmentKindTypes.value;
    });
    return;
  }

  update(() => {
    equipmentKindTypes_search.value = equipmentKindTypes.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => {
  fnSearch("");
});

const fnSearch = equipmentKindCategoryUuid => {
  equipmentKindTypes_search.value = [];

  if (equipmentKindCategoryUuid) {
    ajaxGetEquipmentKindTypes({
      ...ajaxParams,
      equipment_kind_category_uuid: equipmentKindCategoryUuid,
    })
      .then((res) => {
        if (res.content.equipment_kind_types.length > 0) {
          equipmentKindTypes.value = res.content.equipment_kind_types
            .map(equipmentKindType => {
              return {
                label: equipmentKindType.name,
                value: equipmentKindType.uuid,
              }
            });
        }
      })
      .catch((e) => errorNotify(e.msg));
  }
};
</script>
