<script lang="ts">
  import Package from "lucide-svelte/icons/package";
  import Home from "lucide-svelte/icons/home";
  import ShoppingCart from "lucide-svelte/icons/shopping-cart";
  import Users from "lucide-svelte/icons/users";
  import { link } from "svelte-spa-router";
  import { location } from "svelte-spa-router";
  import { onMount } from "svelte";

  export let navActiveClass: string =
    "flex items-center gap-3 rounded-lg bg-muted px-3 py-2 text-primary transition-all hover:text-primary";
  export let navInactiveClass: string =
    "flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary";

  function setActive(location: string) {
    let navElements: NodeListOf<Element> = document.querySelectorAll("nav a");
    navElements.forEach((element) => {
      if (element.attributes.getNamedItem("href")?.value == "#" + location) {
        element.className = navActiveClass;
      } else {
        element.className = navInactiveClass;
      }
    });
  }

  onMount(() => {
    setActive($location)
  });

  $: {
    setActive($location)
  }
</script>

<nav class="grid items-start px-2 text-sm font-medium lg:px-4">
  <a href="/containers" use:link class={navInactiveClass}>
    <Home class="h-4 w-4" />
    Containers
  </a>
  <a href="/networks" use:link class={navInactiveClass}>
    <ShoppingCart class="h-4 w-4" />
    Networks
  </a>
  <a href="/stacks" use:link class={navInactiveClass}>
    <Package class="h-4 w-4" />
    Stacks
  </a>
  <a href="/users" use:link class={navInactiveClass}>
    <Users class="h-4 w-4" />
    Users
  </a>
</nav>
