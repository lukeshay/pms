import { FC, ReactEventHandler, ReactNode, useRef, useState } from "react";
import { useLoaderData } from "../hooks/use-loader-data";
import { BaseLayout } from "../layouts/base-layout";
import { withUserData } from "../helpers/with-user-data";
import { booksApi } from "../lib/apis";
import { FormControl } from "../components/form-control";
import { Form } from "../components/form";
import { FormButtons } from "../components/form-buttons";
import { useForm } from "react-hook-form";
import { ActionFunction, useRevalidator } from "react-router-dom";
import { cn } from "../lib/cn";
import { ModelsBook } from "../api";

const loader = withUserData(async () => {
	return booksApi.v1BooksGet();
});

type CreateBookInputs = {
	title: string;
	author: string;
	rating: number;
	purchasedAt: Date;
	finishedAt: Date;
};

const UpsertBookSplitPanel: FC<{ onSuccess: () => Promise<void> | void; onCancel: () => Promise<void> | void; book?: ModelsBook }> = ({
	onSuccess,
	onCancel,
	book
}) => {
	const [error, setError] = useState<string | undefined>(undefined);
	const {
		register,
		handleSubmit,
		formState: { errors },
	} = useForm<CreateBookInputs>();

	const onSubmit = handleSubmit(async (data) => {
		try {
			if (book) {
				await booksApi.v1BooksIdPut({id: book.id, book: data})
			} else {
				await booksApi.v1BooksPost({ book: data });
			}

			onSuccess();
		} catch {
			setError(`Failed to ${book ? "update" : "create"} book. Please try again.`);
		}
	});

	return (
		<div>
			<div className="pb-4 pt-2">
				<h3 className="text-xl font-bold">{`${book ? "Update" : "Create"} Book`}</h3>
			</div>
			<Form onSubmit={onSubmit} error={error}>
				<FormControl label="Title" name="title" error={errors.title && "Title is required"}>
					<input
						className="input-bordered input"
						placeholder="title"
						defaultValue={book?.title}
						{...register("title", {
							required: true,
						})}
					/>
				</FormControl>
				<FormControl label="Author" name="author" error={errors.author && "Author is required"}>
					<input
						className="input-bordered input"
						placeholder="author"
						defaultValue={book?.author}
						{...register("author", {
							required: true,
						})}
					/>
				</FormControl>
				<FormControl label="Rating" name="rating" error={errors.rating && "Must be a number between 1 and 5"}>
					<input
						className="input-bordered input"
						placeholder="rating"
						type="number"
						defaultValue={book?.rating}
						{...register("rating", {
							valueAsNumber: true,
							min: 1,
							max: 5,
						})}
					/>
				</FormControl>
				<FormControl label="Purchased Date" name="purchasedAt" error={errors.purchasedAt && "Must be a date"}>
					<input
						className="input-bordered input"
						placeholder="purchasedAt"
						type="date"
						defaultValue={book?.purchasedAt}
						{...register("purchasedAt", {
							valueAsDate: true,
						})}
					/>
				</FormControl>
				<FormControl label="Completed Date" name="finishedAt" error={errors.finishedAt && "Must be a date"}>
					<input
						className="input-bordered input"
						placeholder="finishedAt"
						type="date"
						defaultValue={book?.finishedAt}
						{...register("finishedAt", {
							valueAsDate: true,
						})}
					/>
				</FormControl>
				<FormButtons>
					<button formAction="none" onClick={onCancel} className="btn-outline btn">
						{"Cancel"}
					</button>
					<button type="submit" className="btn-primary btn">
						{book ? "Update" : "Create"}
					</button>
				</FormButtons>
			</Form>
		</div>
	);
};

const Root = () => {
	const [splitPanel, setSplitPanel] = useState<ReactNode | undefined>();
	const { books, claims } = useLoaderData<typeof loader>();
	const [selected, setSelected] = useState<string | undefined>();
	const { revalidate } = useRevalidator();
	const actionsDropdown = useRef<HTMLDetailsElement>();

	const handleChange: ReactEventHandler<HTMLInputElement> = ({ currentTarget }) => {
		const id = currentTarget.value;

		setSelected(selected && id === selected ? undefined : id);
	};

	const handleCreateClick = () => {
		setSplitPanel(
			<UpsertBookSplitPanel
				onCancel={() => {
					setSplitPanel(undefined);
				}}
				onSuccess={() => {
					revalidate();
					setSplitPanel(undefined);
				}}
			/>,
		);
	};

	const handleUpdateClick = () => {
		actionsDropdown.current?.removeAttribute("open");

		if (!selected) {
			return;
		}

		const book = books.find((book) => book.id === selected);

		if (!book) {
			return;
		}

		setSplitPanel(
			<UpsertBookSplitPanel
				book={book}
				onCancel={() => {
					setSplitPanel(undefined);
				}}
				onSuccess={() => {
					revalidate();
					setSelected(undefined);
					setSplitPanel(undefined);
				}}
			/>,
		);
	};

	const handleDeleteClick = async () => {
		actionsDropdown.current?.removeAttribute("open");

		if (!selected) {
			return;
		}


		try {
			await booksApi.v1BooksIdDelete({ id: selected });
			revalidate();
			setSelected(undefined);
		} catch {}
	};

	return (
		<BaseLayout
			claims={claims}
			splitPanel={splitPanel}
			onSplitPanelClose={() => {
				setSplitPanel(undefined);
			}}
		>
			<div className="rounded-box overflow-x-auto bg-base-100 px-3 py-4">
				<div className="flex w-full justify-between px-4 pb-4 pt-2">
					<h2 className="text-3xl font-bold">My Books</h2>
					<div className="flex space-x-2">
						<details
							className="dropdown"
							onClick={(event) => {
								if (!selected) {
									event.preventDefault();
								}
							}}
							ref={(el) => (actionsDropdown.current = el)}
						>
							<summary className={cn("btn m-1", !selected && " btn-disabled")}>{"Actions"}</summary>
							<ul className="dropdown-content menu rounded-box z-10 w-52 border border-base-200 bg-base-100 p-2 shadow">
								<li>
									<button onClick={handleUpdateClick}>{"Edit"}</button>
								</li>
								<li>
									<button onClick={handleDeleteClick}>{"Delete"}</button>
								</li>
							</ul>
						</details>
						<button onClick={handleCreateClick} className="btn-primary btn">
							{"Create"}
						</button>
					</div>
				</div>
				<table className="table">
					<thead>
						<tr>
							<th />
							<th>Title</th>
							<th>Author</th>
							<th>Rating</th>
						</tr>
					</thead>
					<tbody>
						{books.map((book) => (
							<tr key={book.id} className="hover cursor-pointer">
								<td>
									<input
										className="radio"
										type="radio"
										value={book.id}
										onChange={handleChange}
										checked={selected === book.id}
									/>
								</td>
								<td>{book.title}</td>
								<td>{book.author}</td>
								<td>{book.rating}/5</td>
							</tr>
						))}
					</tbody>
				</table>
			</div>
		</BaseLayout>
	);
};

Root.loader = loader;

Root.path = "/";

export { Root };
