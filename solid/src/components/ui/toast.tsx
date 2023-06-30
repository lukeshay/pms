import toaster from "solid-toast";

import { Alert, AlertProperties as AlertProperties } from "./alert";

const toast = {
	error: (properties: Omit<AlertProperties, "variant">) =>
		toaster.custom(() => <Alert variant="error" {...properties} />),
	info: (properties: Omit<AlertProperties, "variant">) =>
		toaster.custom(() => <Alert variant="info" {...properties} />),
	neutral: (properties: Omit<AlertProperties, "variant">) =>
		toaster.custom(() => <Alert variant="neutral" {...properties} />),
	success: (properties: Omit<AlertProperties, "variant">) =>
		toaster.custom(() => <Alert variant="success" {...properties} />),
	warning: (properties: Omit<AlertProperties, "variant">) =>
		toaster.custom(() => <Alert variant="warning" {...properties} />),
};

export const useToast = () => toast;
