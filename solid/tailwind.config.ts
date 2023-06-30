import type { Config } from "tailwindcss";

import daisyui from "daisyui";

const config: Config = {
	content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx,css,md,mdx,html,json,scss}"],
	// darkMode: false, // or 'media' or 'class'
	plugins: [daisyui],
	theme: {
		extend: {},
	},
};

export default config;
