<script lang="ts">
  // packages
  import Router from "svelte-spa-router";
  import { ModeWatcher } from "mode-watcher";

  // local
  import Login from "$lib/components/Login.svelte";
  import Base from "$lib/components/Base.svelte";
  import { routes } from "$app/routes";
  import { cookieExists } from "$services/auth";
  import { isAuthorized } from "$store";

  if (cookieExists("session")) {
    isAuthorized.set(true);
  }
</script>

<ModeWatcher />
{#if $isAuthorized}
  <Base>
    <Router {routes} />
  </Base>
{:else}
  <Login />
{/if}
