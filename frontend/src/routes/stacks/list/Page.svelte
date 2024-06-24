<script lang="ts">
  import DataTable from "./DataTable.svelte";
  import {fade} from 'svelte/transition'
</script>

<div transition:fade={{ duration: 100 }}>
  <DataTable />
</div>
