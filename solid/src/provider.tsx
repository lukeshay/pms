import { Router } from "@solidjs/router";
import { Context, ParentComponent, createContext, useContext } from "solid-js";

import { createGlobalStore } from "./stores/create-global-store";

const GlobalStoreContext = createContext<ReturnType<typeof createGlobalStore>>() as Context<
	ReturnType<typeof createGlobalStore>
>;

export const Provider: ParentComponent = (properties) => {
	const globalStore = createGlobalStore();

	return (
		<Router>
			<GlobalStoreContext.Provider value={globalStore}>{properties.children}</GlobalStoreContext.Provider>
		</Router>
	);
};

export const useGlobalStore = () => useContext(GlobalStoreContext);
