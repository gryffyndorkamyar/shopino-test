import type { PageLoad } from './$types';
import { fetchStores } from '$lib/api/stores';

export const ssr = false;

export const load: PageLoad = async () => {
	try {
		const stores = await fetchStores();
		return { stores, error: null as string | null };
	} catch (e) {
		return {
			stores: [],
			error: e instanceof Error ? e.message : 'خطا در دریافت فروشگاه‌ها'
		};
	}
};
