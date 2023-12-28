<template>
  <div class="q-pa-md">
    <q-card>
      <q-card-section>
        <div class="row q-mb-md">
          <div class="col"><span style="font-size: 20px">搜索</span></div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn
                color="primary"
                label="搜索"
                icon="search"
                @click="fnSearch"
              />
              <q-btn color="primary" label="重置" flat @click="fnResetSearch" />
            </q-btn-group>
          </div>
        </div>
        <div class="row">
          <div class="col">
            <q-form>
              <div class="row q-col-gutter-sm">
                <div class="col-3">
                  <q-input
                    outlined
                    clearable
                    lazy-rules
                    v-model="username_search"
                    label="账号"
                    :rules="[]"
                  />
                </div>
                <div class="col-3">
                  <q-input
                    outlined
                    clearable
                    lazy-rules
                    v-model="nickname_search"
                    label="昵称"
                    :rules="[]"
                  />
                </div>
                <div class="col-3">
                  <sel-rbac-role_search label-name="所属角色" />
                </div>
              </div>
            </q-form>
          </div>
        </div>
      </q-card-section>
    </q-card>

    <q-card class="q-mt-md">
      <q-card-section>
        <div class="row">
          <div class="col">
            <span :style="{ fontSize: '20px' }">用户列表</span>
          </div>
          <div class="col text-right">
            <q-btn-group>
              <q-btn
                color="secondary"
                label="新建用户"
                icon="add"
                @click="fnOpenAlertCreateAccount"
              />
              <q-btn
                color="negative"
                label="批量删除用户"
                icon="delete"
                @click="fnDestroyAccounts"
              />
            </q-btn-group>
          </div>
        </div>
        <div class="row q-mt-md">
          <div class="col">
            <q-table
              flat
              bordered
              title="用列表"
              :rows="rows"
              row-key="uuid"
              :pagination="{ rowsPerPage: 200 }"
              :rows-per-page-options="[50, 100, 200, 0]"
              rows-per-page-label="分页"
              selection="multiple"
              v-model:selected="selected"
            >
              <template v-slot:header="props">
                <q-tr :props="props">
                  <q-th align="left"
                    ><q-checkbox key="allCheck" v-model="props.selected"
                  /></q-th>
                  <q-th align="left">#</q-th>
                  <q-th
                    align="left"
                    name="username"
                    key="username"
                    @click="
                      (event) => fnColumnReverseSort(event, props, sortBy)
                    "
                  >
                    用户名
                  </q-th>
                  <q-th
                    align="left"
                    name="nickname"
                    key="nickname"
                    @click="
                      (event) => fnColumnReverseSort(event, props, sortBy)
                    "
                  >
                    昵称
                  </q-th>
                  <q-th align="left">头像</q-th>
                  <q-th align="left">所属角色</q-th>
                  <q-th align="right"></q-th>
                </q-tr>
              </template>
              <template v-slot:body="props">
                <q-tr :props="props">
                  <q-td
                    ><q-checkbox
                      :key="props.row.uuid"
                      :value="props.row.uuid"
                      v-model="props.selected"
                  /></q-td>
                  <q-td>{{ props.row.index }}</q-td>
                  <q-td>{{ props.row.username }}</q-td>
                  <q-td>{{ props.row.nickname }}</q-td>
                  <q-td>
                    <q-avatar
                      key="avatar"
                      :props="props"
                      v-if="props.row.avatar"
                    >
                      <img size="50px" :src="props.row.avatar" />
                    </q-avatar>
                  </q-td>
                  <q-td>
                    <q-chip
                      v-for="(rbacRole, idx) in props.row.rbac_roles"
                      :key="idx"
                      color="primary"
                      text-color="white"
                    >
                      {{ rbacRole.name }}
                    </q-chip>
                  </q-td>
                  <q-td>
                    <q-btn-group>
                      <q-btn
                        @click="fnOpenAlertEditAccount(props.row.operation)"
                        color="warning"
                        icon="edit"
                      >
                        编辑
                      </q-btn>
                      <q-btn
                        @click="fnOpenAlertEditPassword(props.row.operation)"
                        color="warning"
                        icon="lock"
                      >
                        重置密码
                      </q-btn>
                      <q-btn
                        @click="fnDestroyAccount(props.row.operation)"
                        color="negative"
                        icon="delete"
                      >
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

  <!-- 对话框 -->
  <!-- 新建用户对话框 -->
  <q-dialog v-model="alertCreateAccount">
    <q-card style="width: 800px">
      <q-card-section>
        <div class="text-h6">新建用户</div>
      </q-card-section>
      <q-card-section class="q-pt-none">
        <q-form class="q-gutter-md" @submit.prevent="">
          <div class="row">
            <div class="col">
              <q-input
                outlined
                clearable
                lazy-rules
                v-model="username_alertCreateAccount"
                label="账号"
                :rules="[]"
              />
              <q-input
                outlined
                clearable
                lazy-rules
                v-model="nickname_alertCreateAccount"
                label="昵称"
                :rules="[]"
                class="q-mt-md"
              />
              <q-input
                outlined
                clearable
                lazy-rules
                type="password"
                v-model="password_alertCreateAccount"
                label="密码"
                :rules="[]"
                class="q-mt-md"
              />
              <q-input
                outlined
                clearable
                lazy-rules
                type="password"
                v-model="passwordConfirmation_alertCreateAccount"
                label="确认密码"
                :rules="[]"
                class="q-mt-md"
              />
              <q-file
                outlined
                v-model="avatar_alertCreateAccount"
                label="用户头像"
                accept="image/*"
                max-file-size="1048576"
                counter
                class="q-mt-md"
              />
              <chk-rbac-role_alert-create
                label-name="所属角色"
                class="q-mt-md"
              />
            </div>
          </div>
        </q-form>
      </q-card-section>
      <q-card-actions align="right">
        <q-btn
          type="button"
          label="确定"
          icon="check"
          color="secondary"
          @click="fnStoreAccount"
          v-close-popup
        />
      </q-card-actions>
    </q-card>
  </q-dialog>
  <!-- 编辑用户对话框 -->
  <q-dialog v-model="alertEditAccount">
    <q-card style="width: 800px">
      <q-card-section>
        <div class="text-h6">编辑用户</div>
      </q-card-section>
      <q-card-section class="q-pt-none">
        <q-form class="q-gutter-md" @submit.prevent="">
          <div class="row">
            <div class="col">
              <q-input
                outlined
                clearable
                lazy-rules
                v-model="username_alertEditAccount"
                label="账号"
                :rules="[]"
              />
              <q-input
                outlined
                clearable
                lazy-rules
                v-model="nickname_alertEditAccount"
                label="昵称"
                :rules="[]"
                class="q-mt-md"
              />
              <q-avatar
                key="avatar"
                v-if="showAvatar_alertEditAccount"
                class="q-mt-md"
              >
                <img size="200px" :src="showAvatar_alertEditAccount" />
              </q-avatar>
              <!-- <q-file outlined v-model="avatar_alertEditAccount" label="用户头像" accept="image/*" max-file-size="1048576"
                counter class="q-mt-md" /> -->
              <q-uploader
                :headers="[
                  { name: 'Authorization', value: `JWT ${authToken}` },
                ]"
                field-name="avatar"
                :url="`${settings.apiUrl}/account/${currentUuid}/updateAvatar`"
                label="修改头像"
                class="full-width q-mt-md"
                @uploaded="fnOnUploadedAvatar($event)"
                flat
                bordered
              >
                <template v-slot:list="scope">
                  <q-list separator>
                    <q-item v-for="file in scope.files" :key="file.__key">
                      <q-item-section>
                        <q-item-label class="full-width ellipsis">
                          {{ file.name }}
                        </q-item-label>
                        <q-item-label caption>
                          Status: {{ file.__status }}
                        </q-item-label>
                        <q-item-label caption>
                          {{ file.__sizeLabel }} / {{ file.__progressLabel }}
                        </q-item-label>
                      </q-item-section>
                      <q-item-section v-if="file.__img" thumbnail class="gt-xs">
                        <img :src="file.__img.src" />
                      </q-item-section>
                      <q-item-section top side>
                        <q-btn
                          class="gt-xs"
                          size="12px"
                          flat
                          dense
                          round
                          icon="delete"
                          @click="scope.removeFile(file)"
                        />
                      </q-item-section>
                    </q-item>
                  </q-list>
                </template>
              </q-uploader>
              <chk-rbac-role_alert-edit label-name="所属角色" class="q-mt-md" />
            </div>
          </div>
        </q-form>
      </q-card-section>
      <q-card-actions align="right">
        <q-btn
          type="button"
          label="确定"
          icon="check"
          color="secondary"
          @click="fnUpdateAccount"
          v-close-popup
        />
      </q-card-actions>
    </q-card>
  </q-dialog>
  <!-- 编辑密码对话框 -->
  <q-dialog v-model="alertEditPassword">
    <q-card style="width: 800px">
      <q-card-section>
        <div class="text-h6">编辑密码</div>
      </q-card-section>
      <q-card-section class="q-pt-none">
        <q-form class="q-gutter-md" @submit.prevent="">
          <div class="row">
            <div class="col">
              <q-input
                outlined
                clearable
                lazy-rules
                v-model="oldPassword_alertEditPassword"
                label="旧密码"
                :rules="[]"
                class="q-mb-md"
              />
              <q-input
                outlined
                clearable
                lazy-rules
                v-model="password_alertEditPassword"
                label="旧密码"
                :rules="[]"
                class="q-mb-md"
              />
              <q-input
                outlined
                clearable
                lazy-rules
                v-model="passwordConfirmation_alertEditPassword"
                label="旧密码"
                :rules="[]"
                class="q-mb-md"
              />
            </div>
          </div>
        </q-form>
      </q-card-section>
      <q-card-actions align="right">
        <q-btn
          type="button"
          label="确定"
          icon="check"
          color="negative"
          @click="fnUpdatePassword"
          v-close-popup
        />
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>
<script setup>
import { ref, onMounted, provide } from "vue";
import collect from "collect.js";
import settings from "src/settings";
import {
  ajaxGetAccounts,
  ajaxGetAccount,
  ajaxStoreAccount,
  ajaxUpdateAccount,
  ajaxDestroyAccount,
  ajaxDestroyAccounts,
  ajaxUpdateAccountPassword,
} from "src/apis/account";
import {
  loadingNotify,
  successNotify,
  errorNotify,
  actionNotify,
  destroyActions,
} from "src/utils/notify";
import { fnColumnReverseSort } from "src/utils/common";
import { getBase64 } from "src/utils/file";
import SelRbacRole_search from "src/components/SelRbacRole_search.vue";
import ChkRbacRole_alertCreate from "src/components/ChkRbacRole_alertCreate.vue";
import ChkRbacRole_alertEdit from "src/components/ChkRbacRole_alertEdit.vue";

