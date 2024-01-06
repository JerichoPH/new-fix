<template>
  <q-select outlined use-input clearable v-model="installIndoorRoomTypeUuid_search"
    :options="installIndoorRoomTypes_search" :label="labelName" @filter="fnFilter" emit-value map-options
    :display-value="installIndoorRoomTypesMap[installIndoorRoomTypeUuid_search]" />
</template>
<script setup>
import { inject, defineProps, onMounted, ref, watch } from "vue";
import { ajaxGetInstallIndoorRoomTypes } from "/src/apis/install";
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
const installIndoorRoomTypeUuid_search = inject("installIndoorRoomTypeUuid_search");
const installIndoorRoomTypes_search = ref([]);
const installIndoorRoomTypes = ref([]);
const installIndoorRoomTypesMap = ref({});

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      installIndoorRoomTypes_search.value = installIndoorRoomTypes.value;
    });
    return;
  }

  update(() => {
    installIndoorRoomTypes_search.value = installIndoorRoomTypes.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  installIndoorRoomTypes_search.value = [];

  ajaxGetInstallIndoorRoomTypes(ajaxParams)
    .then((res) => {
      installIndoorRoomTypes.value = collect(res.content.install_indoor_room_types)
        .map(installIndoorRoomType => {
          return {
            label: installIndoorRoomType.name,
            value: installIndoorRoomType.uuid,
          };
        })
        .all();
      installIndoorRoomTypesMap.value = collect(installIndoorRoomTypes.value).pluck("label", "value").all();
    })
    .catch(e => errorNotify(e.msg));
};
</script>
src/utils/notify
