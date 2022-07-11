import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
const path = require('path')


function resolve(dir) {
	return path.join(__dirname, '..', dir)
}

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    AutoImport({
      resolvers: [ElementPlusResolver()],
    }),
    Components({
      resolvers: [ElementPlusResolver()],
    }),
  ],
  server: {
    port: 3001,
    strictPort: true,
    CORS: true,
    proxy: {
        '/api': {
          target: 'http://localhost:9001',
          changeOrigin: true,
        },
    },
    disableHostCheck: true,
  },
  resolve: {
    alias: {
      '@': resolve('src'),
		}
  }
})
