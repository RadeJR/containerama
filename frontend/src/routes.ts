import Containers from "./routes/containers/Page.svelte";
import Home from "./routes/Home.svelte"

export const routes = {
  // Exact path
  "/": Home,
  "/containers": Containers
};