// 用户令牌
let authToken = ref(localStorage.getItem("auth.token"));

// 搜索栏数据
const username_search = ref("");
const nickname_search = ref("");
const rbacRoleUuid_search = ref("");

// 表格数据
const columns = ref([
  {
    name: "username",
    field: "username",
    label: "用户名",
    sortable: true,
    align: "left",
  },
  {
    name: "nickname",
    field: "nickname",
    label: "昵称",
    sortable: true,
    align: "left",
  },
]);
const rows = ref([]);
const selected = ref([]);
const sortBy = ref("");

// 新建用户对话框数据
const alertCreateAccount = ref(false);
const username_alertCreateAccount = ref("");
const nickname_alertCreateAccount = ref("");
const password_alertCreateAccount = ref("");
const passwordConfirmation_alertCreateAccount = ref("");
const avatar_alertCreateAccount = ref(null);
const rbacRoleUuids_alertCreateAccount = ref([]);

// 编辑用户对话框数据
const alertEditAccount = ref(false);
const username_alertEditAccount = ref("");
const nickname_alertEditAccount = ref("");
const showAvatar_alertEditAccount = ref(null);
const avatar_alertEditAccount = ref(null);
const rbacRoleUuids_alertEditAccount = ref([]);
const currentUuid = ref("");

