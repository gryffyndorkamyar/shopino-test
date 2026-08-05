<script lang="ts">
	import {
		formatToman,
		hasDiscount,
		productImage,
		type Product
	} from '$lib/api/products';

	let { product, index = 0 }: { product: Product; index?: number } = $props();
</script>

<a
	class="product"
	href="/products/{product.id}"
	style="animation-delay: {Math.min(index, 8) * 55}ms"
>
	<div class="product__media">
		<img src={productImage(product)} alt={product.name} loading="lazy" />
		{#if hasDiscount(product)}
			<span class="product__discount">{product.discount_percent}٪</span>
		{/if}
	</div>
	<div class="product__body">
		<h3>{product.name}</h3>
		{#if product.description}
			<p class="product__desc">{product.description}</p>
		{/if}
		<div class="product__price">
			<strong>{formatToman(product.price)}</strong>
			{#if hasDiscount(product)}
				<span>{formatToman(product.original_price ?? 0)}</span>
			{/if}
		</div>
	</div>
</a>

<style>
	.product {
		display: block;
		background: var(--surface);
		border: 1px solid var(--line);
		border-radius: var(--radius);
		overflow: hidden;
		animation: rise 0.65s ease both;
		transition: transform 0.25s ease, box-shadow 0.25s ease;
	}

	.product:hover {
		transform: translateY(-4px);
		box-shadow: var(--shadow-soft);
	}

	.product__media {
		position: relative;
		aspect-ratio: 4 / 5;
		overflow: hidden;
		background: var(--paper-deep);
	}

	.product__media img {
		width: 100%;
		height: 100%;
		object-fit: cover;
		transition: transform 0.45s ease;
	}

	.product:hover .product__media img {
		transform: scale(1.04);
	}

	.product__discount {
		position: absolute;
		top: 0.7rem;
		left: 0.7rem;
		padding: 0.25rem 0.5rem;
		border-radius: 8px;
		background: var(--accent);
		color: #fff;
		font-size: 0.78rem;
		font-weight: 700;
	}

	.product__body {
		padding: 0.85rem 0.9rem 1rem;
		display: grid;
		gap: 0.35rem;
	}

	.product__body h3 {
		margin: 0;
		font-size: 0.98rem;
		line-height: 1.45;
	}

	.product__desc {
		margin: 0;
		color: var(--ink-soft);
		font-size: 0.82rem;
		line-height: 1.55;
		display: -webkit-box;
		-webkit-line-clamp: 2;
		line-clamp: 2;
		-webkit-box-orient: vertical;
		overflow: hidden;
	}

	.product__price {
		display: flex;
		align-items: baseline;
		gap: 0.55rem;
		margin-top: 0.25rem;
	}

	.product__price strong {
		font-size: 0.95rem;
	}

	.product__price span {
		color: var(--ink-soft);
		font-size: 0.8rem;
		text-decoration: line-through;
	}
</style>
