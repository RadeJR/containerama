<script lang="ts">
  import { ScrollArea } from "$lib/components/ui/scroll-area/index.js";
  import { Switch } from "$lib/components/ui/switch";
  import { Separator } from "$lib/components/ui/separator/index.js";
  import { onDestroy, onMount } from "svelte";
  import { fade } from "svelte/transition";
  import Label from "$lib/components/ui/label/label.svelte";
  export let params: any = {};

  let logs: string[] = [];
  let source: EventSource;
  let scrarea: HTMLElement | null;
  let scrollToBottom: boolean = true;

  onMount(() => {
    source = new EventSource(`/api/containers/${params.id}/logs`);
    source.onmessage = function (event) {
      logs = [...logs, event.data];
      scrollToBottomFn();
    };
  });

  onDestroy(() => {
    source.close();
  });

  function scrollToBottomFn() {
    scrarea = document.getElementById("scrollarea-viewport");
    if (scrarea && scrollToBottom) {
      scrarea.scrollTop = scrarea.scrollHeight + 100;
    }
  }
</script>

<div transition:fade={{ duration: 100 }}>
  <div class="flex justify-between py-4">
    <h4 class="text-lg font-medium">Logs</h4>
    <div>
      <Switch id="scroll-switch" bind:checked={scrollToBottom} />
      <Label for="scroll-switch" class="text-lg pl-2 font-medium">Scroll to bottom</Label>
    </div>
  </div>
  <ScrollArea class="h-[650px] rounded-md border">
    <div class="p-4">
      {#each logs as log}
        <div class="text-sm">
          {#if log.slice(0, 3) == "OUT"}
            <span class="p-1 mr-1 text-green-500">{log.slice(0, 3)}</span>
          {:else}
            <span class="p-1 mr-1 text-red-500">{log.slice(0, 3)}</span>
          {/if}
          {log.slice(3)}
        </div>
        <Separator class="my-2" />
      {/each}
    </div>
  </ScrollArea>
</div>
