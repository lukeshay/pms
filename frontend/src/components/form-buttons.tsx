import { clsx } from "clsx";
import { HTMLAttributes, forwardRef } from "react";

export const FormButtons = forwardRef<HTMLDivElement, HTMLAttributes<HTMLDivElement>>(
	({ className, ...props }, ref) => (
		<div ref={ref} className={clsx("float-right flex space-x-2 pt-4", className)} {...props} />
	),
);

FormButtons.displayName = "FormButtons";
