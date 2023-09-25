const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    proxy: {
      '^/api': {
        target: 'https://mdm-auth-test.appleseed.kz',
        changeOrigin: true,
        secure: false,
        pathRewrite: { '^/api': '/api' },
        logLevel: 'debug'
      },
      '^/references': {
        target: 'https://mdm-dynamic-api-staging.appleseed.kz',
        changeOrigin: true,
        secure: false,
        pathRewrite: { '^/references': '/api' },
        logLevel: 'debug'
      }
    }
  }
})
