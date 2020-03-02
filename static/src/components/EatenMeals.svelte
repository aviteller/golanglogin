<script>
  import config from "../config";
  import { createEventDispatcher } from "svelte";
  import { onMount } from "svelte";
  import ToggleIcon from "../UI/ToggleIcon.svelte";
  import Table from "../UI/Table.svelte";
  const dispatch = createEventDispatcher();
  import {
    mealTypeToString,
    mealTypeOptions,
    genderOptions,
    cookieUser,
    convertDateToAge,
    convertDateToString
  } from "../helpers";
  export let meals;
  let newEatenMeal = {
    MealID: "",
    Date: new Date().toISOString().split("T")[0]
  };
  let mealsOptGroup = [];
  const makeMealSelect = () => {
    mealsOptGroup = [];
    meals.forEach(meal => {
      if (mealsOptGroup.filter(m => meal.Mealtype === m.id).length == 0)
        mealsOptGroup.push({
          id: meal.Mealtype,
          name: mealTypeToString(meal.Mealtype),
          values: []
        });
    });

    mealsOptGroup.forEach(mealOpt => {
      meals.forEach(m => {
        if (mealOpt.id == m.Mealtype) {
          mealOpt.values.push({
            id: m.ID,
            text: m.Name
          });
        }
      });
    });
  };

  const deleteEatenMeal = id => {
    if(window.confirm("Are you sure want to delete")) {

      fetch(`${config.apiUrl}api/eatenmeal/${id}`, {
        method: "DELETE",
        headers: {
          "x-access-token": cookieUser.token
        }
      }).then(res => {
        eatenmeals = eatenmeals.filter(eatenmeal => eatenmeal.ID !== id);
        if (eatenmeals.length == 0) eatenmeals = [];
        dispatch("updateeatenmeals", eatenmeals);
      });
    }
  };

  const submitAddEatenMealForm = () => {
    newEatenMeal = {
      UserID: cookieUser.id,
      MealID: newEatenMeal.MealID,
      Date: newEatenMeal.Date
    };

    fetch(`${config.apiUrl}api/eatenmeal`, {
      method: "POST",
      headers: {
        "x-access-token": cookieUser.token
      },
      body: JSON.stringify(newEatenMeal)
    })
      .then(res => res.json())
      .then(res => {
        if (res.ID != 0) {
          let eatenMealToAdd = res;
          eatenmeals = [...eatenmeals, eatenMealToAdd];
          dispatch("updateeatenmeals", eatenmeals);
        }
        newEatenMeal = {
          MealID: "",
          Date: new Date().toISOString().split("T")[0]
        };
      });
  };

  export let eatenmeals;

  let showEatenMealForm = false;
  const toggleEatenMealForm = () => (showEatenMealForm = !showEatenMealForm);

  onMount(() => makeMealSelect());
</script>

<h1>Add eaten meal</h1>
<ToggleIcon on:toggled={toggleEatenMealForm} />
{#if showEatenMealForm}
  <div>

    <select id="eatenmeal-id" bind:value={newEatenMeal.MealID}>
      <option value="0">--Please Select Meal--</option>
      {#each mealsOptGroup as optGroup}
        <optgroup label={optGroup.name}>
          {#each optGroup.values as mealOpt}
            <option value={mealOpt.id}>{mealOpt.text}</option>
          {/each}
        </optgroup>
      {/each}
    </select>
    <input type="date" bind:value={newEatenMeal.Date} />
    <button on:click={submitAddEatenMealForm}>Add new eaten meal</button>
  </div>
{/if}
{#if eatenmeals && eatenmeals.length > 0}
  <Table headers={['Name', 'Date', { name: 'Action', col: 2 }]}>

    {#each eatenmeals as meal}
      <tr>

        <td>{meal.Meal.Name} ({mealTypeToString(meal.Meal.Mealtype)})</td>
        <td>{convertDateToString(meal.Date)}</td>
        <td>
          <a href={`/#/eatenmeal/${meal.ID}`}>Edit</a>
        </td>
        <td>
          <button on:click={() => deleteEatenMeal(meal.ID)}>X</button>
        </td>
      </tr>
    {/each}
  </Table>
{:else}
  <h1>Add Eaten meal</h1>
{/if}
