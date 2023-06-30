import { VariantProps, cva } from "class-variance-authority";
import { ClassValue } from "clsx";
import { Component, JSX } from "solid-js";

import { classListToClassValues, cn } from "../../lib/cn";

export const buttonVariants = cva("btn", {
	defaultVariants: {
		active: false,
		block: false,
		disabled: false,
		loading: false,
		outline: false,
		shape: "default",
		size: "md",
		variant: "neutral",
		wide: false,
	},
	variants: {
		active: {
			false: "",
			true: "btn-active",
		},
		block: {
			false: "",
			true: "btn-block",
		},
		disabled: {
			false: "",
			true: "btn-disabled",
		},
		loading: {
			false: "",
			true: "btn-disabled",
		},
		outline: {
			false: "",
			true: "btn-outline",
		},
		shape: {
			circle: "btn-circle",
			default: "",
			square: "btn-square",
		},
		size: {
			lg: "btn-lg",
			md: "btn-md",
			sm: "btn-sm",
			xs: "btn-xs",
		},
		variant: {
			accent: "btn-accent",
			error: "btn-error",
			ghost: "btn-ghost",
			info: "btn-info",
			link: "btn-link",
			neutral: "btn-neutral",
			primary: "btn-primary",
			secondary: "btn-secondary",
			success: "btn-success",
			warning: "btn-warning",
		},
		wide: {
			false: "",
			true: "btn-wide",
		},
	},
});

export type ButtonVariantsProperties = VariantProps<typeof buttonVariants>;

export type ButtonProperties = JSX.ButtonHTMLAttributes<HTMLButtonElement> & ButtonVariantsProperties;

export const buttonClass = (properties: ButtonVariantsProperties, ...className: ClassValue[]) =>
	cn(buttonVariants(properties), className);

export const Button: Component<ButtonProperties> = (properties) => (
	<button
		{...properties}
		class={buttonClass(
			properties as ButtonVariantsProperties,
			properties.class,
			classListToClassValues(properties.classList),
		)}
		classList={undefined}
	>
		{properties.loading && <span class="loading loading-spinner" />}
		{properties.children}
	</button>
);
