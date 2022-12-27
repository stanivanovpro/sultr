<template>
  <q-page>
    <div class="row landing-header">
      <img
        alt="Sultr business logo"
        class="logo"
        src="~/assets/images/logo-light.svg"
      />
      <q-space />
      <SidebarPanelToggleBtn />
    </div>
    <div class="row hero">
      <p class="q-mb-lg">{{ $t('pages.landing.hero_text') }}</p>
      <p class="text-accent">{{ $t('pages.landing.hero_text_accented') }}</p>
      <q-btn class="full-width q-mt-lg cta" to="/signup">{{
        $t('pages.landing.register_cta')
      }}</q-btn>
    </div>
    <div class="hero-image"></div>
  </q-page>
</template>

<script lang="ts">
import SidebarPanelToggleBtn from 'components/SidebarPanelToggleBtn.vue';

import { defineComponent } from 'vue';

export default defineComponent({
  name: 'IndexPage',
  components: {
    SidebarPanelToggleBtn,
  },
});
</script>
