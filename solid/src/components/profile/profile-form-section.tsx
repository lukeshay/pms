import { Input } from "../ui/input";
import { ParentComponent } from "solid-js";
import { Button } from "../ui/button";
import { Form } from "../ui/form";
import { FormControl } from "../ui/form-control";
import profileLoader from "../../routes/profile.loader";
import { useRouteData } from "@solidjs/router";
import { ProfileSection } from "./profile-section";

export type ProfileFormSectionProperties = {};

export const ProfileFormSection: ParentComponent<ProfileFormSectionProperties> = () => {
	const { data } = useRouteData<typeof profileLoader>();

	return (
		<ProfileSection description={`Welcome back, ${data()!.claims.firstName}!`} header="Profile">
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
					<FormControl label="Email" name="email">
						<Input bordered type="text" name="email" value={data()!.claims.email} />
					</FormControl>
					<FormControl label="First Name" name="give-name">
						<Input bordered type="text" name="give-name" value={data()!.claims.firstName} />
					</FormControl>
					<FormControl label="Last Name" name="family-name">
						<Input bordered type="text" name="family-name" value={data()!.claims.lastName} />
					</FormControl>
				</Form>
			</form>
		</ProfileSection>
	);
};
