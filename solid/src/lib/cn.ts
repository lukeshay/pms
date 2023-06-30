import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";

import { ClassList } from "../types/solid-types";

export const cn = (...inputs: ClassValue[]) => twMerge(clsx(inputs));

export const classListToClassValues = (classList?: ClassList) =>
	Object.entries(classList ?? {}).map(([key, value]) => value && key);
