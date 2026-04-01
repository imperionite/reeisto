<script>
  import { theme } from "$lib/stores/theme.svelte";
</script>

<footer
  class="mt-20 py-10 border-t border-dark-surface/10 dark:border-light-bg/10"
>
  <div
    class="container mx-auto px-6 flex flex-col md:flex-row justify-between items-center gap-6"
  >
    <div class="text-xs opacity-50 font-mono flex flex-col items-start gap-1">
      <span
        >REETIS Industrial Materials Division // {new Date().getFullYear()}.{String(
          new Date().getMonth() + 1,
        ).padStart(2, "0")}</span
      >
      <span
        class="opacity-40 text-[10px] font-light tracking-wider leading-tight"
      >
        Academic POC & Simulation Project. No real-world trading or funds use.
      </span>
    </div>

    <div class="flex items-center gap-4">
      <!-- Theme Toggle -->
      <button
        onclick={() => theme.toggle()}
        class="flex items-center gap-3 px-4 py-2 rounded-full bg-dark-surface/5 dark:bg-light-bg/5 hover:bg-primary/10 transition-colors border border-transparent hover:border-primary/30"
        title="Toggle Theme"
      >
        <span class="text-sm font-medium">
          {theme.current === "light" ? "Daylight" : "Night Ops"}
        </span>
        <div
          class="w-10 h-5 bg-dark-bg dark:bg-primary rounded-full relative transition-colors"
        >
          <div
            class={`absolute top-1 w-3 h-3 rounded-full bg-white transition-all ${theme.current === "dark" ? "left-6" : "left-1"}`}
          ></div>
        </div>
      </button>

      <!-- GitHub Link -->
      <a
        href="https://github.com/YOUR_USERNAME/reetis-poc"
        target="_blank"
        rel="noopener noreferrer"
        class="p-2 hover:bg-dark-surface/20 dark:hover:bg-light-bg/20 rounded-lg transition-all group"
        title="View Source on GitHub"
      >
        <svg
          viewBox="0 0 16 16"
          class="w-5 h-5 opacity-70 group-hover:opacity-100 transition-all"
          fill="none"
          aria-label="GitHub"
        >
          <path
            d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.72.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.012 8.012 0 0 0 16 8c0-4.42-3.58-8-8-8z"
          />
        </svg>
      </a>
    </div>
  </div>
</footer>

<style>
  /* Direct CSS custom property access */
  svg path {
    fill: var(--color-dark-bg);
  }

  [data-theme="dark"] svg path {
    fill: var(--color-primary) !important;
  }
</style>
