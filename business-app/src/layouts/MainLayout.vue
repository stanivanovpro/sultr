<template>
  <q-layout view="hHh lpR fFf">
    <SidebarPanel />
    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import SidebarPanel from 'components/SidebarPanel.vue';

export default defineComponent({
  name: 'MainLayout',
  components: {
    SidebarPanel,
  },
});
</script>
