<script lang="ts">
	import {
		formatToman,
		hasDiscount,
		productImage
	} from '$lib/api/products';

	let { data } = $props();
	const product = $derived(data.product);
</script>

<svelte:head>
	<title>{product.name} | شاپینو</title>
</svelte:head>

<section class="detail">
	<div class="container detail__grid">
		<div class="media">
			<img src={productImage(product)} alt={product.name} />
			{#if hasDiscount(product)}
				<span class="badge">{product.discount_percent}٪ تخفیف</span>
			{/if}
		</div>

		<div class="info">
			<a class="back" href="/products">بازگشت به محصولات</a>
			<h1>{product.name}</h1>
			{#if product.description}
				<p class="desc">{product.description}</p>
			{/if}

			<div class="price">
				<strong>{formatToman(product.price)}</strong>
				{#if hasDiscount(product)}
					<span>{formatToman(product.original_price ?? 0)}</span>
				{/if}
			</div>

			<ul class="meta">
				<li><span>موجودی</span><strong>{product.stock ?? 0}</strong></li>
				<li><span>فروشگاه</span><strong>#{product.store_id ?? '-'}</strong></li>
				<li><span>دسته</span><strong>#{product.category_id ?? '-'}</strong></li>
			</ul>

			<button type="button" class="buy" disabled={(product.stock ?? 0) <= 0}>
				{(product.stock ?? 0) > 0 ? 'افزودن به سبد' : 'ناموجود'}
			</button>
		</div>
	</div>
</section>

<style>
	.detail {
		padding-block: clamp(1.5rem, 4vw, 3rem) 2rem;
	}

	.detail__grid {
		display: grid;
		gap: 1.5rem;
	}

	.media {
		position: relative;
		border-radius: calc(var(--radius) + 4px);
		overflow: hidden;
		background: var(--paper-deep);
		aspect-ratio: 4 / 5;
	}

	.media img {
		width: 100%;
		height: 100%;
		object-fit: cover;
	}

	.badge {
		position: absolute;
		top: 1rem;
		left: 1rem;
		padding: 0.4rem 0.7rem;
		border-radius: 10px;
		background: var(--accent);
		color: #fff;
		font-weight: 700;
		font-size: 0.88rem;
	}

	.info {
		display: grid;
		align-content: start;
		gap: 0.85rem;
	}

	.back {
		color: var(--ink-soft);
		font-size: 0.9rem;
		width: fit-content;
	}

	.back:hover {
		color: var(--ink);
	}

	.info h1 {
		margin: 0;
		font-family: var(--font-display);
		font-size: clamp(1.7rem, 3vw, 2.4rem);
		letter-spacing: -0.03em;
		line-height: 1.25;
	}

	.desc {
		margin: 0;
		color: var(--ink-soft);
		line-height: 1.85;
		max-width: 42ch;
	}

	.price {
		display: flex;
		align-items: baseline;
		gap: 0.7rem;
	}

	.price strong {
		font-size: 1.35rem;
	}

	.price span {
		color: var(--ink-soft);
		text-decoration: line-through;
	}

	.meta {
		list-style: none;
		margin: 0.4rem 0 0;
		padding: 0;
		display: grid;
		gap: 0.55rem;
	}

	.meta li {
		display: flex;
		justify-content: space-between;
		gap: 1rem;
		padding: 0.7rem 0;
		border-bottom: 1px solid var(--line);
		color: var(--ink-soft);
	}

	.meta strong {
		color: var(--ink);
	}

	.buy {
		margin-top: 0.6rem;
		border: 0;
		border-radius: 12px;
		padding: 0.95rem 1.2rem;
		background: var(--ink);
		color: #fff;
		font: inherit;
		font-weight: 700;
		width: min(100%, 280px);
		transition: background 0.2s ease, transform 0.2s ease;
	}

	.buy:hover:not(:disabled) {
		background: var(--accent);
		transform: translateY(-1px);
	}

	.buy:disabled {
		opacity: 0.55;
		cursor: not-allowed;
	}

	@media (min-width: 900px) {
		.detail__grid {
			grid-template-columns: 1.05fr 0.95fr;
			gap: 2.5rem;
			align-items: start;
		}

		.media {
			aspect-ratio: 5 / 6;
		}
	}
</style>
