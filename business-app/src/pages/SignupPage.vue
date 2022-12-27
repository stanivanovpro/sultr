<template>
  <div>
    <SignupForm />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import SignupForm from 'components/SignupForm.vue';

export default defineComponent({
  name: 'SignupPage',
  components: { SignupForm },
});
</script>
