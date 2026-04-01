import { writable } from 'svelte/store';
import { browser } from '$app/environment';

// Initialize from localStorage if available
const storedToken = browser ? localStorage.getItem('reetis_token') : null;
const storedUser = browser ? JSON.parse(localStorage.getItem('reetis_user') || 'null') : null;

export const auth = writable({
    token: storedToken,
    user: storedUser,
    isAuthenticated: !!storedToken
});

export const login = (token, user) => {
    localStorage.setItem('reetis_token', token);
    localStorage.setItem('reetis_user', JSON.stringify(user));
    auth.set({ token, user, isAuthenticated: true });
};

export const logout = () => {
    localStorage.removeItem('reetis_token');
    localStorage.removeItem('reetis_user');
    auth.set({ token: null, user: null, isAuthenticated: false });
    window.location.href = '/login';
};