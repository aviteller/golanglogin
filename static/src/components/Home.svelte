<script>
  import { onMount } from "svelte";
  import config from "../config";
  import Cookies from "../Cookies";
  import { push } from "svelte-spa-router";

  let userLoaded = false;
  let mode = "";

  let c = new Cookies();

  let cookieUser = {
    id: JSON.parse(c.getCookie("jwt")).id,
    name: JSON.parse(c.getCookie("jwt")).name,
    token: JSON.parse(c.getCookie("jwt")).token
  };

  let newChild = {
    Name: "",
    Age: 0,
    Gender: "Male"
  };

  let newMeal = {
    Name: "",
    MealType: "Breakfast"
  };

  let newEatenMeal = {
    MealID: "",
    Date: new Date().toISOString().split('T')[0]
  };

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
  let mealsOptGroup = [];
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
        makeMealSelect();

        userLoaded = true;
      });
  };

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

  const deletePerson = id => {
    fetch(`${config.apiUrl}api/person/${id}`, {
      method: "DELETE",
      headers: {
        "x-access-token": cookieUser.token
      }
    }).then(res => {
      children = children.filter(child => child.ID !== id);
      if (children.length == 0) children = [];
    });
  };

  const deleteMeal = id => {
    fetch(`${config.apiUrl}api/meal/${id}`, {
      method: "DELETE",
      headers: {
        "x-access-token": cookieUser.token
      }
    }).then(res => {
      meals = meals.filter(meal => meal.ID !== id);
      if (meals.length == 0) meals = [];
      makeMealSelect();
    });
  };
  const deleteEatenMeal = id => {
    fetch(`${config.apiUrl}api/eatenmeal/${id}`, {
      method: "DELETE",
      headers: {
        "x-access-token": cookieUser.token
      }
    }).then(res => {
      eatenmeals = eatenmeals.filter(eatenmeal => eatenmeal.ID !== id);
      if (eatenmeals.length == 0) eatenmeals = [];
 
    });
  };

  const submitAddChildForm = () => {
    newChild = {
      UserID: cookieUser.id,
      Name: newChild.Name,
      Age: newChild.Age,
      Gender: newChild.Gender
    };

    fetch(`${config.apiUrl}api/person`, {
      method: "POST",
      headers: {
        "x-access-token": cookieUser.token
      },
      body: JSON.stringify(newChild)
    })
      .then(res => res.json())
      .then(res => {
        if (res.Value.ID != 0) {
          let childToAdd = {
            Name: res.Value.Name,
            Age: res.Value.Age,
            Gender: res.Value.Gender,
            ID: res.Value.ID
          };
          children = [...children, childToAdd];
        }

        newChild = {
          Name: "",
          Age: 0,
          Gender: "Male"
        };
      });
  };
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
          makeMealSelect();
        }
        newMeal = {
          Name: "",
          MealType: "Breakfast"
        };
      });
  };
  const submitAddEatenMealForm = () => {
    newEatenMeal = {
      UserID: cookieUser.id,
      MealID: newEatenMeal.MealID,
      Date: newEatenMeal.Date,
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
          let eatenMealToAdd = res
          eatenmeals = [...eatenmeals, eatenMealToAdd];
        
        }
        newEatenMeal = {
    MealID: "",
    Date: new Date().toISOString().split('T')[0]
        };
      });
  };

  let genderOptions = [
    { id: "Male", text: "Male" },
    { id: "Female", text: "Female" }
  ];

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
    loadUser();
  });
</script>

{#if userLoaded}

  <h1>Hello {user.Name}</h1>
  <button on:click={logout}>Logout</button>
  <br />
  <button on:click={toggleMode}>{mode || 'Default'}</button>

  {#if mode == 'AddChild'}
    <h1>Add Child</h1>
    <div class="add-child-from">
      <label for="name">child name</label>
      <input type="text" name="name" id="name" bind:value={newChild.Name} />
      <label for="Gender">Gender</label>
      <select id="Gender" bind:value={newChild.Gender}>
        {#each genderOptions as goption}
          <option value={goption.id}>{goption.text}</option>
        {/each}
      </select>
      <label for="Age">Age</label>
      <input type="number" name="Age" id="Age" bind:value={newChild.Age} />
      <br />
      <button on:click={submitAddChildForm}>Submit</button>
    </div>
    {#if children.length > 0}
      {#each children as child}
        {child.Name} {child.Age} {child.Gender}
        <button on:click={() => deletePerson(child.ID)}>X</button>
        <br />
      {/each}
    {/if}
  {:else if mode == 'AddMeal'}
    <h1>Add Meal</h1>
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
    {#if meals.length > 0}
      {#each meals as meal}
        {meal.Name} - {meal.Calories} - {mealTypeToString(meal.Mealtype)}
        <a href={`/#/meal/${meal.ID}`}>Edit</a>
        <button on:click={() => deleteMeal(meal.ID)}>X</button>
        <br />
      {/each}
    {:else}
      <h1>Add meal</h1>
    {/if}
  {:else}
    <h1>Create meal</h1>
    <div>
    
    <select id="eatenmeal-id" bind:value={newEatenMeal.MealID} >
    <option value="0">--Please Select Meal--</option>
      {#each mealsOptGroup as optGroup}
        <optgroup label={optGroup.name}>
          {#each optGroup.values as mealOpt}
            <option value={mealOpt.id}>{mealOpt.text}</option>
          {/each}
        </optgroup>
      {/each}
    </select>
    <input type="date" bind:value={newEatenMeal.Date}>
    <button on:click={submitAddEatenMealForm}>Add new eaten meal</button>
    </div>
    {#if eatenmeals && eatenmeals.length > 0}
      {#each eatenmeals as meal}
        {meal.ID} - {meal.Meal.Name} ({mealTypeToString(meal.Meal.Mealtype)})- {meal.Date} 
        <a href={`/#/eatenmeal/${meal.ID}`}>Edit</a>
        <button on:click={() => deleteEatenMeal(meal.ID)}>X</button>
        <br />
      {/each}
    {:else}
      <h1>Add Eaten meal</h1>
    {/if}
  {/if}
{:else}
  <h1>Loading...</h1>
{/if}
