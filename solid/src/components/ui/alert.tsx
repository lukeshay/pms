import { VariantProps, cva } from "class-variance-authority";
import { ClassValue } from "clsx";
import { Component, JSX } from "solid-js";

import { classListToClassValues, cn } from "../../lib/cn";
import { ClassList } from "../../types/solid-types";

export const alertVariants = cva("alert w-fit max-w-3xl", {
	defaultVariants: {
		variant: "neutral",
	},
	variants: {
		variant: {
			error: "alert-error",
			info: "alert-info",
			neutral: "",
			success: "alert-success",
			warning: "alert-warning",
		},
	},
});

export type AlertVariantsProperties = VariantProps<typeof alertVariants>;

export type AlertProperties = AlertVariantsProperties & {
	buttons?: JSX.Element;
	class?: string;
	classList?: ClassList;
	description: string;
	title?: string;
};

export const alertClass = (properties: AlertVariantsProperties, ...className: ClassValue[]) =>
	cn(alertVariants(properties), className);

const ALERT_VARIANT_ICON_MAP: Record<NonNullable<AlertVariantsProperties["variant"]>, JSX.Element> = {
	error: (
		<svg class="h-6 w-6 shrink-0 stroke-current" fill="none" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
			<path
				d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
				stroke-linecap="round"
				stroke-linejoin="round"
				stroke-width="2"
			/>
		</svg>
	),
	info: (
		<svg class="h-6 w-6 shrink-0 stroke-current" fill="none" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
			<path
				d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
				stroke-linecap="round"
				stroke-linejoin="round"
				stroke-width="2"
			/>
		</svg>
	),
	neutral: (
		<svg class="h-6 w-6 shrink-0 stroke-info" fill="none" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
			<path
				d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
				stroke-linecap="round"
				stroke-linejoin="round"
				stroke-width="2"
			/>
		</svg>
	),
	success: (
		<svg class="h-6 w-6 shrink-0 stroke-current" fill="none" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
			<path
				d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
				stroke-linecap="round"
				stroke-linejoin="round"
				stroke-width="2"
			/>
		</svg>
	),
	warning: (
		<svg class="h-6 w-6 shrink-0 stroke-current" fill="none" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
			<path
				d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
				stroke-linecap="round"
				stroke-linejoin="round"
				stroke-width="2"
			/>
		</svg>
	),
};

export const Alert: Component<AlertProperties> = (properties) => (
	<div
		class={alertClass(
			properties as AlertVariantsProperties,
			properties.class,
			classListToClassValues(properties.classList),
		)}
	>
		<svg class="h-6 w-6 shrink-0 stroke-current" fill="none" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
			{properties.variant ? ALERT_VARIANT_ICON_MAP[properties.variant] : ALERT_VARIANT_ICON_MAP.neutral}
		</svg>
		{properties.title ? (
			<div>
				<h3 class="font-bold">{properties.title}</h3>
				<div class="text-xs">{properties.description}</div>
			</div>
		) : (
			<span>{properties.description}</span>
		)}
		{properties.buttons && <div class="pl-4">{properties.buttons}</div>}
	</div>
);
