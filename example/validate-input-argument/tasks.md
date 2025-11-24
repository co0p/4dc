# Tasks: Validate Input Argument

## Relevant Files

- `epicsum.sh` - Main bash script with argument validation logic

**Note:** Single script file keeps it simple per constitution's "Simplicity Over Flexibility" principle.

## Progress Tracking

**As you complete each task, check it off by changing `- [ ]` to `- [x]`.**

## Implementation Tasks

- [x] **Setup**
  - [x] Create feature branch: `git checkout -b feature/validate-input-argument`

- [x] **Argument validation check**
  - [x] Add check at script start: if `$1` is empty, print "please specify an input csv" to stderr and exit 1
  - [x] Verify: Running `./epicsum.sh` with no arguments shows error and exits with code 1

- [x] **File existence check**
  - [x] Add check after argument validation: if file at `$1` doesn't exist, print error to stderr and exit 1
  - [x] Verify: Running `./epicsum.sh nonexistent.csv` shows file not found error and exits with code 1

- [x] **Success path verification**
  - [x] Verify: Running `./epicsum.sh input.csv` (with valid file) proceeds without validation errors
  - [x] Verify: Exit code is 0 when validation passes (script continues to main logic)

- [x] **Integration & Verification**
  - [x] Test all acceptance criteria:
    - No argument → error message displayed
    - Nonexistent file → error message displayed
    - Valid file → no error, script continues
  - [x] Verify error messages are clear and actionable

- [x] **Deploy**
  - [x] Commit: "Add input argument validation"
  - [x] Push to branch (or merge to main per constitution's direct deployment approach)
