<!-- Login.svelte -->
<script>
  import { isAuthorized } from "../store";

  let username = "";
  let password = "";

  async function login() {
    try {
      // Perform authentication logic here, for example, by making an API request
      const response = await fetch("/api/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ username, password }),
      });

      if (response.ok) {
        // Authentication successful, redirect or perform desired action
        console.log("Login successful");
        isAuthorized.set(true)
      } else {
        // Authentication failed, handle error
        console.error("Login failed");
      }
    } catch (error) {
      console.error("Error logging in:", error);
    }
  }
</script>

<h1>Login</h1>

<form on:submit|preventDefault={login}>
  <label>
    Username:
    <input type="text" bind:value={username} required />
  </label>
  <label>
    Password:
    <input type="password" bind:value={password} required />
  </label>
  <button type="submit">Login</button>
</form>
