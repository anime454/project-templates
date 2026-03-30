@echo off
:: setup.bat – Bootstrap the development environment on Windows.
:: Run this script from the repository root in an elevated Command Prompt or
:: PowerShell (run as Administrator) for best results.

setlocal EnableDelayedExpansion

:: ---------------------------------------------------------------------------
:: Helpers
:: ---------------------------------------------------------------------------
set "INFO=[setup]"
set "OK=[setup] OK"
set "WARN=[setup] WARN"
set "ERR=[setup] ERROR"

:: ---------------------------------------------------------------------------
:: Check for winget (Windows Package Manager)
:: ---------------------------------------------------------------------------
where winget >nul 2>&1
if %ERRORLEVEL% neq 0 (
    echo %ERR% winget is not available.
    echo        Install the App Installer from the Microsoft Store:
    echo        https://apps.microsoft.com/detail/9nblggh4nns1
    pause
    exit /b 1
)
echo %OK% winget is available.

:: ---------------------------------------------------------------------------
:: Git
:: ---------------------------------------------------------------------------
where git >nul 2>&1
if %ERRORLEVEL% neq 0 (
    echo %WARN% Git not found. Installing via winget...
    winget install --id Git.Git -e --source winget
    if %ERRORLEVEL% neq 0 (
        echo %ERR% Failed to install Git. Download it from https://git-scm.com/downloads
        pause
        exit /b 1
    )
    :: Refresh PATH for current session
    set "PATH=%PATH%;C:\Program Files\Git\bin"
)
for /f "tokens=*" %%v in ('git --version 2^>nul') do echo %OK% %%v

:: ---------------------------------------------------------------------------
:: Go
:: ---------------------------------------------------------------------------
where go >nul 2>&1
if %ERRORLEVEL% neq 0 (
    echo %WARN% Go not found. Installing via winget...
    winget install --id GoLang.Go -e --source winget
    if %ERRORLEVEL% neq 0 (
        echo %ERR% Failed to install Go. Download it from https://go.dev/dl/
        pause
        exit /b 1
    )
    :: Refresh PATH for current session
    set "PATH=%PATH%;C:\Program Files\Go\bin"
)
for /f "tokens=*" %%v in ('go version 2^>nul') do echo %OK% %%v

:: ---------------------------------------------------------------------------
:: Docker Desktop
:: ---------------------------------------------------------------------------
where docker >nul 2>&1
if %ERRORLEVEL% neq 0 (
    echo %WARN% Docker not found. It is required to run the database locally.
    set /p INSTALL_DOCKER="Install Docker Desktop now? [y/N]: "
    if /i "!INSTALL_DOCKER!"=="y" (
        echo %INFO% Installing Docker Desktop via winget...
        winget install --id Docker.DockerDesktop -e --source winget
        if %ERRORLEVEL% neq 0 (
            echo %ERR% Failed to install Docker Desktop.
            echo        Download it from https://www.docker.com/products/docker-desktop/
        ) else (
            echo %WARN% Start Docker Desktop manually before running containers.
        )
    ) else (
        echo %WARN% Skipping Docker installation. You will need it to start PostgreSQL.
    )
) else (
    for /f "tokens=1-3" %%a in ('docker --version 2^>nul') do echo %OK% %%a %%b %%c
)

:: ---------------------------------------------------------------------------
:: Make (via winget – GnuWin32 or Git for Windows SDK)
:: ---------------------------------------------------------------------------
where make >nul 2>&1
if %ERRORLEVEL% neq 0 (
    echo %WARN% make not found. Installing via winget (GnuWin32)...
    winget install --id GnuWin32.Make -e --source winget
    if %ERRORLEVEL% neq 0 (
        echo %WARN% Could not install make automatically.
        echo        Alternatively, you can use the Git Bash terminal which includes make,
        echo        or install it via Chocolatey: choco install make
    ) else (
        :: GnuWin32 installs to C:\Program Files (x86)\GnuWin32\bin
        set "PATH=%PATH%;C:\Program Files (x86)\GnuWin32\bin"
    )
)
where make >nul 2>&1
if %ERRORLEVEL% equ 0 (
    for /f "tokens=1-3" %%a in ('make --version 2^>nul ^| findstr /i "make"') do echo %OK% %%a %%b %%c
)

:: ---------------------------------------------------------------------------
:: golangci-lint
:: ---------------------------------------------------------------------------
where golangci-lint >nul 2>&1
if %ERRORLEVEL% neq 0 (
    echo %WARN% golangci-lint not found. Installing...
    winget install --id golangci.golangci-lint -e --source winget 2>nul
    if %ERRORLEVEL% neq 0 (
        echo %INFO% Falling back to go install...
        go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
    )
)
where golangci-lint >nul 2>&1
if %ERRORLEVEL% equ 0 (
    for /f "tokens=*" %%v in ('golangci-lint --version 2^>nul ^| findstr /i "golangci"') do echo %OK% %%v
)

:: ---------------------------------------------------------------------------
:: Configure Git hooks
:: ---------------------------------------------------------------------------
echo %INFO% Configuring Git hooks...
git config core.hooksPath .githooks
if %ERRORLEVEL% neq 0 (
    echo %ERR% Failed to configure Git hooks.
) else (
    echo %OK% Git hooks configured ^(core.hooksPath=.githooks^)
)

:: ---------------------------------------------------------------------------
:: Download Go module dependencies
:: ---------------------------------------------------------------------------
echo %INFO% Downloading Go module dependencies...

set "REPO_ROOT=%~dp0.."
set "GO_PROJECTS=go\hexagonal\car-parking-system"

for %%p in (%GO_PROJECTS%) do (
    set "PROJECT_PATH=%REPO_ROOT%\%%p"
    if exist "!PROJECT_PATH!\go.mod" (
        echo %INFO%   -> %%p
        pushd "!PROJECT_PATH!"
        go mod download
        if %ERRORLEVEL% equ 0 (
            echo %OK%   %%p dependencies downloaded
        ) else (
            echo %ERR%  Failed to download dependencies for %%p
        )
        popd
    )
)

:: ---------------------------------------------------------------------------
:: Done
:: ---------------------------------------------------------------------------
echo.
echo %OK% Setup complete!
echo.
echo   Next steps:
echo   1. Start the database:  cd go\hexagonal\car-parking-system ^& docker compose up -d
echo   2. Run the API:         make run
echo   3. Run tests:           make test
echo.
pause
endlocal
