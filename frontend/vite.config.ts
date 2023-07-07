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
			sass: {
				additionalData: `
					@import '$styles/global'
				`,
			}
		},
	},
	resolve: {
		alias: {
			$: resolve(__dirname, "./src"),
			$styles: resolve(__dirname, "./src/styles"),
			$assets: resolve(__dirname, "./src/assets"),
			$components: resolve(__dirname, "./src/components"),
		},
	},
});
