<script>
  export let params = {};
  import { onMount } from "svelte";
  import config from "../config";
  import Cookies from "../Cookies";
  import { mealTypeToString, mealTypeOptions, cookieUser } from "../helpers";
  import { push } from "svelte-spa-router";
  import ToggleIcon from "../UI/ToggleIcon.svelte";
  import Table from "../UI/Table.svelte";

  let meal = {};
  let children = [];
  let mealRatings = [];
  let loaded = false;

  let mergedChildrenArray = [];

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
        mergedChildrenArray = [];

        for (let i = 0; i < children.length; i++) {
          mergedChildrenArray.push({
            ...children[i],
            ...mealRatings.find(mr => mr.PersonID === children[i].ID)
          });
        }
        loaded = true;
      });
  };

  const addRating = childID => {
    loaded = false;
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
        loadEatenMeal();
      });
  };

  const toggleRating = (childID, ate) => {
    loaded = false;
    let newRating = {
      Ate: !ate
    };

    fetch(`${config.apiUrl}api/rating/${childID}`, {
      method: "PUT",
      headers: {
        "x-access-token": cookieUser.token
      },
      body: JSON.stringify(newRating)
    }).then(res => {
      loadEatenMeal();
    });
  };

  const updateRating = (e, ratingID) => {
    let newRating = {
      Rating: parseInt(e.target.value)
    };

    if (0 <= newRating.Rating && newRating.Rating <= 10) {
      loaded = false;
      fetch(`${config.apiUrl}api/rating/${ratingID}`, {
        method: "PUT",
        headers: {
          "x-access-token": cookieUser.token
        },
        body: JSON.stringify(newRating)
      }).then(res => {
        loadEatenMeal();
      });
    }
  };

  const getTotalCalorieCount = mealRatings => {
    let totalCalories = mealRatings.reduce(
      (a, c) => a + c.EatenMeal.Meal.Calories,
      0
    );
    return totalCalories;
  };

  const getAverageRating = mergedChildrenArray => {
    let totalEaten = 0;
    let totalRatings = mergedChildrenArray.reduce((a, c) => {
      if (c.Rating && c.Ate) {
        totalEaten++;
        return a + c.Rating;
      } else {
        return a;
      }
    }, 0);

    let average = totalRatings / totalEaten;

    return average.toFixed(2)
  };

  onMount(() => {
    loadEatenMeal();
  });
</script>

<a href="/#/">Back</a>
{#if loaded}
  <h1>{meal.Name} - {mealTypeToString(meal.Mealtype)}</h1>
  <h4>{meal.Date}</h4>
  <!-- <h2>{meal.Calories}</h2> -->
  <h5>Average Rating: {getAverageRating(mergedChildrenArray)}</h5>

  <Table headers={['Child Name', 'Eaten?', 'Rating(0-10)']}>

    {#each mergedChildrenArray as child}
      <tr>
        {#if !child.PersonID}
          <td>{child.Name}</td>
          <td>
            <ToggleIcon on:toggled={() => addRating(child.ID)} />
            <!-- <button on:click={() => addRating(child.ID)}>NO</button> -->
          </td>
          <td>0</td>
        {:else}
          <td>{child.Name}</td>
          <td>
            <ToggleIcon
              on:toggled={() => toggleRating(child.ID, child.Ate)}
              toggledMode={child.Ate ? true : false} />
          </td>
          <td>
            {#if child.Ate}
              <input
                style="width:90px"
                type="number"
                min="0"
                max="10"
                bind:value={child.Rating}
                on:blur={e => updateRating(e, child.ID)} />
            {:else}0{/if}
          </td>
        {/if}
      </tr>
    {/each}
  </Table>
{:else}
  <h1>LOADING</h1>
{/if}
