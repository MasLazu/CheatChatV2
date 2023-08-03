<script lang="ts">
	import { Avatar } from '@skeletonlabs/skeleton'
	export let name: string
	export let email: string
	export let photo: number
</script>

<div class="preview-chat flex items-center rounded-2xl hover:bg-slate-50 p-2.5 cursor-pointer">
	<Avatar src="https://i.pravatar.cc/?img={photo}" width="w-14" />
	<div class="dec ml-3 overflow-hidden w-[80%]">
		<h5 class="font-semibold mb-1 truncate">{name}</h5>
		<p class="text-sm text-slate-500 truncate">{email}</p>
	</div>
</div>
