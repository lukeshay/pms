import { createForm } from "@felte/solid";
import { validator } from "@felte/validator-zod";
import { useRouteData } from "@solidjs/router";
import { ParentComponent } from "solid-js";

import { useGlobalStore } from "../../provider";
import profileLoader from "../../routes/profile.loader";
import { UserUpdateSchema, userUpdateSchema } from "../../schemas/user-schema";
import { Form } from "../ui/form";
import { FormControl } from "../ui/form-control";
import { Input } from "../ui/input";
import { useToast } from "../ui/toast";
import { ProfileSection } from "./profile-section";

export type ProfileFormSectionProperties = never;

export const ProfileFormSection: ParentComponent<ProfileFormSectionProperties> = () => {
	const { usersApi } = useGlobalStore();
	const { data, refetch } = useRouteData<typeof profileLoader>();
	const toast = useToast();

	const { errors, form } = createForm<UserUpdateSchema>({
		extend: validator({ schema: userUpdateSchema }),
		initialValues: {
			email: data()?.claims.email,
			firstName: data()?.claims.firstName,
			lastName: data()?.claims.lastName,
		},
		onSubmit: async (values) => {
			try {
				await usersApi.v1UsersIdPut({ id: data()!.claims.id, user: values });
				await refetch();

				toast.success("Profile updated successfully");
			} catch (error) {
				toast.error(`Something went wrong: ${(error as Error).message}`);
			}
		},
	});

	return (
		<ProfileSection description={`Welcome back, ${data()!.claims.firstName}!`} header="Profile">
			<form use:form>
				<Form>
					<FormControl error={errors("email")} label="Email" name="email">
						<Input bordered name="email" type="text" value={data()!.claims.email} />
					</FormControl>
					<FormControl error={errors("firstName")} label="First Name" name="firstName">
						<Input autocomplete="given-name" bordered name="firstName" type="text" value={data()!.claims.firstName} />
					</FormControl>
					<FormControl error={errors("lastName")} label="Last Name" name="lastName">
						<Input autocomplete="family-name" bordered name="lastName" type="text" value={data()!.claims.lastName} />
					</FormControl>
				</Form>
			</form>
		</ProfileSection>
	);
};
