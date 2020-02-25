<script>
  import { onMount } from "svelte";
  import config from "../config";
  import Cookies from "../Cookies";
  import { push } from "svelte-spa-router";

  let userLoaded = false;

  let c = new Cookies();

  let cookieUser = {
    id: JSON.parse(c.getCookie("jwt")).id,
    name: JSON.parse(c.getCookie("jwt")).name,
    token: JSON.parse(c.getCookie("jwt")).token
  };

  let newChild = {
    Name: "",
    Age: 0,
    Gender: "Male"
  };

  let user = {};
  let children = [];

  const logout = () => {
    c.eraseCookie("jwt")
     push("/login");
  }

  const loadUser = () => {
    fetch(`${config.apiUrl}api/user/${cookieUser.id}`, {
      method: "GET",
      headers: {
        "x-access-token": cookieUser.token
      }
    })
      .then(res => res.json())
      .then(res => {
        user = res;
        children = user.Person;
        userLoaded = true;
      });
  };

  const deletePerson = id => {
    fetch(`${config.apiUrl}api/person/${id}`, {
      method: "DELETE",
      headers: {
        "x-access-token": cookieUser.token
      }
    }).then(res => {
      children = children.filter(child => child.ID !== id);
      if(children.length == 0) children=[]
    });
  };

  const submitAddChildForm = () => {
    newChild = {
      UserID: cookieUser.id,
      Name: newChild.Name,
      Age: newChild.Age,
      Gender: newChild.Gender
    };

    fetch(`${config.apiUrl}api/person`, {
      method: "POST",
      headers: {
        "x-access-token": cookieUser.token
      },
      body: JSON.stringify(newChild)
    })
      .then(res => res.json())
      .then(res => {

        if(res.Value.ID != 0) {

          let childToAdd = {
            Name: res.Value.Name,
            Age: res.Value.Age,
            Gender: res.Value.Gender,
            ID: res.Value.ID
          };
  
          children = [...children, childToAdd];
        }


        newChild = {
          Name: "",
          Age: 0,
          Gender: "Male"
        };
      });


  };

  let genderOptions = [
    { id: "Male", text: "Male" },
    { id: "Female", text: "Female" }
  ];

  let selected;

  onMount(() => {
    loadUser();
  });
</script>

{#if userLoaded}

  <h1>Hello {user.Name}</h1>
  <button on:click={logout}>Logout</button>

  <div class="add-child-from">
    <label for="name">child name</label>
    <input type="text" name="name" id="name" bind:value={newChild.Name} />
    <label for="Gender">Gender</label>
    <select id="Gender" bind:value={newChild.Gender}>
      {#each genderOptions as goption}
        <option value={goption.id}>{goption.text}</option>
      {/each}
    </select>
    <label for="Age">Age</label>
    <input type="number" name="Age" id="Age" bind:value={newChild.Age} />
    <br />
    <button on:click={submitAddChildForm}>Submit</button>
  </div>

  {#each children as child}
    {child.Name} {child.Age} {child.Gender}
    <button on:click={() => deletePerson(child.ID)}>X</button>
    <br />
  {/each}
{:else}
  <h1>Loading</h1>
{/if}
