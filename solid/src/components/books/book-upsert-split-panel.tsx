import { createForm } from "@felte/solid";
import { validator } from "@felte/validator-zod";
import { ModelsBook } from "@pms/api";
import { Component } from "solid-js";
import { BookUpsertSchema, bookUpsertSchema } from "../../schemas/book-schema";
import { Form } from "../ui/form";
import { FormButtons } from "../ui/form-buttons";
import { FormControl } from "../ui/form-control";
import { SplitPanel } from "../ui/split-panel";
import { Button } from "../ui/button";

export type BookUpsertSplitPanelProperties = {
	book?: ModelsBook;
	error?: string;
	onCancel: () => Promise<void> | void;
	onSubmit: (values: BookUpsertSchema) => Promise<void> | void;
};

export const BookUpsertSplitPanel: Component<BookUpsertSplitPanelProperties> = (properties) => {
	const { errors, form } = createForm<BookUpsertSchema>({
		extend: validator({ schema: bookUpsertSchema }),
		initialValues: properties.book ?? {},
		onSubmit: async (values) => {
			await properties.onSubmit({
				...values,
				finishedAt: values.finishedAt === "" ? undefined : values.finishedAt,
				purchasedAt: values.purchasedAt === "" ? undefined : values.purchasedAt,
				rating: Number(values.rating) === 0 || Number.isNaN(Number(values.rating)) ? undefined : values.rating,
			});
		},
	});

	return (
		<SplitPanel onSplitPanelClose={properties.onCancel} header={`${properties.book ? "Update" : "Create"} Book`}>
			<form use:form>
				<Form
					actions={
						<>
							<Button
								formAction="none"
								onClick={async () => {
									await properties.onCancel();
								}}
							>
								{"Cancel"}
							</Button>
							<Button type="submit" variant="primary">
								{properties.book ? "Update" : "Create"}
							</Button>
						</>
					}
					error={properties.error}
					secondaryActions={properties.book && <Button formAction="none">{"Delete"}</Button>}
				>
					<FormControl error={errors("title")} label="Title" name="title">
						<input class="input-bordered input" name="title" placeholder="title" />
					</FormControl>
					<FormControl error={errors("author")} label="Author" name="author">
						<input class="input-bordered input" name="author" placeholder="author" />
					</FormControl>
					<FormControl error={errors("rating")} label="Rating" name="rating">
						<input class="input-bordered input" name="rating" placeholder="rating" type="number" />
					</FormControl>
					<FormControl error={errors("purchasedAt")} label="Purchased Date" name="purchasedAt">
						<input class="input-bordered input" name="purchasedAt" placeholder="purchasedAt" type="date" />
					</FormControl>
					<FormControl error={errors("finishedAt")} label="Completed Date" name="finishedAt">
						<input class="input-bordered input" name="finishedAt" placeholder="finishedAt" type="date" />
					</FormControl>
				</Form>
			</form>
		</SplitPanel>
	);
};
