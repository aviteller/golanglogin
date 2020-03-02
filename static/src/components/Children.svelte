<script>
  import { createEventDispatcher } from "svelte";
  import ToggleIcon from "../UI/ToggleIcon.svelte";
  import Table from "../UI/Table.svelte";
  const dispatch = createEventDispatcher();
  export let children;
  import config from "../config";
  let newChild = {
    Name: "",
    DateOfBirth: 0,
    Gender: "Male"
  };

  import {
    mealTypeToString,
    mealTypeOptions,
    genderOptions,
    cookieUser,
    convertDateToAge,
    convertDateToString
  } from "../helpers";

  let showAddChildForm = false;

  const deletePerson = id => {
    fetch(`${config.apiUrl}api/person/${id}`, {
      method: "DELETE",
      headers: {
        "x-access-token": cookieUser.token
      }
    }).then(res => {
      children = children.filter(child => child.ID !== id);
      if (children.length == 0) children = [];
      dispatch("updatechildren", children);
    });
  };
  const submitAddChildForm = () => {
    newChild = {
      UserID: cookieUser.id,
      Name: newChild.Name,
      DateOfBirth: newChild.DateOfBirth,
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
        if (res.Value.ID != 0) {
          let childToAdd = {
            Name: res.Value.Name,
            DateOfBirth: res.Value.DateOfBirth,
            Gender: res.Value.Gender,
            ID: res.Value.ID
          };
          children = [...children, childToAdd];
          dispatch("updatechildren", children);
        }

        newChild = {
          Name: "",
          DateOfBirth: "",
          Gender: "Male"
        };
      });
  };

  const toggleAddChildForm = () => (showAddChildForm = !showAddChildForm);
</script>

<h1>Add Child</h1>
<ToggleIcon on:toggled={toggleAddChildForm} />
{#if showAddChildForm}
  <div class="add-child-from">
    <label for="name">child name</label>
    <input type="text" name="name" id="name" bind:value={newChild.Name} />
    <label for="Gender">Gender</label>
    <select id="Gender" bind:value={newChild.Gender}>
      {#each genderOptions as goption}
        <option value={goption.id}>{goption.text}</option>
      {/each}
    </select>
    <label for="DateOfBirth">DateOfBirth</label>
    <input
      type="date"
      name="DateOfBirth"
      id="DateOfBirth"
      bind:value={newChild.DateOfBirth} />
    <br />
    <button on:click={submitAddChildForm}>Submit</button>
  </div>
{/if}
{#if children.length > 0}
  <Table headers={['Name', 'Age', 'Gender', 'Action']}>

    {#each children as child}
      <tr>
        <td>
          <a href={`/#/person/${child.ID}`}>{child.Name}</a>
        </td>
        <td>{convertDateToAge(child.DateOfBirth)}</td>
        <td>{child.Gender}</td>
        <td>
          <button on:click={() => deletePerson(child.ID)}>X</button>
        </td>
      </tr>

  
    {/each}
  </Table>
{/if}
