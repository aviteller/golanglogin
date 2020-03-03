<script>
  import { push } from "svelte-spa-router";
  import config from "../config";
  import { waitFor } from "../helpers";
  import Cookies from "../Cookies";
  let c = new Cookies();
  let user = {
    name: "",
    password: ""
  };

  let errors = "";

  const loginUser = async user => {
    let res = await fetch(`${config.apiUrl}api/login`, {
      method: "POST",
      body: JSON.stringify(user)
    });
    let data = await res.json();
    return data;
  };

  const submitLoginForm = () => {
    if (user.name == "" || user.password == "") {
      errors = "Please enter all fields";
    } else {
      errors = "";
      loginUser(user)
        .then(res => {
          let userForCookie = {
            id: res.user.id,
            name: res.user.name,
            token: res.token,
            loggedIn: true
          };

          c.setCookie("jwt", JSON.stringify(userForCookie), 1);
        })
        .then(res => location.assign("/"));
    }
  };
</script>

<h1>Login</h1>
<!-- <a href="#/Register">Register</a> -->

<form class="auth-from" on:submit|preventDefault>
  <div style="color:red">{errors}</div>
  <label for="name">Username</label>
  <input type="text" name="name" id="name" bind:value={user.name} />
  <label for="password">Password</label>
  <input
    type="password"
    name="password"
    id="password"
    bind:value={user.password} />
  <button on:click={submitLoginForm}>Submit</button>
</form>
