import Cookies from "./Cookies";
import Home from "./components/Home.svelte";
import Login from "./components/Login.svelte";
import Register from "./components/Register.svelte";

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
    "/Login": Login,
    "/Register": Register,
  };
}


export default routes;
