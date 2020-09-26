import Vue from 'vue';
import { AxiosStatic } from 'axios';
import VueRouter, { Route } from 'vue-router';
import { Validator, VeeValidateComponentOptions } from 'vee-validate';

declare module '*.vue' {
  const _default: Vue;
  export default _default;
}

declare module 'vue/types/vue' {
  interface Vue {
    $axios: AxiosStatic;
    $router: VueRouter;
    $route: Route;
    $validator: Validator;
  }
}

declare module 'vue/types/options' {
  interface ComponentOptions<V extends Vue> {
    middleware?: string | string[];
    layout?: string | string[];
    $_veeValidate?: VeeValidateComponentOptions;
  }
}
