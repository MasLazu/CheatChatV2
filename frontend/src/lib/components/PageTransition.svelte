<script lang="ts">
	import { fly } from 'svelte/transition'
	import { backInOut } from 'svelte/easing'
	export let trigger: any
</script>

{#key trigger}
	<div
		in:fly={{ y: -50, duration: 250, delay: 300, easing: backInOut }}
		out:fly={{ y: -50, duration: 250, easing: backInOut }}
	>
		<slot />
	</div>
{/key}
