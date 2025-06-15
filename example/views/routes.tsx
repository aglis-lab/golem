import type { RouteObject } from "react-router";
import React, { Suspense } from "react";

const Layout = React.lazy(() => import("./layout"));
const Home = React.lazy(() => import("./pages/home"));
const Testing = React.lazy(() => import("./pages/testing"));

const routes: RouteObject[] = [
  {
    path: "/testing",
    hydrateFallbackElement: <div>Loading 1...</div>,
    element: (
      <Suspense fallback={<div>Loading...</div>}>
        <Layout>
          <Testing />
        </Layout>
      </Suspense>
    ),
  },
  {
    path: "/",
    hydrateFallbackElement: <div>Loading 1...</div>,
    element: (
      <Suspense fallback={<div>Loading...</div>}>
        <Layout>
          <Home />
        </Layout>
      </Suspense>
    ),
    loader: async () => {
      let rnd = Math.random();
      console.log("Loader called, random number:", rnd);
      // Simulate a delay for demonstration purposes
      await new Promise((resolve) => setTimeout(resolve, 1000));
      // Return data to be used in the component
      console.log("Loader completed, returning data");
      // You can return any data you want here, for example:
      return { title: `Welcome to Golem!, random value is ${rnd}` };
      // For now, just returning a simple object
    },
  },
];

export default routes;
