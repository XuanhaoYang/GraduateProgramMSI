module.exports = {
  transpileDependencies: [
    'vuetify'
  ],
  lintOnSave: process.env.NODE_ENV !== 'production',
  assetsDir: 'static',
  publicPath:'./',
  chainWebpack: config => {
    config.plugin('html').tap(args => {
      args[0].title = 'Welcome to MSIsensor-RNA'
      return args
    })
  },
  productionSourceMap: false

}
