import { defineConfig } from "@solidjs/start/config";
import deno from "@deno/vite-plugin";

export default defineConfig({
  //ssr: false,
  vite: {
    plugins: [deno()],
  },
  server: {
    prerender: {
      routes: ["/", "/ja", "/en"],
      crawlLinks: true,
    },
  },
});
