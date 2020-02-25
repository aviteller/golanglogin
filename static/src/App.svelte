<script>
  import routes from "./routes";
  import Router from "svelte-spa-router";

	let title = "App";
</script>

<main>
  <Router {routes} />
</main>
