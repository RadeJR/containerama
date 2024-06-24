import Containers from "./routes/containers/list/Page.svelte";
import Home from "./routes/Home.svelte"
import Logging from "./routes/containers/logging/Page.svelte"
import ContainerCreate from "./routes/containers/create/Page.svelte"
import Stacks from "./routes/stacks/list/Page.svelte"
import StacksCreate from "./routes/stacks/create/Page.svelte"

export const routes = {
  // Exact path
  "/": Home,
  "/containers": Containers,
  "/containers/:id/logs": Logging,
  "/containers/create": ContainerCreate,
  "/stacks": Stacks,
  "/stacks/create": StacksCreate,
};
