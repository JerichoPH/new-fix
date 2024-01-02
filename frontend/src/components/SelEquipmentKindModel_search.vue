<template>
  <q-select outlined use-input clearable v-model="equipmentKindModelUuid_search" :options="equipmentKindModels_search"
    :display-value="collect(equipmentKindModels_search).pluck('value', 'label').all()[
      equipmentKindModelUuid_search
    ]
      " :label="labelName" @filter="fnFilter" emit-value map-options :disable="!equipmentKindTypeUuid_search" />
</template>

<script setup>
import { ref, onMounted, inject, defineProps, watch } from "vue";
import { ajaxGetEquipmentKindModels } from "/src/apis/equipmentKind";
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
const equipmentKindModels = ref([]);

watch(equipmentKindTypeUuid_search, () => fnSearch());

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      equipmentKindModels_search.value = equipmentKindModels.value;
    });
    return;
  }

  update(() => {
    equipmentKindModels_search.value = equipmentKindModels.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  equipmentKindModels_search.value = [];

  if (!equipmentKindTypeUuid_search.value) {
    equipmentKindModelUuid_search.value = "";
    return;
  }

  ajaxGetEquipmentKindModels({
    ...ajaxParams,
    equipment_kind_category_uuid: equipmentKindTypeUuid_search.value,
  })
    .then((res) => {
      if (res.content.equipment_kind_models.length > 0) {
        equipmentKindModels.value = res.content.equipment_kind_models.map(
          (equipmentKindModel) => {
            return {
              label: equipmentKindModel.name,
              value: equipmentKindModel.uuid,
            };
          }
        );
      }
    })
    .catch((e) => errorNotify(e.msg));
};
</script>
