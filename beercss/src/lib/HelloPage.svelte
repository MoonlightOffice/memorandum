<script lang="ts">
	import Hello from '$lib/Hello.svelte';
</script>

{#snippet foo(name: string, age: number)}
	<p>Hi, I am {name}, {age} years old!</p>
{/snippet}

{@render foo('Ichigo', 22)}

<Hello>
	{#snippet content(name: string, age: number)}
		<p>Hello under world, {name}, {age} years old</p>
	{/snippet}
</Hello>
