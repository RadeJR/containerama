<script lang="ts">
  import Ellipsis from "lucide-svelte/icons/ellipsis";
  import * as DropdownMenu from "$lib/components/ui/dropdown-menu";
  import { Button } from "$lib/components/ui/button";
  import type { AxiosResponse } from "axios";
  import { getAxios } from "$conf/axios";
  import { push } from "svelte-spa-router";

  export let id: string;
  export let state: string;
  export let updateTable: Function;

  async function stopContainer() {
    let result: AxiosResponse = await getAxios().put(
      `/api/containers/${id}/stop`,
    );
    if (result.status == 204) {
      updateTable();
    }
  }
  function editContainer(e: Event) {
    e.preventDefault();
    push(`/containers/${id}/edit`);
  }
  async function startContainer() {
    let result: AxiosResponse = await getAxios().put(
      `/api/containers/${id}/start`,
    );
    if (result.status == 204) {
      updateTable();
    }
  }
  async function removeContainer() {
    let result: AxiosResponse = await getAxios().delete(
      `/api/containers/${id}`,
    );
    if (result.status == 204) {
      updateTable();
    }
  }
  function showLogs(e: Event) {
    e.preventDefault();
    push(`/containers/${id}/logs`);
  }
</script>

<DropdownMenu.Root>
  <DropdownMenu.Trigger asChild let:builder>
    <Button
      variant="ghost"
      builders={[builder]}
      size="icon"
      class="relative h-8 w-8 p-0"
      on:click={(e) => {
        e.stopPropagation();
      }}
    >
      <span class="sr-only">Open menu</span>
      <Ellipsis class="h-4 w-4" />
    </Button>
  </DropdownMenu.Trigger>
  <DropdownMenu.Content>
    <DropdownMenu.Group>
      <DropdownMenu.Label>Actions</DropdownMenu.Label>
      <DropdownMenu.Item on:click={() => navigator.clipboard.writeText(id)}>
        Copy container ID
      </DropdownMenu.Item>
      {#if state == "running"}
        <DropdownMenu.Item on:click={stopContainer}>
          Stop container
        </DropdownMenu.Item>
      {:else}
        <DropdownMenu.Item on:click={startContainer}>
          Start container
        </DropdownMenu.Item>
        <DropdownMenu.Item on:click={removeContainer}>
          Remove container
        </DropdownMenu.Item>
      {/if}
      <DropdownMenu.Item on:click={editContainer}>
        Edit container
      </DropdownMenu.Item>
      <DropdownMenu.Item on:click={showLogs}>Show logs</DropdownMenu.Item>
    </DropdownMenu.Group>
  </DropdownMenu.Content>
</DropdownMenu.Root>
