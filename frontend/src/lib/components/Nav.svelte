<script lang="ts">
  // icons
  import Package from "lucide-svelte/icons/package";
  import Network from "lucide-svelte/icons/network";
  import Layers from "lucide-svelte/icons/layers"
  // packages
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

<nav class="pt-2 grid items-start px-2 text-sm font-medium lg:px-4">
  <a href="/containers" use:link class={navInactiveClass}>
    <Package class="h-4 w-4" />
    Containers
  </a>
  <a href="/networks" use:link class={navInactiveClass}>
    <Network class="h-4 w-4" />
    Networks
  </a>
  <a href="/stacks" use:link class={navInactiveClass}>
    <Layers class="h-4 w-4" />
    Stacks
  </a>
</nav>
