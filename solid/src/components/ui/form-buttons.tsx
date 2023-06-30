import { JSX, ParentComponent } from "solid-js";

import { classListToClassValues, cn } from "../../lib/cn";

export const FormButtons: ParentComponent<JSX.HTMLAttributes<HTMLDivElement>> = (properties) => (
	<div
		{...properties}
		class={cn("float-right flex space-x-2 pt-4", properties.class, classListToClassValues(properties.classList))}
		classList={undefined}
	/>
);
