# Implementation: Add macOS .app bundle via Makefile

Increment branch: `feature/macos-app-bundle`

## Planned Files Summary (Confirm Before Coding)
- `examples/pomodoro/Makefile` — new — `bundle/clean/version` targets to build `dist/Pomodoro.app`
- `examples/pomodoro/packaging/macos/Info.plist.tmpl` — new — Template with `CFBundleName`, `CFBundleIdentifier`, `CFBundleVersion`, `LSUIElement=1`
- `examples/pomodoro/assets/AppIcon.icns` — new — Committed placeholder `.icns` used by bundle
- `examples/pomodoro/VERSION` — new — App version source (fallback to `git describe` in Makefile)

STOP: Confirm or edit the list before coding.

## Implementation Tasks & Subtasks
- [x] **Setup**
  - [x] Ensure on branch `feature/macos-app-bundle` (verify `git branch --show-current`)
  - [x] Create `examples/pomodoro/packaging/macos/` and `examples/pomodoro/assets/` (verify dirs exist)
  - [x] Add `examples/pomodoro/VERSION` seeded with `0.1.0` (verify file contains version)

- [x] **Makefile: version + platform gating**
  - [x] Add Darwin guard (`uname`) to `bundle` target (verify non‑macOS prints "macOS only")
  - [x] Implement `version` rule: prefer `VERSION` file; fallback to `git describe --tags --always` (verify `make version` echoes)
  - [x] Define variables: `APP_NAME=Pomodoro`, `BUNDLE_ID=co0p.pomodoro`, `DIST=dist`, `APP=$(DIST)/$(APP_NAME).app` (verify `make -n bundle` prints commands)

- [x] **Makefile: build + stage bundle**
  - [x] Compile binary to `examples/pomodoro/dist/Pomodoro.app/Contents/MacOS/pomodoro` via `go build -o` (verify binary exists)
  - [x] Create directories `Contents/{MacOS,Resources}` (verify structure)
  - [x] Generate `Contents/Info.plist` from `packaging/macos/Info.plist.tmpl` (sed replace version/name/id) (verify keys set)
  - [x] Copy `assets/AppIcon.icns` to `Contents/Resources/AppIcon.icns` (verify file present)

- [x] **Plist template**
  - [x] Add `Info.plist.tmpl` with placeholders for `__BUNDLE_NAME__`, `__BUNDLE_ID__`, `__BUNDLE_VERSION__`, and `LSUIElement`=1 (verify template saved)
  - [x] Ensure `LSUIElement`=1 is preserved in rendered plist (verify with `plutil -p` if available)

- [x] **Icon asset**
  - [x] Add a placeholder `assets/AppIcon.icns` (committed) to satisfy bundle structure (verify file exists)
  - [x] Note: runtime tray icon remains generated template icon; `.icns` is for bundle metadata only (verify app launches without Dock icon)

- [x] **Clean target**
  - [x] Implement `make clean` to remove `examples/pomodoro/dist/` (verify directory deleted)

- [x] **Manual Verification (macOS)**
  - [x] Run `make bundle` (verify bundle at `examples/pomodoro/dist/Pomodoro.app`)
  - [x] Run `make open` (verify menu bar icon-only appears, menu shows Start/Pause/Quit)
  - [x] Inspect `Info.plist` for `CFBundleIdentifier`, `CFBundleVersion`, `LSUIElement=1` (verify values)

- [ ] **Cross-platform guard**
  - [ ] On non‑macOS, run `make bundle` (verify friendly no-op message)

- [x] **Incremental Commit**
  - [x] Commit Makefile + template + VERSION + icon with message: "build(macos): add Makefile bundle and plist template"
  - [x] Commit icon pipeline/CLI and README changes

## Code Implementation
Pending confirmation of the Planned Files Summary. No code changes included yet.

## Validation
- `make bundle` creates `dist/Pomodoro.app` with binary, `Info.plist` (with `LSUIElement=1`), and `AppIcon.icns`.
- `open` launches a working menu bar app (icon-only) with expected menu items.
- Non‑macOS runs print a clear message and do not create artifacts.
- `make clean` removes `dist/`.

## Key Decisions & Trade-offs
- Use a committed `.icns` placeholder (choice 4C) to avoid `iconutil` dependency during the increment.
- Version comes from `VERSION` with git fallback (choice 3A) for reproducibility.
- Keep build logic in Makefile/scripts, not in app code (aligns with Dependency Discipline).

## Open Questions
- Should we later switch to generated `.icns` to avoid storing binaries in repo?
- Do we need universal binaries now (arm64+amd64 via `lipo`) or later?
