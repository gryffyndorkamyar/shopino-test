export type Store = {
	id: number;
	name: string;
	slug: string;
	description: string;
	is_active: boolean;
	owner_username: string;
	created_at: string;
};

async function readError(res: Response): Promise<string> {
	try {
		const data = await res.json();
		if (typeof data?.detail === 'string') return data.detail;
	} catch {
		/* ignore */
	}
	return `خطا در دریافت فروشگاه‌ها (${res.status})`;
}

export async function fetchStores(): Promise<Store[]> {
	const res = await fetch('/api/stores/');
	if (!res.ok) throw new Error(await readError(res));
	const data = await res.json();
	if (Array.isArray(data)) return data;
	if (Array.isArray(data?.results)) return data.results;
	return [];
}
