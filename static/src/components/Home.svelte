<script>
  export let params = {};
  import { onMount } from "svelte";
  import config from "../config";
  import Cookies from "../Cookies";
  import {
    mealTypeToString,
    mealTypeOptions,
    genderOptions,
    cookieUser,
    convertDateToAge,
    convertDateToString
  } from "../helpers";
  import { push } from "svelte-spa-router";
  import Meals from "./Meals.svelte";
  import Children from "./Children.svelte";
  import EatenMeals from "./EatenMeals.svelte";
  let c = new Cookies();
  let userLoaded = false;
  let mode = params.mode || "";

  let user = {};
  let children = [];
  let meals = [];
  let eatenmeals = [];

  const toggleMode = () => {
    if (mode == "") {
      mode = "AddChild";
    } else if (mode == "AddChild") {
      mode = "AddMeal";
    } else {
      mode = "";
    }
  };

  const logout = () => {
    c.eraseCookie("jwt");
    push("/login");
  };

  const loadUser = () => {
    fetch(`${config.apiUrl}api/user/${cookieUser.id}/0`, {
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
        user = res.user;
        children = user.People;
        meals = user.Meals;
        eatenmeals = res.eatenMeals;

        userLoaded = true;
      });
  };

  const updateChildrenArray = e => (children = e.detail);
  const updateMealsArray = e => (meals = e.detail);
  const updateEatenMealsArray = e => (eatenmeals = e.detail);
  onMount(() => loadUser());
</script>

{#if userLoaded}

  <h1>Hello {user.Name}</h1>
  <button on:click={logout}>Logout</button>
  <br />
  <button on:click={toggleMode}>Toggle View</button>

  {#if mode == 'AddChild'}
    <Children {children} on:updatechildren={updateChildrenArray} />
  {:else if mode == 'AddMeal'}
    <Meals {meals} on:updatemeals={updateMealsArray} />
  {:else}
    <EatenMeals
      {eatenmeals}
      {meals}
      on:updateeatenmeals={updateEatenMealsArray} />
  {/if}
{:else}
  <h1>Loading...</h1>
{/if}
