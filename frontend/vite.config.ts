import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  resolve: {
    alias: {
      $lib: path.resolve("./src/lib"),
      $conf: path.resolve("./src/conf"),
      $store: path.resolve("./src/store"),
      $app: path.resolve("./src"),
      $services: path.resolve("./src/services"),
    },
  },
})
