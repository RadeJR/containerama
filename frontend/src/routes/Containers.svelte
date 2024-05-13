<script lang="ts">
  import { getAxios } from "$conf/axios";
  import type { AxiosResponse } from "axios";
  import { onMount } from "svelte";

  type Container = {
    Id: string;
    Names: string[];
    Image: string;
    Created: string;
    Ports: Port[];
    State: string;
  };

  type Port = {
    IP: string;
    PrivatePort: number;
    PublicPort: number;
    Type: string;
  };

  let containers: Container[] = [];

  async function getData() {
    try {
      let response: AxiosResponse = await getAxios().get("/api/containers");
      containers = response.data;
    } catch (err) {}
  }
  onMount(() => {
    getData();
  });
</script>

{#each containers as d}
  <p>{d.Id}</p>
{/each}
