import { useEffect, useState } from "react";

export default function Home() {
  const [health, setHealth] = useState("loading...");
  const [info, setInfo] = useState(null);
  const [error, setError] = useState("");

  useEffect(() => {
    async function load() {
      try {
        const h = await fetch("/api/health");
        const healthText = await h.text();
        setHealth(healthText);

        const i = await fetch("/api/info");
        const infoJson = await i.json();
        setInfo(infoJson);
      } catch (e) {
        setError("Failed to connect to backend");
      }
    }

    load();
  }, []);

  return (
    <main style={{ fontFamily: "Arial, sans-serif", padding: 40, maxWidth: 900, margin: "0 auto" }}>
      <h1>РАБОТАЕТ</h1>
      <p>Frontend: Next.js</p>
      <p>Backend: Go</p>
      <p>Requests go through Next.js proxy: <code>/api/*</code></p>

      <hr />

      <h2>Health</h2>
      <p>{health}</p>

      <h2>Info</h2>
      {error && <p style={{ color: "red" }}>{error}</p>}
      {info ? (
        <pre
          style={{
            background: "#f4f4f4",
            padding: 16,
            borderRadius: 8,
            overflowX: "auto"
          }}
        >
          {JSON.stringify(info, null, 2)}
        </pre>
      ) : (
        <p>Loading info...</p>
      )}
    </main>
  );
}
