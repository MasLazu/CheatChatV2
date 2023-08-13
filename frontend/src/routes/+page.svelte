<script lang="ts">
	import { Avatar } from '@skeletonlabs/skeleton'
	import OtherChatBubble from '$lib/components/chatBubble/OtherChatBubble.svelte'
	import OurChatBubble from '$lib/components/chatBubble/OurChatBubble.svelte'
	import PreviewMessage from '$lib/components/PreviewMessage.svelte'
	import PageTransition from '$lib/components/PageTransition.svelte'
	import { contactsStore } from '$lib/store/contact'
	import { groupsStore } from '$lib/store/group'
	import { previewChatStore } from '$lib/store/previewChat'
	import ContactPreview from '$lib/components/ContactPreview.svelte'
	import GroupPriview from '$lib/components/GroupPriview.svelte'
	import { flip } from 'svelte/animate'
	import { quintOut } from 'svelte/easing'
	import { crossfade } from 'svelte/transition'
	import { useForm, email, validators, required, Hint } from 'svelte-use-form'
	import type { AxiosResponse, AxiosError } from 'axios'
	import type { apiResponse, messageApiResponse } from '$lib/interfaces/apiResponse'
	import type { contact } from '$lib/interfaces/contact'
	import type { group } from '$lib/interfaces/group'
	import axios from 'axios'
	import { currentChatStore } from '$lib/store/currentChat'
	import { userStore } from '$lib/store/user'
	import { extractDate } from '$lib/utils/helper/extractDate'
	import { nameIfSaved } from '$lib/utils/helper/nameIfSaved'
	import { browser } from '$app/environment'
	import { websocketStore } from '$lib/store/websocket'
	import OtherGroupChatBubble from '$lib/components/chatBubble/OtherGroupChatBubble.svelte'
	import { cacheChatStore } from '$lib/store/cacheChat'

	let heightTopBar: number

	//popup state
	let popupForm: 'addContact' | 'makeGroup' | null = null

	//tab slider logic
	let slidebar = [{ id: 'selectedSLede', mark: 'messages' }]
	$: selectedSidebar = slidebar[0].mark

	const [send, receive] = crossfade({
		duration: (d) => Math.sqrt(d * 200),

		fallback(node, params) {
			const style = getComputedStyle(node)
			const transform = style.transform === 'none' ? '' : style.transform

			return {
				duration: 600,
				easing: quintOut,
				css: (t) => `
					transform: ${transform} scale(${t});
					opacity: ${t}
				`
			}
		}
	})

	//add contact logic
	const addContactForm = useForm()

	let errorAddContact: null | string = null

	interface AddContactApiResponse extends apiResponse {
		data: contact
	}

	const handleAddContact = async () => {
		if ($addContactForm.valid) {
			console.log($addContactForm.values)
			try {
				const result: AxiosResponse<AddContactApiResponse> = await axios.post(
					`http://${import.meta.env.VITE_BACKEND_DOMAIN}/api/login/contact`,
					$addContactForm.values,
					{ withCredentials: true }
				)
				contactsStore.update((prev) => [
					...prev,
					{ name: result.data.data.name, email: result.data.data.email }
				])
			} catch (err) {
				const errors = err as AxiosError<messageApiResponse>
				if (errors.response?.data.data.message) {
					errorAddContact = errors.response?.data.data.message
				}
				setTimeout(() => (errorAddContact = null), 5000)
			}
		}
	}

	const makeGroupForm = useForm()

	let errorMakeGroup: null | string = null

	interface MakeGroupApiResponse extends apiResponse {
		data: group
	}

	const handleMakeGroup = async () => {
		if ($makeGroupForm.valid) {
			try {
				const result: AxiosResponse<MakeGroupApiResponse> = await axios.post(
					`http://${import.meta.env.VITE_BACKEND_DOMAIN}/api/login/group`,
					$makeGroupForm.values,
					{ withCredentials: true }
				)
				groupsStore.update((prev) => [
					...prev,
					{ id: result.data.data.id, name: result.data.data.name }
				])
				console.log(result.data.data)
			} catch (err) {
				const errors = err as AxiosError<messageApiResponse>
				if (errors.response?.data.data.message) {
					errorMakeGroup = errors.response?.data.data.message
				}
				setTimeout(() => (errorMakeGroup = null), 5000)
			}
		}
	}

	const sendMessageForm = useForm()

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

		setTimeout(() => {
			$sendMessageForm.reset()
			chatsWindow.scrollTop = bottomChat.offsetTop
		}, 100)
	}

	$: {
		if ($currentChatStore && chatsWindow && bottomChat) {
			setTimeout(() => {
				if (chatsWindow && bottomChat) chatsWindow.scrollTop = bottomChat.offsetTop
			}, 100)
		} else if (!$currentChatStore) {
			$sendMessageForm.reset()
		}
	}

	let windowWidth: number
</script>

<svelte:window bind:innerWidth={windowWidth} />

