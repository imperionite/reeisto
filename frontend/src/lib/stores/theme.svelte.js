import { browser } from '$app/environment';

class ThemeManager {
    current = $state(browser ? (localStorage.getItem('reetis_theme') || 'light') : 'light');

    constructor() {
        $effect.root(() => {
            $effect(() => {
                if (browser) {
                    localStorage.setItem('reetis_theme', this.current);
                    document.documentElement.setAttribute('data-theme', this.current);
                    // Add/remove dark class for Tailwind's dark: modifier
                    if (this.current === 'dark') {
                        document.documentElement.classList.add('dark');
                    } else {
                        document.documentElement.classList.remove('dark');
                    }
                }
            });
        });
    }

    toggle() {
        this.current = this.current === 'light' ? 'dark' : 'light';
    }
}

export const theme = new ThemeManager();