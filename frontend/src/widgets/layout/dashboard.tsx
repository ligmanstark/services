import { Outlet } from "react-router";
import { useState, type FC } from "react";
import "@app/styles/index.css";
import { SidebarProvider, SidebarTrigger } from "@/shared/ui/shadcn/sidebar";
import { AppSidebar } from "./ui/sidebar";

const DashboardLayout: FC = () => {
  const [open, setOpen] = useState(false);
  return (
    <SidebarProvider open={open} onOpenChange={setOpen}>
      <AppSidebar />
      <main className="p-4">
        <SidebarTrigger />
        <Outlet />
      </main>
    </SidebarProvider>
  );
};

export default DashboardLayout;
