<template>
  <q-select outlined use-input clearable v-model="equipmentKindModelUuid_search" :options="equipmentKindModels_search"
    :display-value="collect(equipmentKindModels_search).pluck('value', 'label').all()[equipmentKindModelUuid_search]"
    :label="labelName" @filter="fnFilter" emit-value map-options :disable="!equipmentKindTypeUuid_search" />
</template>

<script setup>
import { ref, onMounted, inject, defineProps, watch } from "vue";
import { ajaxEquipmentKindModelList } from "/src/apis/equipmentKind";
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
const equipmentKindTypeUuid_search = inject("equipmentKindTyypeUuid_search");
const equipmentKindModelUuid_search = inject("equipmentKindTypeUuid_search");
const equipmentKindModels_search = ref([]);

watch(equipmentKindTypeUuid_search, newValue => {
  fnSearch(newValue);
});

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      equipmentKindModels_search.value = equipmentKindModels_search.value;
    });
    return;
  }

  update(() => {
    equipmentKindModels_search.value = equipmentKindModels_search.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

const fnSearch = equipmentKindTypeUuid => {
  equipmentKindModels_search.value = [];
  equipmentKindTypeUuid_search.value = "";

  if (equipmentKindTypeUuid) {
    ajaxEquipmentKindModelList({
      ...ajaxParams,
      equipment_kind_category_uuid: equipmentKindTypeUuid,
    })
      .then((res) => {
        if (res.content.equipment_kind_models.length > 0) {
          equipmentKindModels_search.value = res.content.equipment_kind_models
            .map(equipmentKindModel => {
              return {
                label: equipmentKindModel.name,
                value: equipmentKindModel.uuid,
              }
            });
        }
      })
      .catch((e) => {
        errorNotify(e.msg);
      });
  }
};

onMounted(() => {
  fnSearch("");
});
</script>
