import { create } from "zustand";
import { persist, createJSONStorage } from "zustand/middleware";
import { ControllersAuthTokens } from "../api";

type Tokens = {
	tokens: ControllersAuthTokens | undefined;
	setTokens: (tokens?: ControllersAuthTokens) => void;
};

export const useTokens = create(
	persist<Tokens>(
		(set) => ({
			tokens: undefined as ControllersAuthTokens | undefined,
			setTokens: (tokens: ControllersAuthTokens) => set({ tokens }),
		}),
		{
			name: "tokens", // unique name
			storage: createJSONStorage(() => sessionStorage),
		},
	),
);
