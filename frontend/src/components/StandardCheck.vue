<template>
  <template v-if="!isEmpty(groupByItem)">
    <q-card flat bordered>
      <q-card-section>
        <p class="text-body1">{{ labelName }}</p>
        <q-separator class="q-mb-md" />
        <div class="row" v-for="(item, idx) in l_com" :key="idx"
          v-show="(parentValue === getVal(item.group, groupByKey)) || (!parentValue)">
          <div class="col-12">
            <p class="text-subtitle2 q-mb-none">{{ getVal(item.group, groupByName) }}</p>
          </div>
          <template v-for="(subs2, idx2) in item.subs" :key="idx2">
            <div class="col-4" v-for="sub in subs2" :key="getVal(sub, valueField)">
              <q-checkbox v-model="currentValue" :val="getVal(sub, valueField)" :key="getVal(sub, valueField)"
                :label="getVal(sub, labelField)" />
            </div>
          </template>
        </div>
      </q-card-section>
    </q-card>
  </template>
  <template v-else>
    <q-card flat bordered>
      <q-card-section>
        <p class="text-body1">{{ labelName }}</p>
        <q-separator class="q-mb-md" />
        <div class="row">
          <template v-for="(items, idx) in l_com" :key="idx">
            <div class="col-4" v-for="item in items" :key="getVal(item, valueField)">
              <q-checkbox v-model="currentValue" :val="getVal(item, valueField)" :key="getVal(item, valueField)"
                :label="getVal(item, labelField)" />
            </div>
          </template>
        </div>
      </q-card-section>
    </q-card>
  </template>
</template>
<script setup>
import { ref, onMounted, inject, defineProps, watch } from "vue";
import collect from "collect.js";
import { getVal, isEmpty } from "src/utils/common";
import { errorNotify } from "src/utils/notify";

const props = defineProps({
  labelName: {
    type: String,
    default: "",
  },
  dataSource: {
    type: Function,
    required: true,
  },
  dataSourceField: {
    type: String,
    required: true,
  },
  currentField: {
    type: String,
    required: true,
  },
  sechma: {
    type: String,
    required: true,
  },
  parentField: {
    type: String,
    default: "",
  },
  groupByItem: {
    type: String,
    default: "",
  },
  groupByKey: {
    type: String,
    default: "uuid",
  },
  groupByName: {
    type: String,
    default: "name",
  },
  labelField: {
    type: String,
    default: "name",
  },
  valueField: {
    type: String,
    default: "uuid",
  },
  chunk: {
    type: Number,
    default: 4,
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
const dataSource = props.dataSource;
const dataSourceField = props.dataSourceField;
const currentField = props.currentField;
const sechma = props.sechma;
const parentField = props.parentField;
const groupByItem = props.groupByItem;
const groupByKey = props.groupByKey;
const groupByName = props.groupByName;
const labelField = props.labelField;
const valueField = props.valueField;
const chunk = props.chunk;
const l_com = ref([]);
const currentValue = inject(`${currentField}_${sechma}`);

let parentValue = null;
if (parentField) {
  parentValue = inject(`${parentField}_${sechma}`);
  watch(parentValue, () => fnSearch());
}

onMounted(() => fnSearch());

const fnSearch = () => {
  l_com.value = [];

  dataSource(ajaxParams)
    .then(res => {
      if (!isEmpty(groupByItem)) {
        const tmp = {};
        collect(getVal(res, `content.${dataSourceField}`))
          .each(item => {
            if (!tmp[getVal(item, `${groupByItem}.${groupByKey}`)]) {
              tmp[getVal(item, `${groupByItem}.${groupByKey}`)] = { group: getVal(item, groupByItem), subs: [] };
            }
            tmp[getVal(item, `${groupByItem}.${groupByKey}`)].subs.push(item);
          });
        l_com.value = collect(tmp)
          .map(item => {
            item.subs = collect(item.subs).chunk(chunk).all();
            return item;
          })
          .all();
      } else {
        l_com.value = collect(getVal(res, `content.${dataSourceField}`)).chunk(chunk).all();
      }
    })
    .catch(e => errorNotify(e.msg));
};
</script>
