<script lang="ts">
  import axios from "axios";
  import Router, { push } from "svelte-spa-router";
  import Login from "./routes/Login.svelte";
  import { routes } from "./routes";
  import { isAuthorized } from "./store";

  function cookieExists(name: string): boolean {
    // Split document.cookie string into individual cookies
    const cookies: string[] = document.cookie.split(";");

    // Iterate over cookies to find the one with the specified name
    for (let i = 0; i < cookies.length; i++) {
      const cookie: string = cookies[i].trim();

      // Check if cookie starts with the specified name
      if (cookie.startsWith(name + "=")) {
        return true; // Cookie exists
      }
    }

    return false; // Cookie does not exist
  }

  if (cookieExists("session")) {
    console.log("Cookie exists!");
    isAuthorized.set(true);
    push("/login")
  }
  axios.interceptors.response.use(null, function (error) {
    console.log("INTERCEPTED!!!")
    console.log(error)
    if (error.response.status == 401) {
      isAuthorized.set(false);
      push("/login")
    }
    return Promise.reject(error);
  });

  async function logout() {
    const response = await axios.get("/api/containers", {
      withCredentials: true
    });
  }
</script>

<main>
  {#if $isAuthorized}
    <Router {routes} />
  {:else}
    <Login />
  {/if}
  <button on:click={logout}></button>
</main>

<style>
</style>
