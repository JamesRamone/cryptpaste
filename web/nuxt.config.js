const pkg = require('./package')
import purgecss from '@fullhuman/postcss-purgecss'


module.exports = {
  mode: 'spa',
  /*
  ** Headers of the page
  */
  head: {
    title: pkg.name,
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: pkg.description }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },

  /*
  ** Customize the progress-bar color
  */
  loading: { color: '#fff' },

  /*
  ** Global CSS
  */
  css: [
    '~/assets/css/tailwind.css'
  ],

  /*
  ** Plugins to load before mounting the App
  */
  plugins: [
  ],

  /*
  ** Nuxt.js modules
  */
  modules: [
    // Doc: https://github.com/nuxt-community/axios-module#usage
    '@nuxtjs/axios',
    '@nuxtjs/proxy',
    'nuxt-purgecss',
  ],
  /*
  ** Axios module configuration
  */
  axios: {
    // See https://github.com/nuxt-community/axios-module#options
    // proxy: true
    // baseUrl: 'api.cryptpaste.xyz',
    browserBaseUrl: 'https://api.cryptpaste.xyz'
  },
  proxy: {
    // '/api': 'https://api.cryptpaste.xyz/'
  },
  
  purgeCSS: {
    content: ['./pages/**/*.vue', './layouts/**/*.vue', './components/**/*.vue'],
    whitelist: ['html', 'body'],
  },
  /*
  ** Build configuration
  */
  build: {
    /*
    ** You can extend webpack config here
    */
    extend(config, ctx) {
     
    }
  },

}
