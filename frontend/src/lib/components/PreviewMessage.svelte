<script lang="ts">
	import { Avatar } from '@skeletonlabs/skeleton'
	import { extractDate } from '$lib/utils/helper/extractDate'
	import { userStore } from '../store/user'
	import { loadFullChat } from '$lib/utils/helper/loadFullChat'
	import { toastStore, type ToastSettings } from '@skeletonlabs/skeleton'
	import { nameIfSaved } from '$lib/utils/helper/nameIfSaved'

	export let senderEmail: string
	export let message: string
	export let createdAt: Date
	//null if not group
	export let groupId: number | null
	export let groupName: string | null
	//null if not personal
	export let email: string | null
	export let photo: number = 1

	//toast
	const triggerToast = (message: string) => {
		const t: ToastSettings = {
			message: message,
			timeout: 10000,
			background: 'variant-filled-error'
		}
		toastStore.trigger(t)
	}

	const getGroupMessage = () => {
		if ($userStore && senderEmail === $userStore.email) {
			return `you : ${message}`
		} else {
			return `${nameIfSaved(senderEmail)} : ${message}`
		}
	}

	const handleHandleOpenFUllChat = async (key: number | string) => {
		try {
			await loadFullChat(key)
		} catch (err) {
			triggerToast('Failed to get chat data')
		}
	}
</script>

{#if groupId}
	<!-- svelte-ignore a11y-no-static-element-interactions -->
	<!-- svelte-ignore a11y-click-events-have-key-events -->
	<div
		class="preview-chat flex items-center rounded-2xl hover:bg-slate-50 p-2.5 cursor-pointer"
		on:click={async () => await handleHandleOpenFUllChat(groupId ?? 0)}
	>
		<Avatar src="https://i.pravatar.cc/?img={photo}" width="w-14" />
		<div class="dec ml-3 overflow-hidden w-[80%]">
			<div class="flex justify-between">
				<h5 class="font-semibold mb-1 truncate">{groupName}</h5>
				<p class="text-slate-400 text-xs">{extractDate(createdAt)}</p>
			</div>
			<p class="text-sm text-slate-500 truncate">{getGroupMessage()}</p>
		</div>
	</div>
{:else}
	<!-- svelte-ignore a11y-no-static-element-interactions -->
	<!-- svelte-ignore a11y-click-events-have-key-events -->
	<div
		class="preview-chat flex items-center rounded-2xl hover:bg-slate-50 p-2.5 cursor-pointer"
		on:click={() => handleHandleOpenFUllChat(email ?? '')}
	>
		<Avatar src="https://i.pravatar.cc/?img={photo}" width="w-14" />
		<div class="dec ml-3 overflow-hidden w-[80%]">
			<div class="flex justify-between">
				<h5 class="font-semibold mb-1 truncate">{nameIfSaved(email ?? '')}</h5>
				<p class="text-slate-400 text-xs">{extractDate(createdAt)}</p>
			</div>
			<p class="text-sm text-slate-500 truncate">{message}</p>
		</div>
	</div>
{/if}
