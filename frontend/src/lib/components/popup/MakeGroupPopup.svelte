<script lang="ts">
	import { Hint } from 'svelte-use-form'
	import { useForm } from 'svelte-use-form'
	import axios, { type AxiosResponse, AxiosError } from 'axios'
	import type { apiResponse } from '$lib/interfaces/apiResponse'
	import type { group } from '$lib/interfaces/group'
	import { groupsStore } from '$lib/store/group'
	import type { messageApiResponse } from '$lib/interfaces/apiResponse'
	import PageTransition from '$lib/components/PageTransition.svelte'
	import { createEventDispatcher } from 'svelte'
	import { required } from 'svelte-use-form/validators'
	import { validators } from 'svelte-use-form'

	const makeGroupForm = useForm()

	let popupState: 'makeGroup' | null

	const dispatch = createEventDispatcher()

	const updatePopupState = (state: 'makeGroup' | null) => {
		popupState = state
		dispatch('updatePopupState', popupState)
	}

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
</script>

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
			<Hint for="name" on="required" class="text-sm text-red-500 h-0">Name is required</Hint>
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
			on:click={() => updatePopupState(null)}>Cancel</button
		>
		<button type="submit" class="btn variant-filled rounded-2xl variant-filled-primary text-white"
			>Add</button
		>
	</div>
</form>
