<script>
  import "../app.css";
  import Nav from "$lib/components/Nav.svelte";
  import Footer from "$lib/components/Footer.svelte";
  import { theme } from '$lib/stores/theme.svelte';
  import { auth } from "$lib/stores/auth";
  import { onMount } from "svelte";
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";

  // REQUIRED in runes mode
  let { children } = $props();

  onMount(() => {
    theme.init();

    if (!$auth.isAuthenticated && $page.url.pathname !== "/login") {
      goto("/login");
    }
  });
</script>

<div
  class="min-h-screen flex flex-col bg-[var(--bg-dark)] text-[var(--text-primary)]"
>
  <Nav />

  <main class="grow max-w-7xl mx-auto w-full px-6 py-8">
    {@render children()}
  </main>

  <Footer />
</div>
