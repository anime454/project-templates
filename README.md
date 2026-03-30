# project-templates
This repository contains reusable project structure templates that I frequently use.

## Setup Guide

### Prerequisites

The setup scripts will check for and install the following tools if they are missing:

| Tool | Purpose |
|---|---|
| [Git](https://git-scm.com/downloads) | Version control |
| [Go 1.25+](https://go.dev/dl/) | Build and run Go projects |
| [Docker](https://www.docker.com/products/docker-desktop/) | Run local databases via Docker Compose |
| [make](https://www.gnu.org/software/make/) | Run project tasks |
| [golangci-lint](https://golangci-lint.run/) | Go linter |

---

### macOS

```bash
# Clone the repository
git clone https://github.com/anime454/project-templates.git
cd project-templates

# Run the setup script
bash scripts/setup.sh
```

> **Note:** The script uses [Homebrew](https://brew.sh) to install missing tools on macOS.
> Install Homebrew first if you haven't already:
> ```bash
> /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
> ```

---

### Linux

```bash
# Clone the repository
git clone https://github.com/anime454/project-templates.git
cd project-templates

# Run the setup script
bash scripts/setup.sh
```

> **Note:** The script uses `apt-get` and official installer scripts to install missing tools.
> You may be prompted for your `sudo` password.

---

### Windows

Open **Command Prompt** or **PowerShell** as Administrator, then run:

```bat
REM Clone the repository
git clone https://github.com/anime454/project-templates.git
cd project-templates

REM Run the setup script
scripts\setup.bat
```

> **Note:** The script uses [winget](https://learn.microsoft.com/windows/package-manager/winget/) (Windows Package Manager)
> to install missing tools. Make sure the **App Installer** is up to date in the Microsoft Store.

---

### After Setup

Once setup completes, you can start working with any template. For example, to run the
Go Hexagonal / Car Parking System template:

```bash
cd go/hexagonal/car-parking-system

# Start PostgreSQL via Docker Compose
docker compose up -d

# Run the API server
make run

# Run tests
make test
```
