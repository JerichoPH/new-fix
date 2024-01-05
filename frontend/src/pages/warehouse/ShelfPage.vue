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
        <div class="row q-mt-md">
          <div class="col">
            <q-form>
              <div class="row q-col-gutter-sm">
                <div class="col">
                  <q-input outlined clearable lazy-rules v-model="name_search" label="名称" :rules="[]" class="q-mb-md" />
                </div>
                <div class="col">
                  <sel-organization-railway_search label-name="所属路局" />
                </div>
                <div class="col">
                  <sel-organization-paragraph_search label-name="所属站段" />
                </div>
                <div class="col">
                  <sel-organization-workshop_search label-name="所属车间" />
                </div>
                <div class="col">
                  <sel-organization-work-area_search label-name="所属工区" />
                </div>
              </div>
              <div class="row q-col-gutter-sm">
                <div class="col">
                  <sel-warehouse-storehouse_search label-name="所属库房" />
                </div>
                <div class="col">
                  <sel-warehouse-area_search label-name="所属库区" />
                </div>
                <div class="col">
                  <sel-warehouse-platoon_search label-name="所属排" />
                </div>
                <div class="col"></div>
                <div class="col"></div>
              </div>
            </q-form>
          </div>
        </div>
      </q-card-section>
    </q-card>

    <q-card class="q-mt-md">
      <q-card-section>
        <div class="row">
          <div class="col"><span :style="{ fontSize: '20px' }">仓库-架列表</span></div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn color="secondary" label="新建仓库-架" icon="add" @click="fnOpenAlertCreateWarehouseShelf" />
              <q-btn color="negative" label="删除仓库-架" icon="deleted" @click="fnDestroyCreateWarehouseShelves" />
            </q-btn-group>
          </div>
        </div>
        <div class="row q-mt-md">
          <div class="col">
            <q-table flat bordered title="" :rows="rows" row-key="uuid" :pagination="{ rowsPerPage: 200 }"
              :rows-per-page-options="[50, 100, 200, 0]" rows-per-page-label="分页" :selected-rows-label="() => { }"
              selection="multiple" v-model:selected="selected">
              <template v-slot:header="props">
                <q-tr :props="props">
                  <q-th align="left"><q-checkbox key="allCheck" v-model="props.selected" /></q-th>
                  <q-th align="left">#</q-th>
                  <q-th align="left" key="name" @click="(event) => fnColumnReverseSort(event, props, sortBy)">
                    名称
                  </q-th>
                  <q-th align="left" key="warehousePlatoon" @click="(event) => fnColumnReverseSort(event, props, sortBy)">
                    所属仓库-排
                  </q-th>
                  <q-th align="right"></q-th>
                </q-tr>
              </template>
              <template v-slot:body="props">
                <q-tr :props="props">
                  <q-td><q-checkbox :key="props.row.uuid" :value="props.row.uuid" v-model="props.selected" /></q-td>
                  <q-td>{{ props.row.index }}</q-td>
                  <q-td key="name" :props="props">
                    <a href="javascript:" @click="fnOpenAlertEditWarehouseShelfSubs(props.row.operation)">{{
                      props.row.name }}</a>
                  </q-td>
                  <q-td key="warehousePlatoon" :props="props">
                    <join-string :values="[
                      props.row.warehouseStorehouse.name,
                      props.row.warehouseArea.name,
                      props.row.warehousePlatoon.name,
                    ]" />
                  </q-td>
                  <q-td key="operation" :props="props">
                    <q-btn-group>
                      <q-btn @click="fnOpenAlertEditWarehouseShelf(props.row.operation)" color="warning" icon="edit">
                        编辑
                      </q-btn>
                      <q-btn @click="fnDestroyCreateWarehouseShelf(props.row.operation)" color="negative" icon="delete">
                        删除
                      </q-btn>
                    </q-btn-group>
                  </q-td>
                </q-tr>
              </template>
            </q-table>
          </div>
        </div>
      </q-card-section>
    </q-card>
  </div>

  <!-- 弹窗 -->
  <!-- 新建仓库-架弹窗 -->
  <q-dialog v-model="alertCreateWarehouseShelf" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">新建仓库-架</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnStoreWarehouseShelf">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertCreateWarehouseShelf" label="名称" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-railway_alert-create label-name="所属路局" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-paragraph_alert-create label-name="所属站段" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-workshop_alert-create label-name="所属车间" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-work-area_alert-create label-name="所属工区" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-warehouse-storehouse_alert-create label-name="所属库房" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-warehouse-area_alert-create label-name="所属库区" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-warehouse-platoon_alert-create label-name="所属排" />
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn type="button" label="关闭" v-close-popup />
          <q-btn type="submit" label="确定" icon="check" color="secondary" />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
  <!-- 编辑仓库-架弹窗 -->
  <q-dialog v-model="alertEditWarehouseShelf" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">编辑仓库-架</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnUpdateWarehouseShelf">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertEditWarehouseShelf" label="名称" :rules="[]" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-railway_alert-edit label-name="所属路局" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-paragraph_alert-edit label-name="所属站段" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-workshop_alert-edit label-name="所属车间" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-organization-work-area_alert-edit label-name="所属工区" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-warehouse-storehouse_alert-edit label-name="所属库房" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-warehouse-area_alert-edit label-name="所属库区" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <sel-warehouse-platoon_alert-edit label-name="所属排" />
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn type="button" label="关闭" v-close-popup />
          <q-btn type="submit" label="确定" icon="check" color="warning" />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
  <!-- 编辑仓库-架正视图 -->
  <q-dialog v-model="alertEditWarehouseShelfSubs" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '550px' }">
      <q-card-section>
        <div class="row">
          <div class="col-4">
            <div class="text-h6">编辑架结构</div>
          </div>
          <div class="col-8">
            <q-btn-group>
              <q-btn @click="fnOpenAlertCreateWarehouseTiers" color="secondary" icon="add" label="新建层"></q-btn>
              <q-btn @click="fnDestoryWarehouseTiers" color="negative" icon="delete" label="删除层"></q-btn>
              <q-btn @click="fnOpenAlertCreateWarehouseCells" color="secondary" icon="add" label="新建位"></q-btn>
            </q-btn-group>
          </div>
        </div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertEditWarehouseShelfSubs" label="名称" :rules="[]"
                @keydown="fnUpdateWarehouseShelfName($event)" />
            </div>
          </div>
          <div class="row q-mt-md">
            <div class="col">
              <template v-if="warehouseShelf_alertEditWarehouseShelfSubs">
                <div class="q-gutter-y-sm column">
                  <q-toolbar flat round dense class="bg-grey-3" style="white-space: nowrap; overflow-x: scroll;"
                    v-for="(warehouseTier, idx) in collect(warehouseShelf_alertEditWarehouseShelfSubs.warehouse_tiers || []).reverse().all()"
                    :key="idx">
                    <q-checkbox v-model="warehouseTierUuids_alertEditWarehouseShelfSubs" :label="warehouseTier.name"
                      :val="warehouseTier.uuid" :key="warehouseTier.uuid" />
                    &emsp;
                    <q-btn
                      @click="fnOpenAlertEditWarehouseTier(warehouseShelf_alertEditWarehouseShelfSubs, warehouseTier)"
                      color="warning" icon="edit" />
                    <q-separator vertical />
                    &emsp;
                    <q-btn v-for="(warehouseCell, idx2) in warehouseTier.warehouse_cells" :key="idx2"
                      :label="warehouseCell.name" color="primary"
                      @click="fnOpenAlertEditWarehouseCell(warehouseTier, warehouseCell)" />
                  </q-toolbar>
                </div>
              </template>
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn type="button" label="关闭" v-close-popup />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
  <!-- 新建仓库-层弹窗 -->
  <q-dialog v-model="alertCreateWarehouseTiers" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">新建仓库-层</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnStoreWarehouseTiers">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="number_alertCreateWarehouseTiers" label="数量" :rules="[]"
                type="number" />
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn type="button" label="关闭" v-close-popup />
          <q-btn type="submit" label="确定" icon="check" color="secondary" />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
  <!-- 编辑仓库-层弹窗 -->
  <q-dialog v-model="alertEditWarehouseTier" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">编辑仓库-层</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnUpdateWarehouseTier">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertEditWarehouseTier" label="名称" :rules="[]" />
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn type="button" label="关闭" v-close-popup />
          <q-btn type="submit" label="确定" icon="check" color="warning" />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
  <!-- 新建仓库-位弹窗 -->
  <q-dialog v-model="alertCreateWarehouseCells" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">新建仓库-位</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnStoreWarehouseCells">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="number_alertCreateWarehouseCells" label="数量" :rules="[]"
                type="number" />
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn type="button" label="关闭" v-close-popup />
          <q-btn type="submit" label="确定" icon="check" color="secondary" />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
  <q-dialog v-model="alertEditWarehouseCell" no-backdrop-dismiss>
    <q-card :style="{ minWidth: '450px' }">
      <q-card-section>
        <div class="text-h6">编辑仓库-位</div>
      </q-card-section>
      <q-form class="q-gutter-md" @submit.prevent="fnUpdateWarehouseCell">
        <q-card-section class="q-pt-none">
          <div class="row">
            <div class="col">
              <q-input outlined clearable lazy-rules v-model="name_alertEditWarehouseCell" label="名称" :rules="[]" />
            </div>
          </div>
        </q-card-section>
        <q-card-actions align="right">
              <q-btn color="negative" label="删除" icon="delete" @click="fnDestroyWarehouseCell" />
              <q-btn type="button" label="关闭" v-close-popup />
              <q-btn type="submit" label="确定" icon="check" color="warning" />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
