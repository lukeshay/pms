import { useState } from "react";
import { BooksApi, Configuration } from "../api";
import { useLoaderData } from "../hooks/use-loader-data";
import { BaseLayout } from "../layouts/base-layout";
import { withUserData } from "../helpers/with-user-data";

const loader = withUserData(async () => {
	const booksApi = new BooksApi(
		new Configuration({
			basePath: [import.meta.env.VITE_API_URL, "api"].join("/"),
		}),
	);

	return booksApi.v1BooksGet();
});

const Root = () => {
	const { books } = useLoaderData<typeof loader>();
	const [selected, setSelected] = useState<string | undefined>();

	const handleSelect = (id: string) => {
		setSelected(selected && id === selected ? undefined : id);
	};

	return (
		<BaseLayout>
			<div className="rounded-box overflow-x-auto bg-base-100 px-3 py-4">
				<div className="flex w-full justify-between px-4 pb-4 pt-2">
					<h2 className="text-3xl font-bold">My Books</h2>
					<div>
						<a href="/books/create" className="btn-primary btn">
							Add Book
						</a>
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
									<input className="radio" onChange={() => handleSelect(book.id)} checked={selected === book.id} />
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

export { Root };
