<script>
  import REEBackground from "$lib/components/REEBackground.svelte";
  import {siteConfig} from "$lib/config";
</script>

<svelte:head>
  <title>Privacy Policy — REEISTO</title>
  <meta
    name="description"
    content="Privacy practices for REEISTO, an academic inventory simulation platform."
  />
  <link rel="canonical" href={`${siteConfig.siteUrl}/privacy`} />
</svelte:head>

<div class="min-h-screen py-16 px-4">
  <REEBackground opacity={0.05} />
  <div class="max-w-3xl mx-auto space-y-10">
    <h1 class="text-3xl font-mono font-semibold">Privacy Policy</h1>

    <section class="space-y-4 text-sm opacity-80">
      <p>
        REEISTO does not collect or process personal data for commercial
        purposes. The platform is designed as a closed academic simulation.
      </p>
    </section>

    <section class="space-y-4 text-sm opacity-80">
      <h2 class="section-title">Local Storage</h2>
      <p>
        Authentication tokens and user session data are stored locally in the
        browser using localStorage. This data is not transmitted to third-party
        services.
      </p>
    </section>

    <section class="space-y-4 text-sm opacity-80">
      <h2 class="section-title">No Tracking</h2>
      <p>
        REEISTO does not use analytics, tracking scripts, or cookies for user
        profiling.
      </p>
    </section>

    <section class="space-y-4 text-sm opacity-80">
      <h2 class="section-title">Data Nature</h2>
      <p>
        All data within the system is synthetic and intended solely for
        demonstration.
      </p>
    </section>
  </div>
</div>
