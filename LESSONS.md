# shomer lesson plan

Sessions are about 90 minutes. Each lesson ends when its **Done when** check passes. Stop there,
even if there's time left. Stretch goals are optional and never required for the next lesson.

Concept tags: **[Go]** language and idioms · **[CLI]** Unix/command-line conventions (any
language) · **[Security]** the reason shomer exists.

**Session ritual**
- *Start:* read the current lesson below and run the last lesson's Done-when check.
- *End:* tick the lesson, then update "Current lesson" here and "Current step" in CLAUDE.md.

**Current lesson: 2**

---

## Stage 1: Dependency scan (SCA)

### Lesson 1: Read the lockfile safely ✅ *(done 2026-09-30)*
**Goal:** shomer takes a repo path and reads its `package-lock.json`, failing cleanly if it can't.

**Concepts**
- [Go] `go mod init`, `go.mod`, `package main`, `func main()`, grouped imports, `go fmt`
- [Go] `os.Args`, slices, `len()`, index-out-of-range panics
- [Go] `path/filepath.Join`, `os.ReadFile`, multiple return values, `if err != nil`
- [CLI] exit codes (0 or non-zero), `os.Exit`, stdout vs stderr, `usage:` messages
- [CLI] the shell expands `~` before your program sees it; `go run` hides your exit code
- [Security] input validation; don't leak internals in error messages

**Outcome:** reads the lockfile and prints 409272 bytes for manifest (matches `wc -c`); missing
lockfile prints Go's real error to stderr and exits 1. Caught and fixed a fail-open bug (error
printed but exit 0).

**Done when**
- `shmr ~/Desktop/manifest` prints a byte count and exits 0
- `shmr ~/Desktop` (no lockfile) prints an error to stderr and exits 1

**Stretch:** change the error message to `usage: shomer <path-to-repo>`; move the banner off stdout.

### Lesson 2: Parse the lockfile
**Goal:** list every package in the lockfile with its version.

**Concepts**
- [Go] structs, struct tags (`json:"version"`), `encoding/json.Unmarshal`, maps, `for range`
- [Go] pointers, lightly (why `Unmarshal` takes `&x`)
- [Security] what a lockfile pins; direct vs transitive dependencies; why the exact version matters

**Done when:** `shmr ~/Desktop/manifest` prints one `name@version` per line plus a total count,
with `node_modules/` stripped from the names.

**Stretch:** sort the output (`sort` / `slices` package).

### Lesson 3: Ask OSV about the packages
**Goal:** send every package to OSV.dev and get back the IDs of known vulnerabilities.

**Concepts**
- [Go] building request structs, `json.Marshal`, `net/http` POST, `defer resp.Body.Close()`,
  checking the status code
- [Go] batching a slice (querybatch accepts a limited number of queries per request)
- [Go] move lockfile parsing into its own package (`internal/lockfile`); exported names are
  capitalized; `main` only wires the pieces together
- [Security] CVE vs GHSA IDs; what a vulnerability database is; ecosystems (`npm`)

**Done when:** shomer prints each affected package followed by its vulnerability IDs
(e.g. `GHSA-xxxx-xxxx-xxxx`), and prints nothing for clean packages.

**Stretch:** add a timeout to the HTTP client (`http.Client{Timeout: ...}`), and understand why.

### Lesson 4: Enrich the results
**Goal:** turn the IDs into a readable report: package, version, ID, severity, fixed version.

**Concepts**
- [Go] a second API call per ID (`GET /v1/vulns/{id}`), nested structs, handling missing fields
- [Go] `text/tabwriter` or `fmt` width verbs for aligned output
- [Security] CVSS scores and severity levels; affected ranges vs fixed versions

**Done when:** `shmr ~/Desktop/manifest` prints an aligned table with all five columns.

**Stretch:** skip duplicate IDs across packages with a `map[string]bool` "seen" set.

### Lesson 5: Check the answer key, and set the exit code
**Goal:** trust shomer's results, and let a script or CI act on them.

**Concepts**
- [Go] exiting non-zero when findings exist; `go install` so `shomer` runs from any folder
- [Go] a first `go test` for the lockfile parser
- [Security] false positives and false negatives; why you check a tool against an answer key

**Done when:** shomer's findings on manifest and pcx-pipeline match `npm audit` (or every
difference is explained in writing), and `shomer <repo>; echo $?` returns 1 when vulns are found.

---

## Stages 2–6
Broken into lessons when we reach each one (see the Roadmap in CLAUDE.md).

Stage 2 is shomer's second job, so it starts with **subcommands** (`shomer deps <path>`,
`shomer supply-chain <path>`): a subcommand picks the job, a flag (`--json`) adjusts it.
