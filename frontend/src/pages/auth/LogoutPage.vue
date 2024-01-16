<template>
  <q-page class="flex flex-center">
    <img alt="Quasar logo" src="~assets/quasar-logo-vertical.svg" style="width: 200px; height: 200px" />
    <h3>已安全退出</h3>
  </q-page>
</template>

<script setup>
import { onMounted } from "vue";
import { ajaxLogout } from "src/apis/auth";

onMounted(() => {
  ajaxLogout();
  localStorage.removeItem("auth.token");
});
</script>
