import type { PageLoad } from './$types';
import { error } from '@sveltejs/kit';
import { fetchProduct } from '$lib/api/products';

export const ssr = false;

export const load: PageLoad = async ({ params }) => {
	try {
		const product = await fetchProduct(params.id);
		return { product };
	} catch (e) {
		const message = e instanceof Error ? e.message : 'محصول پیدا نشد';
		const lower = message.toLowerCase();
		if (lower.includes('not found') || lower.includes('404') || message.includes('پیدا نشد')) {
			throw error(404, 'محصول پیدا نشد');
		}
		throw error(500, message);
	}
};
