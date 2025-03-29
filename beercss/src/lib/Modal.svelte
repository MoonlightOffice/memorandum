<script lang="ts">
	let showDialog = $state(false);
</script>

<button onclick={() => (showDialog = true)}>Dialog</button>

<div class="overlay" class:active={showDialog}></div>
<dialog class:active={showDialog} class="modal">
	<h5>Title</h5>
	<p>Content of dialog</p>
	<nav>
		<button onclick={() => (showDialog = false)}>Cancel</button>
		<button onclick={() => (showDialog = false)}>Confirm</button>
	</nav>
</dialog>

<style>
	dialog {
		position: relative;
		top: 200px;
	}
</style>
