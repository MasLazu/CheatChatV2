<script lang="ts">
	import { Avatar } from '@skeletonlabs/skeleton'
	export let name = ''
	export let timestamp = ''
	export let message = ''
	export let photo: number
</script>

<div class="preview-chat flex items-center rounded-2xl hover:bg-slate-50 p-2.5 cursor-pointer">
	<Avatar src="https://i.pravatar.cc/?img={photo}" width="w-14" />
	<div class="dec ml-3 overflow-hidden w-[80%]">
		<div class="flex justify-between">
			<h5 class="font-semibold mb-1 truncate">{name}</h5>
			<p class="text-slate-400 text-sm">{timestamp}</p>
		</div>
		<p class="text-sm text-slate-500 truncate">{message}</p>
	</div>
</div>
