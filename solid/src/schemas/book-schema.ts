import { z } from "zod";

export const bookSchema = z.object({
	author: z.string().nonempty(),
	createdAt: z.string().nonempty(),
	createdBy: z.string().nonempty(),
	deletedAt: z.string().optional(),
	deletedBy: z.string().optional(),
	finishedAt: z.string().optional(),
	id: z.string().uuid().ulid(),
	purchasedAt: z.string().optional(),
	rating: z.number().min(1).max(5).optional(),
	title: z.string().nonempty(),
	updatedAt: z.string().nonempty(),
	updatedBy: z.string().nonempty(),
	userId: z.string().uuid().ulid().nonempty(),
});

export const bookUpsertSchema = bookSchema.omit({
	createdAt: true,
	createdBy: true,
	deletedAt: true,
	deletedBy: true,
	id: true,
	updatedAt: true,
	updatedBy: true,
	userId: true,
});

export type BookUpsertSchema = z.infer<typeof bookUpsertSchema>;
