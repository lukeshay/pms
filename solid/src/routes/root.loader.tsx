import { ModelsBook } from "@pms/api";
import { useNavigate } from "@solidjs/router";
import { createResource } from "solid-js";

import { useGlobalStore } from "../provider";

const rootLoader = () => {
	const { booksApi } = useGlobalStore();
	const navigate = useNavigate();

	const [data, { refetch }] = createResource(async () => {
		try {
			const { books } = await booksApi.v1BooksGet();

			return { books };
		} catch {
			navigate("/sign-in");
		}

		return {
			books: [] as ModelsBook[],
		};
	});

	return {
		data,
		refetch,
	};
};

export default rootLoader;
