# 🧱 Golem

**Golem** is a batteries-included fullstack framework for Go — built to bring Rails-style productivity into the Go ecosystem, powered by modern React and data loaders.

Like its mythical namesake, Golem is a powerful servant: structured, dependable, and brought to life by your code.

---

## ⚡️ Why Golem?

Traditional Go frameworks leave too much up to you — routing, database setup, SSR, frontend boilerplate. **Golem** solves this by:

- 🧠 Enforcing structure via convention-over-configuration
- ⚙️ Integrating deeply with React Router and its data loader API
- 🔌 Providing seamless Go ↔ JSX loader wiring
- 🧱 Including DB migrations and model layers
- 🎯 Ensuring your SPA is SEO-friendly and fast

---

## 🔋 Features

| Feature                | Description                                                     |
| ---------------------- | --------------------------------------------------------------- |
| 🧭 Nested Routing      | Go-based route tree, mirrored into React Router                 |
| 📡 Data Loaders        | First-class Go ↔ React loader syncing                           |
| ⚡ SSR Ready           | Initial data is preloaded and hydrated without JavaScript delay |
| 🧱 ORM Included        | Simple built-in ORM with migrations                             |
| 🔧 CLI Tools           | `golem generate route`, `golem generate model`, `golem dev`     |
| 🔐 Auth Middleware     | Session & user helpers out of the box                           |
| 📦 Bundler Flexibility | Works with **Vite**, **Bun**, or custom setups                  |

---

## 🏗 Project Structure

my-app/
├── app/
│ ├── routes/ # Go route handlers and loader functions
│ ├── models/ # Database models (auto-migrated)
│ ├── views/ # JSX/TSX components for frontend
│ └── assets/ # Static files
├── frontend/ # React app (Vite or Bun)
│ ├── routes/ # React Router route components
│ └── entry.tsx # Hydration entry point
├── golem.config.json # Shared route metadata between Go ↔ React
├── go.mod
└── main.go

---

## 🚀 Quickstart

### 1. Install Golem CLI

```bash
go install github.com/your-org/golem/cmd/golem@latest

```

### 2. Create a new project

```bash
golem new my-app
cd my-app
```

### 3. Run the dev server

```bash
golem dev
```

This runs both:

Go backend (with loader and database support)

React frontend (via Vite or Bun)

## ✨ Example: Go ↔ JSX Loader

app/routes/dashboard.go

```go
package routes

func DashboardLoader(ctx Context) (any, error) {
    user := ctx.Session.User()
    stats := db.GetStatsFor(user.ID)
    return map[string]any{
        "username": user.Name,
        "messages": stats.UnreadCount,
    }, nil
}

var Dashboard = Route{
    Path: "/dashboard",
    Loader: DashboardLoader,
}
```

frontend/routes/dashboard.tsx

```tsx
import { useLoaderData } from "react-router-dom";

export default function Dashboard() {
  const { username, messages } = useLoaderData() as {
    username: string;
    messages: number;
  };

  return (
    <div>
      Welcome, {username}. You have {messages} unread messages.
    </div>
  );
}
```

## 🔗 Integration Flow

| Step                              | Description                          |
| --------------------------------- | ------------------------------------ |
| ✅ Define Go route + loader       | `app/routes/dashboard.go`            |
| ✅ Golem generates route metadata | Written to `golem.config.json`       |
| ✅ React uses that metadata       | Automatically hydrated on first load |
| ✅ Data stays in sync             | No boilerplate REST endpoints needed |

## 📦 Built-in ORM

Example:

```bash
golem generate model User name:string email:string age:int
```

Creates:

```go
type User struct {
    ID    int
    Name  string
    Email string
    Age   int
}
```

You can then run:

```bash
golem migrate
```

## 🔌 Compatible Bundlers

You can use either:

- **Vite**: Recommended for best DX
- **Bun**: Use with golem dev --bun
- **Custom**: Just hook into the loader JSON injection + hydration

## 🧩 Advanced Plans

- [ ] Nested loader hydration
- [ ] HMR with DB/model reload
- [ ] GraphQL or RPC integration (optional)
- [ ] File-system based routes
- [ ] Plugin system for auth, queue, mail, etc

## 👥 Community

We’re building Golem for developers who love speed and clarity. Whether you're a Go backend dev, a React enthusiast, or someone who just wants one tool that just works — you're welcome here.

- GitHub: [github.com/aglis-lab/golem](github.com/aglis-lab/golem)
- Discussions: [github.com/aglis-lab/golem/discussions](https://github.com/aglis-lab/golem)
- Twitter: @golem

## 📜 License

MIT © 2025–present — The Golem Contributors

```yaml
Let me know if you want:

- A working starter template zipped up
- Codegen files for `golem generate model` or `golem generate route`
- Integrated DB migration scaffolding
- A logo concept

Or if you want to rebrand again — just say the word.
```
