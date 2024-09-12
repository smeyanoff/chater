const { defineConfig } = require('@vue/cli-service')

module.exports = defineConfig({
  transpileDependencies: true,
  publicPath: '/',  // Убедитесь, что путь к ресурсам правильный
  outputDir: 'dist',  // Директория, куда будут складываться сгенерированные файлы
  assetsDir: 'static',  // Директория для хранения статических файлов, таких как JS и CSS
  productionSourceMap: false,  // Отключаем карты исходников для продакшена
})