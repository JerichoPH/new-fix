<template>
  <q-select outlined use-input clearable v-model="installIndoorRoomUuid_com"
    :options="installIndoorRooms_com" :label="labelName" @filter="fnFilter" emit-value map-options
    :display-value="installIndoorRoomsMap[installIndoorRoomUuid_com]" :disable="!installIndoorRoomTypeUuid_com"/>
</template>
<script setup>
import { inject, defineProps, onMounted, ref, watch } from "vue";
import { ajaxGetInstallIndoorRooms } from "/src/apis/install";
import collect from "collect.js";
import { errorNotify } from "src/utils/notify";

const props = defineProps({
  labelName: {
    type: String,
    default: "",
    required: true,
  },
  sechma: {
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
const sechma = props.sechma;
const installIndoorRoomTypeUuid_com = inject(`installIndoorRoomTypeUuid_${sechma}`);
const installIndoorRoomUuid_com = inject(`installIndoorRoomUuid_${sechma}`);
const installIndoorRooms_com = ref([]);
const installIndoorRooms = ref([]);
const installIndoorRoomsMap = ref({});

watch(installIndoorRoomTypeUuid_com,()=>fnSearch());

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

  if(!installIndoorRoomTypeUuid_com.value){
    installIndoorRoomUuid_com.value = "";
    return;
  }

  ajaxGetInstallIndoorRooms(ajaxParams)
    .then((res) => {
      installIndoorRooms.value = collect(res.content.install_indoor_rooms)
        .map(installIndoorRoom => {
          return {
            label: installIndoorRoom.name,
            value: installIndoorRoom.uuid,
          };
        })
        .all();
      installIndoorRoomsMap.value = collect(installIndoorRooms.value).pluck("label", "value").all();
    })
    .catch(e => errorNotify(e.msg));
};
</script>
src/utils/notify
