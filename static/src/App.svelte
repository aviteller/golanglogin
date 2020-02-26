<script>
  import routes from "./routes";
  import Router from "svelte-spa-router";

	let title = "App";
</script>

<main style="width:40%;margin-left:30%;border:1px solid; padding:1rem;">
  <Router {routes} />
</main>
