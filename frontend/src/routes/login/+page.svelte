<script lang="ts">
	import axios from 'axios'
	import type { AxiosResponse } from 'axios'
	import { useForm, email, required, Hint, validators, minLength } from 'svelte-use-form'
	import type { apiResponse } from '../../lib/interfaces/apiResponse'
	import { goto } from '$app/navigation'
	import { userStore } from '$lib/store'

	interface LoginApiResponse extends apiResponse {
		data: {
			email: string
			username: string
			message: string
		}
	}

	const form = useForm()

	const handleLogin = async () => {
		if ($form.valid) {
			try {
				await axios.post(import.meta.env.VITE_BACKEND_DOMAIN + '/api/guest/login', $form.values, {
					withCredentials: true
				})
				const result: AxiosResponse<LoginApiResponse> = await axios.post(
					import.meta.env.VITE_BACKEND_DOMAIN + '/api/login/current',
					$form.values,
					{ withCredentials: true }
				)
				userStore.set(result.data.data)
				goto('/')
			} catch (err) {
				console.log(err)
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
				class="bg-slate-100 border-[1px] border-slate-200 outline-none px-4 py-2.5 rounded-2xl"
				type="text"
				placeholder="Input"
				use:validators={[required, email]}
			/>
			<Hint for="email" on="required" class="text-sm text-red-500">Email is required</Hint>
			<Hint for="email" on="email" hideWhenRequired class="text-sm text-red-500"
				>Email must be a valid email</Hint
			>
		</label>
		<label class="label my-3">
			<span class="ml-1">Password</span>
			<input
				name="password"
				class="bg-slate-100 border-[1px] border-slate-200 outline-none px-4 py-2.5 rounded-2xl"
				type="password"
				placeholder="Input"
				use:validators={[required, minLength(6)]}
			/>
			<Hint for="password" on="required" class="text-sm text-red-500">Password required</Hint>
			<Hint for="password" on="minLength" hideWhenRequired class="text-sm text-red-500"
				>Password at least 6 character</Hint
			>
		</label>
		<button
			type="submit"
			class="btn btn-lg variant-filled w-full mt-10 rounded-2xl variant-filled-primary text-white"
			>Login</button
		>
		<div class="text-center text-sm my-2">
			not have an account ? <a href="/register" class="text-sky-500" data-sveltekit-preload-data
				>Register</a
			>
		</div>
	</form>
</div>
