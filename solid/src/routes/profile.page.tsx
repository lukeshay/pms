import { type Component } from "solid-js";

import { PasswordFormSection } from "../components/profile/password-form-section";
import { ProfileFormSection } from "../components/profile/profile-form-section";
import { BaseLayout } from "../layouts/base-layout";

const ProfilePage: Component = () => (
	<BaseLayout>
		<div class="space-y-12">
			<ProfileFormSection />
			<PasswordFormSection />
		</div>
	</BaseLayout>
);

export default ProfilePage;
