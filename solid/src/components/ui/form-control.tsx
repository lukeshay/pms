import { JSX, ParentComponent } from "solid-js";
import { ClassList } from "../../types/solid-types";
import { classListToClassValues, cn } from "../../lib/cn";

export type FormControlProperties = {
	error?: null | string | string[];
	label: JSX.Element;
	name: string;
	class?: string;
	classList?: ClassList;
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
