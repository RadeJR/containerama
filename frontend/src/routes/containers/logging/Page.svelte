<script lang="ts">
  import { ScrollArea } from "$lib/components/ui/scroll-area/index.js";
  import { Separator } from "$lib/components/ui/separator/index.js";
  import { onDestroy, onMount } from "svelte";
  import { fade } from "svelte/transition";
  export let params: any = {};

  let logs: string[] = [];
  let source: EventSource;
  let scrarea = document.getElementById("scrollarea-viewport");
  let scrollToBottom: boolean = true

  onMount(() => {
    source = new EventSource(`/api/containers/${params.id}/logs`);
    source.onmessage = function (event) {
      logs = [...logs, event.data];
      if (scrarea && scrollToBottom) {
        scrarea.scrollTop = scrarea.scrollHeight + 100;
      }
    };
  });

  onDestroy(() => {
    source.close();
  });
</script>

<div transition:fade={{ duration: 100 }}>
  <ScrollArea class="h-[850px] rounded-md border">
    <div class="p-4">
      <h4 class="mb-4 text-sm font-medium leading-none">Logs</h4>
      {#each logs as log}
        <div class="text-sm">
          {#if log.slice(0, 3) == "OUT"}
            <span class="p-1 mr-1 text-green-500">{log.slice(0, 3)}</span>
          {:else}
            <span class="p-1 mr-1 text-red-500">{log.slice(0, 3)}</span>
          {/if}
          {log.slice(7)}
        </div>
        <Separator class="my-2" />
      {/each}
    </div>
  </ScrollArea>
</div>
