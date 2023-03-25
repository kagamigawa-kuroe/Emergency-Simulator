const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true
})

// module.exports={
//   devServer:{
//   proxy:'http://localhost:8082/'// 配置访问的服务器地址
//   // ["/getinfo"]:{
//   //   target:'http://localhost:8082',
//   //   changeOrigin:true,
//   //   pathRewrite: {
//   //     '^/geinfo': ''
//   //   }
//   }
//
// }
