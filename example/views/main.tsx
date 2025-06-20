import ReactDOM from "react-dom/client";
import {
  createBrowserRouter,
  RouterProvider,
  type RouteObject,
} from "react-router";
import routes from "./routes";
import React from "react";

const router = createBrowserRouter(routes);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
