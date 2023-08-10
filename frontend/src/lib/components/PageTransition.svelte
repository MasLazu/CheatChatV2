<script lang="ts">
	import { fly } from 'svelte/transition'
	import { backInOut } from 'svelte/easing'
	export let trigger: any
	export let amount: number
</script>

{#key trigger}
	<div
		in:fly={{ y: amount, duration: 250, delay: 300, easing: backInOut }}
		out:fly={{ y: amount, duration: 250, easing: backInOut }}
	>
		<slot />
	</div>
{/key}
