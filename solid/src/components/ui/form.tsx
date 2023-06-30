import { JSX, ParentComponent } from "solid-js";

import { classListToClassValues, cn } from "../../lib/cn";

export type FormProperties = JSX.HTMLAttributes<HTMLDivElement> & {
	error?: JSX.Element;
	actions?: JSX.Element;
	secondaryActions?: JSX.Element;
};

export const Form: ParentComponent<FormProperties> = (properties) => (
	<div
		{...properties}
		class={cn("w-full max-w-3xl space-y-4", properties.class, classListToClassValues(properties.classList))}
		classList={undefined}
	>
		{properties.error && (
			<div class="alert alert-error">
				<svg class="h-6 w-6 shrink-0 stroke-current" fill="none" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
					<path
						d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
					/>
				</svg>
				<span>{properties.error}</span>
			</div>
		)}
		{properties.children}
		<div class={cn("flex items-center pt-8", properties.secondaryActions ? "justify-between" : "justify-end")}>
			{properties.secondaryActions && <div class="flex items-center space-x-2">{properties.secondaryActions}</div>}
			<div class={"flex items-center space-x-2"}>{properties.actions}</div>
		</div>
	</div>
);
