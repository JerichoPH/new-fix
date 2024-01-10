<template>
  <q-page class="flex flex-center">
    <!-- <img alt="Quasar logo" src="~assets/quasar-logo-vertical.svg" style="width: 200px; height: 200px" /> -->
    <q-form
      @submit="fnLogin"
      @reset="fnReset"
      class="q-gutter-md"
      style="width: 600px"
    >
      <q-input
        filled
        v-model="username"
        label="账号 *"
        hint=""
        lazy-rules
        :rules="[(val) => (val && val.length >= 2) || '不能小于2位']"
      />

      <q-input
        filled
        type="password"
        v-model="password"
        label="密码 *"
        lazy-rules
        :rules="[(val) => (val && val.length >= 6) || '不能小于6位']"
      />

      <q-toggle v-model="accept" label="接受条款" />

      <div>
        <q-btn label="登录" type="submit" color="primary" />
        <q-btn
          label="重置表单"
          type="reset"
          color="primary"
          flat
          class="q-ml-sm"
        />
        <q-btn
          label="注册"
          color="primary"
          flat
          class="q-ml-sm"
          @click="fnHrefTo('/auth/register')"
        />
      </div>
    </q-form>
  </q-page>
</template>

<script>
import { defineComponent } from "vue";
import { loadingNotify, successNotify, errorNotify } from "src/utils/notify.js";
import { ajaxLogin } from "src/apis/auth.js";

export default defineComponent({
  name: "LoginPage",

  data() {
    return {
      username: "",
      password: "",
      accept: false,
    };
  },

  methods: {
    /**
     * 登录
     */
    fnLogin() {
      const loading = loading("登录中……");

      ajaxLogin({
        username: this.username,
        password: this.password,
      })
        .then(res => {
          if (res.content.token) {
            loading();
            localStorage.setItem("auth.token", res.content.token);
            localStorage.setItem("auth.account", res.content.account);
            successNotify(
              res.msg,
              500,
              (router) => {
                router.push("/");
              },
              this.$router
            );
          }
        })
        .catch((e) => {
          loading();
          errorNotify(e.msg, 5000);
        });
    },
    /**
     * 重置表单
     */
    fnReset() {
      this.username = "";
      this.password = "";
      this.accept = false;
    },
    /**
     * 跳转页面
     * @param {string} path
     */
    fnHrefTo(path) {
      if (path) this.$router.push(path);
    },
  },
});
</script>
../../utils/notify.js
