export const withUserData = async <T>(loader: (...args: any) => T | Promise<T>): Promise<T> => {
	const result = await loader();

	return result;
};
