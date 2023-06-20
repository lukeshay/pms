import React from "react";
import ReactDOM from "react-dom/client";
import { createBrowserRouter, LoaderFunction, RouterProvider } from "react-router-dom";

import "./index.css";

import { Root } from "./routes/root";
import { Head } from "./components/head";
import { API_URL } from "./lib/constants";
import { SignIn } from "./routes/sign-in";

const router = createBrowserRouter([
	{
		path: Root.path,
		loader: Root.loader as unknown as LoaderFunction,
		element: <Root />,
	},
	{
		path: SignIn.path,
		element: <SignIn />,
	},
]);

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
	<React.StrictMode>
		<Head>
			{/* <link as="fetch" crossOrigin="anonymous" href={`${API_URL}/auth/`} rel="preload" /> */}
			<title>{"PMS"}</title>
		</Head>
		<RouterProvider router={router} />
	</React.StrictMode>,
);
