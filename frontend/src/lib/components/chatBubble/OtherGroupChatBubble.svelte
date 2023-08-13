<script lang="ts">
	import { Avatar } from '@skeletonlabs/skeleton'
	export let message: string = ''
	export let timestamp: string = ''
	export let photo: number = 1
	export let name: string = ''
	export let variant: 'default' | 'compact' = 'default'
</script>

{#if variant === 'default'}
	<div class="flex justify-start mt-2">
		<div class="grid grid-cols-[auto_1fr] gap-2 max-w-[70%]">
			<Avatar src="https://i.pravatar.cc/?img={photo}" width="lg:w-12 md:w-10 w-8" />
			<div class="max-w-full overflow-hidden">
				<header class="flex gap-6 items-center p-1">
					<p class="font-bold text-sm">{name}</p>
					<small class="opacity-50 md:text-sm text-xs">{timestamp}</small>
				</header>
				<div
					class="card lg:px-5 lg:py-4 py-3 px-4 bg-white border-slate-300 rounded-2xl rounded-tl-none space-y-2 max-w-full overflow-hidden"
				>
					<p class="text-sm lg:text-base text-ellipsis overflow-hidden">{message}</p>
				</div>
			</div>
		</div>
	</div>
{:else}
	<div class="flex justify-start mt-1">
		<div class="grid grid-cols-[auto_1fr] gap-2 max-w-[70%]">
			<div
				class="card lg:px-5 lg:py-4 py-3 px-4 bg-white border-slate-300 rounded-2xl rounded-tl-none space-y-2"
			>
				<p class="text-sm lg:text-base text-ellipsis overflow-hidden">{message}</p>
			</div>
		</div>
	</div>
{/if}
