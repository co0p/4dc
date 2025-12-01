# How to Get Started

1. **Install dependencies:**
	```sh
	npm install
	```


## Development Workflow (Hot-Reload)

1. **Install dependencies:**
	```sh
	npm install
	```

2. **Start the Go backend (API only):**
	```sh
	cd backend
	go run ./cmd/shareit/main.go
	```

3. **Start the Svelte frontend (hot-reload):**
	```sh
	npm run dev
	```

4. **Access the app:**
	Open [http://localhost:5173](http://localhost:5173) in your browser.

The frontend will fetch catalog data from the backend via `/api/catalog` (proxy is configured in Vite for development).

## Production Workflow (Single Deployable)

1. **Build the Svelte frontend:**
	```sh
	npm run build
	```

2. **Start the Go backend (serves API and static frontend):**
	```sh
	cd backend
	go run ./cmd/shareit/main.go
	```

3. **Access the app:**
	Open [http://localhost:8080](http://localhost:8080) in your browser.

The backend will serve both API endpoints and static frontend assets from the production build.

# sv

Everything you need to build a Svelte project, powered by [`sv`](https://github.com/sveltejs/cli).


## Developing

Once you've created a project and installed dependencies with `npm install` (or `pnpm install` or `yarn`), start a development server:

```sh
npm run dev

# or start the server and open the app in a new browser tab
npm run dev -- --open
```

## Building

To create a production version of your app:

```sh
npm run build
```

You can preview the production build with `npm run preview`.

> To deploy your app, you may need to install an [adapter](https://svelte.dev/docs/kit/adapters) for your target environment.
