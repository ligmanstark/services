import type { FC } from "react";
import { useHead } from "@unhead/react";

const DashboardPage: FC = () => {
  useHead({
    title: "Дашборд - Idea Garden",
    meta: [{ name: "description", content: "Самый лучший дашборд Idea Garden" }],
  });
  return (
    <>
      <div>
        <h1>Dashboard</h1>
      </div>
    </>
  );
};

export default DashboardPage;
