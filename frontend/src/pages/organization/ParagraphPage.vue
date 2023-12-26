<template>
  <div class="q-pa-md">
    <q-card>
      <q-card-section>
        <div class="row">
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
                <div class="col">
                  <q-input outlined clearable lazy-rules v-model="name_search" label="名称" :rules="[]" class="q-mb-md" />
                </div>
                <div class="col">
                  <sel-organization-railway_search label-name="所属路局" :ajax-params="{}" />
                </div>
                <div class="col">
                  <sel-organization-paragraph_search label-name="所属站段" :ajax-params="{}" />
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
import { ref, onMounted, provide } from 'vue';
import collect from 'collect.js';
import {
  ajaxGetOrganizationPragraphs,
  ajaxGetOrganizationPragraph,
  ajaxStoreOrganizationPragraph,
  ajaxUpdateOrganizationPragraph,
  ajaxDestroyOrganizationPragraph,
  ajaxDestroyOrganizationPragraphs,
} from 'src/apis/organization';
import {
  loadingNotify,
  successNotify,
  errorNotify,
  actionNotify,
  destroyActions,
} from 'src/utils/notify';
import SelOrganizationRailway_search from 'src/components/SelOrganizationRailway_search.vue';
import SelOrganizationParagraph_search from 'src/components/SelOrganizationParagraph_search.vue';

// 搜索栏数据
const name_search = ref('');
const organizationRailwayUuid_search = ref('');
provide('organizationRailwayUuid_search', organizationRailwayUuid_search);
const organizationParagraphUuid_search = ref('');
provide('organizationParagraphUuid_search', organizationParagraphUuid_search);

onMounted(() => fnInit());

const fnInit = () => fnSearch();

const fnResetSearch = () => { };

const fnSearch = () => { };

</script>
src/apis/organization
