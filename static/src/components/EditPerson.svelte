<script>
  export let params = {};
  import { onMount } from "svelte";
  import config from "../config";
  import Cookies from "../Cookies";
  import { push } from "svelte-spa-router";

  let c = new Cookies();
 
  let cookieUser = {
    id: JSON.parse(c.getCookie("jwt")).id,
    name: JSON.parse(c.getCookie("jwt")).name,
    token: JSON.parse(c.getCookie("jwt")).token
  };

   const loadPerson = () => {
    fetch(`${config.apiUrl}api/person/${params.id}`, {
      method: "GET",
      headers: {
        "x-access-token": cookieUser.token
      }
    })
      .then(res => {
   
          return res.json();
        
      })
      .then(res => {
        console.log('res', res)
      });
  };

  onMount(() => loadPerson())

</script>