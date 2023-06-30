import { VariantProps, cva } from "class-variance-authority";
import { ClassValue } from "clsx";
import { Component, JSX } from "solid-js";

import { classListToClassValues, cn } from "../../lib/cn";

export const inputVariants = cva("input", {
	defaultVariants: {
		bordered: false,
		size: "md",
		variant: "neutral",
	},
	variants: {
		bordered: {
			false: "",
			true: "input-bordered",
		},
		size: {
			lg: "input-lg",
			md: "input-md",
			sm: "input-sm",
			xs: "input-xs",
		},
		variant: {
			accent: "input-accent",
			error: "input-error",
			ghost: "input-ghost",
			info: "input-info",
			link: "input-link",
			neutral: "input-neutral",
			primary: "input-primary",
			secondary: "input-secondary",
			success: "input-success",
			warning: "input-warning",
		},
	},
});

export type InputVariantsProperties = VariantProps<typeof inputVariants>;

export type InputProperties = JSX.InputHTMLAttributes<HTMLInputElement> & InputVariantsProperties;

export const inputClass = (properties: InputVariantsProperties, ...className: ClassValue[]) =>
	cn(inputVariants(properties), className);

export const Input: Component<InputProperties> = (properties) => (
	<input
		{...properties}
		class={inputClass(
			properties as InputVariantsProperties,
			properties.class,
			classListToClassValues(properties.classList),
		)}
		classList={undefined}
	/>
);
