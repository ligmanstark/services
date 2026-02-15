import { Outlet } from "react-router";
import type { FC } from "react";
import "@app/styles/index.css";
import { SidebarProvider, SidebarTrigger } from "@/shared/ui/shadcn/sidebar";
import { AppSidebar } from "./ui/sidebar";

const DashboardLayout: FC = () => {
  return (
    <SidebarProvider>
      <AppSidebar />
      <main className="p-4">
        <SidebarTrigger />
        <Outlet />
      </main>
    </SidebarProvider>
  );
};

export default DashboardLayout;
