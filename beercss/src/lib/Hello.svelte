<script lang="ts">
	import type { Snippet } from 'svelte';

	let { content }: { content: Snippet<[string, number]> } = $props();
</script>

<p>Block start</p>
{@render content('Mizuki Kanzaki', 24)}
<p>Block end</p>