</template>
<script setup>
import { ref, onMounted, provide, watch } from "vue";
import collect from "collect.js";
import { fnColumnReverseSort } from "src/utils/common";
import {
  ajaxGetWarehouseShelves,
  ajaxGetWarehouseShelf,
  ajaxStoreWarehouseShelf,
  ajaxUpdateWarehouseShelf,
  ajaxDestroyWarehouseShelf,
  ajaxDestroyWarehouseShelves,
  ajaxStoreWarehouseTiers,
  ajaxUpdateWarehouseTier,
  ajaxDestroyWarehouseTiers,
  ajaxStoreWarehouseCells,
  ajaxUpdateWarehouseCell,
  ajaxDestroyWarehouseCell,
} from "src/apis/warehouse";
import {
  loadingNotify,
  successNotify,
  errorNotify,
  confirmNotify,
  destroyActions,
} from "src/utils/notify";
import JoinString from "src/components/JoinString.vue";
import SelOrganizationRailway_search from "src/components/SelOrganizationRailway_search.vue";
import SelOrganizationParagraph_search from "src/components/SelOrganizationParagraph_search.vue";
import SelOrganizationWorkshop_search from "src/components/SelOrganizationWorkshop_search.vue";
import SelOrganizationWorkArea_search from "src/components/SelOrganizationWorkArea_search.vue";
import SelWarehouseStorehouse_search from "src/components/SelWarehouseStorehouse_search.vue";
import SelWarehouseArea_search from "src/components/SelWarehouseArea_search.vue";
import SelWarehousePlatoon_search from "src/components/SelWarehousePlatoon_search.vue";
import SelOrganizationRailway_alertCreate from "src/components/SelOrganizationRailway_alertCreate.vue";
import SelOrganizationParagraph_alertCreate from "src/components/SelOrganizationParagraph_alertCreate.vue";
import SelOrganizationWorkshop_alertCreate from "src/components/SelOrganizationWorkshop_alertCreate.vue";
import SelOrganizationWorkArea_alertCreate from "src/components/SelOrganizationWorkArea_alertCreate.vue";
import SelWarehouseStorehouse_alertCreate from "src/components/SelWarehouseStorehouse_alertCreate.vue";
import SelWarehouseArea_alertCreate from "src/components/SelWarehouseArea_alertCreate.vue";
import SelWarehousePlatoon_alertCreate from "src/components/SelWarehousePlatoon_alertCreate.vue";
import SelOrganizationRailway_alertEdit from "src/components/SelOrganizationRailway_alertEdit.vue";
import SelOrganizationParagraph_alertEdit from "src/components/SelOrganizationParagraph_alertEdit.vue";
import SelOrganizationWorkshop_alertEdit from "src/components/SelOrganizationWorkshop_alertEdit.vue";
import SelOrganizationWorkArea_alertEdit from "src/components/SelOrganizationWorkArea_alertEdit.vue";
import SelWarehouseStorehouse_alertEdit from "src/components/SelWarehouseStorehouse_alertEdit.vue";
import SelWarehouseArea_alertEdit from "src/components/SelWarehouseArea_alertEdit.vue";
import SelWarehousePlatoon_alertEdit from "src/components/SelWarehousePlatoon_alertEdit.vue";

