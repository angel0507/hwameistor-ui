import type { App } from 'vue';
import i18nInstall, { loadLocaleMessages, i18n, loadLanguageAsync } from './vue-i18n';
import daoStyleInstall from './dao-style';

export default function install<T>(app: App<T>) {
  app.use(daoStyleInstall);
  app.use(i18nInstall);
}

export { loadLocaleMessages, i18n, loadLanguageAsync };
