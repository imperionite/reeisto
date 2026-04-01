<svelte:head>
  <title>About — REEISTO</title>
  <meta
    name="description"
    content="REEISTO is an academic proof-of-concept inventory system simulating rare earth material logistics and trading operations."
  />
</svelte:head>

<div class="min-h-screen py-16 px-4">
  <div class="max-w-4xl mx-auto space-y-14">

    <!-- HEADER -->
    <section class="text-center border-b border-current/10 pb-10">
      <h1 class="text-4xl md:text-5xl font-mono font-semibold tracking-tight">
        REEISTO
      </h1>

      <p class="mt-4 text-lg opacity-80">
        Rare Earth Elements Inventory System for Trading Operations
      </p>
    </section>

    <!-- INTRO -->
    <section class="space-y-6">
      <div class="panel p-8">
        <p class="leading-relaxed text-base opacity-90">
          <strong class="text-primary">REEISTO</strong> is an internal operational platform designed to simulate inventory tracking and transaction workflows within a rare earth materials trading environment.
        </p>

        <p class="leading-relaxed mt-4 text-base opacity-80">
          Traders interact with inventory distributed across multiple warehouse locations. Buy and sell actions are recorded as material movements—allocation, acquisition, or release—without financial execution.
        </p>
      </div>
    </section>

    <!-- CORE -->
    <section class="space-y-6">
      <h2 class="section-title">Core Functionality</h2>

      <div class="grid md:grid-cols-2 gap-6">
        <div class="panel p-6">
          <p class="label mb-3">Scope</p>
          <ul class="space-y-2 text-sm opacity-85">
            <li>Multi-location inventory tracking</li>
            <li>Material allocation and release</li>
            <li>Transaction logging</li>
          </ul>
        </div>

        <div class="panel p-6">
          <p class="label mb-3">Excluded</p>
          <ul class="space-y-2 text-sm opacity-85">
            <li>Payment processing</li>
            <li>Order matching</li>
            <li>Market execution</li>
          </ul>
        </div>
      </div>

      <p class="text-sm opacity-70 max-w-2xl">
        The system reflects industrial workflows where pricing and settlement occur outside operational platforms. REEISTO focuses strictly on logistics and inventory state.
      </p>
    </section>

    <!-- CONTEXT -->
    <section class="space-y-6">
      <h2 class="section-title">Industry Context</h2>

      <div class="panel p-8 space-y-6">
        <ul class="grid md:grid-cols-3 gap-4 text-sm opacity-85">
          <li class="card">Inventory buffering by traders</li>
          <li class="card">Distributed warehouse storage</li>
          <li class="card">Contract-based allocation</li>
        </ul>

        <p class="text-sm opacity-70">
          These patterns are typical in rare earth and specialty materials supply chains, particularly in B2B and B2G environments.
        </p>
      </div>
    </section>

    <!-- ACADEMIC -->
    <section class="space-y-6">
      <h2 class="section-title">Academic Scope</h2>

      <div class="grid md:grid-cols-2 gap-6">
        <div class="panel p-6">
          <p class="label mb-3">Nature</p>
          <ul class="text-sm space-y-2 opacity-85">
            <li>Fictitious data and users</li>
            <li>Simulation-only environment</li>
            <li>No real-world execution</li>
          </ul>
        </div>

        <div class="panel p-6">
          <p class="label mb-3">Focus</p>
          <ul class="text-sm space-y-2 opacity-85">
            <li>Service-layer design</li>
            <li>State consistency</li>
            <li>API-driven frontend</li>
          </ul>
        </div>
      </div>

      <p class="text-center text-sm opacity-60 border-t border-current/10 pt-6">
        This system is an academic proof-of-concept and should not be used for real trading or financial operations.
      </p>
    </section>

  </div>
</div>

<style>
  .panel {
    background: rgba(var(--rgb-dark-surface), 0.25);
    border: 1px solid rgba(255, 255, 255, 0.05);
  }

  [data-theme="light"] .panel {
    background: rgba(var(--rgb-light-surface), 0.7);
    border: 1px solid rgba(0, 0, 0, 0.06);
  }

  .section-title {
    font-size: 1.4rem;
    font-weight: 600;
    letter-spacing: 0.04em;
  }

  .label {
    font-size: 0.7rem;
    text-transform: uppercase;
    letter-spacing: 0.12em;
    opacity: 0.6;
  }

  .card {
    padding: 1rem;
    border: 1px solid rgba(255,255,255,0.04);
    background: rgba(255,255,255,0.02);
  }

  [data-theme="light"] .card {
    border: 1px solid rgba(0,0,0,0.05);
    background: rgba(0,0,0,0.02);
  }
</style>