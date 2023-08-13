<script lang="ts">
	import axios from 'axios'
	import type { AxiosResponse, AxiosError } from 'axios'
	import { useForm, email, required, Hint, validators, minLength } from 'svelte-use-form'
	import type { apiResponse, messageApiResponse } from '$lib/interfaces/apiResponse'
	import { goto } from '$app/navigation'
	import { userStore } from '$lib/store/user'
	import PageTransition from '$lib/components/PageTransition.svelte'

	interface LoginApiResponse extends apiResponse {
		data: {
			email: string
			username: string
			message: string
		}
	}

	let errorLogin: string | null = null

	const form = useForm()

	const handleLogin = async () => {
		if ($form.valid) {
			try {
				await axios.post(
					`http://${import.meta.env.VITE_BACKEND_DOMAIN}/api/guest/login`,
					$form.values,
					{ withCredentials: true }
				)
				const result: AxiosResponse<LoginApiResponse> = await axios.get(
					`http://${import.meta.env.VITE_BACKEND_DOMAIN}/api/login/current`,
					{ withCredentials: true }
				)
				userStore.set(result.data.data)
				goto('/')
			} catch (err) {
				const errors = err as AxiosError<messageApiResponse>
				if (errors.response?.data.data.message) {
					errorLogin = errors.response?.data.data.message
				}
				setTimeout(() => (errorLogin = null), 5000)
			}
		}
	}
</script>

<div class="flex justify-center items-center h-screen w-screen">
	<form
		use:form
		class="w-96 bg-white rounded-2xl px-7 py-8 border-[1px] border-slate-200"
		on:submit|preventDefault={handleLogin}
	>
		<h1 class="text-4xl font-semibold text-center mb-10">Welecome Back</h1>
		<label class="label my-3">
			<span class="ml-1">Email</span>
			<input
				name="email"
				class="bg-slate-100 border-[1px] border-slate-200 outline-none px-4 py-2.5 rounded-2xl w-full"
				type="text"
				placeholder="Input"
				use:validators={[required, email]}
			/>
			<Hint for="email" on="required" class="text-sm text-red-500">Email is required</Hint>
			<Hint for="email" on="email" hideWhenRequired class="text-sm text-red-500"
				>Email must be a valid email</Hint
			>
		</label>
		<label class="label mt-3 mb-5">
			<span class="ml-1">Password</span>
			<input
				name="password"
				class="bg-slate-100 border-[1px] border-slate-200 outline-none px-4 py-2.5 rounded-2xl w-full"
				type="password"
				placeholder="Input"
				use:validators={[required, minLength(6)]}
			/>
			<Hint for="password" on="required" class="text-sm text-red-500">Password required</Hint>
			<Hint for="password" on="minLength" hideWhenRequired class="text-sm text-red-500"
				>Password at least 6 character</Hint
			>
		</label>
		<PageTransition trigger={errorLogin} amount={-8}>
			{#if errorLogin}
				<div
					class="bg-error-100 p-2.5 w-full text-error-600 rounded-2xl border-2 border-error-300 flex justify-center mt-5"
				>
					<i class="bi bi-exclamation-triangle-fill mr-2" />{errorLogin}
				</div>
			{/if}
		</PageTransition>
		<button
			type="submit"
			class="btn btn-lg variant-filled w-full mt-8 rounded-2xl variant-filled-primary text-white"
			>Login</button
		>
		<div class="text-center text-sm my-2">
			not have an account ? <a href="/register" class="text-sky-500" data-sveltekit-preload-data
				>Register</a
			>
		</div>
	</form>
</div>
