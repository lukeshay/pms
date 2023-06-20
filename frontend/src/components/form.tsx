import clsx from "clsx";
import { forwardRef } from "react";
import { Form as RRForm, FormProps } from "react-router-dom";

export const Form = forwardRef<HTMLFormElement, FormProps & { error?: string }>(
	({ className, error, children, ...props }, ref) => (
		<RRForm ref={ref} className={clsx("w-full max-w-3xl space-y-4", className)} {...props}>
			{error && (
				<div className="alert alert-error">
					<svg
						xmlns="http://www.w3.org/2000/svg"
						className="h-6 w-6 shrink-0 stroke-current"
						fill="none"
						viewBox="0 0 24 24"
					>
						<path
							strokeLinecap="round"
							strokeLinejoin="round"
							strokeWidth="2"
							d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
						/>
					</svg>
					<span>{error}</span>
				</div>
			)}
			{children}
		</RRForm>
	),
);

Form.displayName = "Form";
