<script lang="ts">
	export let message: string = ''
	export let timestamp: string = ''
	export let variant: 'default' | 'compact' = 'default'
</script>

{#if variant === 'default'}
	<div class="flex justify-start mt-2">
		<div class="grid grid-cols-[auto_1fr] gap-2 max-w-[70%]">
			<div class="max-w-full overflow-hidden">
				<header class="flex gap-6 items-center p-1">
					<small class="opacity-50 md:text-sm text-xs">{timestamp}</small>
				</header>
				<div
					class="card lg:px-5 lg:py-4 py-3 px-4 bg-white border-slate-300 rounded-2xl rounded-tl-none space-y-2 max-w-full"
				>
					<p class="text-sm lg:text-base text-ellipsis overflow-hidden max-w-full">{message}</p>
				</div>
			</div>
		</div>
	</div>
{:else}
	<div class="flex justify-start mt-1">
		<div class="grid grid-cols-[auto_1fr] gap-2 max-w-[70%]">
			<div
				class="card lg:px-5 lg:py-4 py-3 px-4 bg-white border-slate-300 rounded-2xl rounded-tl-none space-y-2 max-w-full"
			>
				<p class="text-sm lg:text-base text-ellipsis overflow-hidden max-w-full">{message}</p>
			</div>
		</div>
	</div>
{/if}
