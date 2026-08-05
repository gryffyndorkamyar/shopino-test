<script lang="ts">
	import ProductCard from '$lib/components/ProductCard.svelte';
	import StateBox from '$lib/components/StateBox.svelte';

	let { data } = $props();
</script>

<svelte:head>
	<title>محصولات | شاپینو</title>
</svelte:head>

<section class="page">
	<div class="container">
		<header class="page__head">
			<h1>همه محصولات</h1>
			<p>لیست زنده از سرویس محصول گولنگ</p>
		</header>

		{#if data.error}
			<StateBox
				tone="error"
				title="خطا در دریافت محصولات"
				message={data.error}
				hint="مطمئن شو سرویس golang و gateway روی ۸۰۸۸ بالا هستند."
			/>
		{:else if data.products.length === 0}
			<StateBox title="محصولی پیدا نشد" message="هنوز هیچ محصولی در دیتابیس نیست." />
		{:else}
			<div class="product-grid">
				{#each data.products as product, i (product.id)}
					<ProductCard {product} index={i} />
				{/each}
			</div>
		{/if}
	</div>
</section>

<style>
	.page {
		padding-block: clamp(2rem, 5vw, 3.5rem) 1rem;
	}

	.page__head {
		margin-bottom: 1.6rem;
	}

	.page__head h1 {
		margin: 0;
		font-family: var(--font-display);
		font-size: clamp(1.8rem, 3vw, 2.4rem);
		letter-spacing: -0.03em;
	}

	.page__head p {
		margin: 0.45rem 0 0;
		color: var(--ink-soft);
	}

	.product-grid {
		display: grid;
		grid-template-columns: repeat(2, minmax(0, 1fr));
		gap: 0.9rem;
	}

	@media (min-width: 760px) {
		.product-grid {
			grid-template-columns: repeat(3, minmax(0, 1fr));
			gap: 1.15rem;
		}
	}

	@media (min-width: 1024px) {
		.product-grid {
			grid-template-columns: repeat(4, minmax(0, 1fr));
		}
	}
</style>
