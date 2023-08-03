<script lang="ts">
	import axios from 'axios'
	import type { AxiosResponse } from 'axios'
	import { useForm, email, required, Hint, validators, minLength } from 'svelte-use-form'
	import type { apiResponse } from '../../lib/interfaces/apiResponse'

	interface RegisterApiResponse extends apiResponse {
		data: {
			message: string
		}
	}

	const form = useForm()

	const handleLogin = async () => {
		if ($form.valid) {
			console.log($form.values)
			try {
				const result: AxiosResponse<RegisterApiResponse> = await axios.post(
					import.meta.env.VITE_BACKEND_DOMAIN + '/api/guest/login',
					$form.values,
					{ withCredentials: true }
				)
				console.log(result.data.data)
			} catch (err) {
				console.log(err)
			}
		}
	}
</script>

<div class="flex justify-center items-center h-screen w-screen">
	<form use:form class="w-96 bg-white rounded-xl px-7 py-8 shadow-lg">
		<h1 class="text-4xl font-semibold text-center mb-10">Register</h1>
		<label class="label my-3">
			<span class="ml-1">Username</span>
			<input
				name="username"
				class="input py-2 px-3 outline-none rounded-xl"
				type="text"
				placeholder="Input"
				use:validators={[required]}
			/>
			<Hint for="username" on="required" class="text-sm text-red-500">Username is required</Hint>
		</label>
		<label class="label my-3">
			<span class="ml-1">Email</span>
			<input
				name="email"
				class="input py-2 px-3 outline-none rounded-xl"
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
				class="input py-2 px-3 outline-none rounded-xl"
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
			type="button"
			on:click|preventDefault={handleLogin}
			class="btn btn-lg variant-filled w-full mt-10 rounded-xl">Register</button
		>
		<div class="text-center text-sm my-2">
			not have an account ? <a href="/login" class="text-sky-500" data-sveltekit-preload-data
				>Login</a
			>
		</div>
	</form>
</div>