//  编辑密码对话框数据
const alertEditPassword = ref(false);
const oldPassword_alertEditPassword = ref("");
const password_alertEditPassword = ref("");
const passwordConfirmation_alertEditPassword = ref("");

provide("rbacRoleUuid_search", rbacRoleUuid_search);
provide("checkedRbacRoleUuids_alertCreate", rbacRoleUuids_alertCreateAccount);
provide("checkedRbacRoleUuids_alertEdit", rbacRoleUuids_alertEditAccount);

onMounted(() => {
  fnInit();
});

/**
 * 初始化页面
 */
const fnInit = () => {
  fnSearch();
};

/**
 * 重置搜索栏条件
 */
const fnResetSearch = () => {
  username_search.value = "";
  nickname_search.value = "";
  rbacRoleUuid_search.value = "";
};

/**
 * 搜索
 */
const fnSearch = () => {
  rows.value = [];
  selected.value = [];

  ajaxGetAccounts({
    "@~[]": ["Avatar", "RbacRoles"],
    username: username_search.value,
    nickname: nickname_search.value,
    rbac_role_uuid_search: rbacRoleUuid_search.value,
  })
    .then((res) => {
      rows.value = [];

      if (res.content.accounts.length > 0) {
        collect(res.content.accounts).each((account, idx) => {
          rows.value.push({
            index: idx + 1,
            uuid: account.uuid,
            username: account.username,
            nickname: account.nickname,
            avatar: account.avatar
              ? `${settings.baseUrl}/${account.avatar.url}`
              : "",
            rbac_roles: account.rbac_roles || [],
            operation: { uuid: account.uuid },
          });
        });
      }
    })
    .catch((e) => {
      errorNotify(e.msg);
    });
};

