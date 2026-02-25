import { useHead } from "@unhead/react";
import type { FC } from "react";

const ServicesPage: FC = () => {
  useHead({
    title: "Сервисы - Idea Garden",
    meta: [{ name: "description", content: "Самый лучший дашборд Idea Garden" }],
  });
  return (
    <div>
      <h1>Services</h1>
    </div>
  );
};

export default ServicesPage;
