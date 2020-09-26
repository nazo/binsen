export default {
  /*
   ** Nuxt rendering mode
   ** See https://nuxtjs.org/api/configuration-mode
   */
  mode: 'universal',
  /*
   ** Nuxt target
   ** See https://nuxtjs.org/api/configuration-target
   */
  target: 'server',
  /*
   ** Headers of the page
   ** See https://nuxtjs.org/api/configuration-head
   */
  head: {
    titleTemplate: '%s - ' + process.env.npm_package_name,
    title: process.env.npm_package_name || '',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      {
        hid: 'description',
        name: 'description',
        content: process.env.npm_package_description || '',
      },
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }],
  },
  /*
   ** Global CSS
   */
  css: [],
  /*
   ** Plugins to load before mounting the App
   ** https://nuxtjs.org/guide/plugins
   */
  router: {
    middleware: 'i18n'
  },
  plugins: [
    '~/plugins/axios',
    '~/plugins/vee-validate',
    '~/plugins/i18n',
  ],
  modules: [
    '@nuxtjs/pwa',
    '@nuxtjs/axios'
  ],
  manifest: {
    name: 'binsen',
    lang: 'ja'
  },
  env: {
    API_CLIENT_BASE_URI: process.env.API_CLIENT_BASE_URI,
    API_SERVER_BASE_URI: process.env.API_SERVER_BASE_URI,
  },
  /*
   ** Auto import components
   ** See https://nuxtjs.org/api/configuration-components
   */
  components: true,
  /*
   ** Nuxt.js dev-modules
   */
  buildModules: ['@nuxt/typescript-build', '@nuxtjs/vuetify'],
  /*
   ** Nuxt.js modules
   */
  modules: [
    // Doc: https://axios.nuxtjs.org/usage
    '@nuxtjs/axios',
  ],
  /*
   ** Axios module configuration
   ** See https://axios.nuxtjs.org/options
   */
  axios: {
    proxyHeaders: true,
    baseURL: process.env.API_SERVER_BASE_URI,
    browserBaseURL: process.env.API_CLIENT_BASE_URI,
    credentials: true,
    proxyHeadersIgnore: ['host', 'accept'],
  },
  /*
   ** vuetify module configuration
   ** https://github.com/nuxt-community/vuetify-module
   */
  vuetify: {
  },
  /*
   ** Build configuration
   ** See https://nuxtjs.org/api/configuration-build/
   */
  build: {
    transpile: [
      "vee-validate/dist/rules"
    ],
  },
}