// 搜索栏数据
const name_search = ref("");
const organizationRailwayUuid_search = ref("");
provide("organizationRailwayUuid_search", organizationRailwayUuid_search);
const organizationParagraphUuid_search = ref("");
provide("organizationParagraphUuid_search", organizationParagraphUuid_search);
const organizationWorkshopUuid_search = ref("");
provide("organizationWorkshopUuid_search", organizationWorkshopUuid_search);
const organizationWorkAreaUuid_search = ref("");
provide("organizationWorkAreaUuid_search", organizationWorkAreaUuid_search);
const warehouseStorehouseUuid_search = ref("");
provide("warehouseStorehouseUuid_search", warehouseStorehouseUuid_search);
const warehouseAreaUuid_search = ref("");
provide("warehouseAreaUuid_search", warehouseAreaUuid_search);
const warehousePlatoonUuid_search = ref("");
provide("warehousePlatoonUuid_search", warehousePlatoonUuid_search);

// 表格数据
const rows = ref([]);
const selected = ref([]);
const sortBy = ref("");

// 新建仓库-架弹窗
const alertCreateWarehouseShelf = ref(false);
const name_alertCreateWarehouseShelf = ref("");
const organizationRailwayUuid_alertCreateWarehouseShelf = ref("");
provide("organizationRailwayUuid_alertCreate", organizationRailwayUuid_alertCreateWarehouseShelf);
const organizationParagraphUuid_alertCreateWarehouseShelf = ref("");
provide("organizationParagraphUuid_alertCreate", organizationParagraphUuid_alertCreateWarehouseShelf);
const organizationWorkshopUuid_alertCreateWarehouseShelf = ref("");
provide("organizationWorkshopUuid_alertCreate", organizationWorkshopUuid_alertCreateWarehouseShelf);
const organizationWorkAreaUuid_alertCreateWarehouseShelf = ref("");
provide("organizationWorkAreaUuid_alertCreate", organizationWorkAreaUuid_alertCreateWarehouseShelf);
const warehouseStorehouseUuid_alertCreateWarehouseShelf = ref("");
provide("warehouseStorehouseUuid_alertCreate", warehouseStorehouseUuid_alertCreateWarehouseShelf);
const warehouseAreaUuid_alertCreateWarehouseShelf = ref("");
provide("warehouseAreaUuid_alertCreate", warehouseAreaUuid_alertCreateWarehouseShelf);
const warehousePlatoonUuid_alertCreateWarehouseShelf = ref("");
provide("warehousePlatoonUuid_alertCreate", warehousePlatoonUuid_alertCreateWarehouseShelf);

