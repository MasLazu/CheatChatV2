<script lang="ts">
	import { slide } from 'svelte/transition'
	import { backInOut } from 'svelte/easing'
	export let trigger: any
</script>

{#key trigger}
	<div
		in:slide={{ duration: 250, delay: 300, easing: backInOut, axis: 'x' }}
		out:slide={{ duration: 250, easing: backInOut, axis: 'x' }}
	>
		<slot />
	</div>
{/key}
