<div align="center">

# ssh-portfolio

**A fully interactive terminal portfolio, served over SSH.**
Browse my work without leaving your shell.

[![Go](https://img.shields.io/badge/Go-1.25-00ADD8?style=flat-square&logo=go)](https://go.dev)
[![Bubbletea](https://img.shields.io/badge/Bubbletea-TUI-FF69B4?style=flat-square)](https://github.com/charmbracelet/bubbletea)
[![Docker](https://img.shields.io/badge/Docker-ready-2496ED?style=flat-square&logo=docker)](https://hub.docker.com)
[![CI/CD](https://github.com/MoveFastAndBreakThings-dot/ssh-portfolio/actions/workflows/ci-cd.yml/badge.svg)](https://github.com/MoveFastAndBreakThings-dot/ssh-portfolio/actions/workflows/ci-cd.yml)
[![Fly.io](https://img.shields.io/badge/Deployed-Fly.io-8B5CF6?style=flat-square)](https://fly.io)
[![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)](LICENSE)

</div>

---

## Try it live

```bash
ssh -p 23234 samar-personal-portfolio.fly.dev
```

> No login. No password. Just SSH in and explore.

---

## What is this?

Most portfolios are websites. This one is an SSH server.

You connect with a standard SSH client, and instead of a shell, you land in a rich terminal UI — tabs for About, Experience, Projects, Skills, and Links. Navigate with your keyboard, cycle through color themes, scroll through content. It's a portfolio that lives where developers already are: the terminal.

The About tab renders a true-color portrait photo using half-block Unicode characters (`▄`) with 24-bit ANSI color — no ASCII art, actual pixels.

Built entirely in Go using the [Charmbracelet](https://charm.sh) stack:

| Package | Role |
|---|---|
| [`wish`](https://github.com/charmbracelet/wish) | SSH server middleware |
| [`bubbletea`](https://github.com/charmbracelet/bubbletea) | Elm-architecture TUI framework |
| [`lipgloss`](https://github.com/charmbracelet/lipgloss) | Terminal layout & styling |
| [`bubbles`](https://github.com/charmbracelet/bubbles) | Pre-built TUI components |

---

## Key bindings

| Key | Action |
|---|---|
| `→` / `l` / `Tab` | Next tab |
| `←` / `h` / `Shift+Tab` | Previous tab |
| `↑` / `k` &nbsp; `↓` / `j` | Scroll |
| `1` – `5` | Jump directly to tab |
| `t` | Cycle theme: Dark → Light → Hacker |
| `q` / `Ctrl+C` | Quit |

---

## Run locally

**Prerequisites:** Go 1.21+

```bash
git clone https://github.com/your-username/ssh-portfolio
cd ssh-portfolio
go run .
```

Then in a second terminal:

```bash
ssh -p 23234 localhost
```

Accept the host-key fingerprint on first connection. The key is generated once and saved to `.ssh/host_key`.

---

## Makefile

| Command | Does |
|---|---|
| `make build` | Compile binary |
| `make run` | Dev server on `:23234` |
| `make connect` | `ssh -p 23234 localhost` |
| `make test` | Run all tests (verbose) |
| `make test-race` | Tests with race detector |
| `make cover` | HTML coverage report |
| `make docker-up` | Build + start container |
| `make docker-down` | Stop container |
| `make docker-logs` | Tail container logs |
| `make clean` | Remove build artifacts |
| **`make ship`** | **fmt → vet → build → test → deploy** |

---

## CI/CD (GitHub Actions)

Every push to `main`:
1. Runs tests + race detector
2. If tests pass → auto-deploys to Fly.io

PRs only run tests (no deploy).

### One-time setup

**Get Fly.io API token:**
```bash
fly tokens create deploy -x 999999h
```
Copy the output token.

**Add to GitHub:**
1. Repo → **Settings → Secrets and variables → Actions**
2. **New repository secret**
3. Name: `FLY_API_TOKEN`, Value: paste token
4. Save

Now every `git push` deploys automatically.

---

## Deploy to Fly.io (live setup)

This project is deployed on [Fly.io](https://fly.io) free tier. To redeploy after changes:

```bash
fly deploy
```

### First-time setup

```bash
# install flyctl
curl -L https://fly.io/install.sh | sh

# login
fly auth login

# create app (lowercase, dashes only)
fly apps create your-app-name

# create volume for SSH host key persistence
fly volumes create ssh_keys --region ord --size 1

# deploy
fly deploy

# allocate public IPs
fly ips allocate-v4 --shared
fly ips allocate-v6
```

Users connect via:
```bash
ssh -p 23234 your-app-name.fly.dev
```

> SSH uses raw TCP — it cannot be proxied through nginx, Caddy, or Cloudflare. The port must be open directly. Fly.io handles this natively via `[[services]]` TCP config in `fly.toml`.

The `fly.toml` mounts a persistent volume at `/app/.ssh` so the host key survives redeploys. Without this, clients get "host key changed" warnings on every deploy.

---

## Deploy with Docker (self-hosted)

```bash
docker compose up --build -d
```

The `.ssh/` volume persists the host key across restarts.

---

## Customise for yourself

All personal data lives in one file:

```
content/data.go
```

Edit the structs there — all five views update automatically.

```go
var MyProfile = Profile{
    Name: "Your Name",
    Role: "Your Role",
    Bio:  []string{"paragraph one", "paragraph two"},
}
```

### Replace the portrait photo

Swap `tui/portrait.jpg` with your own JPEG. The renderer automatically:
- Decodes the JPEG
- Scales to 40 cols wide preserving aspect ratio
- Renders using half-block `▄` chars with 24-bit ANSI true color

No tools needed — just replace the file.

Then redeploy:
```bash
fly deploy
```

---

## Environment variables

| Variable | Default | Description |
|---|---|---|
| `SSH_PORT` | `23234` | Port the SSH server listens on |

---

## Testing

Tests across 5 files in `tests/`:

| File | Covers |
|---|---|
| `content_test.go` | Data validation — no blank fields, valid URLs, no duplicates |
| `model_test.go` | Tab nav, wrap-around, jump keys, quit, resize, unknown keys |
| `views_test.go` | Every tab's rendered content, bullets, header, footer |
| `styles_test.go` | Themes, colors, portrait rendering, `NewStyles` |
| `keygen_test.go` | Host key file creation, permissions, PEM validity, RSA parse |

```bash
make test
make test-race
make cover
```

---

## Project structure

```
ssh-portfolio/
├── .github/
│   └── workflows/
│       └── ci-cd.yml    # test → deploy on push to main
├── main.go              # SSH server setup (wish + bubbletea middleware)
├── Makefile             # build / test / deploy targets
├── fly.toml             # Fly.io deployment config
├── Dockerfile           # two-stage build → ~7MB Alpine image
├── docker-compose.yml   # local Docker setup
├── content/
│   └── data.go          # ← edit this to personalise
├── tui/
│   ├── model.go         # Bubbletea model, update loop, keybindings
│   ├── views.go         # Tab renderers (About, Experience, Projects, Skills, Links)
│   ├── styles.go        # Themes and lipgloss style definitions
│   ├── ascii.go         # Portrait renderer (JPEG → half-block true color)
│   └── portrait.jpg     # Your photo (swap to personalise)
└── tests/
    ├── content_test.go
    ├── model_test.go
    ├── views_test.go
    ├── styles_test.go
    └── keygen_test.go
```

---

## Tech stack

```
Go  ──▶  wish  ──▶  SSH server
                       │
                  bubbletea  ──▶  TUI event loop
                       │
                  lipgloss   ──▶  layout + colour
                       │
                  image/jpeg ──▶  portrait → half-block pixels
                       │
                  content/   ──▶  your data
```

---

## License

MIT — fork it, adapt it, make it yours.

---

<div align="center">

**Built by [Samardeep Singh](https://github.com/MoveFastAndBreakThings-dot)**

*"Why have a website when you can have an SSH server?"*

</div>
