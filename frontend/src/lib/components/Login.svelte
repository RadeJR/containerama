<script lang="ts">
  import { Button } from "$lib/components/ui/button/index.js";
  import { Input } from "$lib/components/ui/input/index.js";
  import { Label } from "$lib/components/ui/label/index.js";
  import type { AxiosResponse } from "axios";
  import LightSwitch from "./LightSwitch.svelte";
  import { getAxios } from "$conf/axios";
  import { isAuthorized } from "$store";

  let username: string;
  let password: string;

  async function login() {
    console.log("Sending req");
    console.log(JSON.stringify({ username, password }));
    let response: AxiosResponse = await getAxios().post(
      "/api/login",
      JSON.stringify({ username, password }),
    );
    if (response.status == 204) {
      isAuthorized.set(true);
    } else {
      console.log(response.data);
    }
  }
</script>

<div class="p-1 fixed top-1 left-1">
  <LightSwitch />
</div>
<div class="flex items-center justify-center py-12">
  <div class="mx-auto grid w-[350px] gap-6">
    <div class="grid gap-2 text-center">
      <h1 class="text-3xl font-bold">Login</h1>
      <p class="text-balance text-muted-foreground">
        Enter your email below to login to your account
      </p>
    </div>
    <div class="grid gap-4">
      <div class="grid gap-2">
        <Label for="username">Email</Label>
        <Input id="username" bind:value={username} required />
      </div>
      <div class="grid gap-2">
        <div class="flex items-center">
          <Label for="password">Password</Label>
          <a href="##" class="ml-auto inline-block text-sm underline">
            Forgot your password?
          </a>
        </div>
        <Input id="password" type="password" bind:value={password} required />
      </div>
      <Button on:click={login} class="w-full">Login</Button>
    </div>
  </div>
</div>
