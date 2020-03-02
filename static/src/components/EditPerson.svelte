<script>
  export let params = {};
  import { onMount } from "svelte";
  import config from "../config";
  import { push } from "svelte-spa-router";
  import {
    mealTypeToString,
    mealTypeOptions,
    cookieUser,
    genderOptions,
    convertDateToAge,
    convertDateToString
  } from "../helpers";
  import Table from "../UI/Table.svelte";
  import ToggleIcon from "../UI/ToggleIcon.svelte";
  let loaded = false;
  let mealRatings = [];
  let person = {};
  let editMode = false;

  const toggleEdit = () => (editMode = !editMode);

  const loadPerson = () => {
    loaded = false;
    fetch(`${config.apiUrl}api/person/${params.id}`, {
      method: "GET",
      headers: {
        "x-access-token": cookieUser.token
      }
    })
      .then(res => res.json())
      .then(res => {
        person = res;
        mealRatings = person.MealRatings;
        // person.totalCalories = getTotalCalorieCount(mealRatings)

        loaded = true;
      });
  };

  const getTotalCalorieCount = mealRatings => {
    let totalCalories = mealRatings.reduce(
      (a, c) => a + c.EatenMeal.Meal.Calories,
      0
    );
    return totalCalories;
  };

  const submitUpdatePersonForm = () => {
    fetch(`${config.apiUrl}api/person/${params.id}`, {
      method: "PUT",
      headers: {
        "x-access-token": cookieUser.token
      },
      body: JSON.stringify(person)
    })
      .then(res => res.json())
      .then(res => {
        toggleEdit();
      });
  };

  onMount(() => loadPerson());
</script>

<a href="/#/home/AddChild">Back</a>
{#if loaded}
  <ToggleIcon on:toggled={toggleEdit} />

  {#if editMode}
    <div class="add-child-from">
      <label for="name">child name</label>
      <input type="text" name="name" id="name" bind:value={person.Name} />
      <label for="Gender">Gender</label>
      <select id="Gender" bind:value={person.Gender}>
        {#each genderOptions as goption}
          <option value={goption.id}>{goption.text}</option>
        {/each}
      </select>
      <label for="DateOfBirth">DateOfBirth</label>
      <input
        type="date"
        name="DateOfBirth"
        id="DateOfBirth"
        bind:value={person.DateOfBirth} />
      <br />
      <button on:click={submitUpdatePersonForm}>Submit</button>
    </div>
  {:else}
    <h1>{person.Name}</h1>
    <h3>{convertDateToAge(person.DateOfBirth)}</h3>
  {/if}
  <h4>Meals Eaten: {mealRatings.length}</h4>

  {#if mealRatings && mealRatings.length > 0}
    <Table headers={['Name', 'Date', 'Rating']}>

      {#each mealRatings as meal}
        <tr>
          <td>
            <a href={`/#/meal/${meal.EatenMeal.Meal.ID}`}>
              {meal.EatenMeal.Meal.Name} ({mealTypeToString(meal.EatenMeal.Meal.Mealtype)})
            </a>
          </td>
          <td>{convertDateToString(meal.EatenMeal.Date)}</td>
          <td>{meal.Rating}</td>
        </tr>
      {/each}
    </Table>
  {/if}
{:else}
  <h1>Loading</h1>
{/if}
