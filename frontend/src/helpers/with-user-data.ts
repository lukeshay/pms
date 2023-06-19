export const withUserData =
	<T>(loader: (...args: any) => T | Promise<T>): ((...args: any) => Promise<T>) =>
	async (...args: any) => {
		const result = await loader(...args);

		return result;
	};
