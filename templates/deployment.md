# Deployment

Document on deployment practices and runbooks for this project. Updated whenever deployment procedure, target environment, or rollback strategy changes.

---

## Deployment Overview

<!--
One-sentence summary of what gets deployed and where.
Example:
- This is a Node.js backend service deployed to AWS Lambda
- This is a React frontend deployed to Vercel
- This is a Go CLI tool distributed via Homebrew and GitHub Releases
-->

---

## Release Triggers and Versioning

<!--
How is a release triggered? How is it versioned?
Example:
- Manual: Tag a commit with `v<major>.<minor>.<patch>` and push to origin; CI builds and publishes
- Automatic: Merge to `main` triggers a release with semantic versioning based on conventional commits
- Versioning scheme: Semantic Versioning (SemVer) for libraries, CalVer for applications
-->

---

## Deployment Target(s)

<!--
Where does code run in production?
Example:
- Production: `https://api.example.com` (AWS ECS, us-east-1)
- Staging: `https://staging-api.example.com` (AWS ECS, us-east-1)
- Development: Local development environment only

Or for CLI:
- macOS: Homebrew tap `example/tap/tool`
- Linux: GitHub Releases, apt repository
- Windows: GitHub Releases, Scoop bucket
-->

---

## Deployment Procedure

<!--
Step-by-step runbook for deploying a release.
Example:

### Prerequisites
- [ ] All tests pass locally and in CI
- [ ] Code reviewed and merged to `main`
- [ ] Tag pushed with `git tag -a v<version> -m "Release <version>" && git push origin v<version>`

### Deploy Steps
1. CI pipeline is triggered by the tag push
2. Build artifacts are created: `dist/` for frontend, Docker image for backend
3. Artifact is published to target registry (npm, Docker Hub, GitHub Releases)
4. If deployment to production is automatic:
   - ECS task definition is updated with new image tag
   - Service is updated and old tasks are drained gracefully (30s drain timeout)
   - Health checks pass before considering deployment complete
5. If manual approval is needed:
   - Ops team receives notification in [Slack/email/deployment dashboard]
   - Approval triggers the above steps

### Verification
- [ ] Health checks pass on target environment
- [ ] Smoke tests pass (e.g., `curl https://api.example.com/health`)
- [ ] Logs show no errors
- [ ] Key metrics (latency, error rate) are nominal
-->

---

## Rollback Procedure

<!--
How do you undo a bad deployment?
Example:
- Automatic rollback: If health checks fail, ECS automatically reverts to previous task definition
- Manual rollback: `git revert <commit>`, tag with `v<version>-hotfix.1`, push; CI redeploys
- Database migrations: Forward-only; data rollback requires [process]
- Feature flags: Bad behavior can be disabled without redeployment via [system]

For CLI releases:
- Yanked versions: Tag with `v<version>` and mark as yanked in release notes
- Users on old version: Keep supporting previous major version for [N] months
-->

---

## Deployment Checklist

<!--
Copy and use before each deployment:

- [ ] Code committed and pushed
- [ ] All tests pass in CI
- [ ] Code reviewed
- [ ] Changelog updated (docs/CHANGELOG.md or RELEASES.md)
- [ ] Version bumped and tagged
- [ ] Staging deployment succeeds
- [ ] Smoke tests pass on staging
- [ ] Approval given for production deployment
- [ ] Production deployment succeeds
- [ ] Health checks pass
- [ ] Metrics are nominal
- [ ] Announcement posted to [team channel]
-->

---

## Environment Configuration

<!--
How are environment variables, secrets, and configuration managed?
Example:
- Secrets are stored in [AWS Secrets Manager / HashiCorp Vault / GitHub Secrets]
- Environment variables are injected at deployment time from [source]
- Configuration file: `.env.production` (not checked in), managed by [process]
- Database connection string: Retrieved from [secrets manager] at startup
-->

---

## Monitoring and Alerts

<!--
What happens after code is deployed? How do you know if it's broken?
Example:
- Error rates are monitored in Datadog; alert if error rate > 1% for 5 min
- Latency p99 is tracked; alert if > [threshold]
- Database connection pool is monitored; alert if exhausted
- Disk space is monitored; alert if < 10% free
- Deployment notifications sent to [Slack channel]
-->

---

## Known Deployment Limitations or Risks

<!--
Document any deployment constraints or gotchas.
Example:
- Database migrations are applied separately; code must be backward-compatible
- Deployment is not atomic: old and new code may run simultaneously for [duration]
- Secrets rotation requires [manual step]
- Large deployments > 100MB take [N] minutes; monitor for timeout
-->
