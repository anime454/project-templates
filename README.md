## Commit message convention

This repo follows **Conventional Commits** to keep history readable and make changelogs/releases easier.

### Format

```text
<type>(<scope>): <short summary>

<body: what/why/how (optional, but recommended for non-trivial changes)>

<footer: references/issues/breaking changes (optional)>
```

- **type**: what kind of change this is (required)
- **scope**: area/module affected (optional, but recommended)
- **short summary**: imperative, present tense (required)
- **body**: explain context, motivation, and how to test (optional)
- **footer**: references like issue IDs, or breaking changes (optional)

### Types

Use one of these types:

- `feat`: new feature / new template capability
- `fix`: bug fix
- `chore`: tooling, housekeeping, dependency bumps (no functional change)
- `refactor`: code change that neither fixes a bug nor adds a feature
- `docs`: documentation only
- `test`: add/update tests
- `ci`: CI/CD changes
- `perf`: performance improvements
- `build`: build system changes (Makefile, build scripts)

### Type details & examples

#### `feat` — new feature
Use when you add **new functionality** or a new template capability.

Examples:
- `feat(hexagonal): add health check endpoint template`
- `feat(hooks): add root pre-commit runner for templates`
- `feat(car-parking-system): add docker compose for local postgres`

Example with body:

```text
feat(hooks): add root pre-commit runner for templates

Runs template hooks once per folder based on staged paths.
How to test: stage a file under go/hexagonal/car-parking-system and run .githooks/pre-commit.
```

#### `fix` — bug fix
Use when you **correct behavior that was wrong**.

Examples:
- `fix(hooks): handle filenames with spaces in staged paths`
- `fix(make): use 'makefile' when Makefile is lowercase`
- `fix(car-parking-system): correct module path in go.mod`

#### `chore` — maintenance
Use for housekeeping changes that don’t add features or fix bugs.

Examples:
- `chore(deps): bump Go version in templates`
- `chore(make): simplify local dev targets`

#### `refactor` — restructure (no behavior change)
Use when you restructure code/scripts without changing behavior.

Examples:
- `refactor(hooks): extract run_hook_if_exists into helper`
- `refactor(car-parking-system): reorganize internal packages`

#### `docs` — documentation only
Use when you change only docs.

Examples:
- `docs(readme): document commit message rules`
- `docs(car-parking-system): add local setup steps`

#### `test` — tests only
Use when you add/update tests only.

Examples:
- `test(car-parking-system): add handler tests for parking flow`
- `test: add hook runner tests`

#### `ci` — CI/CD changes
Use for changes to GitHub Actions/workflows/pipelines.

Examples:
- `ci: add golangci-lint workflow`
- `ci: run hooks on pull requests`

#### `perf` — performance
Use when you improve performance.

Examples:
- `perf(hooks): skip scan when no staged files`
- `perf(make): parallelize lint targets`

#### `build` — build system
Use for changes to build tooling (Makefile, build scripts, docker build).

Examples:
- `build(make): add pre-commit target for lint and test`
- `build(docker): update build args`

### Scopes (examples)

Scopes are optional, but recommended. Examples:
- `hooks`, `make`, `docs`, `hexagonal`, `car-parking-system`, `docker`, `config`

### Notes

- Keep the subject line under ~72 characters when possible.
- Use the imperative mood (e.g., "add", "fix", "update").
- For breaking changes, add a footer like:

```text
BREAKING CHANGE: <describe what changed and migration steps>
```