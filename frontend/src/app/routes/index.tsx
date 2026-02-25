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
        path: "services",
        lazy: async () => {
          const { default: ServicesPage } = await import("@pages/services/page");
          return { element: <ServicesPage /> };
        },
      },
      {
        path: "login",
        lazy: async () => {
          const { default: LoginPage } = await import("@pages/login/page");
          return { element: <LoginPage /> };
        },
      },
      {
        path: "registration",
        lazy: async () => {
          const { default: RegistrationPage } = await import("@pages/registration/page");
          return { element: <RegistrationPage /> };
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
