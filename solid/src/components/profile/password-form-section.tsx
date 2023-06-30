import { createForm } from "@felte/solid";
import { validator } from "@felte/validator-zod";
import { useRouteData } from "@solidjs/router";
import { ParentComponent } from "solid-js";

import { useGlobalStore } from "../../provider";
import profileLoader from "../../routes/profile.loader";
import { UserPasswordUpdateSchema, userPasswordUpdateSchema } from "../../schemas/user-schema";
import { Form } from "../ui/form";
import { FormControl } from "../ui/form-control";
import { Input } from "../ui/input";
import { useToast } from "../ui/toast";
import { ProfileSection } from "./profile-section";

export type PasswordFormSectionProperties = never;

export const PasswordFormSection: ParentComponent<PasswordFormSectionProperties> = () => {
	const { usersApi } = useGlobalStore();
	const { data, refetch } = useRouteData<typeof profileLoader>();
	const toast = useToast();

	const { errors, form } = createForm<UserPasswordUpdateSchema>({
		extend: validator({ schema: userPasswordUpdateSchema }),
		initialValues: {},
		onSubmit: async (values) => {
			try {
				console.log({ values });
				await refetch();

				toast.success("Profile updated successfully");
			} catch (error) {
				toast.error(`Something went wrong: ${(error as Error).message}`);
			}
		},
	});

	return (
		<ProfileSection header="Password">
			<form use:form>
				<Form>
					<input formNoValidate name="username" type="hidden" value={data()?.claims.aud} />
					<FormControl error={errors("currentPassword")} label="Current Password" name="currentPassword">
						<Input autocomplete="current-password" bordered name="currentPassword" type="password" />
					</FormControl>
					<FormControl error={errors("newPassword")} label="New Password" name="newPassword">
						<Input autocomplete="new-password" bordered name="newPassword" type="password" />
					</FormControl>
					<FormControl error={errors("repeatNewPassword")} label="Repeat Password" name="repeatNewPassword">
						<Input autocomplete="repeat-new-password" bordered name="repeatNewPassword" type="password" />
					</FormControl>
				</Form>
			</form>
		</ProfileSection>
	);
};
