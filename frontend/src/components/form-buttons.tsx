import { HTMLAttributes, forwardRef } from "react";
import { cn } from "../lib/cn";

export const FormButtons = forwardRef<HTMLDivElement, HTMLAttributes<HTMLDivElement>>(
	({ className, ...props }, ref) => (
		<div ref={ref} className={cn("float-right flex space-x-2 pt-4", className)} {...props} />
	),
);

FormButtons.displayName = "FormButtons";
