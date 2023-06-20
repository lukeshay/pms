import { FC, ReactNode } from "react";

export type FormControlProps = {
	children: ReactNode;
	label: ReactNode;
	name: string;
	error?: string;
};

export const FormControl: FC<FormControlProps> = ({ children, name, label, error }) => (
	<div className="form-control">
		<label className="label" htmlFor={name}>
			<span className="label-text">{label}</span>
		</label>
		{children}
		{error && (
			<label className="label" htmlFor={name}>
				<span className="label-text-alt text-error">{error}</span>
			</label>
		)}
	</div>
);
