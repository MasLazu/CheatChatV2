<script lang="ts">
	import { Avatar } from '@skeletonlabs/skeleton'
	export let message: string = ''
	export let timestamp: string = ''
	export let photo: string = ''
</script>

<div class="flex justify-end">
	<div class="grid grid-cols-[auto_1fr] gap-2 max-w-[70%]">
		<div>
			<header class="flex items-center justify-end p-1">
				<small>{timestamp}</small>
			</header>
			<div class="bg-primary-500 text-white rounded-xl rounded-tr-none space-y-2 p-4">
				<p class="text-sm lg:text-base">{message}</p>
			</div>
		</div>
		<Avatar src="https://i.pravatar.cc/?img={photo}" width="w-12" />
	</div>
</div>
