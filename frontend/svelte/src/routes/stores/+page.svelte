<script lang="ts">
	import StateBox from '$lib/components/StateBox.svelte';

	let { data } = $props();
</script>

<svelte:head>
	<title>فروشگاه‌ها | شاپینو</title>
</svelte:head>

<section class="page">
	<div class="container">
		<header class="page__head">
			<h1>برترین فروشگاه‌ها</h1>
			<p>لیست فروشگاه‌های فعال از API جنگو</p>
		</header>

		{#if data.error}
			<StateBox tone="error" title="خطا در دریافت فروشگاه‌ها" message={data.error} />
		{:else if data.stores.length === 0}
			<StateBox
				title="فروشگاهی ثبت نشده"
				message="بعد از seed دیتابیس جنگو، فروشگاه‌ها اینجا دیده می‌شوند."
			/>
		{:else}
			<div class="store-grid">
				{#each data.stores as store, i (store.id)}
					<article class="store" style="animation-delay: {Math.min(i, 8) * 50}ms">
						<div class="store__mark">{store.name.slice(0, 1)}</div>
						<div>
							<h2>{store.name}</h2>
							<p class="store__owner">@{store.owner_username}</p>
							{#if store.description}
								<p class="store__desc">{store.description}</p>
							{/if}
						</div>
					</article>
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

	.store-grid {
		display: grid;
		gap: 0.9rem;
	}

	.store {
		display: grid;
		grid-template-columns: auto 1fr;
		gap: 1rem;
		align-items: start;
		padding: 1.1rem 1.15rem;
		border: 1px solid var(--line);
		border-radius: var(--radius);
		background: var(--surface);
		animation: rise 0.6s ease both;
	}

	.store__mark {
		width: 3rem;
		height: 3rem;
		border-radius: 12px;
		display: grid;
		place-items: center;
		background: var(--ink);
		color: #fff;
		font-family: var(--font-display);
		font-weight: 700;
		font-size: 1.2rem;
	}

	.store h2 {
		margin: 0;
		font-size: 1.1rem;
	}

	.store__owner {
		margin: 0.25rem 0 0;
		color: var(--accent-deep);
		font-size: 0.88rem;
	}

	.store__desc {
		margin: 0.55rem 0 0;
		color: var(--ink-soft);
		line-height: 1.7;
		font-size: 0.92rem;
	}

	@media (min-width: 760px) {
		.store-grid {
			grid-template-columns: repeat(2, minmax(0, 1fr));
			gap: 1rem;
		}
	}
</style>
