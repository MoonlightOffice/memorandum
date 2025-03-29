<script lang="ts">
	import 'beercss';
	import '$lib/globals.css';
	import type { Snippet } from 'svelte';

	let { children }: { children: Snippet<[]> } = $props();

	$effect(() => {
		const body = document.querySelector('body');
		body?.classList.remove('light')
		body?.classList.add('dark')
	})
</script>

{@render children()}
