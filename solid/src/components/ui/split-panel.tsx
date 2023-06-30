import { TbX } from "solid-icons/tb";
import { JSX, ParentComponent } from "solid-js";

import { Button } from "./button";

export type SplitPanelProperties = {
	header: JSX.Element;
	onSplitPanelClose?: () => void;
	width?: string;
};

export const SplitPanel: ParentComponent<SplitPanelProperties> = (properties) => {
	return (
		<div class="min-w-[30rem] border-l border-base-content/10 px-4 py-6">
			<div class="flex items-center justify-between border-b border-base-content/10 pb-6">
				<h3 class="text-xl font-bold">{properties.header}</h3>
				<Button
					onClick={() => {
						properties.onSplitPanelClose?.();
					}}
					class="flex items-center justify-center"
					shape="circle"
					size="sm"
					variant="ghost"
				>
					<TbX size={20} />
				</Button>
			</div>
			<div class="pt-8">{properties.children}</div>
		</div>
	);
};
