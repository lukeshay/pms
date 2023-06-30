import { JSX, ParentComponent, Show } from "solid-js";

import { Container } from "../ui/container";

export type ProfileSectionProperties = {
	description?: JSX.Element;
	header: JSX.Element;
};

export const ProfileSection: ParentComponent<ProfileSectionProperties> = (properties) => (
	<Container class="grid grid-cols-2">
		<div class="col-span-1">
			<h1 class="text-2xl font-bold">{properties.header}</h1>
			<Show when={properties.description}>
				<p class="pt-2 text-gray-500">{properties.description}</p>
			</Show>
		</div>
		<div class="col-span-1">{properties.children}</div>
	</Container>
);