// 编辑仓库-架弹窗数据
const alertEditWarehouseShelf = ref(false);
const currentWarehouseShelfUuid = ref("");
const name_alertEditWarehouseShelf = ref("");
const organizationRailwayUuid_alertEditWarehouseShelf = ref("");
provide("organizationRailwayUuid_alertEdit", organizationRailwayUuid_alertEditWarehouseShelf);
const organizationParagraphUuid_alertEditWarehouseShelf = ref("");
provide("organizationParagraphUuid_alertEdit", organizationParagraphUuid_alertEditWarehouseShelf);
const organizationWorkshopUuid_alertEditWarehouseShelf = ref("");
provide("organizationWorkshopUuid_alertEdit", organizationWorkshopUuid_alertEditWarehouseShelf);
const organizationWorkAreaUuid_alertEditWarehouseShelf = ref("");
provide("organizationWorkAreaUuid_alertEdit", organizationWorkAreaUuid_alertEditWarehouseShelf);
const warehouseStorehouseUuid_alertEditWarehouseShelf = ref("");
provide("warehouseStorehouseUuid_alertEdit", warehouseStorehouseUuid_alertEditWarehouseShelf);
const warehouseAreaUuid_alertEditWarehouseShelf = ref("");
provide("warehouseAreaUuid_alertEdit", warehouseAreaUuid_alertEditWarehouseShelf);
const warehousePlatoonUuid_alertEditWarehouseShelf = ref("");
provide("warehousePlatoonUuid_alertEdit", warehousePlatoonUuid_alertEditWarehouseShelf);

