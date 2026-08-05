<script lang="ts">
	import { page } from '$app/stores';
</script>

<section class="err">
	<div class="container">
		<p class="code">{$page.status}</p>
		<h1>{$page.error?.message ?? 'خطایی رخ داد'}</h1>
		<a href="/">بازگشت به صفحه اصلی</a>
	</div>
</section>

<style>
	.err {
		padding-block: clamp(4rem, 12vw, 7rem);
	}

	.code {
		margin: 0;
		font-family: var(--font-display);
		font-size: 3rem;
		font-weight: 800;
		color: var(--accent);
		letter-spacing: -0.04em;
	}

	h1 {
		margin: 0.4rem 0 1.2rem;
		font-family: var(--font-display);
		font-size: clamp(1.4rem, 3vw, 2rem);
	}

	a {
		color: var(--accent-deep);
		font-weight: 600;
	}
</style>
