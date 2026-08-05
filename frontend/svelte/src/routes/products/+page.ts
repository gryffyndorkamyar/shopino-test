import type { PageLoad } from './$types';
import { fetchProducts } from '$lib/api/products';

export const ssr = false;

export const load: PageLoad = async () => {
	try {
		const products = await fetchProducts();
		return { products, error: null as string | null };
	} catch (e) {
		return {
			products: [],
			error: e instanceof Error ? e.message : 'خطا در دریافت محصولات'
		};
	}
};
