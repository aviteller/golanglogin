<script>
  import { createEventDispatcher } from "svelte";
  import ToggleIcon from "../UI/ToggleIcon.svelte";
  import Table from "../UI/Table.svelte";
  import config from "../config";
  import {
    mealTypeToString,
    mealTypeOptions,
    genderOptions,
    cookieUser,
    convertDateToAge,
    convertDateToString
  } from "../helpers";

  export let meals;

  const dispatch = createEventDispatcher();

  let newMeal = { Name: "", MealType: "Breakfast" };

  const submitAddMealForm = () => {
    newMeal = {
      UserID: cookieUser.id,
      Name: newMeal.Name,
      MealType: newMeal.MealType
    };

    fetch(`${config.apiUrl}api/meal`, {
      method: "POST",
      headers: {
        "x-access-token": cookieUser.token
      },
      body: JSON.stringify(newMeal)
    })
      .then(res => res.json())
      .then(res => {
        if (res.Value.ID != 0) {
          let mealToAdd = {
            Name: res.Value.Name,
            Mealtype: res.Value.Mealtype,
            Calories: 0,
            ID: res.Value.ID
          };
          meals = [...meals, mealToAdd];

          dispatch("updatemeals", meals);
        }
        newMeal = {
          Name: "",
          MealType: "Breakfast"
        };
      });
  };

  const deleteMeal = id => {
    if (window.confirm("Are you sure want to delete")) {
      fetch(`${config.apiUrl}api/meal/${id}`, {
        method: "DELETE",
        headers: {
          "x-access-token": cookieUser.token
        }
      }).then(res => {
        meals = meals.filter(meal => meal.ID !== id);
        if (meals.length == 0) meals = [];

        dispatch("updatemeals", meals);
      });
    }
  };

  let showAddMealForm = false;

  const toggleAddMealForm = () => (showAddMealForm = !showAddMealForm);
</script>

<h1>Add Meal</h1>
<ToggleIcon on:toggled={toggleAddMealForm} />
{#if showAddMealForm}
  <div class="add-child-from">
    <label for="name">meal name</label>
    <input type="text" name="name" id="name" bind:value={newMeal.Name} />
    <label for="MealType">MealType</label>
    <select id="MealType" bind:value={newMeal.MealType}>
      {#each mealTypeOptions as goption}
        <option value={goption.id}>{goption.text}</option>
      {/each}
    </select>

    <br />
    <button on:click={submitAddMealForm}>Submit</button>
  </div>
{/if}
{#if meals.length > 0}
  <Table headers={['Name', 'Calories', 'Type', { name: 'Action', col: 2 }]}>

    {#each meals as meal}
      <tr>
        <td>{meal.Name}</td>
        <td>{meal.Calories}</td>
        <td>{mealTypeToString(meal.Mealtype)}</td>
        <td>
          <a href={`/#/meal/${meal.ID}`}>Edit</a>

        </td>
        <td>
          <button on:click={() => deleteMeal(meal.ID)}>X</button>
        </td>
      </tr>
    {/each}
  </Table>
{:else}
  <h1>Add meal</h1>
{/if}
