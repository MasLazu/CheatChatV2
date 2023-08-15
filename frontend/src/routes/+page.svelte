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
