export function cookieExists(name: string): boolean {
  // Split document.cookie string into individual cookies
  const cookies: string[] = document.cookie.split(";");

  // Iterate over cookies to find the one with the specified name
  for (let i = 0; i < cookies.length; i++) {
    const cookie: string = cookies[i].trim();

    // Check if cookie starts with the specified name
    if (cookie.startsWith(name + "=")) {
      return true; // Cookie exists
    }
  }

  return false; // Cookie does not exist
}

