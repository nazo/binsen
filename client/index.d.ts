declare module '*.vue' {
  import Vue from 'vue'
  import '@nuxtjs/pwa'
  import '@nuxtjs/axios'
  import ''
  const _default: Vue
  export default _default
}

declare module "*.json" {
  const value: any;
  export default value;
}
