<script lang="ts">
	import { Avatar } from '@skeletonlabs/skeleton'
	import OtherChatBubble from '$lib/components/OtherChatBubble.svelte'
	import OurChatBubble from '$lib/components/OurChatBubble.svelte'
	import PreviewMessage from '$lib/components/PreviewMessage.svelte'
	import PageTransition from '$lib/components/PageTransition.svelte'
	import { contactsStore } from '$lib/store'
	import ContactPreview from '$lib/components/ContactPreview.svelte'
	import GroupPriview from '$lib/components/GroupPriview.svelte'
	import { flip } from 'svelte/animate'
	import { quintOut } from 'svelte/easing'
	import { crossfade } from 'svelte/transition'

	let previewMessageDummy = [
		{
			name: 'Sinta',
			timestamp: '09:26 PM',
			message:
				'Impedit vero recusandae iste assumenda nesciunt. Praesentium nobis maxime quo blanditiis nam.'
		},
		{
			name: 'Yanto',
			timestamp: '09:26 PM',
			message: 'Guys cek Figma dong, minta feedbackenyaa'
		},
		{
			name: 'Handoko',
			timestamp: '09:26 PM',
			message: 'Praesentium nobis maxime quo blanditiis nam'
		}
	]

	let bubbles = [
		{
			photo: '48',
			name: 'yanto',
			timestamp: 'Yesterday 2:30pm',
			message:
				'Impedit vero recusandae iste assumenda nesciunt. Praesentium nobis maxime quo blanditiis nam.'
		},
		{
			photo: '6',
			name: 'supri',
			timestamp: 'Yesterday 2:56pm',
			message:
				'Impedit vero recusandae iste assumenda nesciunt. Praesentium nobis maxime quo blanditiis nam. Nostrum aliquam non.'
		},
		{
			photo: '10',
			name: 'supri',
			timestamp: 'Yesterday 3:20pm',
			message: 'Impedit vero recusandae iste assumenda nesciunt.'
		}
	]

	let heightTopBar: number

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
</script>

<div class="w-screen h-screen flex overflow-hidden bg-slate-100">
	<aside class="sidebar h-screen lg:w-[27%] md:w-[33%] w-full gap-3">
		<div class="bg-white px-1 flex flex-col h-screen overflow-hidden pb-4">
			<div class="pb-3 px-4">
				<div
					class="top-bar py-5 block text-primary-500 text-3xl leading-10 font-medium"
					bind:clientHeight={heightTopBar}
				>
					<div class="flex items-start">
						<h1>Cheat Chat</h1>
						<i class="fa-solid fa-comments text-base" />
					</div>
				</div>

				<input
					class="input bg-slate-100 border-[1px] border-slate-200 outline-none px-4 py-2.5 rounded-2xl"
					type="text"
					placeholder="Search"
				/>

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
				<PageTransition trigger={selectedSidebar}>
					{#if selectedSidebar === 'messages'}
						<!-- messages tab -->
						<h4 class="text-slate-400 text-sm p-2.5 pt-0">
							<i class="bi bi-geo-alt-fill mr-2" />Pinned Message
						</h4>
						{#each previewMessageDummy as message}
							<PreviewMessage {...message} photo={Math.floor(Math.random() * (40 - 1 + 1)) + 1} />
						{/each}
						<h4 class="text-slate-400 text-sm p-2.5">
							<i class="bi bi-chat-left-text-fill mr-2" />All Message
						</h4>
						{#each previewMessageDummy as message}
							<PreviewMessage {...message} photo={Math.floor(Math.random() * (40 - 1 + 1)) + 1} />
						{/each}
						{#each previewMessageDummy as message}
							<PreviewMessage {...message} photo={Math.floor(Math.random() * (40 - 1 + 1)) + 1} />
						{/each}
					{:else if selectedSidebar === 'contacts'}
						<!-- contacts tab -->
						<div class="flex justify-between items-center text-slate-400 p-2.5 pt-0">
							<h4 class="text-sm">
								<i class="bi bi-house-fill mr-2" />Group
							</h4>
							<button
								class="h-6 grid place-content-center hover:bg-primary-500 hover:text-white rounded-full cursor-pointer transition-all duration-200"
							>
								<i class="bi bi-plus text-2xl" />
							</button>
						</div>
						<GroupPriview />
						<div class="flex justify-between items-center text-slate-400 p-2.5 pt-0">
							<h4 class="text-sm">
								<i class="bi bi-house-fill mr-2" />Group
							</h4>
							<button
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

	<main class="lg:w-[73%] md:w-[67%] md:flex hidden flex-col bg-white h-screen">
		<div
			class="top-bar py-4 px-6 flex justify-between items-center"
			style="height: {heightTopBar}px;"
		>
			<div class="profil flex gap-4 items-center">
				<Avatar src="https://i.pravatar.cc/?img=26" width="w-[3.2rem]" />
				<h2 class="text-2xl font-semibold">Yanto Kucul</h2>
			</div>
			<div class="btn-group variant-ghost-primary rounded-2xl h-full text-slate-600">
				<button><i class="fa-solid fa-phone mr-3" />Call</button>
				<button><i class="fa-solid fa-camera mr-3" />Video Call</button>
			</div>
		</div>
		<div
			class="md:ml-1 mx-3 rounded-2xl grow bg-slate-100 overflow-y-auto border-[1px] border-slate-200"
		>
			<div class="p-6 flex flex-col gap-6 rounded-2xl overflow-y-auto h-full">
				{#each bubbles as bubble}
					<OtherChatBubble {...bubble} />
					<OurChatBubble {...bubble} />
				{/each}
			</div>
		</div>
		<div class="top-bar py-3 px-3 flex items-center gap-3">
			<button class="h-full w-auto px-1"
				><i class="bi bi-paperclip text-slate-400 text-3xl" />
			</button>
			<textarea
				rows="1"
				class="textarea rounded-2xl bg-slate-100 p-3 h-full border-slate-200 resize-none flex-grow outline-none"
				placeholder="Type a message"
			/>
			<button type="button" class="btn-icon variant-filled bg-primary-500 h-full w-auto"
				><i class="fa-solid fa-paper-plane text-2xl" /></button
			>
		</div>
	</main>
</div>
