<script>
  import "../app.css";
  import Nav from "$lib/components/Nav.svelte";
  import Footer from "$lib/components/Footer.svelte";
  import { auth } from "$lib/stores/auth";
  import { onMount } from "svelte";
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";

  import { siteConfig } from "$lib/config";

  // REQUIRED in runes mode
  let { children, data } = $props();

  const BASE_URL = siteConfig.siteUrl;

  onMount(() => {
    if (!$auth.isAuthenticated && $page.url.pathname !== "/login") {
      goto("/login");
    }
  });
</script>

<svelte:head>
  <title>{data.meta?.title ?? "REEISTO"}</title>
  <meta name="description" content={data.meta.description} />
  <link rel="canonical" href={`${BASE_URL}${data.meta.path}`} />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <meta name="robots" content="index, follow" />
  <meta
    name="keywords"
    content="academic project, pre-trade, inventory tracking, software development, non-commercial, educational, demonstration"
  />
  <meta name="author" content={siteConfig?.siteName} />

  <meta name="application-name" content={siteConfig?.siteName} />
  <meta
    name="category"
    content="Academic Project, Frontend Software Development Simulation"
  />
  <meta
    name="application-purpose"
    content="This website is a non-commercial academic project intended for demonstration and educational use only. It does not provide real inventory of REE or financial transactions."
  />
  <!-- Favicon -->
  <link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon.png" />
  <link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png" />
  <link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png" />
  <link rel="shortcut icon" href="/favicon.ico" />
  <link rel="manifest" href="/site.webmanifest" />

  <!-- OpenGraph -->
  <meta property="og:title" content={data.meta.title} />
  <meta property="og:description" content={data.meta.description} />
  <meta property="og:url" content={`${BASE_URL}${data.meta.path}`} />
  <meta property="og:type" content="website" />
  <meta property="og:image" content={siteConfig?.siteUrl + "/og-image.png"} />
  <meta name="classification" content="Academic Project" />

  <!-- Twitter -->
  <meta name="twitter:card" content="summary_large_image" />
  <meta name="twitter:title" content={data.meta.title} />
  <meta name="twitter:description" content={data.meta.description} />
  <meta name="twitter:description" content={data.meta.description} />
  <meta name="twitter:image" content={siteConfig?.siteUrl + "/og-image.png"} />
</svelte:head>

<div class="min-h-screen flex flex-col bg-(--bg-dark) text-(--text-primary)">
  <div
    class="sticky top-0 z-50 bg-dark-surface/95 backdrop-blur-md border-b border-current/10"
  >
    <Nav />
  </div>

  <main class="grow pt-16 max-w-7xl mx-auto w-full px-6 pb-8">
    <!-- pt-16 = ~64px nav height -->
    {@render children()}
  </main>

  <Footer />
</div>
