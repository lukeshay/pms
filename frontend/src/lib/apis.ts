import { AuthApi, BooksApi, Configuration } from "../api";
import { useTokens } from "../hooks/use-tokens";
import { API_URL } from "./constants";

const configuration = new Configuration({
	basePath: API_URL,
	middleware: [
		{
			pre: async (context) => {
				context.init.headers = {
					...context.init.headers,
					Authorization: `Bearer ${useTokens.getState().tokens?.authorization}`,
				};

				return context;
			},
		},
	],
});

export const booksApi = new BooksApi(configuration);

export const authApi = new AuthApi(configuration);
