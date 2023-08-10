<script lang="ts">
	import { Avatar } from '@skeletonlabs/skeleton'
	export let message: string = ''
	export let timestamp: string = ''
	export let photo: number = 1
	export let name: string = ''
</script>

<div class="flex justify-start">
	<div class="grid grid-cols-[auto_1fr] gap-2 max-w-[70%]">
		<Avatar src="https://i.pravatar.cc/?img={photo}" width="lg:w-12 md:w-10 w-8" />
		<div>
			<header class="flex gap-6 items-center p-1">
				<p class="font-bold text-sm">{name}</p>
				<small class="opacity-50 text-sm">{timestamp}</small>
			</header>
			<div class="card lg:p-4 p-3 bg-white border-slate-300 rounded-2xl rounded-tl-none space-y-2">
				<p class="text-sm lg:text-base">{message}</p>
			</div>
		</div>
	</div>
</div>