<div class="w-screen h-screen flex overflow-hidden">
	<!-- background absolute -->
	<PageTransition amount={-50} trigger={popupForm}>
		{#if popupForm}
			<div
				class="backdrop-brightness-50 absolute top-0 left-0 w-screen h-screen z-40 grid place-content-center"
			>
				{#if popupForm === 'addContact'}
					<!-- add contact form -->
					<form
						on:submit|preventDefault={handleAddContact}
						use:addContactForm
						class="bg-white rounded-2xl px-7 py-8 border-[1px] border-slate-200"
					>
						<h1 class="text-xl font-semibold text-center mb-4">Add Contact</h1>
						<label class="label my-6 flex items-center gap-5">
							<span class="ml-1 w-13">Name</span>
							:
							<div>
								<input
									name="name"
									class="bg-slate-100 border-[1px] border-slate-200 outline-none px-4 py-2.5 rounded-2xl grow"
									type="text"
									placeholder="Input"
									use:validators={[required]}
								/>
								<Hint for="name" on="required" class="text-sm text-red-500 h-0"
									>Name is required</Hint
								>
							</div>
						</label>
						<label class="label my-6 flex items-center gap-5">
							<span class="ml-1 w-13">Email</span>
							:
							<div>
								<input
									name="email"
									class="bg-slate-100 border-[1px] border-slate-200 outline-none px-4 py-2.5 rounded-2xl grow"
									type="text"
									placeholder="Input"
									use:validators={[required, email]}
								/>
								<Hint for="email" on="required" class="text-sm text-red-500 h-0"
									>Email is required</Hint
								>
								<Hint for="email" on="email" hideWhenRequired class="text-sm text-red-500 h-0"
									>Email must be a valid email</Hint
								>
							</div>
						</label>
						<PageTransition trigger={errorAddContact} amount={-8}>
							{#if errorAddContact}
								<div
									class="bg-error-100 p-2.5 w-full text-error-600 rounded-2xl border-2 border-error-300 flex justify-center mt-5"
								>
									<i class="bi bi-exclamation-triangle-fill mr-2" />{errorAddContact}
								</div>
							{/if}
						</PageTransition>
						<div class="w-full flex justify-end gap-3 mt-10">
							<button
								class="btn variant-filled rounded-2xl bg-slate-400 text-white"
								on:click={() => (popupForm = null)}>Cancel</button
							>
							<button
								type="submit"
								class="btn variant-filled rounded-2xl variant-filled-primary text-white">Add</button
							>
						</div>
					</form>
				{:else if popupForm === 'makeGroup'}
					<!--make group form -->
					<form
						on:submit|preventDefault={handleMakeGroup}
						use:makeGroupForm
						class="bg-white rounded-2xl px-7 py-8 border-[1px] border-slate-200"
					>
						<h1 class="text-xl font-semibold text-center mb-4">Make Group</h1>
						<label class="label my-6 flex items-center gap-5">
							<span class="ml-1 w-13">Name</span>
							:
							<div>
								<input
									name="name"
									use:validators={[required]}
									class="bg-slate-100 border-[1px] border-slate-200 outline-none px-4 py-2.5 rounded-2xl grow"
									type="text"
									placeholder="Input"
								/>
								<Hint for="name" on="required" class="text-sm text-red-500 h-0"
									>Name is required</Hint
								>
							</div>
						</label>
						<PageTransition trigger={errorMakeGroup} amount={-8}>
							{#if errorMakeGroup}
								<div
									class="bg-error-100 p-2.5 w-full text-error-600 rounded-2xl border-2 border-error-300 flex justify-center mt-5"
								>
									<i class="bi bi-exclamation-triangle-fill mr-2" />{errorMakeGroup}
								</div>
							{/if}
						</PageTransition>
						<div class="w-full flex justify-end gap-3 mt-10">
							<button
								class="btn variant-filled rounded-2xl bg-slate-400 text-white"
								on:click={() => (popupForm = null)}>Cancel</button
							>
							<button
								type="submit"
								class="btn variant-filled rounded-2xl variant-filled-primary text-white">Add</button
							>
						</div>
					</form>
				{/if}
			</div>
		{/if}
	</PageTransition>
	<!-- sidebar -->
	{#if !$currentChatStore || windowWidth > 768}
		<aside class="sidebar gap-3 h-screen lg:w-[27%] md:w-[33%] w-full">
			<div class="bg-white px-1 flex flex-col h-screen overflow-hidden pb-4">
				<div class="pb-3 px-4">
					<header
						class="top-bar py-5 block text-primary-500 text-3xl leading-10 font-medium"
						bind:clientHeight={heightTopBar}
					>
						<div class="flex items-start">
							<h1>Cheat Chat</h1>
							<i class="fa-solid fa-comments text-base" />
						</div>
					</header>

					<input
						class="input bg-slate-100 border-[1px] border-slate-200 outline-none px-4 py-2.5 rounded-2xl"
						type="text"
						placeholder="Search"
					/>

					<!-- slider sidebar -->
					<div class="nav grid grid-cols-2 px-1.5 mt-6">
						<button class="text-center pb-2.5" on:click={() => (slidebar[0].mark = 'contacts')}
							>Contacts</button
						>
						<button class="text-center pb-2.5" on:click={() => (slidebar[0].mark = 'messages')}
							>Messages</button
						>
						<div class="h-[3px] bg-slate-300 w-full col-span-2 grid grid-cols-2">
							<div class="slide-contacts">
								{#each slidebar.filter((t) => t.mark === 'contacts') as slide (slide.id)}
									<div
										class="bg-primary-500 w-full h-full"
										animate:flip
										in:receive={{ key: slide.id }}
										out:send={{ key: slide.id }}
									/>
								{/each}
							</div>
							<div class="slide-messages">
								{#each slidebar.filter((t) => t.mark === 'messages') as slide (slide.id)}
									<div
										class="bg-primary-500 w-full h-full"
										animate:flip
										in:receive={{ key: slide.id }}
										out:send={{ key: slide.id }}
									/>
								{/each}
							</div>
						</div>
					</div>
				</div>
				<div
					class="schat-list flex-grow flex flex-col gap-0.5 py-2 overflow-y-auto px-1.5 rounded-2xl"
					style="-ms-overflow-style: none;  scrollbar-width: none;"
				>
					<PageTransition trigger={selectedSidebar} amount={-50}>
						{#if selectedSidebar === 'messages'}
							<!-- messages tab -->
							{#each $previewChatStore as previewChat}
								<PreviewMessage
									{...previewChat}
									photo={Math.floor(Math.random() * (40 - 1 + 1)) + 1}
								/>
							{/each}
						{:else if selectedSidebar === 'contacts'}
							<!-- contacts tab -->
							<div class="flex justify-between items-center text-slate-400 p-2.5 pt-0">
								<h4 class="text-sm">
									<i class="bi bi-house-fill mr-2" />Group
								</h4>
								<button
									on:click={() => (popupForm = 'makeGroup')}
									class="h-6 grid place-content-center hover:bg-primary-500 hover:text-white rounded-full cursor-pointer transition-all duration-200"
								>
									<i class="bi bi-plus text-2xl" />
								</button>
							</div>
							{#each $groupsStore as group}
								<GroupPriview
									id={group.id}
									name={group.name}
									photo={Math.floor(Math.random() * (40 - 1 + 1)) + 1}
								/>
							{/each}
							<div class="flex justify-between items-center text-slate-400 p-2.5 pt-0">
								<h4 class="text-sm">
									<i class="bi bi-person-fill mr-2" />Person
								</h4>
								<button
									on:click={() => (popupForm = 'addContact')}
									class="h-6 grid place-content-center hover:bg-primary-500 hover:text-white rounded-full cursor-pointer transition-all duration-200"
								>
									<i class="bi bi-plus text-2xl" />
								</button>
							</div>
							{#each $contactsStore as contact}
								<ContactPreview
									name={contact.name}
									email={contact.email}
									photo={Math.floor(Math.random() * (40 - 1 + 1)) + 1}
								/>
							{/each}
						{/if}
					</PageTransition>
				</div>
			</div>
		</aside>
	{/if}

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
				<div class="profil flex gap-4 items-center">
					<h2 class="lg:text-2xl md:text-xl md:hidden text-lg font-semibold">
						{$currentChatStore?.groupId
							? $currentChatStore.groupName
							: nameIfSaved($currentChatStore?.email ?? '')}
					</h2>
					<Avatar src="https://i.pravatar.cc/?img=26" width="w-[3.2rem]" />
					<h2 class="lg:text-2xl md:text-xl text-lg hidden md:block font-semibold">
						{$currentChatStore?.groupId
							? $currentChatStore.groupName
							: nameIfSaved($currentChatStore?.email ?? '')}
					</h2>
				</div>
				<!-- {#if $currentChatStore?.email}
							<div class="btn-group variant-ghost-primary rounded-2xl h-full text-slate-600">
								<button><i class="fa-solid fa-phone mr-3" />Call</button>
								<button><i class="fa-solid fa-camera mr-3" />Video Call</button>
							</div>
						{/if} -->
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
								variant={Math.floor(chat.createdAt.getTime() / 1000) ===
								Math.floor($currentChatStore.chats[index - 1]?.createdAt.getTime() / 1000)
									? 'compact'
									: 'default'}
							/>
						{:else if $currentChatStore?.groupId}
							<OtherGroupChatBubble
								message={chat.message}
								name={nameIfSaved(chat.senderEmail)}
								timestamp={extractDate(chat.createdAt)}
								photo={Math.floor(Math.random() * (40 - 1 + 1)) + 1}
								variant={Math.floor(chat.createdAt.getTime() / 1000) ===
								Math.floor($currentChatStore.chats[index - 1]?.createdAt.getTime() / 1000)
									? 'compact'
									: 'default'}
							/>
						{:else}
							<OtherChatBubble
								message={chat.message}
								timestamp={extractDate(chat.createdAt)}
								variant={Math.floor(chat.createdAt.getTime() / 1000) ===
								Math.floor($currentChatStore.chats[index - 1]?.createdAt.getTime() / 1000)
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
</div>
