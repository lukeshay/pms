import { AuthClaims } from "@pms/api";
import { useNavigate } from "@solidjs/router";
import { createResource } from "solid-js";

import { useGlobalStore } from "../provider";

const profileLoader = () => {
	const { authApi, setClaims } = useGlobalStore();
	const navigate = useNavigate();

	const [data, { refetch }] = createResource(async () => {
		try {
			const { claims } = await authApi.v1AuthGet();

			setClaims(claims);

			return { claims };
		} catch {
			navigate("/sign-in");
		}

		return {
			claims: {} as AuthClaims,
		};
	});

	return {
		data,
		refetch,
	};
};

export default profileLoader;
