<script>
  import routes from "./routes";
  import Router from "svelte-spa-router";
  import Header from "./UI/Header.svelte";


</script>

<style>
  main {
    width: 600px;
    /* margin-left: 30%; */
    /* border: 1px solid; */
    padding: 1rem;
    /* overflow-y: scroll; */
    max-height: 600px;
  }

  /* @media (max-width: 770px) {
    main {
      width: 80%;
      margin-left: 5%;
      margin-right: 10%;
      border: 1px solid;
      padding: 1rem;
    }
  } */
</style>

<main style=" ">
  <Header />
  <section>
    <Router {routes} />
  </section>
</main>
