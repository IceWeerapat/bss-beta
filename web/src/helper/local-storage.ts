export const setObject = (key: string, value: unknown) => {
	localStorage.setItem(key, JSON.stringify(value));
};

export const getObject = <T>(key: string): T => {
	const value = localStorage.getItem(key);
	return value ? JSON.parse(value) : null;
};
