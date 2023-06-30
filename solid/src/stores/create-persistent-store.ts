import { createEffect } from "solid-js";
import { createStore } from "solid-js/store";

export const createPersistentStore = <T extends object>(key: string, value: T) => {
	// load stored todos on init
	const stored = localStorage.getItem(key),
		[state, setState] = createStore<T>(stored ? JSON.parse(stored) : value);

	// JSON.stringify creates deps on every iterable field
	createEffect(() => localStorage.setItem(key, JSON.stringify(state)));

	return [state, setState] as const;
};