/**
 * 重置新建用户对话框
 */
const fnResetAlertCreateAccount = () => {
  username_alertCreateAccount.value = "";
  nickname_alertCreateAccount.value = "";
  password_alertCreateAccount.value = "";
  passwordConfirmation_alertCreateAccount.value = "";
  rbacRoleUuids_alertCreateAccount.value = [];
};

/**
 * 打开新建用户对话框
 */
const fnOpenAlertCreateAccount = () => {
  alertCreateAccount.value = true;
};

/**
 * 新建用户
 */
const fnStoreAccount = async () => {
  const loading = loadingNotify();

  const avatar = {
    original_filename: avatar_alertCreateAccount.value.name,
    original_extension: collect(
      avatar_alertCreateAccount.value.name.split(".")
    ).last(),
    base64: null,
  };
  if (avatar_alertCreateAccount.value) {
    avatar.base64 = await getBase64(avatar_alertCreateAccount.value);
    avatar.base64 = collect(avatar.base64.split("base64,")).last();
  }

  ajaxStoreAccount({
    username: username_alertCreateAccount.value,
    nickname: nickname_alertCreateAccount.value,
    password: password_alertCreateAccount.value,
    password_confirmation: passwordConfirmation_alertCreateAccount.value,
    rbac_role_uuids: rbacRoleUuids_alertCreateAccount.value,
    avatar,
  })
    .then((res) => {
      successNotify(res.msg);
      fnSearch();
      fnResetAlertCreateAccount();
    })
    .catch((e) => {
      errorNotify(e.msg);
    })
    .finally(() => {
      loading();
    });
};

/**
 * 重置编辑用户对话框
 */
const fnResetAlertEditAccount = () => {
  username_alertEditAccount.value = "";
  nickname_alertEditAccount.value = "";
  rbacRoleUuids_alertEditAccount.value = [];
};

/**
 * 打开编辑用户对话框
 */
