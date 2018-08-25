import Vue from 'vue';
import VueI18n, { LocaleMessages } from 'vue-i18n';

Vue.use(VueI18n);

/**
 * load Locale Messages
 * @returns {LocaleMessages} locale messages
 */
function loadLocaleMessages (): LocaleMessages {
  const locales = require.context('../locales', true, /[A-Za-z0-9-_,\s]+\.json$/i);
  const messages: LocaleMessages = {};
  locales.keys().forEach(key => {
    const matched = key.match(/([a-z0-9]+)\./i);
    if (matched && matched.length > 1) {
      const locale = matched[1];
      messages[locale] = locales(key);
    }
  });
  return messages;
}

export default ({ app, store }) => {
  // Set i18n instance on app
  // This way we can use it in middleware and pages asyncData/fetch
  app.i18n = new VueI18n({
    locale: store.state.locale,
    fallbackLocale: 'en',
    messages: loadLocaleMessages()
  });
};
