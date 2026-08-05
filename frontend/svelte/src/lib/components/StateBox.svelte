<script lang="ts">
	let {
		title,
		message,
		hint = '',
		tone = 'neutral'
	}: {
		title: string;
		message: string;
		hint?: string;
		tone?: 'neutral' | 'error';
	} = $props();
</script>

<div class="state" class:state--error={tone === 'error'} role="status">
	<strong>{title}</strong>
	<p>{message}</p>
	{#if hint}
		<p class="state__hint">{hint}</p>
	{/if}
</div>

<style>
	.state {
		padding: 1.4rem 1.2rem;
		border: 1px dashed var(--line);
		border-radius: var(--radius);
		background: rgba(255, 255, 255, 0.72);
	}

	.state--error {
		border-color: rgba(255, 45, 111, 0.35);
	}

	.state strong {
		display: block;
		margin-bottom: 0.35rem;
	}

	.state p {
		margin: 0.25rem 0;
		color: var(--ink-soft);
		line-height: 1.7;
	}

	.state__hint {
		font-size: 0.88rem;
	}
</style>
