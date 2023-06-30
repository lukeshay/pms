import { Input } from "../ui/input";
import { ParentComponent } from "solid-js";
import { Button } from "../ui/button";
import { Form } from "../ui/form";
import { FormControl } from "../ui/form-control";
import { ProfileSection } from "./profile-section";

export type PasswordFormSectionProperties = {};

export const PasswordFormSection: ParentComponent<PasswordFormSectionProperties> = () => {
	return (
		<ProfileSection header="Password">
			<form>
				<Form
					actions={
						<>
							<Button formAction="reset">{"Cancel"}</Button>
							<Button formAction="submit" variant="primary">
								{"Save"}
							</Button>
						</>
					}
				>
					<FormControl label="Current Password" name="current-password">
						<Input bordered name="current-password" autocomplete="current-password" type="password" />
					</FormControl>
					<FormControl label="New Password" name="new-password">
						<Input bordered type="password" name="new-password" autocomplete="new-password" />
					</FormControl>
					<FormControl label="Repeat Password" name="repeat-new-password">
						<Input bordered type="password" name="repeat-new-password" autocomplete="repeat-new-password" />
					</FormControl>
				</Form>
			</form>
		</ProfileSection>
	);
};
