# Deployment

Guide to releasing and operating this project. Explain the deployment model, why it is appropriate, how to execute it safely, and how to recover. Update whenever the release or operational model changes.

Describe the release unit, the target, ownership, and the operational assumptions. Explain why this deployment shape is used.
- This is a Node.js backend service deployed to AWS Lambda
- This is a React frontend deployed to Vercel
- This is a Go CLI tool distributed via Homebrew and GitHub Releases
-->

Explain the release trigger and versioning decision, including who can release and what evidence is required first. Avoid a changelog or release-history list here.
- Manual: Tag a commit with `v<major>.<minor>.<patch>` and push to origin; CI builds and publishes
- Automatic: Merge to `main` triggers a release with semantic versioning based on conventional commits
- Versioning scheme: Semantic Versioning (SemVer) for libraries, CalVer for applications
-->
Describe the environments and their purpose, including the differences that matter for safe verification. Do not use this section as an environment inventory without explaining the deployment model.
Or for CLI:
- macOS: Homebrew tap `example/tap/tool`
- Linux: GitHub Releases, apt repository
- Windows: GitHub Releases, Scoop bucket
-->

---

## Deployment Procedure

Document the actual release procedure as a short runbook, with prerequisites, commands or links, verification signals, and ownership. Explain why the ordering protects users and data. Keep the checklist operational; put rationale in surrounding prose.
Describe rollback triggers, the recovery action, data implications, and who decides. Explain any forward-only or irreversible operation and the recovery alternative.

For CLI releases:
- Yanked versions: Tag with `v<version>` and mark as yanked in release notes
- Users on old version: Keep supporting previous major version for [N] months
-->

---

## Deployment Checklist

Keep only the small set of release decisions and checks that are specific to this project. Do not turn this into a repeated list of every test, deployment, or release ever performed.

Explain configuration ownership, secret handling, safe defaults, and the reason for separating deploy-time configuration from source code. Never record secret values.
- Environment variables are injected at deployment time from [source]
- Configuration file: `.env.production` (not checked in), managed by [process]
- Database connection string: Retrieved from [secrets manager] at startup
Explain how operators know a release is healthy, which signals matter, and what action an alert should trigger. Record thresholds only when they are real, justified, and maintained.
- Error rates are monitored in Datadog; alert if error rate > 1% for 5 min
- Latency p99 is tracked; alert if > [threshold]
- Database connection pool is monitored; alert if exhausted
- Disk space is monitored; alert if < 10% free
- Deployment notifications sent to [Slack channel]
-->

Document constraints that materially affect release safety and the mitigation or follow-up needed. Do not preserve obsolete procedures as historical reference.
- Deployment is not atomic: old and new code may run simultaneously for [duration]
- Secrets rotation requires [manual step]
- Large deployments > 100MB take [N] minutes; monitor for timeout
-->
