import { JSX, ParentComponent } from "solid-js";

import { classListToClassValues, cn } from "../../lib/cn";
import { ClassList } from "../../types/solid-types";

export type FormControlProperties = {
	class?: string;
	classList?: ClassList;
	error?: null | string | string[];
	label: JSX.Element;
	name: string;
};

export const FormControl: ParentComponent<FormControlProperties> = (properties) => (
	<div class={cn("form-control", properties.class, classListToClassValues(properties.classList))}>
		<label class="label" for={properties.name}>
			<span class="label-text">{properties.label}</span>
		</label>
		{properties.children}
		{properties.error && (
			<label class="label" for={properties.name}>
				<span class="label-text-alt text-error">
					{typeof properties.error === "string" ? properties.error : properties.error[0]}
				</span>
			</label>
		)}
	</div>
);
