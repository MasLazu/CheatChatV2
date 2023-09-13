<script lang="ts">
	import PageTransition from '$lib/components/PageTransition.svelte'
	import AddContactPopup from '$lib/components/popup/AddContactPopup.svelte'
	import MakeGroupPopup from '$lib/components/popup/MakeGroupPopup.svelte'
	import FullChatTab from '$lib/components/mainPage/FullChatTab.svelte'
	import Sidebar from '$lib/components/mainPage/Sidebar.svelte'

	let heightTopBar: number

	//popup state
	let popupForm: 'addContact' | 'makeGroup' | null = null

	const updatePopupState = (event: CustomEvent<'addContact' | 'makeGroup' | null>) => {
		popupForm = event.detail
	}

	const updateheightTopBar = (event: CustomEvent<number>) => {
		heightTopBar = event.detail
	}
</script>

<div class="w-screen h-screen flex overflow-hidden">
	<!-- experimental -->
	<!-- <div
		class="backdrop-brightness-50 absolute top-0 left-0 w-screen h-screen z-40 grid place-content-center"
	>
		<div class="md:p-6 p-5 bg-white rounded-2xl w-96 border border-slate-200">
			<h3 class="text-2xl text-primary-500 font-semibold mb-4">
				Invite to group <i class="bi bi-person-plus-fill ml-1" />
			</h3>
			<input
				class="input bg-slate-100 border-[1px] border-slate-200 outline-none px-4 py-2.5 rounded-2xl"
				type="text"
				placeholder="Search"
			/>
			<div
				class="p-2 bg-slate-100 rounded-2xl md:mt-6 mt-5 border border-slate-200 max-h-80 overflow-auto"
			>
				<div
					class="preview-chat flex items-center rounded-2xl hover:bg-slate-50 p-2.5 cursor-pointer"
				>
					<Avatar src="https://i.pravatar.cc/?img=9" width="w-14" />
					<div class="dec ml-3 overflow-hidden w-[80%]">
						<h5 class="font-semibold mb-1 truncate">suprianto</h5>
						<p class="text-sm text-slate-500 truncate">suprianto@gmail.com</p>
					</div>
					<button class="text-success-500 p-1">
						<i class="bi bi-check-lg text-2xl" />
					</button>
				</div>
				{#each $contactsStore as contact}
					<div
						class="preview-chat flex items-center rounded-2xl hover:bg-slate-50 p-2.5 cursor-pointer"
					>
						<Avatar src="https://i.pravatar.cc/?img=9" width="w-14" />
						<div class="dec ml-3 overflow-hidden w-[80%]">
							<h5 class="font-semibold mb-1 truncate">{contact.name}</h5>
							<p class="text-sm text-slate-500 truncate">{contact.email}</p>
						</div>
					</div>
				{/each}
			</div>
			<div class="flex md:gap-6 gap-4">
				<button
					type="button"
					class="btn btn-lg bg-slate-400 rounded-2xl text-white md:mt-6 mt-5 flex-1">Cancel</button
				>
				<button
					type="button"
					class="btn btn-lg variant-filled-primary rounded-2xl text-white md:mt-6 mt-5 flex-1"
					>Invite</button
				>
			</div>
		</div>
	</div> -->
	<!-- background absolute -->
	<PageTransition amount={-50} trigger={popupForm}>
		{#if popupForm}
			<div
				class="backdrop-brightness-50 absolute top-0 left-0 w-screen h-screen z-40 grid place-content-center"
			>
				{#if popupForm === 'addContact'}
					<AddContactPopup on:updatePopupState={updatePopupState} />
				{:else if popupForm === 'makeGroup'}
					<MakeGroupPopup on:updatePopupState={updatePopupState} />
				{/if}
			</div>
		{/if}
	</PageTransition>
	<!-- sidebar -->
	<Sidebar on:updatePopupState={updatePopupState} on:updateHeightTopBar={updateheightTopBar} />
	<!-- full chat tab -->
	<FullChatTab {heightTopBar} />
</div>
