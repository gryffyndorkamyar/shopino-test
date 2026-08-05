export type AuthTokens = {
	access: string;
	refresh: string;
};

export type User = {
	id: number;
	username: string;
	email: string;
	phone?: string;
};

export type RegisterPayload = {
	username: string;
	email: string;
	password: string;
	phone?: string;
};

async function readError(res: Response): Promise<string> {
	try {
		const data = await res.json();
		if (typeof data?.detail === 'string') return data.detail;
		if (Array.isArray(data?.detail)) {
			return data.detail.map((d: { msg?: string }) => d.msg ?? JSON.stringify(d)).join('، ');
		}
		const first = Object.values(data as Record<string, unknown>)[0];
		if (Array.isArray(first)) return String(first[0]);
		if (typeof first === 'string') return first;
	} catch {
		/* ignore */
	}
	return `خطای احراز هویت (${res.status})`;
}

export async function login(username: string, password: string): Promise<AuthTokens> {
	const res = await fetch('/api/auth/login/', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ username, password })
	});
	if (!res.ok) throw new Error(await readError(res));
	return res.json();
}

export async function register(payload: RegisterPayload): Promise<User> {
	const res = await fetch('/api/auth/register/', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify(payload)
	});
	if (!res.ok) throw new Error(await readError(res));
	return res.json();
}

export async function fetchMe(accessToken: string): Promise<User> {
	const res = await fetch('/api/auth/me/', {
		headers: { Authorization: `Bearer ${accessToken}` }
	});
	if (!res.ok) throw new Error(await readError(res));
	return res.json();
}
