# Validation Checklist

Use this checklist to validate the fresh 4dc prompt suite.

## Prompt Set

Expected generated files:
- `constitution.prompt.md`
- `increment.prompt.md`
- `plan.prompt.md`
- `implement.prompt.md`
- `promote.prompt.md`

Check command:

```bash
./templates/generate-all.sh
ls -1 *.prompt.md
```

## Phase Requirements

### constitution
- Produces `.4dc/constitution-review.html` before writing `CONSTITUTION.md`
- Defines engineering guardrails and SDLC artifact policy

### increment
- Produces `.4dc/increment-review.html` before writing `.4dc/increment.md`
- Stays at WHAT/WHY and avoids technical design detail

### plan
- Produces `.4dc/plan-review.html` before writing `.4dc/plan.md`
- Converts requirements to ordered, verifiable technical subtasks

### implement
- Produces `.4dc/implementation-review.html`
- Maintains `.4dc/implementation.md` and `.4dc/promote.md`
- Records objective verification evidence

### promote
- Produces `.4dc/promotion-report.html`
- Applies only approved updates to permanent artifacts
- Suggests emptying `.4dc/` after each promote
- Confirms `.4dc` cleanup decision

## Review-First Rule

For every phase, verify:
1. HTML report exists in `.4dc/`
2. Status is pending approval before final write
3. Final Markdown write occurs only after explicit approval

## Artifact Policy

Permanent artifacts to evaluate every cycle:
- OpenAPI specification
- ADRs
- Visual design guide
- Personas
- Deployment strategy
- Testing decisions
- Observability docs
- C4 diagrams (system, container, component)
- Roadmap (`docs/roadmap.md`)

Transient artifacts:
- `.4dc/increment.md`
- `.4dc/plan.md`
- `.4dc/implementation.md`
- `.4dc/promote.md`
- phase review HTML files

## Installer Validation

Run installer in a test repository and confirm:
- prompt files are copied with `4dc-` prefix
- `.4dc/` is created
- `.4dc` is present in `.gitignore`
