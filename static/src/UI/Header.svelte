<script>
  import Cookies from "../Cookies";
  import { cookieUser, waitFor } from "../helpers";
  import { push } from "svelte-spa-router";
  // import {waitFor} from "../helpers";
  let c = new Cookies();

  const logout = async () => {
    await c.eraseCookie("jwt");

    window.location.reload();
  };
</script>

<style>
  .header {
    display: flex;
    justify-content: space-between;
  }
</style>

{#if cookieUser.name}
  <div class="header">
  
  <h1>Hello {cookieUser.name}</h1>
  <button on:click={logout}>Logout</button>
  </div>
{:else}
  <a href="#/Register">Register</a>
  or
  <a href="#/Login">Login</a>
{/if}
<hr />
