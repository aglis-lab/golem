import { useLoaderData } from "react-router";
import { Welcome } from "./welcome";

export function meta() {
  return [
    { title: "New React Router App" },
    { name: "description", content: "Welcome to React Router!" },
  ];
}

export default function Home() {
  const data = useLoaderData<{ title: string }>();

  return <Welcome title={data.title} />;
}
