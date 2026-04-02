<script>
  export let opacity = 0.07;
</script>

<div class="ree-bg" style={`opacity:${opacity}`}>
  <svg
    viewBox="0 0 800 300"
    class="w-full h-full"
    preserveAspectRatio="xMidYMid slice"
  >
    <g class="ree-grid">

      <!-- Row 1 -->
      <g transform="translate(0,0)">
        {#each ["La","Ce","Pr","Nd","Pm","Sm","Eu"] as el, i}
          <g transform={`translate(${i * 110},0)`}>
            <rect width="100" height="80" rx="4" class="cell" />
            <text x="50" y="48" text-anchor="middle" class="symbol">{el}</text>
          </g>
        {/each}
      </g>

      <!-- Row 2 -->
      <g transform="translate(0,100)">
        {#each ["Gd","Tb","Dy","Ho","Er","Tm","Yb","Lu"] as el, i}
          <g transform={`translate(${i * 110},0)`}>
            <rect width="100" height="80" rx="4" class="cell" />
            <text x="50" y="48" text-anchor="middle" class="symbol">{el}</text>
          </g>
        {/each}
      </g>

      <!-- Y + Sc (offset slightly) -->
      <g transform="translate(250,200)">
        <g transform="translate(0,0)">
          <rect width="100" height="80" rx="4" class="cell" />
          <text x="50" y="48" text-anchor="middle" class="symbol">Y</text>
        </g>
        <g transform="translate(120,0)">
          <rect width="100" height="80" rx="4" class="cell" />
          <text x="50" y="48" text-anchor="middle" class="symbol">Sc</text>
        </g>
      </g>

    </g>
  </svg>
</div>

<style>
  .ree-bg {
    position: absolute;
    inset: 0;
    pointer-events: none;
    z-index: 0;
  }

  .cell {
    fill: transparent;
    stroke: currentColor;
    stroke-width: 1;
    opacity: 0.25;
  }

  .symbol {
    font-size: 18px;
    font-family: monospace;
    fill: currentColor;
    opacity: 0.30;
  }

  /* Light mode tuning */
  :global([data-theme="light"]) .cell {
    opacity: 0.50;
  }

  :global([data-theme="light"]) .symbol {
    opacity: 0.50;
  }
</style>