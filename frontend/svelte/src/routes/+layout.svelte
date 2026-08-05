<script lang="ts">
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import '../app.css';
	import { auth } from '$lib/stores/auth';

	let { children } = $props();

	onMount(() => {
		auth.hydrate();
	});
</script>

<svelte:head>
	<link rel="preconnect" href="https://fonts.googleapis.com" />
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="anonymous" />
	<link
		href="https://fonts.googleapis.com/css2?family=Syne:wght@600;700;800&family=Vazirmatn:wght@400;500;600;700&display=swap"
		rel="stylesheet"
	/>
	<title>شاپینو | مارکت‌پلیس هوشمند پوشاک</title>
</svelte:head>

<div class="shell">
	<header class="topbar">
		<div class="container topbar__inner">
			<a class="brand" href="/">شاپینو</a>
			<nav class="nav" aria-label="منوی اصلی">
				<a href="/products" class:active={$page.url.pathname.startsWith('/products')}>محصولات</a>
				<a href="/stores" class:active={$page.url.pathname.startsWith('/stores')}>فروشگاه‌ها</a>
			</nav>
			<div class="topbar__actions">
				{#if $auth.ready && $auth.user}
					<span class="user">{$auth.user.username}</span>
					<button type="button" class="ghost" onclick={() => auth.logout()}>خروج</button>
				{:else}
					<a class="ghost" href="/login">ورود</a>
					<a class="topbar__cta" href="/register">ثبت‌نام</a>
				{/if}
			</div>
		</div>
	</header>

	<main>
		{@render children()}
	</main>

	<footer class="footer">
		<div class="container footer__grid">
			<div>
				<strong>شاپینو</strong>
				<p>تجربه‌ای نو در خرید آنلاین پوشاک — نسخه تستی متصل به Django و Golang</p>
			</div>
			<div class="footer__links">
				<a href="/products">محصولات</a>
				<a href="/stores">فروشگاه‌ها</a>
				<a href="/login">ورود</a>
			</div>
		</div>
	</footer>
</div>

<style>
	.shell {
		min-height: 100vh;
		display: flex;
		flex-direction: column;
	}

	main {
		flex: 1;
	}

	.topbar {
		position: sticky;
		top: 0;
		z-index: 20;
		backdrop-filter: blur(14px);
		background: rgba(244, 245, 247, 0.82);
		border-bottom: 1px solid var(--line);
	}

	.topbar__inner {
		display: flex;
		align-items: center;
		justify-content: space-between;
		gap: 1rem;
		padding-block: 0.95rem;
	}

	.brand {
		font-family: var(--font-display);
		font-weight: 800;
		font-size: clamp(1.45rem, 2.4vw, 1.85rem);
		letter-spacing: -0.03em;
	}

	.nav {
		display: none;
		gap: 1.4rem;
		color: var(--ink-soft);
		font-size: 0.95rem;
	}

	.nav a:hover,
	.nav a.active {
		color: var(--ink);
	}

	.topbar__actions {
		display: flex;
		align-items: center;
		gap: 0.65rem;
	}

	.user {
		display: none;
		color: var(--ink-soft);
		font-size: 0.88rem;
	}

	.ghost {
		border: 0;
		background: transparent;
		color: var(--ink-soft);
		font: inherit;
		padding: 0.45rem 0.6rem;
		border-radius: 10px;
		cursor: pointer;
	}

	.ghost:hover {
		color: var(--ink);
	}

	.topbar__cta {
		display: inline-flex;
		align-items: center;
		justify-content: center;
		padding: 0.55rem 1rem;
		border-radius: 10px;
		background: var(--ink);
		color: #fff;
		font-size: 0.9rem;
		font-weight: 600;
		transition: transform 0.2s ease, background 0.2s ease;
	}

	.topbar__cta:hover {
		background: var(--accent);
		transform: translateY(-1px);
	}

	.footer {
		border-top: 1px solid var(--line);
		margin-top: 4rem;
		padding: 2rem 0 2.5rem;
	}

	.footer__grid {
		display: flex;
		flex-wrap: wrap;
		justify-content: space-between;
		gap: 1.2rem;
	}

	.footer strong {
		font-family: var(--font-display);
		font-size: 1.2rem;
	}

	.footer p {
		margin: 0.35rem 0 0;
		color: var(--ink-soft);
		font-size: 0.92rem;
		max-width: 36ch;
		line-height: 1.7;
	}

	.footer__links {
		display: flex;
		gap: 1rem;
		align-items: center;
		color: var(--ink-soft);
		font-size: 0.92rem;
	}

	.footer__links a:hover {
		color: var(--ink);
	}

	@media (min-width: 760px) {
		.nav {
			display: flex;
		}

		.user {
			display: inline;
		}
	}
</style>
