<script>
    import { theme } from '$lib/stores/theme.svelte';
    import { auth, logout } from '$lib/stores/auth';
    import { page } from '$app/stores';
</script>

<nav class="border-b border-dark-surface/10 dark:border-light-bg/10 py-4 px-6 flex justify-between items-center bg-white dark:bg-dark-surface">
    <div class="flex items-center gap-8">
        <a href="/" class="text-xl font-bold tracking-tighter text-primary">
            REE<span class="text-dark-bg dark:text-light-bg">ISTO</span>
        </a>
        
        {#if $auth.isAuthenticated}
            <div class="hidden md:flex gap-6 text-sm font-medium">
                <a href="/market" class="hover:text-primary transition-colors">Market</a>
                <a href="/inventory" class="hover:text-primary transition-colors">Inventory</a>
                {#if $auth.user?.role === 'admin'}
                    <a href="/admin" class="hover:text-primary transition-colors italic">Control Panel</a>
                {/if}
            </div>
        {/if}
    </div>

    <div class="flex items-center gap-4">
        {#if $auth.isAuthenticated}
            <span class="text-xs opacity-60 uppercase">{$auth.user?.username}</span>
            <button 
                onclick={logout}
                class="text-xs px-3 py-1 border border-primary text-primary hover:bg-primary hover:text-white transition-all rounded"
            >
                Logout
            </button>
        {/if}
    </div>
</nav>