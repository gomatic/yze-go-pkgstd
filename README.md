# yze-pkgstd

A [`yze`](https://github.com/gomatic/yze) analyzer (category `structure`) enforcing the **per-package** standards of the gomatic three-tier CLI layout. For a command package (`internal/app/commands/<cmd>/`) it checks:

- the first declaration of the command file (the one defining `Command()`) is a **const block** — conventionally the command's metadata (`name`, `usage`, `argUsage`, `description`), though the exact identifiers are not enforced because valid commands legitimately vary (e.g. omitting `argUsage` when there are no positional arguments);
- a **`Command()`** entry point exists;
- the **domain** package is imported under the **`domain`** alias.

Cross-package correspondence (every command has a matching `internal/domain/<cmd>`, and vice versa) is enforced separately by the [`yze-layout`](https://github.com/gomatic/yze-layout) analyzer.

- **Rule:** `yze/pkgstd`
- **Library:** exports `Analyzer` and `Registration` for the [`yze`](https://github.com/gomatic/yze) aggregator and [`stickler`](https://github.com/gomatic/stickler) runner.
- **Binary:** `cmd/yze-pkgstd` runs it standalone (`text`/`-json`, and as a `go vet -vettool`).

Built on the [`go-yze`](https://github.com/gomatic/go-yze) framework.