// 编辑仓库-架正视图数据
const alertEditWarehouseShelfSubs = ref(false);
const name_alertEditWarehouseShelfSubs = ref("");
const warehouseShelf_alertEditWarehouseShelfSubs = ref(null);
const warehouseTierUuids_alertEditWarehouseShelfSubs = ref([]);

// 新建仓库-层弹窗数据
const alertCreateWarehouseTiers = ref(false);
const number_alertCreateWarehouseTiers = ref(0);

// 编辑仓库-层弹窗数据
const alertEditWarehouseTier = ref(false);
const currentUuid_alertEditWarehouseTier = ref("");
const name_alertEditWarehouseTier = ref("");
const warehouseShelfUuid_alertEditWarehouseTier = ref("");

// 新建仓库-位弹窗数据
const alertCreateWarehouseCells = ref(false);
const number_alertCreateWarehouseCells = ref(0);

// 编辑仓库-位弹窗数据
const alertEditWarehouseCell = ref(false);
const currentUuid_alertEditWarehouseCell = ref("");
const name_alertEditWarehouseCell = ref("");
const warehouseTierUuid_alertEditWarehouseCell = ref("");


onMounted(() => fnInit());

const fnInit = () => fnSearch();

const fnResetSearch = () => {
  name_search.value = "";
  organizationRailwayUuid_search.value = "";
  organizationParagraphUuid_search.value = "";
  organizationWorkshopUuid_search.value = "";
  organizationWorkAreaUuid_search.value = "";
  warehouseStorehouseUuid_search.value = "";
  warehouseAreaUuid_search.value = "";
  warehousePlatoonUuid_search.value = "";
};

const fnSearch = () => {
  rows.value = [];
  selected.value = [];

  ajaxGetWarehouseShelves({
    "@~[]": [
      "WarehousePlatoon",
      "WarehousePlatoon.WarehouseArea",
      "WarehousePlatoon.WarehouseArea.WarehouseStorehouse",
      "WarehousePlatoon.WarehouseArea.WarehouseStorehouse.OrganizationWorkshop",
      "WarehousePlatoon.WarehouseArea.WarehouseStorehouse.OrganizationWorkArea",
      "WarehousePlatoon.WarehouseArea.WarehouseStorehouse.OrganizationWorkshop.OrganizationParagraph",
      "WarehousePlatoon.WarehouseArea.WarehouseStorehouse.OrganizationWorkshop.OrganizationParagraph.OrganizationRailway",
    ],
    name: name_search.value,
    organizationRailwayUuid: organizationRailwayUuid_search.value,
    organizationParagraphUuid: organizationParagraphUuid_search.value,
    organizationWorkshopUuid: organizationWorkshopUuid_search.value,
    organizationWorkAreaUuid: organizationWorkAreaUuid_search.value,
    warehouseStorehouseUuid: warehouseStorehouseUuid_search.value,
    warehouseAreaUuid: warehouseAreaUuid_search.value,
    warehousePlatoonUuid: warehousePlatoonUuid_search.value,
  })
    .then((res) => {
      rows.value = collect(res.content.warehouse_shelves)
        .map((warehouseShelf, idx) => {
          return {
            ...warehouseShelf,
            index: idx + 1,
            warehouseStorehouse: warehouseShelf.warehouse_platoon.warehouse_area.warehouse_storehouse,
            warehouseArea: warehouseShelf.warehouse_platoon.warehouse_area,
            warehousePlatoon: warehouseShelf.warehouse_platoon,
            operation: { uuid: warehouseShelf.uuid },
          }
        })
        .all();
    })
    .catch((e) => errorNotify(e.msg));
};

