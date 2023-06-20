import { clsx } from "clsx";
import { HTMLAttributes, forwardRef } from "react";

export const Container = forwardRef<HTMLDivElement, HTMLAttributes<HTMLDivElement>>(({ className, ...props }, ref) => (
	<div ref={ref} className={clsx("rounded-box overflow-x-auto bg-base-100 p-6", className)} {...props} />
));

Container.displayName = "Container";
