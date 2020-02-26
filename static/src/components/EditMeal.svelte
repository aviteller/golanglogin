<script>
  export let params = {};
  import { onMount } from "svelte";
  import config from "../config";
  import Cookies from "../Cookies";
  import { push } from "svelte-spa-router";

  let c = new Cookies();
  let mealLoaded = false;
  let meal = {};
  let ingredients = [];
  let newIngredient = {
    Name: "",
    Calories: 0
  };

  let cookieUser = {
    id: JSON.parse(c.getCookie("jwt")).id,
    name: JSON.parse(c.getCookie("jwt")).name,
    token: JSON.parse(c.getCookie("jwt")).token
  };

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
        meal = res;
        ingredients = meal.Ingredients;
        mealLoaded = true;
      });
  };

  let mealTypeOptions = [
    { id: 0, text: "Breakfast" },
    { id: 1, text: "Lunch" },
    { id: 2, text: "Supper" },
    { id: 3, text: "Snack" }
  ];

  const mealTypeToString = mealType => {
    let mealReturn = mealTypeOptions.filter(meal => meal.id == mealType);
    return mealReturn[0].text;
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
  };
</script>

<a href="/#/">Back</a>
{#if mealLoaded}
  <h1>{meal.Name}</h1>
  <h3>{mealTypeToString(meal.Mealtype)}</h3>
  <h4>Total Calories : {meal.Calories}</h4>
  <h4>Number of Ingredients : {ingredients.length}</h4>

  <div class="add-child-from">
    <label for="nam e">Ingredient name</label>
    <input type="text" name="name" id="name" bind:value={newIngredient.Name} />
    <label for="Calories">Calories</label>
    <input
      type="number"
      name="Calories"
      id="Calories"
      bind:value={newIngredient.Calories} />

    <br />
    <button on:click={submitAddIngredientForm}>Submit</button>
  </div>

  {#if ingredients.length > 0}
    {#each ingredients as ingredient}
      {ingredient.Name} {ingredient.Calories}
      <button on:click={() => deleteIngredient(ingredient.ID)}>X</button>
      <br />
    {/each}
  {:else}
    <h1>Add ingredient</h1>
  {/if}
{:else}
  <h1>Loading Meal</h1>
{/if}
