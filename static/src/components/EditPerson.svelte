<script>
  export let params = {};
  import { onMount } from "svelte";
  import config from "../config";
  import Cookies from "../Cookies";
  import { push } from "svelte-spa-router";

  let c = new Cookies();

  let loaded = false;
  let mealRatings = [];
  let person = {};

  let cookieUser = {
    id: JSON.parse(c.getCookie("jwt")).id,
    name: JSON.parse(c.getCookie("jwt")).name,
    token: JSON.parse(c.getCookie("jwt")).token
  };

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
        mealRatings = person.MealRatings
        // person.totalCalories = getTotalCalorieCount(mealRatings)
       

        loaded = true;
      });
  };

    let mealTypeOptions = [
    { id: 0, text: "Breakfast" },
    { id: 1, text: "Lunch" },
    { id: 2, text: "Supper" },
    { id: 3, text: "Snack" }
  ];

  const getTotalCalorieCount = (mealRatings) => {
    let totalCalories = mealRatings.reduce((a, c) => a + c.EatenMeal.Meal.Calories, 0);
    return totalCalories;
  }

  const mealTypeToString = mealType => {
    let mealReturn = mealTypeOptions.filter(meal => meal.id == mealType);
    return mealReturn[0].text;
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
        {meal.EatenMeal.Meal.Name}({mealTypeToString(meal.EatenMeal.Meal.Mealtype)}) - {meal.EatenMeal.Date} - Rating : {meal.Rating} <br>
    {/each}
  {/if}

{:else}
  <h1>Loading</h1>
{/if}
