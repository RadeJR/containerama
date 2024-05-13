import Containers from "./routes/Containers.svelte";
import Home from "./routes/Home.svelte"

export const routes = {
  // Exact path
  "/": Home,
  "/containers": Containers
};
