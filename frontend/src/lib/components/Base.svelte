<script lang="ts">
  import CircleUser from "lucide-svelte/icons/circle-user";
  import Menu from "lucide-svelte/icons/menu";
  import Package2 from "lucide-svelte/icons/package-2";
  import Search from "lucide-svelte/icons/search";

  import { Button } from "$lib/components/ui/button/index.js";
  import * as Card from "$lib/components/ui/card/index.js";
  import * as DropdownMenu from "$lib/components/ui/dropdown-menu/index.js";
  import { Input } from "$lib/components/ui/input/index.js";
  import * as Sheet from "$lib/components/ui/sheet/index.js";
  import Nav from "$lib/components/Nav.svelte";
    import LightSwitch from "./LightSwitch.svelte";
</script>

<div
  class="grid min-h-screen w-full md:grid-cols-[220px_1fr] lg:grid-cols-[280px_1fr]"
>
  <div class="hidden border-r bg-muted/40 md:block">
    <div class="flex h-full max-h-screen flex-col gap-2">
      <div class="flex h-14 items-center border-b px-4 lg:h-[60px] lg:px-6">
        <a href="/" class="flex items-center gap-2 font-semibold">
          <Package2 class="h-6 w-6" />
          <span class="">Containerama</span>
        </a>
      </div>
      <div class="flex-1">
        <Nav />
      </div>
      <div class="mt-auto p-4">
        <Card.Root>
          <Card.Header class="p-2 pt-0 md:p-4">
            <Card.Title>Upgrade to Pro</Card.Title>
            <Card.Description>
              Unlock all features and get unlimited access to our support team.
            </Card.Description>
          </Card.Header>
          <Card.Content class="p-2 pt-0 md:p-4 md:pt-0">
            <Button size="sm" class="w-full">Upgrade</Button>
          </Card.Content>
        </Card.Root>
      </div>
    </div>
  </div>
  <div class="flex flex-col">
    <header
      class="flex h-14 items-center gap-4 border-b bg-muted/40 px-4 lg:h-[60px] lg:px-6"
    >
      <Sheet.Root>
        <Sheet.Trigger asChild let:builder>
          <Button
            variant="outline"
            size="icon"
            class="shrink-0 md:hidden"
            builders={[builder]}
          >
            <Menu class="h-5 w-5" />
            <span class="sr-only">Toggle navigation menu</span>
          </Button>
        </Sheet.Trigger>
        <Sheet.Content side="left" class="flex flex-col">
          <Nav />
          <div class="mt-auto">
            <Card.Root>
              <Card.Header>
                <Card.Title>Upgrade to Pro</Card.Title>
                <Card.Description>
                  Unlock all features and get unlimited access to our support
                  team.
                </Card.Description>
              </Card.Header>
              <Card.Content>
                <Button size="sm" class="w-full">Upgrade</Button>
              </Card.Content>
            </Card.Root>
          </div>
        </Sheet.Content>
      </Sheet.Root>
      <div class="w-full flex-1">
        <form>
          <div class="relative">
            <Search
              class="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground"
            />
            <Input
              type="search"
              placeholder="Search products..."
              class="w-full appearance-none bg-background pl-8 shadow-none md:w-2/3 lg:w-1/3"
            />
          </div>
        </form>
      </div>
      <LightSwitch />
      <DropdownMenu.Root>
        <DropdownMenu.Trigger asChild let:builder>
          <Button
            builders={[builder]}
            variant="secondary"
            size="icon"
            class="rounded-full"
          >
            <CircleUser class="h-5 w-5" />
            <span class="sr-only">Toggle user menu</span>
          </Button>
        </DropdownMenu.Trigger>
        <DropdownMenu.Content align="end">
          <DropdownMenu.Label>My Account</DropdownMenu.Label>
          <DropdownMenu.Separator />
          <DropdownMenu.Item>Settings</DropdownMenu.Item>
          <DropdownMenu.Item>Support</DropdownMenu.Item>
          <DropdownMenu.Separator />
          <DropdownMenu.Item>Logout</DropdownMenu.Item>
        </DropdownMenu.Content>
      </DropdownMenu.Root>
    </header>
    <main class="flex flex-1 flex-col gap-4 p-4 lg:gap-6 lg:p-6">
      <slot />
    </main>
  </div>
</div>
