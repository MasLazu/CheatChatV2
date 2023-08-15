<script lang="ts">
	import { Hint } from 'svelte-use-form'
	import { useForm, required, email, validators } from 'svelte-use-form'
	import axios, { type AxiosResponse, AxiosError } from 'axios'
	import type { apiResponse } from '$lib/interfaces/apiResponse'
	import type { contact } from '$lib/interfaces/contact'
	import { contactsStore } from '$lib/store/contact'
	import type { messageApiResponse } from '$lib/interfaces/apiResponse'
	import PageTransition from '$lib/components/PageTransition.svelte'
	import { createEventDispatcher } from 'svelte'

	let popupState: 'addContact' | null

	const dispatch = createEventDispatcher()

	const updatePopupState = (state: 'addContact' | null) => {
		popupState = state
		dispatch('updatePopupState', popupState)
	}

	let errorAddContact: string | null = null

	const addContactForm = useForm()

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
</script>

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
			<Hint for="name" on="required" class="text-sm text-red-500 h-0">Name is required</Hint>
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
			<Hint for="email" on="required" class="text-sm text-red-500 h-0">Email is required</Hint>
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
			on:click={() => updatePopupState(null)}>Cancel</button
		>
		<button type="submit" class="btn variant-filled rounded-2xl variant-filled-primary text-white"
			>Add</button
		>
	</div>
</form>