const fnResetAlertCreateWarehouseShelf = () => {
  name_alertCreateWarehouseShelf.value = "";
  organizationRailwayUuid_alertCreateWarehouseShelf.value = "";
  organizationParagraphUuid_alertCreateWarehouseShelf.value = "";
  organizationWorkshopUuid_alertCreateWarehouseShelf.value = "";
  organizationWorkAreaUuid_alertCreateWarehouseShelf.value = "";
  warehouseStorehouseUuid_alertCreateWarehouseShelf.value = "";
  warehouseAreaUuid_alertCreateWarehouseShelf.value = "";
  warehousePlatoonUuid_alertCreateWarehouseShelf.value = "";
};

const fnOpenAlertCreateWarehouseShelf = () => {
  alertCreateWarehouseShelf.value = true;
};

const fnStoreWarehouseShelf = () => {
  const loading = loadingNotify();

  ajaxStoreWarehouseShelf({
    name: name_alertCreateWarehouseShelf.value,
    warehouse_platoon_uuid: warehousePlatoonUuid_alertCreateWarehouseShelf.value,
  })
    .then((res) => {
      successNotify(res.msg);
      fnResetAlertCreateWarehouseShelf();
      fnSearch();

      alertCreateWarehouseShelf.value = false;
    })
    .catch((e) => errorNotify(e.msg))
    .finally(loading());
};

const fnOpenAlertEditWarehouseShelf = params => {
  if (!params["uuid"]) return;
  currentWarehouseShelfUuid.value = params.uuid;

  ajaxGetWarehouseShelf(currentWarehouseShelfUuid.value, {
    "@~[]": [
      "WarehousePlatoon",
      "WarehousePlatoon.WarehouseArea",
      "WarehousePlatoon.WarehouseArea.WarehouseStorehouse",
      "WarehousePlatoon.WarehouseArea.WarehouseStorehouse.OrganizationWorkshop",
      "WarehousePlatoon.WarehouseArea.WarehouseStorehouse.OrganizationWorkArea",
      "WarehousePlatoon.WarehouseArea.WarehouseStorehouse.OrganizationWorkshop.OrganizationParagraph",
      "WarehousePlatoon.WarehouseArea.WarehouseStorehouse.OrganizationWorkshop.OrganizationParagraph.OrganizationRailway",
    ],
  })
    .then((res) => {
      name_alertEditWarehouseShelf.value = res.content.warehouse_shelf.name;
      organizationRailwayUuid_alertEditWarehouseShelf.value = res.content.warehouse_shelf.warehouse_platoon.warehouse_area.warehouse_storehouse.organization_workshop.organization_paragraph.organization_railway.uuid;
      organizationParagraphUuid_alertEditWarehouseShelf.value = res.content.warehouse_shelf.warehouse_platoon.warehouse_area.warehouse_storehouse.organization_workshop.organization_paragraph.uuid;
      organizationWorkshopUuid_alertEditWarehouseShelf.value = res.content.warehouse_shelf.warehouse_platoon.warehouse_area.warehouse_storehouse.organization_workshop.uuid;
      organizationWorkAreaUuid_alertEditWarehouseShelf.value = res.content.warehouse_shelf.warehouse_platoon.warehouse_area.warehouse_storehouse.organization_work_area.uuid;
      warehouseStorehouseUuid_alertEditWarehouseShelf.value = res.content.warehouse_shelf.warehouse_platoon.warehouse_area.warehouse_storehouse.uuid;
      warehouseAreaUuid_alertEditWarehouseShelf.value = res.content.warehouse_shelf.warehouse_platoon.warehouse_area.uuid;
      warehousePlatoonUuid_alertEditWarehouseShelf.value = res.content.warehouse_shelf.warehouse_platoon.uuid;

      alertEditWarehouseShelf.value = true;
    })
    .catch((e) => errorNotify(e.msg));
};

