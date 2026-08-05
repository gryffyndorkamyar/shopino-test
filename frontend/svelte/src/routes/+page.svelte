<script lang="ts">
	import ProductCard from '$lib/components/ProductCard.svelte';
	import StateBox from '$lib/components/StateBox.svelte';

	let { data } = $props();

	const categories = [
		'لباس',
		'کیف',
		'کفش',
		'اکسسوری',
		'ساعت',
		'زیورآلات',
		'ورزش',
		'خانه'
	];

	const featured = $derived(data.products.slice(0, 8));
</script>

<section class="hero" aria-label="معرفی شاپینو">
	<div class="hero__media" aria-hidden="true">
		<img
			src="https://images.unsplash.com/photo-1490481651871-ab68de25d43d?auto=format&fit=crop&w=1800&q=80"
			alt=""
		/>
		<div class="hero__veil"></div>
	</div>

	<div class="container hero__content">
		<p class="hero__brand">شاپینو</p>
		<h1>استایل امروزت، از بهترین فروشگاه‌ها</h1>
		<p class="hero__lead">
			مارکت‌پلیس پوشاک برای خریدارانی که کیفیت، تخفیف واقعی و تجربه سریع می‌خواهند.
		</p>
		<div class="hero__actions">
			<a class="hero__cta" href="/products">مشاهده محصولات</a>
			<a class="hero__secondary" href="/stores">فروشگاه‌ها</a>
		</div>
	</div>
</section>

<section id="categories" class="section">
	<div class="container">
		<div class="section__head">
			<h2>دسته‌بندی‌ها</h2>
			<p>مسیر سریع به دنیای استایل</p>
		</div>
		<nav class="category-rail" aria-label="دسته‌بندی محصولات">
			{#each categories as category, i}
				<a class="category-item" href="/products" style="animation-delay: {i * 40}ms">{category}</a>
			{/each}
		</nav>
	</div>
</section>

<section id="products" class="section">
	<div class="container">
		<div class="section__head section__head--row">
			<div>
				<h2>محصولات منتخب</h2>
				<p>مستقیم از API سرویس گولنگ</p>
			</div>
			<a class="see-all" href="/products">مشاهده همه</a>
		</div>

		{#if data.error}
			<StateBox
				tone="error"
				title="اتصال به API برقرار نشد"
				message={data.error}
				hint="Gateway را روی پورت ۸۰۸۸ بالا بیاور، بعد صفحه را رفرش کن."
			/>
		{:else if featured.length === 0}
			<StateBox
				title="هنوز محصولی ثبت نشده"
				message="محصولات دمو باید بعد از بالا آمدن سرویس گولنگ ساخته شوند."
			/>
		{:else}
			<div class="product-grid">
				{#each featured as product, i (product.id)}
					<ProductCard {product} index={i} />
				{/each}
			</div>
		{/if}
	</div>
</section>

<style>
	.hero {
		position: relative;
		min-height: min(92vh, 820px);
		display: grid;
		align-items: end;
		overflow: hidden;
		color: #fff;
	}

	.hero__media {
		position: absolute;
		inset: 0;
	}

	.hero__media img {
		width: 100%;
		height: 100%;
		object-fit: cover;
		animation: kenburns 14s ease-out forwards;
	}

	.hero__veil {
		position: absolute;
		inset: 0;
		background:
			linear-gradient(90deg, rgba(10, 10, 12, 0.78) 8%, rgba(10, 10, 12, 0.35) 58%, rgba(10, 10, 12, 0.2) 100%),
			linear-gradient(0deg, rgba(10, 10, 12, 0.55), transparent 45%);
	}

	.hero__content {
		position: relative;
		z-index: 1;
		padding-block: clamp(4rem, 12vh, 7rem) 4.5rem;
		max-width: 720px;
		margin-inline: 0 auto 0;
		animation: rise 0.8s ease both;
	}

	.hero__brand {
		margin: 0 0 0.85rem;
		font-family: var(--font-display);
		font-size: clamp(2.6rem, 7vw, 4.6rem);
		font-weight: 800;
		letter-spacing: -0.04em;
		line-height: 0.95;
	}

	.hero h1 {
		margin: 0;
		font-family: var(--font-display);
		font-size: clamp(1.55rem, 3.4vw, 2.35rem);
		font-weight: 700;
		line-height: 1.25;
		max-width: 14ch;
	}

	.hero__lead {
		margin: 1rem 0 1.6rem;
		font-size: 1.05rem;
		line-height: 1.8;
		color: rgba(255, 255, 255, 0.86);
		max-width: 34ch;
	}

	.hero__actions {
		display: flex;
		flex-wrap: wrap;
		gap: 0.75rem;
	}

	.hero__cta,
	.hero__secondary {
		display: inline-flex;
		align-items: center;
		justify-content: center;
		padding: 0.9rem 1.35rem;
		border-radius: 12px;
		font-weight: 700;
		transition: background 0.2s ease, transform 0.2s ease, color 0.2s ease;
	}

	.hero__cta {
		background: var(--accent);
		color: #fff;
	}

	.hero__cta:hover {
		background: var(--accent-deep);
		transform: translateY(-2px);
	}

	.hero__secondary {
		background: rgba(255, 255, 255, 0.12);
		border: 1px solid rgba(255, 255, 255, 0.28);
		color: #fff;
	}

	.hero__secondary:hover {
		background: rgba(255, 255, 255, 0.2);
	}

	.section {
		padding-block: clamp(2.8rem, 6vw, 4.5rem);
	}

	.section__head {
		margin-bottom: 1.5rem;
		animation: rise 0.7s ease both;
	}

	.section__head--row {
		display: flex;
		align-items: end;
		justify-content: space-between;
		gap: 1rem;
	}

	.section__head h2 {
		margin: 0;
		font-family: var(--font-display);
		font-size: clamp(1.55rem, 2.6vw, 2rem);
		letter-spacing: -0.03em;
	}

	.section__head p {
		margin: 0.4rem 0 0;
		color: var(--ink-soft);
	}

	.see-all {
		color: var(--accent-deep);
		font-weight: 600;
		white-space: nowrap;
	}

	.category-rail {
		display: grid;
		grid-auto-flow: column;
		grid-auto-columns: max-content;
		gap: 0.75rem;
		overflow-x: auto;
		padding-bottom: 0.4rem;
		scrollbar-width: thin;
	}

	.category-item {
		padding: 0.7rem 0.2rem;
		border-bottom: 2px solid transparent;
		color: var(--ink-soft);
		font-weight: 600;
		white-space: nowrap;
		animation: rise 0.55s ease both;
		transition: color 0.2s ease, border-color 0.2s ease;
	}

	.category-item:hover {
		color: var(--ink);
		border-color: var(--accent);
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

		.hero__content {
			padding-bottom: 5.5rem;
		}
	}
</style>