const fnOpenAlertEditAccount = (params = {}) => {
  if (!params["uuid"]) return;

  currentUuid.value = params.uuid;

  ajaxGetAccount(currentUuid.value, { "@~[]": ["Avatar", "RbacRoles"] })
    .then((res) => {
      username_alertEditAccount.value = res.content.account.username;
      nickname_alertEditAccount.value = res.content.account.nickname;
      showAvatar_alertEditAccount.value = res.content.account.avatar
        ? `${settings.baseUrl}/${res.content.account.avatar.url}`
        : "";
      rbacRoleUuids_alertEditAccount.value = collect(
        res.content.account.rbac_roles
      )
        .pluck("uuid")
        .all();
      alertEditAccount.value = true;
    })
    .catch((e) => {
      errorNotify(e.msg);
    });
};

/**
 * 编辑用户
 */
const fnUpdateAccount = async () => {
  if (!currentUuid.value) return;

  const avatar = {
    original_filename: avatar_alertEditAccount.value.name,
    original_extension: collect(
      avatar_alertEditAccount.value.name.split(".")
    ).last(),
    base64: null,
  };
  if (avatar_alertEditAccount.value) {
    avatar.base64 = await getBase64(avatar_alertEditAccount.value);
    avatar.base64 = collect(avatar.base64.split("base64,")).last();
  }

  const loading = loadingNotify();

  ajaxUpdateAccount(currentUuid.value, {
    username: username_alertEditAccount.value,
    nickname: nickname_alertEditAccount.value,
    rbac_role_uuids: rbacRoleUuids_alertEditAccount.value,
    avatar,
  })
    .then((res) => {
      successNotify(res.msg);
      fnSearch();
      fnResetAlertEditAccount();
    })
    .catch((e) => {
      errorNotify(e.msg);
    })
    .finally(() => {
      loading();
    });
};

/**
 * 删除用户
 * @param {{*}} params 参数
 */
const fnDestroyAccount = (params = {}) => {
  if (!params["uuid"]) return;

  actionNotify(
    destroyActions(() => {
      const loading = loadingNotify();

      ajaxDestroyAccount(params.uuid)
        .then((res) => {
          successNotify("删除成功");
          fnSearch();
        })
        .catch((e) => {
          errorNotify(e.msg);
        })
        .finally(() => {
          loading();
        });
    })
  );
};

/**
 * 删除多用户
 */
const fnDestroyAccounts = () => {
  if (selected.value.length === 0) return;

  actionNotify(
    destroyActions(() => {
      const loading = loadingNotify();

      ajaxDestroyAccounts(collect(selected.value).pluck("uuid").all())
        .then((res) => {
          successNotify("删除成功");
          fnSearch();
        })
        .catch((e) => {
          errorNotify(e.msg);
        })
        .finally(() => {
          loading();
        });
    })
  );
};

/**
 * 重置密码对话框
 */
const fnResetAlertEditPassword = () => {
  oldPassword_alertEditPassword.value = "";
  password_alertEditPassword.value = "";
  passwordConfirmation_alertEditPassword.value = "";
};

/**
 * 打开编辑密码对话框
 */
const fnOpenAlertEditPassword = (params = {}) => {
  if (!params["uuid"]) return;
  currentUuid.value = params.uuid;
  alertEditPassword.value = true;
};

/**
 * 编辑密码
 */
const fnUpdatePassword = () => {
  if (!currentUuid.value) return;

  const loading = loadingNotify();
  ajaxUpdateAccountPassword(currentUuid.value, {
    old_password: oldPassword_alertEditPassword.value,
    password: password_alertEditPassword.value,
    password_confirmation: passwordConfirmation_alertEditPassword.value,
  })
    .then((res) => {
      successNotify(res.msg);
      fnSearch();
      fnResetAlertEditPassword();
    })
    .catch((e) => {
      errorNotify(e.msg);
    })
    .finally(() => {
      loading();
    });
};

/**
 * 上传头像完成后
 */
const fnOnUploadedAvatar = (e) => {
  const response = JSON.parse(e.xhr.response);
  if (response) successNotify(response.msg, 500, fnSearch);

  // 重置已上传头像
  showAvatar_alertEditAccount.value = `${settings.baseUrl}/${response.content.avatar.save_path}/${response.content.avatar.filename}`;
};
</script>