const fnUpdateWarehouseShelf = () => {
  if (!currentWarehouseShelfUuid.value) return;

  const loading = loadingNotify();
  ajaxUpdateWarehouseShelf(currentWarehouseShelfUuid.value, {
    name: name_alertEditWarehouseShelf.value,
    warehouse_platoon_uuid: warehousePlatoonUuid_alertEditWarehouseShelf.value,
  })
    .then((res) => {
      successNotify(res.msg);
      fnSearch();
      currentWarehouseShelfUuid.value = "";

      alertEditWarehouseShelf.value = false;
    })
    .catch((e) => errorNotify(e.msg))
    .finally(loading());
};

const fnDestroyCreateWarehouseShelf = params => {
  if (!parmas["uuid"]) return;

  confirmNotify(destroyActions(() => {
    const loading = loadingNotify();

    ajaxDestroyWarehouseShelf(params.uuid)
      .then(() => {
        successNotify("删除成功");
        fnSearch();
      })
      .catch((e) => errorNotify(e.msg))
      .finally(loading());
  }));
};

const fnDestroyCreateWarehouseShelves = () => {
  confirmNotify(destroyActions(() => {
    const loading = loadingNotify();

    ajaxDestroyWarehouseShelves(collect(selected.value).pluck("uuid").all())
      .then(() => {
        successNotify("删除成功");
        fnSearch();
      })
      .catch(e => errorNotify(e.msg))
      .finally(loading());
  }));
};

const fnUpdateWarehouseShelfName = (event) => {
  if (event.keyCode !== 13) return;
  if (!currentWarehouseShelfUuid.value) return;
  if (!event.target.value) return;

  const loading = loadingNotify();

  ajaxUpdateWarehouseShelf(currentWarehouseShelfUuid.value, {
    name: event.target.value,
    warehouse_platoon_uuid: warehouseShelf_alertEditWarehouseShelfSubs.value.warehouse_platoon.uuid,
  })
    .then(res => {
      successNotify(res.msg);
      fnSearch();
    })
    .catch(e => errorNotify(e.msg))
    .finally(loading());
};

const fnOpenAlertEditWarehouseShelfSubs = params => {
  if (!params["uuid"]) return;
  currentWarehouseShelfUuid.value = params.uuid;
  warehouseShelf_alertEditWarehouseShelfSubs.value = null;
  warehouseTierUuids_alertEditWarehouseShelfSubs.value = [];

  ajaxGetWarehouseShelf(params.uuid, {
    "@~[]": [
      "WarehousePlatoon",
      "WarehouseTiers",
      "WarehouseTiers.WarehouseCells",
    ],
  })
    .then(res => {
      name_alertEditWarehouseShelfSubs.value = res.content.warehouse_shelf.name;
      warehouseShelf_alertEditWarehouseShelfSubs.value = res.content.warehouse_shelf;
      alertEditWarehouseShelfSubs.value = true;
    })
    .catch(e => errorNotify(e.msg));
};

const fnResetAlertCreateWarehouseTiers = () => {
  number_alertCreateWarehouseTiers.value = 0;
};

const fnOpenAlertCreateWarehouseTiers = () => {
  if (!currentWarehouseShelfUuid.value) return;

  alertCreateWarehouseTiers.value = true;
};

const fnStoreWarehouseTiers = () => {
  const loading = loadingNotify();

  ajaxStoreWarehouseTiers({
    warehouse_shelf_uuid: currentWarehouseShelfUuid.value,
    number: parseInt(number_alertCreateWarehouseTiers.value),
  })
    .then(res => {
      successNotify(res.msg);
      fnOpenAlertEditWarehouseShelfSubs({ uuid: currentWarehouseShelfUuid.value });
      fnResetAlertCreateWarehouseTiers();

      alertCreateWarehouseTiers.value = false;
    })
    .catch(e => errorNotify(e.msg))
    .finally(loading());
};

const fnDestoryWarehouseTiers = () => {
  confirmNotify(destroyActions(() => {
    const loading = loadingNotify();

    ajaxDestroyWarehouseTiers(warehouseTierUuids_alertEditWarehouseShelfSubs.value)
      .then(() => {
        successNotify("删除成功");
        fnOpenAlertEditWarehouseShelfSubs({ uuid: currentWarehouseShelfUuid.value });
      })
      .catch(e => errorNotify(e.msg))
      .finally(loading());
  }));
};

