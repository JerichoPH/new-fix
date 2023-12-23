<template>
  <div class="q-pa-md">
    <q-card>
      <q-card-section>
        <div class="q-mb-md">
          <div class="col"><span :style="{ fontSize: '20px' }">搜索</span></div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="primary" label="搜索" icon="search" @click="fnSearch" />
              <q-btn color="primary" label="重置" flat @click="fnResetSearch" />
            </q-btn-group>
          </div>
        </div>
        <div class="row">
          <div class="col">
            <q-form>
              <div class="row q-pb-sm q-col-gutter-sm">
                <div class="col-3">
                  <q-input outlined clearable lazy-rules v-model="name_search" label="名称" :rules="[]" class="q-mb-md" />
                </div>
                <div class="col-3">
                  <sel-equipment-kind-category_search label-name="器材种类" :ajax-params="{}" />
                </div>
                <div class="col-3">
                  <sel-equipment-kind-type_search label-name="器材类型" :ajax-params="{}" />
                </div>
              </div>
            </q-form>
          </div>
        </div>
      </q-card-section>
    </q-card>

    <q-card class="q-mt-md">
      <q-card-section>

      </q-card-section>
    </q-card>
  </div>
</template>

<script setup>
import { ref, onMounted, provide } from "vue";
import {
  ajaxEquipmentKindModelList,
  ajaxEquipmentKindModelDetail,
  ajaxEquipmentKindModelStore,
  ajaxEquipmentKindModelUpdate,
  ajaxEquipmentKindModelDestroy,
  ajaxEquipmentKindModelDestroyMany,
} from "/src/apis/equipmentKind";
import { notifies } from "/src/utils/notify";
import SelEquipmentKindCategory_search from "src/components/SelEquipmentKindCategory_search.vue";
import SelEquipmentKindType_search from "src/components/SelEquipmentKindType_search.vue";

// 搜索栏数据
const name_search = ref("");
const equipmentKindCategoryUuid_search = ref("");
const equipmentKindTypeUuid_search = ref("");
provide("equipmentKindCategoryUuid_search",equipmentKindCategoryUuid_search);
provide("equipmentKindTypeUuid_search",equipmentKindTypeUuid_search);


</script>
