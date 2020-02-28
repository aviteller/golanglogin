<script>
  export let params = {};
  import { onMount } from "svelte";
  import config from "../config";
  import { push } from "svelte-spa-router";
  import { mealTypeToString, mealTypeOptions, cookieUser } from "../helpers";

 

  let loaded = false;
  let mealRatings = [];
  let person = {};


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

  onMount(() => loadPerson());
</script>

<a href="/#/">Back</a>
{#if loaded}
  <h1>{person.Name}</h1>
  <h3>{person.Age} Years old</h3>
  <h4>Meals Eaten: {mealRatings.length}</h4>

  {#if mealRatings && mealRatings.length > 0}
    {#each mealRatings as meal}
      <a href={`/#/meal/${meal.EatenMeal.Meal.ID}`}>
        {meal.EatenMeal.Meal.Name}({mealTypeToString(meal.EatenMeal.Meal.Mealtype)})
      </a>
      - {meal.EatenMeal.Date} - Rating : {meal.Rating}
      <br />
    {/each}
  {/if}
{:else}
  <h1>Loading</h1>
{/if}
