import { JSX, ParentComponent } from "solid-js";

import { classListToClassValues, cn } from "../../lib/cn";

export const Container: ParentComponent<JSX.HTMLAttributes<HTMLDivElement>> = (properties) => (
	<div
		{...properties}
		class={cn(
			"rounded-box overflow-x-auto bg-base-100 p-6",
			properties.class,
			classListToClassValues(properties.classList),
		)}
		classList={undefined}
	/>
);
