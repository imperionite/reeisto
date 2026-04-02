<script>
  import REEBackground from "$lib/components/REEBackground.svelte";
  import {siteConfig} from "$lib/config";
</script>

<svelte:head>
  <title>Terms of Service — REEISTO</title>
  <meta
    name="description"
    content="Terms governing the use of REEISTO, an academic inventory simulation platform."
  />
  <link rel="canonical" href={`${siteConfig.siteUrl}/terms`} />
</svelte:head>

<div class="min-h-screen py-16 px-4">
  <REEBackground opacity={0.05} />
  <div class="max-w-3xl mx-auto space-y-10">

    <h1 class="text-3xl font-mono font-semibold">Terms of Service</h1>

    <section class="space-y-4 text-sm opacity-80">
      <p>
        REEISTO is an academic proof-of-concept system. By accessing this platform,
        you acknowledge that all functionality is simulated and does not represent
        real-world trading or financial activity.
      </p>
    </section>

    <section class="space-y-4 text-sm opacity-80">
      <h2 class="section-title">No Financial Transactions</h2>
      <p>
        The system does not process payments, execute trades, or facilitate
        financial agreements. All recorded transactions represent simulated
        inventory movements only.
      </p>
    </section>

    <section class="space-y-4 text-sm opacity-80">
      <h2 class="section-title">Fictitious Data</h2>
      <p>
        All users, organizations, and data within REEISTO are fictional and
        created solely for demonstration purposes.
      </p>
    </section>

    <section class="space-y-4 text-sm opacity-80">
      <h2 class="section-title">Limitation of Liability</h2>
      <p>
        The developers of REEISTO are not responsible for any misuse,
        misinterpretation, or reliance on the system. The platform is provided
        "as-is" without guarantees of accuracy or reliability.
      </p>
    </section>

    <section class="space-y-4 text-sm opacity-80">
      <h2 class="section-title">Non-Production Use</h2>
      <p>
        This system is not intended for commercial deployment or operational use
        in real trading environments.
      </p>
    </section>

  </div>
</div>