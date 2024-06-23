import Containers from "./routes/containers/Page.svelte";
import Home from "./routes/Home.svelte"
import Logging from "./routes/containers/logging/Page.svelte"

export const routes = {
  // Exact path
  "/": Home,
  "/containers": Containers,
  "/containers/:id/logs": Logging,
};
