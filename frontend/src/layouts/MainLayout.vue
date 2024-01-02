<template>
  <q-layout view="hHh lpR fFf" class="bg-grey-1">
    <q-header elevated class="bg-primary text-white q-py-xs" height-hint="58">
      <q-toolbar>
        <q-btn flat dense round @click="fnToggleLeftDrawer" aria-label="Menu" icon="menu" />

        <q-btn flat no-caps no-wrap class="q-ml-xs" v-if="$q.screen.gt.xs">
          <q-icon :name="fabYoutube" color="red" size="28px" />
          <q-toolbar-title shrink class="text-weight-bold">
            YouTube
          </q-toolbar-title>
        </q-btn>

        <q-space />

        <div class="YL__toolbar-input-container row no-wrap">
          <q-input dense outlined square v-model="search" placeholder="Search" class="bg-white col" />
          <q-btn class="YL__toolbar-input-btn" color="white" text-color="black" icon="search" unelevated />
        </div>

        <q-space />

        <div class="q-gutter-sm row items-center no-wrap">
          <q-btn round dense flat color="white" icon="video_call" v-if="$q.screen.gt.sm">
            <q-tooltip>Create a video or post</q-tooltip>
          </q-btn>
          <q-btn round dense flat color="white" icon="apps" v-if="$q.screen.gt.sm">
            <q-tooltip>Apps</q-tooltip>
          </q-btn>
          <q-btn round dense flat color="white" icon="message" v-if="$q.screen.gt.sm">
            <q-tooltip>Messages</q-tooltip>
          </q-btn>
          <q-btn round dense flat color="white" icon="notifications">
            <q-badge color="red" text-color="white" floating> 2 </q-badge>
            <q-tooltip>Notifications</q-tooltip>
          </q-btn>
          <q-btn round flat>
            <q-avatar size="26px">
              <img src="https://cdn.quasar.dev/img/boy-avatar.png" />
            </q-avatar>
            <q-tooltip>Account</q-tooltip>
          </q-btn>
        </div>
      </q-toolbar>
    </q-header>

    <q-drawer v-model="leftDrawerOpen" show-if-above bordered :width="240">
      <q-scroll-area class="fit">
        <q-list link dense separator style="min-width: 200px">
          <template v-for="rbacMenu in rbacMenus" :key="rbacMenu.uuid">
            <template v-if="rbacMenu.subs.length > 0">
              <q-expansion-item group="menu" :icon="rbacMenu.icon" :label="rbacMenu.name">
                <q-list dense>
                  <q-item v-for="subMenu in rbacMenu.subs" clickable v-ripple :key="subMenu.uuid"
                    @click="fnHref(subMenu.uri)">
                    <q-item-section avatar>
                      <q-icon :name="subMenu.icon" />
                    </q-item-section>
                    <q-item-section>{{ subMenu.name }}</q-item-section>
                  </q-item>
                </q-list>
              </q-expansion-item>
            </template>
            <template v-else>
              <q-item clickable v-ripple @click="fnHref(rbacMenu.uri)" :key="rbacMenu.uuid">
                <q-item-section avatar>
                  <q-icon :name="rbacMenu.icon" />
                </q-item-section>
                <q-item-section>{{ rbacMenu.name }}</q-item-section>
              </q-item>
            </template>
          </template>
        </q-list>
      </q-scroll-area>
    </q-drawer>
    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import { fabYoutube } from "@quasar/extras/fontawesome-v6";
import { errorNotify } from "src/utils/notify";
import { ajaxGetCurrentAccountMenus } from "src/apis/auth";
import collect from "collect.js";

const router = useRouter();

// 搜索栏数据
const search = ref("");

// 做侧边栏数据
const leftDrawerOpen = ref(true);
const rbacMenus = ref([]);

onMounted(() => fnInit());

/**
 * 初始化
 */
const fnInit = () => {
  if (!localStorage.getItem("auth.token")) {
    errorNotify(
      "未登录",
      router => {
        router.push("/auth/login");
      },
      router
    );
  }

  // 加载当前用户菜单
  ajaxGetCurrentAccountMenus({
    "@~[]": ["Subs"],
    "@== []": ["parent_uuid"],
  })
    .then(res => {
      rbacMenus.value = collect(res.content.rbac_menus).all();
    })
    .catch((e) => errorNotify(e.msg));
};
/**
 * 切换左侧边栏
 */
const fnToggleLeftDrawer = () => {
  leftDrawerOpen.value = !leftDrawerOpen.value;
};

/**
 * 跳转页面
 * @param {string} uri 路由
 */
const fnHref = (uri) => {
  router.push(uri);
};
</script>

<style lang="sass">
.YL

  &__toolbar-input-container
    min-width: 100px
    width: 55%

  &__toolbar-input-btn
    border-radius: 0
    border-style: solid
    border-width: 1px 1px 1px 0
    border-color: rgba(0,0,0,.24)
    max-width: 60px
    width: 100%

  &__drawer-footer-link
    color: inherit
    text-decoration: none
    font-weight: 500
    font-size: .75rem

    &:hover
      color: #000
</style>
