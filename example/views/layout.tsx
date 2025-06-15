import "./app.css";

export default function Layout({ children }: { children?: React.ReactNode }) {
  return (
    <>
      <header className="p-4 shadow">My App</header>
      <main className="container mx-auto p-4">{children}</main>
    </>
  );
}
