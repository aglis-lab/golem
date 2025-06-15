export default function ErrorBoundary({ error }: { error: any }) {
  return (
    <div>
      <h1>Error</h1>
      <pre>{JSON.stringify(error, null, 2)}</pre>
    </div>
  );
}
