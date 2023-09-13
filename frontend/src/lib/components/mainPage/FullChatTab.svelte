<script lang="ts">
	import { currentChatStore } from '$lib/store/currentChat'
	import { browser } from '$app/environment'
	import { nameIfSaved } from '$lib/utils/helper/nameIfSaved'
	import { Avatar } from '@skeletonlabs/skeleton'
	import { userStore } from '$lib/store/user'
	import OurChatBubble from '../chatBubble/OurChatBubble.svelte'
	import { extractDate } from '$lib/utils/helper/extractDate'
	import OtherGroupChatBubble from '../chatBubble/OtherGroupChatBubble.svelte'
	import OtherChatBubble from '../chatBubble/OtherChatBubble.svelte'
	import { useForm } from 'svelte-use-form'
	import { websocketStore } from '$lib/store/websocket'
	import { validators } from 'svelte-use-form'
	import { required } from 'svelte-use-form/validators'

	export let heightTopBar: number

	const sendMessageForm = useForm()

	let windowWidth: number = 0

	let chatsWindow: HTMLDivElement
	let bottomChat: HTMLDivElement

	const handleSendMessage = async () => {
		if ($sendMessageForm.valid) {
			$websocketStore?.send(
				JSON.stringify({
					type: $currentChatStore?.groupId ? 2 : 1,
					body: {
						message: $sendMessageForm.values.message,
						group_id: $currentChatStore?.groupId,
						receiver_email: $currentChatStore?.email,
						sender_email: $userStore?.email
					}
				})
			)
		}
	}

    const reset = () => {
		if ($currentChatStore && chatsWindow && bottomChat) {
			setTimeout(() => {
				chatsWindow.scrollTop = bottomChat.offsetTop
				$sendMessageForm.reset()
			}, 100)
		}
	}

	$: $currentChatStore, reset()
    
</script>

<svelte:window bind:innerWidth={windowWidth} />

{#if $currentChatStore}
	<main
		class="flex {browser ? '' : 'hidden'} lg:w-[73%] md:w-[67%] w-full flex-col bg-white h-screen"
	>
		<div
			class="top-bar md:py-4 py-2.5 md:px-6 px-4 flex justify-between items-center"
			style={windowWidth > 768 ? `height: ${heightTopBar}px;` : ''}
		>
			{#if browser && windowWidth < 768}
				<button class="h-10 grid items-center" on:click={() => currentChatStore.set(null)}
					><i class="bi bi-arrow-left-short text-[3rem] text-slate-600" /></button
				>
			{/if}
			<div class="profil flex gap-4 items-center overflow-hidden justify-between">
				<h2 class="lg:text-2xl md:text-xl md:hidden text-lg font-semibold">
					{$currentChatStore?.groupId
						? $currentChatStore.groupName
						: nameIfSaved($currentChatStore?.email ?? '')}
				</h2>
				<div class="flex items-center gap-4 overflow-hidden">
					<Avatar src="https://i.pravatar.cc/?img=26" width="w-14" />
					<div class="overflow-hidden max-w-[70%]">
						<h2
							class="lg:text-2xl md:text-xl text-lg hidden md:block font-semibold truncate overflow-hidden"
						>
							{$currentChatStore?.groupId
								? $currentChatStore.groupName
								: nameIfSaved($currentChatStore?.email ?? '')}
						</h2>
						{#if $currentChatStore.groupId}
							<h5 class="text-sm truncate overflow-hidden text-slate-600">
								yanto, supri, handoko@gmail.com, gilang@gmail.com, agus@gmail.com, bajigur@gmail.com
							</h5>
						{/if}
					</div>
				</div>
				{#if $currentChatStore.groupId}
					<button>
						<i
							class="bi bi-person-plus-fill text-[2.6rem] text-slate-400 hover:text-primary-500 transition-all duration-150"
						/>
					</button>
				{/if}
			</div>
			{#if $currentChatStore?.email && windowWidth > 768}
				<div class="btn-group variant-ghost-primary rounded-2xl h-full text-slate-600">
					<button><i class="fa-solid fa-phone mr-3" />Call</button>
					<button><i class="fa-solid fa-camera mr-3" />Video Call</button>
				</div>
			{/if}
		</div>
		<div
			class="md:ml-1 mx-3 rounded-2xl grow bg-slate-100 overflow-y-auto border-[1px] border-slate-200"
		>
			<div
				class="lg:p-6 md:p-4 p-3 px-2 flex flex-col rounded-2xl overflow-y-auto h-full"
				bind:this={chatsWindow}
			>
				{#each $currentChatStore?.chats ?? [] as chat, index}
					{#if chat.senderEmail === $userStore?.email}
						<OurChatBubble
							message={chat.message}
							timestamp={extractDate(chat.createdAt)}
							variant={Math.floor(chat.createdAt.getTime() / 60000) ===
							Math.floor($currentChatStore.chats[index - 1]?.createdAt.getTime() / 60000)
								? 'compact'
								: 'default'}
						/>
					{:else if $currentChatStore?.groupId}
						<OtherGroupChatBubble
							message={chat.message}
							name={nameIfSaved(chat.senderEmail)}
							timestamp={extractDate(chat.createdAt)}
							photo={Math.floor(Math.random() * (40 - 1 + 1)) + 1}
							variant={Math.floor(chat.createdAt.getTime() / 60000) ===
							Math.floor($currentChatStore.chats[index - 1]?.createdAt.getTime() / 60000)
								? 'compact'
								: 'default'}
						/>
					{:else}
						<OtherChatBubble
							message={chat.message}
							timestamp={extractDate(chat.createdAt)}
							variant={Math.floor(chat.createdAt.getTime() / 60000) ===
							Math.floor($currentChatStore.chats[index - 1]?.createdAt.getTime() / 60000)
								? 'compact'
								: 'default'}
						/>
					{/if}
				{/each}
				<div bind:this={bottomChat} />
			</div>
		</div>
		<form
			class="top-bar py-3 px-3 flex items-center gap-3"
			on:submit|preventDefault={handleSendMessage}
			use:sendMessageForm
		>
			<button type="button" class="h-full w-auto px-1"
				><i class="bi bi-paperclip text-slate-400 text-3xl" />
			</button>
			<textarea
				rows="1"
				name="message"
				class="textarea rounded-2xl bg-slate-100 p-3 h-full border-slate-200 resize-none flex-grow outline-none"
				placeholder="Type a message"
				use:validators={[required]}
			/>
			<button type="submit" class="btn-icon variant-filled bg-primary-500 h-full w-auto"
				><i class="fa-solid fa-paper-plane text-2xl" /></button
			>
		</form>
	</main>
{:else if windowWidth > 768}
	<div
		class="md:ml-1 ml-3 rounded-l-2xl grow bg-slate-100 overflow-y-auto border-slate-200 h-screen"
	/>
{/if}
