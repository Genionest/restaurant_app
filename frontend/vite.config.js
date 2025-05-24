import { defineConfig } from 'vite'
import path from 'path'

export default defineConfig({
  root: './src',
  publicDir: '../public',
  build: {
    outDir: 'dist',
    assetsDir: 'assets',
    rollupOptions: {
    input: {
        main: path.resolve(__dirname, 'src/index.html')
    }
    }
  },
  server: {
    port: 5173
  }
})
