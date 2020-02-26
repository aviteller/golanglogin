<script>
  export let params = {};
  import { onMount } from "svelte";
  import config from "../config";
  import Cookies from "../Cookies";
  import { push } from "svelte-spa-router";

  let c = new Cookies();

  let meal = {};
  let children = [];
  let mealRatings = [];
  let loaded = false;

  let cookieUser = {
    id: JSON.parse(c.getCookie("jwt")).id,
    name: JSON.parse(c.getCookie("jwt")).name,
    token: JSON.parse(c.getCookie("jwt")).token
  };

  const loadEatenMeal = () => {
    fetch(`${config.apiUrl}api/eatenmeal/${params.id}`, {
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
        meal = {
          Name: res.eatenMeal.Meal.Name,
          Mealtype: res.eatenMeal.Meal.Mealtype,
          Date: res.eatenMeal.Date,
          Calories: res.eatenMeal.Meal.Calories
        };
        mealRatings = res.eatenMeal.MealRatings;
        children = res.people;
        loaded = true;
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

  const addRating = childID => {
    let newRating = {
      EatenMealID: parseInt(params.id),
      PersonID: childID,
      Ate: true
    };

    fetch(`${config.apiUrl}api/rating`, {
      method: "POST",
      headers: {
        "x-access-token": cookieUser.token
      },
      body: JSON.stringify(newRating)
    })
      .then(res => res.json())
      .then(res => {
        console.log("res", res);
      });
  };

  const deleteRating = id => {
    fetch(`${config.apiUrl}api/rating/${id}`, {
      method: "DELETE",
      headers: {
        "x-access-token": cookieUser.token
      }
    }).then(res => {
      console.log("res", res);
    });
  };

  const updateRating = () => {
    
  }
  onMount(() => {
    loadEatenMeal();
  });
</script>

<a href="/#/">Back</a>
{#if loaded}
  <h1>{meal.Name} - {mealTypeToString(meal.Mealtype)}</h1>
  <h4>{meal.Date}</h4>
  <h2>{meal.Calories}</h2>

  <table border="1">
    <tr>
      <th>Child Name</th>
      <th>Eaten?</th>
      <th>Rating</th>
    </tr>

    {#each children as child}
      <tr>
        {#if mealRatings.filter(m => child.ID === m.PersonID).length == 0}
          <td>{child.Name}</td>
          <td>
            <button on:click={() => addRating(child.ID)}>No</button>
          </td>
          <td>0</td>
        {:else}
          <td>{child.Name}</td>
          <td>
            <button on:click={() => deleteRating(child.ID)}>Yes</button>
          </td>
          <td>1</td>
        {/if}
      </tr>
    {/each}
  </table>
{:else}
  <h1>LOADING</h1>
{/if}
