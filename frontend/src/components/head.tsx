import { FC, ReactNode } from "react";
import { createPortal } from "react-dom";

const headRoot = document.head;

export type HeadProps = {
	children?: ReactNode;
};

export const Head: FC<HeadProps> = ({ children }) => createPortal(children, headRoot);
