<script lang="ts">
	let showSnackBar = $state(false);
</script>

<button onclick={() => (showSnackBar = !showSnackBar)}>SnackBar</button>

<div class="snackbar {showSnackBar && 'active'}">Some text here</div>
