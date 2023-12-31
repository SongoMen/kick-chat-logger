import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vitest/config';
import { resolve } from "path";

export default defineConfig({
	plugins: [sveltekit()],
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	},
	css: {
		preprocessorOptions: {
			scss: {
				additionalData: `@import '$styles/global';`,
			}
		},
	},
	resolve: {
		alias: {
			$: resolve(__dirname, "./src"),
			$lib: resolve(__dirname, "./src/lib"),
			$types: resolve(__dirname, "./src/types"),
			$styles: resolve(__dirname, "./src/styles"),
			$assets: resolve(__dirname, "./src/assets"),
			$components: resolve(__dirname, "./src/components"),
		},
	},
});
