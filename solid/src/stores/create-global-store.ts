import { AuthApi, AuthClaims, BooksApi, Configuration } from "@pms/api";

import { API_URL } from "../lib/constants";
import { createPersistentStore } from "./create-persistent-store";

export type Tokens = {
	authorization: string;
};

export type GlobalStore = {
	claims: AuthClaims | undefined;
	tokens: Tokens | undefined;
};

const GLOBAL_STORE_KEY = "store";

export const createGlobalStore = () => {
	const [state, setState] = createPersistentStore<GlobalStore>(GLOBAL_STORE_KEY, {
		claims: undefined,
		tokens: undefined,
	});

	const configuration = new Configuration({
		basePath: API_URL,
		middleware: [
			{
				pre: async (context) => {
					if (state.tokens) {
						const headers = new Headers(context.init.headers);

						headers.set("Authorization", `Bearer ${state.tokens.authorization}`);

						return {
							...context,
							init: {
								...context.init,
								headers,
							},
						};
					}

					return context;
				},
			},
		],
	});

	const booksApi = new BooksApi(configuration);
	const authApi = new AuthApi(configuration);

	const setTokens = (tokens?: Tokens) => {
		setState("tokens", tokens);
	};
	const setClaims = (claims?: AuthClaims) => {
		setState("claims", claims);
	};

	return {
		authApi,
		booksApi,
		setClaims,
		setTokens,
		state,
	};
};
