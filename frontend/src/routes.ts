import Containers from "./routes/containers/list/Page.svelte";
import Home from "./routes/Home.svelte"
import Logging from "./routes/containers/logging/Page.svelte"
import Create from "./routes/containers/create/Page.svelte"

export const routes = {
  // Exact path
  "/": Home,
  "/containers": Containers,
  "/containers/:id/logs": Logging,
  "/containers/create": Create,
};
