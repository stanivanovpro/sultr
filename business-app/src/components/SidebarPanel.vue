<template>
  <q-drawer v-model="store.isSidebarOpen" show-if-above>
    <q-card class="info" flat>
      <q-card-section>
        <div
          class="q-mb-md"
          @click="
            $root.$i18n.locale = 'en-US';
            store.isSidebarOpen = false;
          "
        >
          <span class="fi fi-us q-mr-sm"></span>United States
        </div>
        <div
          @click="
            $root.$i18n.locale = 'ru-RU';
            store.isSidebarOpen = false;
          "
        >
          <span class="fi fi-ru q-mr-sm"></span>Русский
        </div>
      </q-card-section>
    </q-card>
  </q-drawer>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { useSidebarStore } from 'stores/SidebarStore';

export default defineComponent({
  name: 'SidebarPanel',
  setup() {
    const store = useSidebarStore();

    return {
      store,
    };
  },
});
</script>
