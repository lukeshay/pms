module.exports = {
	env: {
		browser: true,
		es2021: true,
	},
	extends: [
		"eslint:recommended",
		"plugin:@typescript-eslint/recommended",
		"plugin:solid/typescript",
		"plugin:import/recommended",
		"plugin:import/typescript",
		"plugin:unicorn/recommended",
		"plugin:promise/recommended",
		"plugin:security/recommended",
		"plugin:perfectionist/recommended-natural",
		// "plugin:@microsoft/sdl/required",
		// "plugin:@microsoft/sdl/typescript",
		// "plugin:@microsoft/sdl/recommended",
		"plugin:regexp/recommended",
		"plugin:sonar/recommended",
	],
	overrides: [
		{
			env: {
				node: true,
			},
			files: [".eslintrc.{js,cjs}"],
			parserOptions: {
				sourceType: "script",
			},
		},
	],
	parser: "@typescript-eslint/parser",
	parserOptions: {
		ecmaVersion: "latest",
		sourceType: "module",
	},
	plugins: [
		"@typescript-eslint",
		"solid",
		"import",
		"unicorn",
		"promise",
		"security",
		"perfectionist",
		"regexp",
		"sonar",
	],
	settings: {
		"import/resolver": {
			node: true,
			typescript: true,
		},
	},
};
