<script>
  export let headers;
</script>

<style>
  table {
    border-collapse: collapse;
    width: 100%;
  }
  th {
    background-color: #f2f2f2;
  }

  th {
    text-align: left;
    padding: 8px;
    border-bottom: 1px black solid;
  }

  tr:nth-child(even) {
    background-color: #f2f2f2;
  }

  table :global(tr:hover) {
    background-color:grey!important;
  }

  table :global(tr:nth-child(even)) {
    background-color: #f2f2f2;
  }

  table :global(td) {
    border-bottom: 1px solid #ddd;
  }
</style>

<table>
  <thead>
    <tr>
      {#each headers as header}
        {#if typeof header == 'object'}
          <th colspan={header.col}>{header.name}</th>
        {:else}
          <th>{header}</th>
        {/if}
      {/each}
    </tr>
  </thead>
  <tbody>
    <slot />
  </tbody>

</table>
