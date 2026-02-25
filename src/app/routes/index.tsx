import type { RouteObject } from "react-router";

const routes: RouteObject[] = [
  {
    path: "/",
    lazy: async () => {
      const { default: DashboardLayout } = await import("@widgets/layout/dashboard");
      return { element: <DashboardLayout /> };
    },
    children: [
      {
        path: "home",
        lazy: async () => {
          const { default: HomePage } = await import("@pages/home/page");
          return { element: <HomePage /> };
        },
      },
      {
        path: "services",
        lazy: async () => {
          const { default: ServicesPage } = await import("@pages/services/page");
          return { element: <ServicesPage /> };
        },
      },
      {
        index: true,
        lazy: async () => {
          const { default: HomePage } = await import("@pages/home/page");
          return { element: <HomePage /> };
        },
      },
    ],
  },
];

export default routes;
