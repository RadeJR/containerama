<script lang="ts">
  import Router from "svelte-spa-router";
  import { routes } from "./routes.js";
  import { cookieExists } from "./services/auth";
  import { isAuthorized } from "./store";
  import Login from "./lib/components/Login.svelte";
  import { ModeWatcher } from "mode-watcher";
  import { getAxios } from "$conf/axios.js";
  import Base from "$lib/components/Base.svelte";

  if (cookieExists("session")) {
    console.log("Cookie exists!");
    isAuthorized.set(true);
  }

  getAxios().interceptors.response.use(null, function (error) {
    if (error.response.status == 401) {
      console.log("Unauthorized, setting to false");
      isAuthorized.set(false);
    }
    return Promise.reject(error);
  });
</script>

<ModeWatcher />
{#if $isAuthorized}
  <!-- <Router {routes} /> -->
  <Base>
    <Router {routes} />
  </Base>
{:else}
  <Login />
{/if}
