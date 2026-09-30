# shomer

A Go command-line tool that scans my own repos for security problems. It runs alongside
[Semgrep](https://semgrep.dev) and covers what Semgrep doesn't: vulnerable dependencies,
npm supply-chain risk, secrets in git history, GitHub Actions misconfigurations and repo hygiene.

## The name

*shomer* (Hebrew שֹׁמֵר) means "watchman" or "keeper": "He who keeps Israel will neither
slumber nor sleep" (Psalm 121:4). The short name `shmr` is the Hebrew root itself (ש־מ־ר,
sh-m-r). Hebrew builds words on consonant roots, and the vowels are added to form each word.

## Status

A practice project for learning Go and application security by building the checks myself.
Currently on stage 1, the dependency scan: reading `package-lock.json` and checking packages
against [OSV.dev](https://osv.dev). See [LESSONS.md](LESSONS.md) for the lesson plan.

## Usage

```bash
go build
./shomer <path-to-repo>
```

## Acknowledgment

I wrote the code myself. I used Claude as a tutor to learn Go and the security concepts behind
each check.
