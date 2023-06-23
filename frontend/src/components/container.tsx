import { HTMLAttributes, forwardRef } from "react";
import { cn } from "../lib/cn";

export const Container = forwardRef<HTMLDivElement, HTMLAttributes<HTMLDivElement>>(({ className, ...props }, ref) => (
	<div ref={ref} className={cn("rounded-box overflow-x-auto bg-base-100 p-6", className)} {...props} />
));

Container.displayName = "Container";
