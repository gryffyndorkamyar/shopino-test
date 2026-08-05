<script lang="ts">
	import { goto } from '$app/navigation';
	import { auth } from '$lib/stores/auth';

	let username = $state('');
	let email = $state('');
	let phone = $state('');
	let password = $state('');
	let error = $state('');
	let loading = $state(false);

	async function onSubmit(e: Event) {
		e.preventDefault();
		error = '';
		loading = true;
		try {
			await auth.register({
				username: username.trim(),
				email: email.trim(),
				password,
				phone: phone.trim()
			});
			await goto('/products');
		} catch (err) {
			error = err instanceof Error ? err.message : 'ثبت‌نام ناموفق بود';
		} finally {
			loading = false;
		}
	}
</script>

<svelte:head>
	<title>ثبت‌نام | شاپینو</title>
</svelte:head>

<section class="auth">
	<div class="container auth__wrap">
		<form class="card" onsubmit={onSubmit}>
			<h1>ساخت حساب</h1>
			<p>ثبت‌نام از طریق API احراز هویت جنگو</p>

			<label>
				<span>نام کاربری</span>
				<input bind:value={username} autocomplete="username" required />
			</label>

			<label>
				<span>ایمیل</span>
				<input bind:value={email} type="email" autocomplete="email" required />
			</label>

			<label>
				<span>موبایل (اختیاری)</span>
				<input bind:value={phone} inputmode="tel" autocomplete="tel" />
			</label>

			<label>
				<span>رمز عبور</span>
				<input
					bind:value={password}
					type="password"
					autocomplete="new-password"
					minlength="8"
					required
				/>
			</label>

			{#if error}
				<p class="error">{error}</p>
			{/if}

			<button type="submit" disabled={loading}>{loading ? 'در حال ثبت...' : 'ثبت‌نام'}</button>
			<p class="switch">قبلاً ثبت‌نام کردی؟ <a href="/login">ورود</a></p>
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
