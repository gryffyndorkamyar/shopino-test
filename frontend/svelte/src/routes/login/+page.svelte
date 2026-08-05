<script lang="ts">
	import { goto } from '$app/navigation';
	import { auth } from '$lib/stores/auth';

	let username = $state('');
	let password = $state('');
	let error = $state('');
	let loading = $state(false);

	async function onSubmit(e: Event) {
		e.preventDefault();
		error = '';
		loading = true;
		try {
			await auth.login(username.trim(), password);
			await goto('/products');
		} catch (err) {
			error = err instanceof Error ? err.message : 'ورود ناموفق بود';
		} finally {
			loading = false;
		}
	}
</script>

<svelte:head>
	<title>ورود | شاپینو</title>
</svelte:head>

<section class="auth">
	<div class="container auth__wrap">
		<form class="card" onsubmit={onSubmit}>
			<h1>ورود به شاپینو</h1>
			<p>با حساب جنگو وارد شو تا بتونی محصول جدید بسازی.</p>

			<label>
				<span>نام کاربری</span>
				<input bind:value={username} autocomplete="username" required />
			</label>

			<label>
				<span>رمز عبور</span>
				<input bind:value={password} type="password" autocomplete="current-password" required />
			</label>

			{#if error}
				<p class="error">{error}</p>
			{/if}

			<button type="submit" disabled={loading}>{loading ? 'در حال ورود...' : 'ورود'}</button>
			<p class="switch">حساب نداری؟ <a href="/register">ثبت‌نام</a></p>
		</form>
	</div>
</section>

<style>
	.auth {
		padding-block: clamp(2.5rem, 8vw, 5rem);
	}

	.auth__wrap {
		display: grid;
		justify-items: center;
	}

	.card {
		width: min(100%, 420px);
		display: grid;
		gap: 0.9rem;
		padding: 1.5rem;
		border: 1px solid var(--line);
		border-radius: calc(var(--radius) + 2px);
		background: var(--surface);
		box-shadow: var(--shadow-soft);
		animation: rise 0.55s ease both;
	}

	.card h1 {
		margin: 0;
		font-family: var(--font-display);
		font-size: 1.7rem;
		letter-spacing: -0.03em;
	}

	.card > p {
		margin: 0;
		color: var(--ink-soft);
		line-height: 1.7;
	}

	label {
		display: grid;
		gap: 0.4rem;
		font-size: 0.92rem;
	}

	input {
		border: 1px solid var(--line);
		border-radius: 10px;
		padding: 0.75rem 0.85rem;
		font: inherit;
		background: #fff;
	}

	input:focus {
		outline: 2px solid rgba(255, 45, 111, 0.25);
		border-color: var(--accent);
	}

	button {
		margin-top: 0.3rem;
		border: 0;
		border-radius: 12px;
		padding: 0.9rem 1rem;
		background: var(--ink);
		color: #fff;
		font: inherit;
		font-weight: 700;
		cursor: pointer;
	}

	button:hover:not(:disabled) {
		background: var(--accent);
	}

	button:disabled {
		opacity: 0.7;
		cursor: wait;
	}

	.error {
		margin: 0;
		color: var(--accent-deep);
		font-size: 0.9rem;
	}

	.switch {
		margin: 0;
		text-align: center;
		color: var(--ink-soft);
		font-size: 0.9rem;
	}

	.switch a {
		color: var(--accent-deep);
		font-weight: 600;
	}
</style>
