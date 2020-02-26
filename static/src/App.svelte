<script>
  import routes from "./routes";
  import Router from "svelte-spa-router";

  let title = "App";
</script>

<style>
  main {
    width: 40%;
    margin-left: 30%;
    border: 1px solid;
    padding: 1rem;
  }

  @media (max-width: 770px) {
    main {
      width: 100%;
      margin-left: 0;
      border: 1px solid;
      padding: 1rem;
    }
  }
</style>

<main style=" ">
  <Router {routes} />
</main>