const fnRestAlertCreateWarehouseCells = () => {
  number_alertCreateWarehouseCells.value = 0;
};

const fnOpenAlertCreateWarehouseCells = () => {
  if (collect(warehouseTierUuids_alertEditWarehouseShelfSubs.value).isEmpty()) {
    errorNotify("请勾选需要添加位的【层】");
    return;
  }

  alertCreateWarehouseCells.value = true;
};

const fnStoreWarehouseCells = () => {
  if (collect(warehouseTierUuids_alertEditWarehouseShelfSubs.value).isEmpty()) {
    errorNotify("请勾选需要添加位的【层】");
    return;
  }

  const loading = loadingNotify();
  collect(warehouseTierUuids_alertEditWarehouseShelfSubs.value)
    .each(warehouseTierUuid => {
      ajaxStoreWarehouseCells({
        warehouse_tier_uuid: warehouseTierUuid,
        number: parseInt(number_alertCreateWarehouseCells.value),
      })
        .then(res => {
          successNotify(res.msg);
          fnOpenAlertEditWarehouseShelfSubs({ uuid: currentWarehouseShelfUuid.value });
          fnRestAlertCreateWarehouseCells();

          alertCreateWarehouseCells.value = false;
        })
        .catch(e => errorNotify(e.msg))
        .finally(loading());
    });
};

const fnOpenAlertEditWarehouseTier = (warehouseShelf, warehouseTier) => {
  if (!warehouseTier) return;
  currentUuid_alertEditWarehouseTier.value = warehouseTier.uuid;

  name_alertEditWarehouseTier.value = warehouseTier.name;
  warehouseShelfUuid_alertEditWarehouseTier.value = warehouseShelf.uuid;
  alertEditWarehouseTier.value = true;
};

const fnUpdateWarehouseTier = () => {
  if (!currentUuid_alertEditWarehouseTier.value) return;

  const loading = loadingNotify();
  ajaxUpdateWarehouseTier(currentUuid_alertEditWarehouseTier.value, {
    name: name_alertEditWarehouseTier.value,
    warehouse_shelf_uuid: warehouseShelfUuid_alertEditWarehouseTier.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnOpenAlertEditWarehouseShelfSubs({ uuid: currentWarehouseShelfUuid.value });
      currentUuid_alertEditWarehouseTier.value = "";

      alertEditWarehouseTier.value = false;
    })
    .catch(e => errorNotify(e.msg))
    .finally(loading());
};

const fnOpenAlertEditWarehouseCell = (warehouseTier, warehouseCell) => {
  if (!warehouseCell) return;
  currentUuid_alertEditWarehouseCell.value = warehouseCell.uuid;

  name_alertEditWarehouseCell.value = warehouseCell.name;
  warehouseTierUuid_alertEditWarehouseCell.value = warehouseTier.uuid;
  alertEditWarehouseCell.value = true;
};

const fnUpdateWarehouseCell = () => {
  const loading = loadingNotify();

  ajaxUpdateWarehouseCell(currentUuid_alertEditWarehouseCell.value, {
    name: name_alertEditWarehouseCell.value,
    warehouse_tier_uuid: warehouseTierUuid_alertEditWarehouseCell.value,
  })
    .then(res => {
      successNotify(res.msg);
      fnOpenAlertEditWarehouseShelfSubs({ uuid: currentWarehouseShelfUuid.value });

      alertEditWarehouseCell.value = false;
    })
    .catch(e => errorNotify(e.msg))
    .finally(loading());
};

const fnDestroyWarehouseCell = () => {
  if (!currentUuid_alertEditWarehouseCell.value) return;

  confirmNotify(destroyActions(() => {
    const loading = loadingNotify();

    ajaxDestroyWarehouseCell(currentUuid_alertEditWarehouseCell.value)
      .then(() => {
        successNotify("删除成功");
        fnOpenAlertEditWarehouseShelfSubs({ uuid: currentWarehouseShelfUuid.value });

        alertEditWarehouseCell.value = false;
      })
      .catch(e => errorNotify(e.msg))
      .finally(loading());
  }));
};
</script>
