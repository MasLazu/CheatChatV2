<script lang="ts">
	export let message: string = ''
	export let timestamp: string = ''
	export let variant: 'default' | 'compact' = 'default'
</script>

{#if variant === 'default'}
	<div class="flex justify-end mt-2">
		<div class="grid grid-cols-[auto_1fr] gap-2 max-w-[70%]">
			<div class="max-w-full overflow-hidden">
				<header class="flex items-center justify-end p-1 text-sm">
					<small class=" md:text-sm text-xs">{timestamp}</small>
				</header>
				<div
					class="bg-primary-500 text-white rounded-xl rounded-tr-none space-y-2 lg:px-5 lg:py-4 py-3 px-4 max-w-full"
				>
					<p class="text-sm lg:text-base text-ellipsis overflow-hidden max-w-full">{message}</p>
				</div>
			</div>
		</div>
	</div>
{:else}
	<div class="flex justify-end mt-1">
		<div class="grid grid-cols-[auto_1fr] gap-2 max-w-[70%]">
			<div
				class="bg-primary-500 text-white rounded-xl rounded-tr-none space-y-2 lg:px-5 lg:py-4 py-3 px-4 max-w-full overflow-hidden"
			>
				<p class="text-sm lg:text-base text-ellipsis overflow-hidden max-w-full">{message}</p>
			</div>
		</div>
	</div>
{/if}
