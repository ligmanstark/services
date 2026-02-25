import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { RouterProvider, createBrowserRouter } from "react-router";
import "@app/styles/index.css";
import routes from "./app/routes";
import { createHead, UnheadProvider } from "@unhead/react/client";

const router = createBrowserRouter(routes);
const head = createHead();

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <UnheadProvider head={head}>
      <RouterProvider router={router} />
    </UnheadProvider>
  </StrictMode>
);
