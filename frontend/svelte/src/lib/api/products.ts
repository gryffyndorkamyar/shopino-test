export type Product = {
	id: number;
	name: string;
	description?: string;
	price: number;
	original_price?: number;
	discount_percent?: number;
	image_url?: string;
	stock?: number;
	is_active?: boolean;
	store_id?: number;
	category_id?: number;
	created_at?: string;
	updated_at?: string;
};

export const FALLBACK_PRODUCT_IMAGE =
	'https://images.unsplash.com/photo-1483985988355-763728e1935b?auto=format&fit=crop&w=900&q=80';

async function readError(res: Response): Promise<string> {
	try {
		const data = await res.json();
		if (typeof data?.detail === 'string') return data.detail;
		if (typeof data?.message === 'string') return data.message;
	} catch {
		/* ignore */
	}
	return `درخواست ناموفق بود (${res.status})`;
}

export async function fetchProducts(): Promise<Product[]> {
	const res = await fetch('/api/products/');
	if (!res.ok) throw new Error(await readError(res));
	const data = await res.json();
	return Array.isArray(data) ? data : [];
}

export async function fetchProduct(id: number | string): Promise<Product> {
	const res = await fetch(`/api/products/${id}`);
	if (!res.ok) throw new Error(await readError(res));
	return res.json();
}

export function productImage(p: Product): string {
	return p.image_url && p.image_url.trim() ? p.image_url : FALLBACK_PRODUCT_IMAGE;
}

export function hasDiscount(p: Product): boolean {
	return (p.discount_percent ?? 0) > 0 && (p.original_price ?? 0) > p.price;
}

export function formatToman(value: number): string {
	return new Intl.NumberFormat('fa-IR').format(value) + ' تومان';
}
