import { type Component } from "solid-js";

import { BaseLayout } from "../layouts/base-layout";
import { ProfileFormSection } from "../components/profile/profile-form-section";
import { PasswordFormSection } from "../components/profile/password-form-section";

const ProfilePage: Component = () => (
	<BaseLayout>
		<div class="space-y-12">
			<ProfileFormSection />
			<PasswordFormSection />
		</div>
	</BaseLayout>
);

export default ProfilePage;
