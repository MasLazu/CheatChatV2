<script lang="ts">
	import { Avatar } from '@skeletonlabs/skeleton'
	import { loadFullChat } from './utils/helper/loadFullChat'
	import { toastStore, type ToastSettings } from '@skeletonlabs/skeleton'

	export let id: number
	export let name: string
	export let photo: number

	//toast
	const triggerToast = (message: string) => {
		const t: ToastSettings = {
			message: message,
			timeout: 10000,
			background: 'variant-filled-error'
		}
		toastStore.trigger(t)
	}

	const handleHandleOpenFUllChat = async (key: number) => {
		try {
			await loadFullChat(key)
		} catch (err) {
			triggerToast('Failed to get chat')
		}
	}
</script>

<!-- svelte-ignore a11y-no-static-element-interactions -->
<!-- svelte-ignore a11y-click-events-have-key-events -->
<div
	class="preview-chat flex items-center rounded-2xl hover:bg-slate-50 p-2.5 cursor-pointer"
	on:click={async () => await handleHandleOpenFUllChat(id)}
>
	<Avatar src="https://i.pravatar.cc/?img={photo}" width="w-14" />
	<div class="dec ml-3 overflow-hidden w-[80%]">
		<h5 class="font-semibold mb-1 truncate">{name}</h5>
	</div>
</div>
