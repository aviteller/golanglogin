<script>
  export let params = {};
  import { onMount } from "svelte";
  import config from "../config";
  import Cookies from "../Cookies";
  import { push } from "svelte-spa-router";
  import {
    mealTypeToString,
    mealTypeOptions,
    cookieUser,
    convertDateToString
  } from "../helpers";
  import ToggleIcon from "../UI/ToggleIcon.svelte";
  import Table from "../UI/Table.svelte";

  let mealLoaded = false;
  let meal = {};
  let ingredients = [];
  let editMode = false;
  let editIngredientMode = false;
  let newIngredient = {
    Name: "",
    Calories: 0
  };

  let eatenMeals = [];
  const toggleEdit = () => (editMode = !editMode);
  const toggleEditIngredient = () => (editIngredientMode = !editIngredientMode);

  const loadMeal = () => {
    fetch(`${config.apiUrl}api/meal/${params.id}`, {
      method: "GET",
      headers: {
        "x-access-token": cookieUser.token
      }
    })
      .then(res => {
        if (res.status === 403) {
          logout();
        } else {
          return res.json();
        }
      })
      .then(res => {
        meal = res.meal;
        ingredients = meal.Ingredients;
        eatenMeals = res.eatenMeals;
        mealLoaded = true;
      });
  };

  onMount(() => {
    loadMeal();
  });

  const submitAddIngredientForm = () => {
    newIngredient = {
      MealID: parseInt(params.id),
      Name: newIngredient.Name,
      Calories: newIngredient.Calories
    };

    fetch(`${config.apiUrl}api/ingredient`, {
      method: "POST",
      headers: {
        "x-access-token": cookieUser.token
      },
      body: JSON.stringify(newIngredient)
    })
      .then(res => res.json())
      .then(res => {
        if (res.Value.ID != 0) {
          let IngredientToAdd = {
            Name: res.Value.Name,
            Calories: res.Value.Calories,
            ID: res.Value.ID
          };

          meal.Calories += IngredientToAdd.Calories;
          ingredients = [...ingredients, IngredientToAdd];
        }

        newIngredient = {
          Name: "",
          Calories: 0
        };
      });
  };

  const deleteIngredient = id => {
    if (window.confirm("Are you sure want to delete")) {
      fetch(`${config.apiUrl}api/ingredient/${id}`, {
        method: "DELETE",
        headers: {
          "x-access-token": cookieUser.token
        }
      }).then(res => {
        ingredients = ingredients.filter(ingredient => {
          if (ingredient.ID === id) {
            meal.Calories -= ingredient.Calories;
          } else {
            return ingredient;
          }
        });
        if (ingredients.length == 0) ingredients = [];
      });
    }
  };

  const submitUpdateMealForm = () => {
    fetch(`${config.apiUrl}api/meal/${params.id}`, {
      method: "PUT",
      headers: {
        "x-access-token": cookieUser.token
      },
      body: JSON.stringify(meal)
    }).then(res => toggleEdit());
  };
</script>

<a href="/#/home/AddMeal">Back</a>
{#if mealLoaded}
  <ToggleIcon on:toggled={toggleEdit} />

  {#if editMode}
    <div class="add-child-from">
      <label for="name">meal name</label>
      <input type="text" name="name" id="name" bind:value={meal.Name} />
      <label for="MealType">MealType</label>
      <select id="MealType" bind:value={meal.MealType}>
        {#each mealTypeOptions as goption}
          <option value={goption.id}>{goption.text}</option>
        {/each}
      </select>

      <br />
      <button on:click={submitUpdateMealForm}>Submit</button>
    </div>
  {:else}
    <h1>{meal.Name}</h1>
    <h3>{mealTypeToString(meal.Mealtype)}</h3>
    <h4>Total Calories : {meal.Calories}</h4>
    <h4>Number of Ingredients : {ingredients.length}</h4>
    <h4>Number of Times eaten : {eatenMeals.length}</h4>
  {/if}
  <hr />
  <ToggleIcon on:toggled={toggleEditIngredient} />
  {#if editIngredientMode}
    <div class="add-child-from">
      <label for="name">Ingredient name</label>
      <input
        type="text"
        name="name"
        id="name"
        bind:value={newIngredient.Name} />
      <label for="Calories">Calories</label>
      <input
        type="number"
        name="Calories"
        id="Calories"
        bind:value={newIngredient.Calories} />

      <br />
      <button on:click={submitAddIngredientForm}>Submit</button>
    </div>
  {/if}
  {#if ingredients.length > 0}
    <Table headers={['Name', 'Calories', 'Action']}>

      {#each ingredients as ingredient}
        <tr>
          <td>{ingredient.Name}</td>
          <td>{ingredient.Calories}</td>
          <td>
            <button on:click={() => deleteIngredient(ingredient.ID)}>X</button>
          </td>
        </tr>
      {/each}
    </Table>
  {:else}
    <h1>Add ingredient</h1>
  {/if}
  <hr />
  {#if eatenMeals.length > 0}
    <Table headers={['Eaten Date', 'Link', 'Average Rating']}>
      {#each eatenMeals as eMeal}
        <tr>
          <td>{convertDateToString(eMeal.Date)}</td>
          <td>
            <a href={`/#/eatenmeal/${eMeal.ID}`}>Go to Meal</a>
          </td>
          <td>{eMeal.AverageRating.toFixed(2)}</td>
        </tr>
      {/each}
    </Table>
  {:else}
    <h1>No eaten Meals</h1>
  {/if}
{:else}
  <h1>Loading Meal</h1>
{/if}
