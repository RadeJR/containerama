<script lang="ts">
  import { Button } from "$lib/components/ui/button/index.js";
  import LightSwitch from "$lib/components/LightSwitch.svelte";

</script>

<div class="p-1 fixed top-1 left-1">
  <LightSwitch />
</div>
<div class="flex items-center justify-center py-12">
  <div class="mx-auto grid w-[350px] gap-6">
    <div class="grid gap-2 text-center">
      <h1 class="text-3xl font-bold">Login</h1>
      <p class="text-balance text-muted-foreground">
        Click this button to login with Zitadel
      </p>
    </div>
    <div class="grid gap-4">
      <Button href="/login" class="w-full">Login with Zitadel</Button>
    </div>
  </div>
</div>
