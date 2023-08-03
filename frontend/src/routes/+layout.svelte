<script lang="ts">
	import '@skeletonlabs/skeleton/themes/theme-skeleton.css'
	import '@skeletonlabs/skeleton/styles/skeleton.css'
	import '../app.postcss'
	import { page } from '$app/stores'
	import PageTransition from '$lib/components/PageTransition.svelte'
	import { beforeNavigate } from '$app/navigation'
	import { onMount } from 'svelte'
	import axios from 'axios'
	import { goto } from '$app/navigation'
	import { userStore, groupsStore } from '$lib/store'
	import type { AxiosResponse } from 'axios'
	import type { apiResponse } from '$lib/interfaces/apiResponse'
	import type { user } from '$lib/interfaces/user'
	import type { group } from '$lib/interfaces/group'
	import { Toast, type ToastSettings } from '@skeletonlabs/skeleton'
	import { toastStore } from '@skeletonlabs/skeleton'
	import type { contact } from '$lib/interfaces/contact'
	import { contactsStore } from '$lib/store'

	$: path = $page.url.pathname

	//toast
	const triggerToast = (message: string) => {
		const t: ToastSettings = {
			message: message,
			timeout: 10000,
			background: 'variant-filled-error'
		}
		toastStore.trigger(t)
	}

	//auth middleware
	const loginRoutes = ['/']
	const guestRoutes = ['/register', '/login']

	const authMiddleware = async () => {
		try {
			await axios.get(import.meta.env.VITE_BACKEND_DOMAIN + '/api/login/current', {
				withCredentials: true
			})
			if (guestRoutes.includes(path)) goto('/')
		} catch (err) {
			if (loginRoutes.includes(path)) goto('/login')
		}
	}

	//init
	interface getUserResponse extends apiResponse {
		data: user
	}
	const getUserData = async () => {
		try {
			const result: AxiosResponse<getUserResponse> = await axios.get(
				import.meta.env.VITE_BACKEND_DOMAIN + '/api/login/current',
				{
					withCredentials: true
				}
			)
			userStore.set(result.data.data)
		} catch (err) {
			if (loginRoutes.includes(path)) triggerToast('Failed to get user data')
		}
	}

	interface getUserGroupResponse extends apiResponse {
		data: group[]
	}
	const getUserGroup = async () => {
		try {
			const result: AxiosResponse<getUserGroupResponse> = await axios.get(
				import.meta.env.VITE_BACKEND_DOMAIN + '/api/login/groups',
				{
					withCredentials: true
				}
			)
			groupsStore.set(result.data.data)
		} catch (err) {
			if (loginRoutes.includes(path)) triggerToast('Faild to get user group')
		}
	}

	interface getUserContactsResponse extends apiResponse {
		data: contact[]
	}
	const getUserContacts = async () => {
		try {
			const result: AxiosResponse<getUserContactsResponse> = await axios.get(
				import.meta.env.VITE_BACKEND_DOMAIN + '/api/login/contacts',
				{
					withCredentials: true
				}
			)
			contactsStore.set(result.data.data)
		} catch (err) {
			if (loginRoutes.includes(path)) triggerToast('Faild to get user group')
		}
	}

	beforeNavigate(authMiddleware)
	onMount(async () => {
		await authMiddleware()
		await getUserData()
		await getUserGroup()
		await getUserContacts()
	})
</script>

<div class="bg-slate-100 min-h-screen w-screen">
	<PageTransition trigger={path}>
		<slot />
	</PageTransition>
	<Toast zIndex="z-[999]" />
</div>
