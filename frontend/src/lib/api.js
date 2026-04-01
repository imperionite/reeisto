import { browser } from '$app/environment';

const API_BASE = import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1';

export async function apiFetch(endpoint, method = 'GET', body = null) {
    const token = browser ? localStorage.getItem('reetis_token') : null;
    
    const headers = {
        'Content-Type': 'application/json',
    };

    if (token) {
        headers['Authorization'] = `Bearer ${token}`;
    }

    const options = {
        method,
        headers,
    };

    if (body) {
        options.body = JSON.stringify(body);
    }

    const response = await fetch(`${API_BASE}${endpoint}`, options);
    const result = await response.json();

    if (!response.ok) {
        throw new Error(result.error || 'Something went wrong');
    }

    return result.data; // Matches your utils.Success(c, data) format
}