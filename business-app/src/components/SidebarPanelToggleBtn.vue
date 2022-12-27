<template>
  <q-btn class="icon" icon="menu" @click="store.toggleSidebar" />
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { useSidebarStore } from 'stores/SidebarStore';

export default defineComponent({
  name: 'SidebarPanelToggleBtn',
  setup() {
    const store = useSidebarStore();

    return {
      store,
    };
  },
});
</script>
