#!/usr/bin/env bash
# setup.sh – Bootstrap the development environment on macOS or Linux.
set -euo pipefail

# ---------------------------------------------------------------------------
# Helpers
# ---------------------------------------------------------------------------
info()    { echo "[setup] $*"; }
success() { echo "[setup] ✓ $*"; }
warn()    { echo "[setup] ⚠ $*"; }
error()   { echo "[setup] ✗ $*" >&2; exit 1; }

require_cmd() {
  command -v "$1" &>/dev/null || error "'$1' is required but not installed. $2"
}

# ---------------------------------------------------------------------------
# Detect OS
# ---------------------------------------------------------------------------
OS="$(uname -s)"
case "$OS" in
  Darwin) PLATFORM="macos" ;;
  Linux)  PLATFORM="linux" ;;
  *)      error "Unsupported OS: $OS" ;;
esac

info "Detected platform: $PLATFORM"

# ---------------------------------------------------------------------------
# Install missing dependencies
# ---------------------------------------------------------------------------
install_go() {
  info "Installing Go..."
  if [[ "$PLATFORM" == "macos" ]]; then
    require_cmd brew "Install Homebrew from https://brew.sh first."
    brew install go
  else
    # Linux – download the official binary
    GO_VERSION="1.25.1"
    ARCH="$(uname -m)"
    case "$ARCH" in
      x86_64)  GO_ARCH="amd64" ;;
      aarch64) GO_ARCH="arm64" ;;
      *)        error "Unsupported architecture: $ARCH" ;;
    esac
    TARBALL="go${GO_VERSION}.linux-${GO_ARCH}.tar.gz"
    info "Downloading $TARBALL..."
    curl -fsSL "https://go.dev/dl/${TARBALL}" -o "/tmp/${TARBALL}"
    sudo rm -rf /usr/local/go
    sudo tar -C /usr/local -xzf "/tmp/${TARBALL}"
    rm "/tmp/${TARBALL}"
    export PATH="$PATH:/usr/local/go/bin"
    info "Go installed to /usr/local/go. Add /usr/local/go/bin to your PATH."
  fi
}

install_docker() {
  info "Installing Docker..."
  if [[ "$PLATFORM" == "macos" ]]; then
    require_cmd brew "Install Homebrew from https://brew.sh first."
    brew install --cask docker
    warn "Start Docker Desktop manually before running containers."
  else
    info "Installing Docker via the official convenience script..."
    curl -fsSL https://get.docker.com | sudo sh
    sudo usermod -aG docker "$USER"
    warn "Log out and back in (or run 'newgrp docker') for group changes to take effect."
  fi
}

install_golangci_lint() {
  info "Installing golangci-lint..."
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
    | sh -s -- -b "$(go env GOPATH)/bin" latest
}

# ---------------------------------------------------------------------------
# Check / install each tool
# ---------------------------------------------------------------------------
info "Checking required tools..."

# git
require_cmd git "Install Git from https://git-scm.com/downloads."
success "git $(git --version | awk '{print $3}')"

# go
if ! command -v go &>/dev/null; then
  warn "Go not found."
  install_go
fi
success "go $(go version | awk '{print $3}')"

# docker (optional but recommended)
if ! command -v docker &>/dev/null; then
  warn "Docker not found. It is required to run the database locally."
  read -r -p "Install Docker now? [y/N] " answer
  if [[ "${answer,,}" == "y" ]]; then
    install_docker
  else
    warn "Skipping Docker installation. You will need it to start PostgreSQL."
  fi
else
  success "docker $(docker --version | awk '{print $3}' | tr -d ',')"
fi

# make
if ! command -v make &>/dev/null; then
  warn "make not found."
  if [[ "$PLATFORM" == "macos" ]]; then
    info "Installing Xcode Command Line Tools (includes make)..."
    xcode-select --install 2>/dev/null || true
  else
    info "Installing build-essential..."
    sudo apt-get update -y && sudo apt-get install -y build-essential
  fi
fi
success "make $(make --version | head -1)"

# golangci-lint
if ! command -v golangci-lint &>/dev/null; then
  warn "golangci-lint not found."
  install_golangci_lint
fi
success "golangci-lint $(golangci-lint --version 2>&1 | head -1)"

# ---------------------------------------------------------------------------
# Set up Git hooks
# ---------------------------------------------------------------------------
info "Configuring Git hooks..."
git -C "$(git rev-parse --show-toplevel)" config core.hooksPath .githooks
chmod +x .githooks/*
success "Git hooks configured (core.hooksPath=.githooks)"

# ---------------------------------------------------------------------------
# Download Go module dependencies
# ---------------------------------------------------------------------------
info "Downloading Go module dependencies..."

GO_PROJECTS=(
  "go/hexagonal/car-parking-system"
)

REPO_ROOT="$(git rev-parse --show-toplevel)"

for project in "${GO_PROJECTS[@]}"; do
  project_path="$REPO_ROOT/$project"
  if [[ -d "$project_path" ]]; then
    info "  → $project"
    (cd "$project_path" && go mod download)
    success "  $project dependencies downloaded"
  fi
done

# ---------------------------------------------------------------------------
# Done
# ---------------------------------------------------------------------------
echo ""
success "Setup complete! 🎉"
echo ""
echo "  Next steps:"
echo "  1. Start the database:  cd go/hexagonal/car-parking-system && docker compose up -d"
echo "  2. Run the API:         make run"
echo "  3. Run tests:           make test"
echo ""
