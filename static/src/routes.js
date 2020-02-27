import Cookies from "./Cookies";
import Home from "./components/Home.svelte";
import Login from "./components/Login.svelte";
import Register from "./components/Register.svelte";
import EditMeal from "./components/EditMeal.svelte";
import EditEatenMeal from "./components/EditEatenMeal.svelte";
import EditPerson from "./components/EditPerson.svelte";

let c = new Cookies();

const userLoggedIn = () => {
  let jwt;
  if ((jwt = c.getCookie("jwt"))) {
    return true;
  } else return false;
};


let routes;
const urlParams = new URLSearchParams(window.location.search);

if (!urlParams.has("routemap")) {
  routes = {
    "/": userLoggedIn() ? Home : Login,
    "/login": Login,
    "/register": Register,
    "/meal/:id": userLoggedIn() ? EditMeal : Login,
    "/eatenmeal/:id": userLoggedIn() ? EditEatenMeal : Login,
    "/person/:id": userLoggedIn() ? EditPerson : Login
  };
}


export default routes;
