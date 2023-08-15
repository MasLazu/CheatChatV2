<script lang="ts">
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
	import { useForm } from 'svelte-use-form'
	import { currentChatStore } from '$lib/store/currentChat'
	import { createEventDispatcher } from 'svelte'

	let heightTopBar: number

	let popupState: 'addContact' | 'makeGroup' | null = null

	const dispatch = createEventDispatcher()

	const updatePopupState = (state: 'addContact' | 'makeGroup' | null) => {
		popupState = state
		dispatch('updatePopupState', popupState)
	}

	const updateheightTopBar = () => {
		dispatch('updateHeightTopBar', heightTopBar)
	}

	$: heightTopBar, updateheightTopBar()

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

	const sendMessageForm = useForm()

	let chatsWindow: HTMLDivElement
	let bottomChat: HTMLDivElement

	const reset = () => {
		if ($currentChatStore && chatsWindow && bottomChat) {
			setTimeout(() => {
				chatsWindow.scrollTop = bottomChat.offsetTop
				$sendMessageForm.reset()
			}, 100)
		}
	}

	$: $currentChatStore, reset()

	let windowWidth: number
</script>

<svelte:window bind:innerWidth={windowWidth} />

{#if !$currentChatStore || windowWidth > 768}
	<aside class="sidebar gap-3 h-screen lg:w-[27%] md:w-[33%] w-full">
		<div class="bg-white px-1 flex flex-col h-screen overflow-hidden pb-4">
			<div class="pb-3 px-4">
				<header
					class="top-bar py-5 text-primary-500 text-[1.65rem] leading-10 font-medium flex items-center justify-between"
					bind:clientHeight={heightTopBar}
				>
					<div class="flex items-start">
						<h1>Cheat Chat</h1>
						<i class="fa-solid fa-comments text-base" />
					</div>
					<button><i class="bi bi-three-dots-vertical text-[1.9rem] text-slate-600" /></button>
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
								on:click={() => updatePopupState('makeGroup')}
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
								on:click={() => updatePopupState('addContact')}
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
