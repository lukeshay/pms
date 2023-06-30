import { createForm } from "@felte/solid";
import { useNavigate } from "@solidjs/router";
import { type Component, batch, createSignal } from "solid-js";

import { Container } from "../components/ui/container";
import { Form } from "../components/ui/form";
import { FormButtons } from "../components/ui/form-buttons";
import { FormControl } from "../components/ui/form-control";
import { BaseLayout } from "../layouts/base-layout";
import { useGlobalStore } from "../provider";

type Data = {
	email: string;
	password: string;
};

const SignInPage: Component = () => {
	const { authApi, setClaims, setTokens } = useGlobalStore();
	const [errorText, setErrorText] = createSignal<string>();
	const navigate = useNavigate();

	const { errors, form } = createForm<Data>({
		onSubmit: async (values) => {
			setErrorText();
			try {
				const { tokens } = await authApi.v1AuthSignInPost({ user: values });

				const headers = new Headers();
				headers.set("Authorization", `Bearer ${tokens.authorization}`);

				const { claims } = await authApi.v1AuthGet({ headers });

				batch(() => {
					setTokens(tokens);
					setClaims(claims);
				});

				navigate("/");
			} catch (error) {
				setErrorText("An unknown error ocurred: " + (error as Error).message);
			}
		},
		validate(values) {
			const errors: { email: string[]; password: string[] } = {
				email: [],
				password: [],
			};
			if (!values.email) errors.email.push("Must not be empty");
			if (!/[A-Za-z][^@]*@[A-Za-z][^.@]*\.[a-z]{2,}/.test(values.email)) errors.email.push("Must be a valid email");
			if (!values.password) errors.password.push("Must not be empty");
			return errors;
		},
	});

	return (
		<BaseLayout>
			<Container>
				<h2 class="mb-8 text-3xl font-bold">{"Sign In"}</h2>
				<form use:form>
					<Form error={errorText()}>
						<fieldset>
							<FormControl error={errors("email")} label="Email" name="email">
								<input
									autocomplete="email"
									class="input-bordered input"
									name="email"
									placeholder="email"
									type="email"
								/>
							</FormControl>
							<FormControl error={errors("password")} label="Password" name="password">
								<input
									autocomplete="current-password"
									class="input-bordered input"
									name="password"
									placeholder="password"
									type="password"
								/>
							</FormControl>
						</fieldset>
						<FormButtons>
							<button class="btn-primary btn" type="submit">
								{"Submit"}
							</button>
						</FormButtons>
					</Form>
				</form>
			</Container>
		</BaseLayout>
	);
};

export default SignInPage;
