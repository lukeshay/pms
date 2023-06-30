import { z } from "zod";

export const userUpdateSchema = z.object({
	email: z.string().email(),
	firstName: z.string().nonempty(),
	lastName: z.string().nonempty(),
});

export type UserUpdateSchema = z.infer<typeof userUpdateSchema>;

export const userPasswordUpdateSchema = z
	.object({
		currentPassword: z.string().nonempty(),
		newPassword: z.string().nonempty(),
		repeatNewPassword: z.string().nonempty(),
	})
	.refine((data) => data.newPassword === data.repeatNewPassword, {
		message: "Passwords do not match",
	});

export type UserPasswordUpdateSchema = z.infer<typeof userPasswordUpdateSchema>;
