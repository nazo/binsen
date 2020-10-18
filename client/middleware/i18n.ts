import Vue from 'vue';
import { Store } from 'vuex';
import { IVueI18n } from 'vue-i18n';
import { Route, Location } from 'vue-router';
import { NuxtAppOptions, NuxtError } from '@nuxt/types';
import { defineNuxtMiddleware } from '@nuxtjs/composition-api';
import { Context } from '@nuxt/types/app';

export default defineNuxtMiddleware(({ isHMR, app, store, route, params, error, redirect }: Context) => {
  const i18n = app.i18n! as IVueI18n;
  const defaultLocale = i18n.fallbackLocale;
  // If middleware is called from hot module replacement, ignore it
  if (isHMR) {
    return;
  }
  // Get locale from params
  const locale = params.lang || defaultLocale;
  if (store.state.locales) {
    if (!store.state.locales.includes(locale)) {
      return error({
        message: 'This page could not be found.',
        statusCode: 404
      });
    }
  }
  // Set locale
  store.commit('setLang', locale);
  i18n.locale = store.state.locale;
  // If route is /<defaultLocale>/... -> redirect to /...
  if (locale === defaultLocale && route.fullPath.indexOf('/' + defaultLocale) === 0) {
    const toReplace = '^/' + defaultLocale + (route.fullPath.indexOf('/' + defaultLocale + '/') === 0 ? '/' : '');
    const re = new RegExp(toReplace);
    return redirect(route.fullPath.replace(re, '/'));
  }
});
