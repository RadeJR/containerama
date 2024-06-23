<script lang="ts">
  // packages
  import Router from "svelte-spa-router";
  import { ModeWatcher } from "mode-watcher";

  // local
  import Login from "$lib/components/Login.svelte";
  import Base from "$lib/components/Base.svelte";
  import { routes } from "$app/routes";
  import { isAuthorized } from "$store";
  import { onMount } from "svelte";
  import { getAxios } from "$conf/axios";
  import { User } from "$app/types/user";
  let user: User;

  async function checkIfLoggedIn() {
    try {
      await getAxios()
        .get("/api/userinfo")
        .then((response) => {
          user = response.data;
          isAuthorized.set(true);
        });
    } catch (err) {
      console.log("error fetching user data: " + err);
    }
  }

  onMount(() => {
    checkIfLoggedIn();
  });
</script>

<ModeWatcher />
{#if $isAuthorized}
  <Base {user}>
    <Router {routes} />
  </Base>
{:else}
  <Login />
{/if}
