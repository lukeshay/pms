/* @refresh reload */
import { Route, Routes } from "@solidjs/router";
import { lazy } from "solid-js";
import { render } from "solid-js/web";
import { Toaster } from "solid-toast";

import "./index.css";
import { Provider } from "./provider";
import rootLoader from "./routes/root.loader";
import profileLoader from "./routes/profile.loader";

const root = document.querySelector("#root");

if (import.meta.env.DEV && !(root instanceof HTMLElement)) {
	throw new Error(
		"Root element not found. Did you forget to add it to your index.html? Or maybe the id attribute got misspelled?",
	);
}

render(
	() => (
		<Provider>
			<Routes>
				<Route component={lazy(async () => import("./routes/root.page"))} data={rootLoader} path="/" />
				<Route component={lazy(async () => import("./routes/profile.page"))} data={profileLoader} path="/profile" />
				<Route component={lazy(async () => import("./routes/sign-in.page"))} path="/sign-in" />
			</Routes>
			<Toaster position="bottom-center" />
		</Provider>
	),
	root!,
);
