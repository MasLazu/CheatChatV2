import type { Handle } from '@sveltejs/kit'
// import axios from 'axios'
// import { redirect } from '@sveltejs/kit'
// import { browser } from '$app/environment'

export const handle: Handle = async ({ event, resolve }) => {
	//auth middleware server
	// const session = event.cookies.get('session')
	// console.log(session)
	// const path = event.url.pathname
	// const loginRoutes = ['/']
	// const guestRoutes = ['/register', '/login']

	// try {
	// 	await axios.get(import.meta.env.VITE_BACKEND_DOMAIN + '/api/session/' + session)
	// 	if (guestRoutes.includes(path)) {
	// 		if (!browser) return new Response('Redirect', { status: 303, headers: { Location: '/' } })
	// 		throw redirect(303, '/')
	// 	}
	// } catch (err: any) {
	// 	if (loginRoutes.includes(path)) {
	// 		if (!browser)
	// 			return new Response('Redirect', { status: 303, headers: { Location: '/login' } })
	// 		throw redirect(303, '/login')
	// 	}
	// }

	const response = await resolve(event)
	return response
}
