<script>
  import { login } from "$lib/stores/auth";
  import { apiFetch } from "$lib/api";
  import { goto } from "$app/navigation";
  import { onMount, onDestroy } from "svelte";

  let username = $state("");
  let password = $state("");
  let loading = $state(false);
  let errorMsg = $state("");

  async function handleSubmit(e) {
    e.preventDefault();
    errorMsg = "";
    loading = true;
    try {
      const data = await apiFetch("/login", "POST", { username, password });
      login(data.token, data.user);
      goto("/market");
    } catch (err) {
      errorMsg = err.message;
    } finally {
      loading = false;
    }
  }

</script>

<!-- Key changes: h-screen + flex + no min-height -->
<div class="flex items-center justify-center mt-24 md:mt-32 p-4">
  <div
    class="w-full max-w-md bg-white dark:bg-dark-surface p-8 rounded-lg shadow-xl border border-dark-surface/5"
  >
    <h2 class="text-2xl font-bold mb-1 text-dark-bg dark:text-light-bg">
      Terminal Login
    </h2>
    <p class="text-sm text-primary mb-8 font-mono">TRADERS ONLY</p>

    <form onsubmit={handleSubmit} class="space-y-6">
      <div>
        <label
          class="block text-xs font-bold uppercase mb-2 text-dark-bg dark:text-light-bg opacity-70"
          for="u"
        >
          Identifier
        </label>
        <input
          id="u"
          type="text"
          bind:value={username}
          class="w-full p-3 text-dark-bg dark:text-light-bg border-none rounded focus:ring-2 ring-primary outline-none"
          placeholder="Username"
          required
        />
      </div>

      <div>
        <label
          class="block text-xs font-bold uppercase mb-2 text-dark-bg dark:text-light-bg opacity-70"
          for="p"
        >
          Access Key
        </label>
        <input
          id="p"
          type="password"
          bind:value={password}
          class="w-full p-3 text-dark-bg dark:text-light-bg border-none rounded focus:ring-2 ring-primary outline-none"
          placeholder="••••••••"
          required
        />
      </div>

      {#if errorMsg}
        <div
          class="p-3 bg-red-500/10 border border-red-500 text-red-500 text-xs rounded"
        >
          {errorMsg}
        </div>
      {/if}

      <button
        type="submit"
        disabled={loading}
        class="w-full py-3 bg-primary text-dark-bg font-bold uppercase tracking-widest hover:brightness-110 transition-all disabled:opacity-50"
      >
        {loading ? "Verifying..." : "Establish Connection"}
      </button>
    </form>
  </div>
</div>
