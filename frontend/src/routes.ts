import Home from './routes/Home.svelte'
import Login from './routes/Login.svelte'
import NotFound from './routes/NotFound.svelte'

export const routes = {
    // Exact path
    '/': Home,
    '/login': Login,

    // Catch-all
    // This is optional, but if present it must be the last
    '*': NotFound,
}
