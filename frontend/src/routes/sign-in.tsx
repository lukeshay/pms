import { BaseLayout } from "../layouts/base-layout";
import { Container } from "../components/container";
import { FormControl } from "../components/form-control";
import { Form } from "../components/form";
import { SubmitHandler, useForm } from "react-hook-form";
import { FormButtons } from "../components/form-buttons";
import { authApi } from "../lib/apis";
import { useTokens } from "../hooks/use-tokens";
import { useNavigate } from "react-router-dom";

type Inputs = {
	email: string;
	password: string;
};

const SignIn = () => {
	const { setTokens } = useTokens();
	const navigate = useNavigate();
	const {
		register,
		handleSubmit,
		formState: { errors },
	} = useForm<Inputs>();

	const onSubmit: SubmitHandler<Inputs> = async (data) => {
		const { tokens } = await authApi.v1AuthSignInPost({ user: data });

		setTokens(tokens);
		navigate("/");
	};

	return (
		<BaseLayout>
			<Container>
				<h2 className="mb-8 text-3xl font-bold">{"Sign In"}</h2>
				<Form onSubmit={handleSubmit(onSubmit)}>
					<FormControl label="Email" name="email" error={errors.email && "Email is required"}>
						<input
							type="email"
							placeholder="email"
							className="input-bordered input"
							{...register("email", {
								required: true,
							})}
						/>
					</FormControl>
					<FormControl label="Password" name="password" error={errors.password && "Password is required"}>
						<input
							type="password"
							placeholder="password"
							className="input-bordered input"
							{...register("password", {
								required: true,
							})}
						/>
					</FormControl>
					<FormButtons>
						<button type="submit" className="btn-primary btn">
							{"Submit"}
						</button>
					</FormButtons>
				</Form>
			</Container>
		</BaseLayout>
	);
};

SignIn.path = "/sign-in";

export { SignIn };
