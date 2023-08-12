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
	import { userStore } from '$lib/components/store/user'
	import { groupsStore } from '$lib/components/store/group'
	import { websocketStore } from '$lib/components/store/websocket'
	import { contactsStore } from '$lib/components/store/contact'
	import { previewChatStore } from '$lib/components/store/previewChat'
	import type { AxiosResponse } from 'axios'
	import type { apiResponse } from '$lib/interfaces/apiResponse'
	import type { user } from '$lib/interfaces/user'
	import type { group } from '$lib/interfaces/group'
	import { Toast, toastStore, type ToastSettings } from '@skeletonlabs/skeleton'
	import type { contact } from '$lib/interfaces/contact'

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
			await axios.get(`http://${import.meta.env.VITE_BACKEND_DOMAIN}/api/login/current`, {
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
				`http://${import.meta.env.VITE_BACKEND_DOMAIN}/api/login/current`,
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
				`http://${import.meta.env.VITE_BACKEND_DOMAIN}/api/login/groups`,
				{
					withCredentials: true
				}
			)
			if (result.data.data) groupsStore.set(result.data.data)
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
				`http://${import.meta.env.VITE_BACKEND_DOMAIN}/api/login/contacts`,
				{
					withCredentials: true
				}
			)
			if (result.data.data) contactsStore.set(result.data.data)
		} catch (err) {
			if (loginRoutes.includes(path)) triggerToast('Faild to get user contact')
		}
	}

	interface getUserChatPreviewResponse extends apiResponse {
		data: {
			group: {
				group_id: number
				group_name: string
				chat_id: number
				sender_email: string
				message: string
				created_at: string
			}[]
			personal: {
				email: string
				chat_id: number
				sender_email: string
				message: string
				created_at: string
			}[]
		}
	}
	const getUserChatPreview = async () => {
		try {
			const result: AxiosResponse<getUserChatPreviewResponse> = await axios.get(
				`http://${import.meta.env.VITE_BACKEND_DOMAIN}/api/login/chats/preview`,
				{
					withCredentials: true
				}
			)
			if (result.data.data) {
				let previewChats
				const groups = result.data.data.group.map((group) => {
					return {
						groupId: group.group_id,
						groupName: group.group_name,
						senderEmail: group.sender_email,
						message: group.message,
						chatId: group.chat_id,
						createdAt: new Date(group.created_at),
						email: null
					}
				})
				previewChats = groups
				const personals = result.data.data.personal.map((personal) => {
					return {
						email: personal.email,
						message: personal.message,
						senderEmail: personal.sender_email,
						chatId: personal.chat_id,
						createdAt: new Date(personal.created_at),
						groupId: null,
						groupName: null
					}
				})
				previewChats = [...previewChats, ...personals]
				previewChats.sort((prev, curr) => curr.createdAt.getTime() - prev.createdAt.getTime())
				previewChatStore.set(previewChats)
			}
		} catch (err) {
			if (loginRoutes.includes(path)) triggerToast('Faild to get preview chats')
			console.log(err)
		}
	}

	//connect to websocket
	$: {
		if ($userStore) {
			const websocketConn = new WebSocket(
				`ws://${import.meta.env.VITE_BACKEND_DOMAIN}/api/login/ws`
			)
			websocketStore.set(websocketConn)
		} else {
			websocketStore.set(null)
		}
	}

	beforeNavigate(authMiddleware)
	onMount(async () => {
		await authMiddleware()
		await getUserData()
		await getUserGroup()
		await getUserContacts()
		await getUserChatPreview()
	})
</script>

<div class="bg-slate-100 min-h-screen w-screen">
	<PageTransition trigger={path} amount={-50}>
		<slot />
	</PageTransition>
	<Toast zIndex="z-[999]" />
</div>
