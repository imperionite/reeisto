<script>
  import REEBackground from "$lib/components/REEBackground.svelte";
  export let error;
  export let status;
</script>

<svelte:head>
  <title>404 — REEISTO</title>
</svelte:head>

<div class="min-h-screen flex items-center justify-center px-4">
  <REEBackground opacity={0.05} />
  <div class="text-center space-y-6 max-w-md">

    <div class="text-5xl font-mono font-semibold tracking-widest opacity-80">
      {status}
    </div>

    <div class="text-lg opacity-70">
      Resource not found
    </div>

    <p class="text-sm opacity-50">
      The requested endpoint or page does not exist within this system.
    </p>

    <a
      href="/"
      class="inline-block mt-4 px-6 py-2 border border-current/20 hover:border-primary transition"
    >
      Return to system
    </a>

  </div>
</div>