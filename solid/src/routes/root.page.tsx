import { useRouteData } from "@solidjs/router";
import { type Component, For, JSX, createSignal } from "solid-js";

import { Button } from "../components/ui/button";
import { useToast } from "../components/ui/toast";
import { BaseLayout } from "../layouts/base-layout";
import { useGlobalStore } from "../provider";
import rootLoader from "./root.loader";
import { BookUpsertSplitPanel } from "../components/books/book-upsert-split-panel";
import { Container } from "../components/ui/container";

const RootPage: Component = () => {
	const [formError, setFormError] = createSignal<string>();
	const [selected, setSelected] = createSignal<string>();
	const [splitPanel, setSplitPanel] = createSignal<JSX.Element>();

	const { booksApi } = useGlobalStore();
	const { data, refetch } = useRouteData<typeof rootLoader>();
	const toast = useToast();

	const handleCreateClick = () => {
		setSplitPanel(
			<BookUpsertSplitPanel
				onCancel={() => {
					setSplitPanel();
				}}
				onSubmit={(book) => {
					setFormError();

					booksApi
						.v1BooksPost({
							book,
						})
						.then(() => {
							refetch();
							setSplitPanel();
							toast.success({ description: "Book created" });

							return true;
						})
						.catch((error) => {
							setFormError(`An unexpected error occurred: ${(error as Error).message}`);
						});
				}}
				error={formError()}
			/>,
		);
	};

	const handleChange: JSX.ChangeEventHandlerUnion<HTMLInputElement, Event> = ({ target }) => {
		setSelected(target.value);

		const book = data()!.books.find((book) => book.id === target.value)!;

		setSplitPanel(
			<BookUpsertSplitPanel
				onCancel={() => {
					setSplitPanel();
				}}
				onSubmit={(book) => {
					setFormError();

					booksApi
						.v1BooksIdPut({
							book,
							id: selected()!,
						})
						.then(() => {
							refetch();
							setSplitPanel();
							toast.success({ description: "Book updated" });

							return true;
						})
						.catch((error) => {
							setFormError(`An unexpected error occurred: ${(error as Error).message}`);
						});
				}}
				book={book}
				error={formError()}
			/>,
		);
	};

	return (
		<BaseLayout
			onSplitPanelClose={() => {
				setSplitPanel();
			}}
			splitPanel={splitPanel()}
		>
			<Container>
				<div class="flex w-full justify-between px-4 pb-4 pt-2">
					<h2 class="text-3xl font-bold">{"My Books"}</h2>
					<Button onClick={handleCreateClick} variant="primary">
						{"Create"}
					</Button>
				</div>
				<table class="table">
					<thead>
						<tr>
							<th />
							<th>{"Title"}</th>
							<th>{"Author"}</th>
							<th>{"Rating"}</th>
						</tr>
					</thead>
					<tbody>
						<For each={data()!.books}>
							{(book) => (
								<tr class="hover cursor-pointer">
									<td>
										<input
											checked={selected() === book.id}
											class="radio"
											onChange={handleChange}
											type="radio"
											value={book.id}
										/>
									</td>
									<td>{book.title}</td>
									<td>{book.author}</td>
									<td>{book.rating}/5</td>
								</tr>
							)}
						</For>
					</tbody>
				</table>
			</Container>
		</BaseLayout>
	);
};

export default RootPage;
