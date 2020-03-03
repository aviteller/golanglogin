<script>
  import { push } from "svelte-spa-router";
  import config from "../config";
  let user = {
    name: "",
    email: "",
    password: ""
  };

  let errors = "";

  const createUser = async user => {
    let res = await fetch(`${config.apiUrl}api/register`, {
      mode: "cors",
      method: "POST",
      body: JSON.stringify(user)
    });
    let data = await res.json();
    console.log(data);
    return data;
  };

  const submitRegisterForm = () => {
    if (user.name == "" || user.email == "" || user.password == "") {
      errors = "Please enter all fields";
    } else {
      errors = "";
      createUser(user).then(res => {
        user = {
          name: "",
          email: "",
          password: ""
        };
        push("/login");
      });
    }
  };
</script>

<h1>Register</h1>
<!-- <a href="#/Login">Login</a> -->

<form class="auth-from" on:submit|preventDefault>
  <div style="color:red">{errors}</div>
  <label for="name">Username</label>
  <input type="text" name="name" id="name" bind:value={user.name} />
  <label for="name">Email</label>
  <input type="text" name="email" id="email" bind:value={user.email} />
  <label for="password">Password</label>
  <input
    type="password"
    name="password"
    id="password"
    bind:value={user.password} />
  <button on:click={submitRegisterForm}>Submit</button>
</form>
