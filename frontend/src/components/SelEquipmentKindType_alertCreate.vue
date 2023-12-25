<template>
  <q-select outlined use-input clearable v-model="equipmentKindTypeUuid_alertCreate" :options="options" :label="labelName"
    :option-disable="collect(equipmentKindTypes).pluck('value','label').all()[equipmentKindTypeUuid_alertCreate]"
    @filter="fnFilter" emit-value map-options />
</template>
<script setup>
import { inject, defineProps, onMounted, ref, watch } from "vue";
import collect from "collect.js"
import { ajaxGetEquipmentKindTypes } from "/src/apis/equipmentKind";
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
const equipmentKindCategoryUuid_alertCreate = inject("equipmentKindCategoryUuid_alertCreate");
const equipmentKindTypeUuid_alertCreate = inject("equipmentKindTypeUuid_alertCreate");
const options = ref([]);
const equipmentKindTypes = ref([]);

watch(equipmentKindCategoryUuid_alertCreate, newValue => {
  fnSearch(newValue);
});

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      options.value = equipmentKindTypes.value;
    });
    return;
  }

  update(() => {
    options.value = equipmentKindTypes.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => {
  fnSearch("");
});

/**
 * 搜索器材类型
 * @param {string} equipmentKindCategoryUuid
 */
const fnSearch = equipmentKindCategoryUuid => {
  equipmentKindTypes.value = [];

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
      .catch((e) => {
        errorNotify(e.msg);
      });
  }
};
</script>
src/utils/notify
