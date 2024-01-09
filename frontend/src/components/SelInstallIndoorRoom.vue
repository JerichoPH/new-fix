<template>
  <q-select outlined use-input clearable v-model="installIndoorRoomUuid_com" :options="installIndoorRooms_com"
    :label="labelName" @filter="fnFilter" emit-value map-options
    :display-value="installIndoorRoomsMap[installIndoorRoomUuid_com]" />
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
  sign: {
    type: String,
    default: "com",
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

const installIndoorRoomUuid_com = inject(`installIndoorRoomTypeUuid_${sign}`);
const installIndoorRooms_com = ref([]);
const installIndoorRooms = ref([]);
const installIndoorRoomsMap = ref({});

const fnFilter = (val, update) => {
  if (val === "") {
    update(() => {
      installIndoorRooms_com.value = installIndoorRooms.value;
    });
    return;
  }

  update(() => {
    installIndoorRooms_com.value = installIndoorRooms.value.filter(
      (v) => v.label.toLowerCase().indexOf(val.toLowerCase()) > -1
    );
  });
};

onMounted(() => fnSearch());

const fnSearch = () => {
  installIndoorRooms_com.value = [];

  ajaxGetInstallIndoorRoomTypes(ajaxParams)
    .then((res) => {
      installIndoorRooms.value = collect(res.content.install_indoor_room_types)
        .map(installIndoorRoomType => {
          return {
            label: installIndoorRoomType.name,
            value: installIndoorRoomType.uuid,
          };
        })
        .all();
      installIndoorRoomsMap.value = collect(installIndoorRooms.value).pluck("label", "value").all();
    })
    .catch(e => errorNotify(e.msg));
};
</script>
src/utils/notify
