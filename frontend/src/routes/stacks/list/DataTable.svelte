<script lang="ts">
  import {
    createTable,
    Render,
    Subscribe,
    createRender,
  } from "svelte-headless-table";
  import {
    addPagination,
    addSortBy,
    addTableFilter,
    addHiddenColumns,
    addSelectedRows,
  } from "svelte-headless-table/plugins";
  import ArrowUpDown from "lucide-svelte/icons/arrow-up-down";
  import * as Table from "$lib/components/ui/table";
  import { writable, type Writable } from "svelte/store";
  import { getAxios } from "$conf/axios";
  import type { AxiosResponse } from "axios";
  import { onMount } from "svelte";
  import DataTableActions from "./DataTableActions.svelte";
  import { Button } from "$lib/components/ui/button";
  import { Input } from "$lib/components/ui/input";
  import ChevronDown from "lucide-svelte/icons/chevron-down";
  import * as DropdownMenu from "$lib/components/ui/dropdown-menu";
  import DataTableCheckbox from "./DataTableCheckbox.svelte";
  import { push } from "svelte-spa-router";
  import { Stack } from "$app/types/stack";

  let stacks: Writable<Stack[]> = writable([]);

  async function getData() {
    try {
      let response: AxiosResponse = await getAxios().get("/api/stacks");
      $stacks = response.data;
    } catch (err) {
      console.log("error fetching stacks: " + err);
    }
  }
  onMount(() => {
    getData();
  });

  const table = createTable(stacks, {
    page: addPagination(),
    sort: addSortBy(),
    filter: addTableFilter({
      fn: ({ filterValue, value }) =>
        value.toLowerCase().includes(filterValue.toLowerCase()),
    }),
    hide: addHiddenColumns(),
    select: addSelectedRows(),
  });

  const compareNames = (leftValue: string[], rightValue: string[]): number => {
    if (leftValue[0].slice(1) > rightValue[0].slice(1)) {
      return 1;
    }
    return -1;
  };

  const columns = table.createColumns([
    table.column({
      accessor: "Id",
      header: (_, { pluginStates }) => {
        const { allPageRowsSelected } = pluginStates.select;
        return createRender(DataTableCheckbox, {
          checked: allPageRowsSelected,
        });
      },
      cell: ({ row }, { pluginStates }) => {
        const { getRowState } = pluginStates.select;
        const { isSelected } = getRowState(row);

        return createRender(DataTableCheckbox, {
          checked: isSelected,
        });
      },
      plugins: {
        sort: {
          disable: true,
        },
        filter: {
          exclude: true,
        },
      },
    }),
    table.column({
      accessor: "Name",
      header: "Name",
      plugins: {
        sort: {
          compareFn: compareNames,
        },
      },
    }),
    table.column({
      accessor: "PathToFile",
      header: "Path To File",
    }),
    table.column({
      accessor: ({ Id }) => Id,
      header: "",
      cell: ({ value }) => {
        return createRender(DataTableActions, {
          id: value,
        });
      },
      plugins: {
        sort: {
          disable: true,
        },
        filter: {
          exclude: true,
        },
      },
    }),
  ]);

  const {
    headerRows,
    pageRows,
    tableAttrs,
    tableBodyAttrs,
    pluginStates,
    flatColumns,
    rows,
  } = table.createViewModel(columns);

  const { hasNextPage, hasPreviousPage, pageIndex } = pluginStates.page;
  const { filterValue } = pluginStates.filter;
  const { hiddenColumnIds } = pluginStates.hide;
  const { selectedDataIds } = pluginStates.select;

  const ids = flatColumns.map((col) => col.id);
  let hideForId = Object.fromEntries(ids.map((id) => [id, true]));

  $: $hiddenColumnIds = Object.entries(hideForId)
    .filter(([, hide]) => !hide)
    .map(([id]) => id);

  const hidableCols = ["PathToFile"];
</script>

<div>
  <div class="flex items-center justify-between py-4">
    <Input
      class="max-w-sm"
      placeholder="Filter table..."
      type="text"
      bind:value={$filterValue}
    />
    <div>
      <DropdownMenu.Root>
        <DropdownMenu.Trigger asChild let:builder>
          <Button variant="outline" class="ml-auto" builders={[builder]}>
            Columns <ChevronDown class="ml-2 h-4 w-4" />
          </Button>
        </DropdownMenu.Trigger>
        <DropdownMenu.Content>
          {#each flatColumns as col}
            {#if hidableCols.includes(col.id)}
              <DropdownMenu.CheckboxItem bind:checked={hideForId[col.id]}>
                {col.header}
              </DropdownMenu.CheckboxItem>
            {/if}
          {/each}
        </DropdownMenu.Content>
      </DropdownMenu.Root>
      <Button
        variant="outline"
        on:click={() => {
          push("/stacks/create");
        }}>Create Stack</Button
      >
    </div>
  </div>
  <div class="rounded-md border">
    <Table.Root {...$tableAttrs}>
      <Table.Header>
        {#each $headerRows as headerRow}
          <Subscribe rowAttrs={headerRow.attrs()}>
            <Table.Row>
              {#each headerRow.cells as cell (cell.id)}
                <Subscribe
                  attrs={cell.attrs()}
                  let:attrs
                  props={cell.props()}
                  let:props
                >
                  <Table.Head {...attrs} class="[&:has([role=checkbox])]:pl-3">
                    {#if cell.id === "Names" || cell.id === "Image" || cell.id === "State"}
                      <Button variant="ghost" on:click={props.sort.toggle}>
                        <Render of={cell.render()} />
                        <ArrowUpDown class={"ml-2 h-4 w-4"} />
                      </Button>
                    {:else}
                      <Render of={cell.render()} />
                    {/if}
                  </Table.Head>
                </Subscribe>
              {/each}
            </Table.Row>
          </Subscribe>
        {/each}
      </Table.Header>
      <Table.Body {...$tableBodyAttrs}>
        {#each $pageRows as row (row.id)}
          <Subscribe rowAttrs={row.attrs()} let:rowAttrs>
            <Table.Row
              {...rowAttrs}
              data-state={$selectedDataIds[row.id] && "selected"}
              class="cursor-pointer"
            >
              {#each row.cells as cell (cell.id)}
                <Subscribe attrs={cell.attrs()} let:attrs>
                  <Table.Cell {...attrs}>
                    <Render of={cell.render()} />
                  </Table.Cell>
                </Subscribe>
              {/each}
            </Table.Row>
          </Subscribe>
        {/each}
      </Table.Body>
    </Table.Root>
  </div>
  <div class="flex items-center justify-end space-x-4 py-4">
    <div class="flex-1 text-sm text-muted-foreground">
      {Object.keys($selectedDataIds).length} of{" "}
      {$rows.length} row(s) selected.
    </div>
    <Button
      variant="outline"
      size="sm"
      on:click={() => ($pageIndex = $pageIndex - 1)}
      disabled={!$hasPreviousPage}>Previous</Button
    >
    <Button
      variant="outline"
      size="sm"
      disabled={!$hasNextPage}
      on:click={() => ($pageIndex = $pageIndex + 1)}>Next</Button
    >
  </div>
</div>
