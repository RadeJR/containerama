<script lang="ts">
  import Ellipsis from "lucide-svelte/icons/ellipsis";
  import * as DropdownMenu from "$lib/components/ui/dropdown-menu";
  import { Button } from "$lib/components/ui/button";
    import type { AxiosResponse } from "axios";
    import { getAxios } from "$conf/axios";
 
  export let id: string;
  export let updateTable: Function;
  
  async function stopContainer() {
    let result: AxiosResponse = await getAxios().put(`/api/containers/${id}`)
    if (result.status == 200) {
      updateTable()
    }
  }
</script>
 
<DropdownMenu.Root>
  <DropdownMenu.Trigger asChild let:builder>
    <Button
      variant="ghost"
      builders={[builder]}
      size="icon"
      class="relative h-8 w-8 p-0"
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
      <DropdownMenu.Item on:click={stopContainer}>
        Stop container
      </DropdownMenu.Item>
      <DropdownMenu.Item on:click={() => navigator.clipboard.writeText(id)}>
        Remove container
      </DropdownMenu.Item>
    </DropdownMenu.Group>
    <DropdownMenu.Separator />
    <DropdownMenu.Item>View customer</DropdownMenu.Item>
    <DropdownMenu.Item>View payment details</DropdownMenu.Item>
  </DropdownMenu.Content>
</DropdownMenu.Root>
