import { boot } from 'quasar/wrappers';
import * as VueI18n from 'vue-i18n';
import messages from 'src/i18n';

export default boot(async ({ app }) => {
  console.log(navigator.language);
  const i18n = VueI18n.createI18n({
    locale: navigator.language,
    fallbackLocale: 'en-US',
    globalInjection: true,
    missingWarn: false,
    fallbackWarn: false,
    legacy: false,
    messages,
  });
  app.use(i18n);
});
