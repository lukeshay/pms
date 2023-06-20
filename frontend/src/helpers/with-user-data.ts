import { redirect } from "react-router-dom";
import { AuthClaims } from "../api";
import { authApi } from "../lib/apis";

export const withUserData =
	<T>(loader: (...args: any) => T | Promise<T>): ((...args: any) => Promise<T & { claims: AuthClaims }>) =>
	async (...args: any) => {
		let claims: AuthClaims;

		try {
			const res = await authApi.v1AuthGet();
			claims = res.claims;
		} catch {
			return redirect("/sign-in") as unknown as T & { claims: AuthClaims };
		}

		const result = await loader(...args);

		return {
			...result,
			claims,
		};
	};
