const nodeExternals = require('webpack-node-externals');
const resolve = (dir) => require('path').join(__dirname, dir);

module.exports = {
  /*
  ** Headers of the page
  */
  head: {
    title: 'binsen',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: 'document management system' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      { rel: 'stylesheet', href: 'https://fonts.googleapis.com/css?family=Roboto:300,400,500,700|Material+Icons' }
    ]
  },
  /*
  ** Customize the progress bar color
  */
  loading: { color: '#3B8070' },
  /*
  ** Build configuration
  */
  build: {
    babel: {
      plugins: [
        'transform-decorators-legacy',
        'transform-class-properties',
        ['transform-imports', {
          'vuetify': {
            'transform': 'vuetify/es5/components/${member}',
            'preventFullImport': true
          }
        }]
      ]
    },
    vendor: [
      'vuex-class',
      'nuxt-class-component',
    ],
    extractCSS: true,
    cssSourceMap: false,
    extend (config, ctx) {
      if (ctx.isDev && ctx.isClient) {
        config.module.rules.push({
          enforce: 'pre',
          test: /\.(ts|vue)$/,
          loader: 'eslint-loader',
          exclude: /(node_modules)/
        });
      }
      if (ctx.isServer) {
        config.externals = [
          nodeExternals({
            whitelist: [/^vuetify/]
          })
        ];
      }
    }
  },
  router: {
    middleware: 'i18n'
  },
  plugins: [
    '~/plugins/axios',
    '~/plugins/vee-validate',
    '~/plugins/vuetify',
    '~/plugins/i18n',
  ],
  modules: [
    '@nuxtjs/pwa',
    '@nuxtjs/axios',
    '~/modules/typescript.js',
  ],
  manifest: {
    name: 'binsen',
    lang: 'ja'
  },
  env: {
    API_CLIENT_BASE_URI: process.env.API_CLIENT_BASE_URI,
    API_SERVER_BASE_URI: process.env.API_SERVER_BASE_URI,
  },
  css: [
    '~/assets/style/app.styl'
  ],
  vendor: [
    '~/plugins/vuetify'
  ],
  axios: {
    proxyHeaders: true,
    baseURL: process.env.API_SERVER_BASE_URI,
    browserBaseURL: process.env.API_CLIENT_BASE_URI,
    credentials: true,
    proxyHeadersIgnore: ['host', 'accept'],
  }
};
