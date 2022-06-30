import { writable } from 'svelte/store'

export const isAuthenticated = writable(undefined)
export const user = writable({})
export const popupOpen = writable(false)
export const error = writable()