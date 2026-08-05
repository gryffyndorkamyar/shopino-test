import { browser } from '$app/environment';
import { writable } from 'svelte/store';
import { fetchMe, login as apiLogin, register as apiRegister, type User } from '$lib/api/auth';

const ACCESS_KEY = 'shopino_access';
const REFRESH_KEY = 'shopino_refresh';

type AuthState = {
	user: User | null;
	access: string | null;
	ready: boolean;
};

function createAuthStore() {
	const { subscribe, update, set } = writable<AuthState>({
		user: null,
		access: null,
		ready: false
	});

	async function hydrate() {
		if (!browser) {
			set({ user: null, access: null, ready: true });
			return;
		}

		const access = localStorage.getItem(ACCESS_KEY);
		if (!access) {
			set({ user: null, access: null, ready: true });
			return;
		}

		try {
			const user = await fetchMe(access);
			set({ user, access, ready: true });
		} catch {
			localStorage.removeItem(ACCESS_KEY);
			localStorage.removeItem(REFRESH_KEY);
			set({ user: null, access: null, ready: true });
		}
	}

	async function login(username: string, password: string) {
		const tokens = await apiLogin(username, password);
		localStorage.setItem(ACCESS_KEY, tokens.access);
		localStorage.setItem(REFRESH_KEY, tokens.refresh);
		const user = await fetchMe(tokens.access);
		set({ user, access: tokens.access, ready: true });
		return user;
	}

	async function register(payload: {
		username: string;
		email: string;
		password: string;
		phone?: string;
	}) {
		await apiRegister(payload);
		return login(payload.username, payload.password);
	}

	function logout() {
		if (browser) {
			localStorage.removeItem(ACCESS_KEY);
			localStorage.removeItem(REFRESH_KEY);
		}
		update((s) => ({ ...s, user: null, access: null, ready: true }));
	}

	return { subscribe, hydrate, login, register, logout };
}

export const auth = createAuthStore();
