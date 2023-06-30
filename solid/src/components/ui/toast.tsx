import toaster from "solid-toast";

import { Alert, AlertProperties as AlertProperties } from "./alert";

const toastToast = (variant: AlertProperties["variant"], properties: Omit<AlertProperties, "variant"> | string) => {
	const alertProperties = typeof properties === "string" ? { description: properties } : properties;

	toaster.custom(() => <Alert variant={variant} {...alertProperties} />);
};

const toast = {
	error: (properties: Omit<AlertProperties, "variant"> | string) => toastToast("error", properties),
	info: (properties: Omit<AlertProperties, "variant"> | string) => toastToast("info", properties),
	neutral: (properties: Omit<AlertProperties, "variant"> | string) => toastToast("neutral", properties),
	success: (properties: Omit<AlertProperties, "variant"> | string) => toastToast("success", properties),
	warning: (properties: Omit<AlertProperties, "variant"> | string) => toastToast("warning", properties),
};

export const useToast = () => toast;
