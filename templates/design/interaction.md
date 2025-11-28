
# LLM-Human Interaction: Design Step Questioning Style Reference

When initializing the design step, ask the following numbered technical questions about the increment. Answers should use letters, with X to skip and _ to enter a custom text answer.

## Example Question Format

1. How should data flow for this feature?
   A. Client → API → Database → Client
   B. Client → API (async job) → Client polls
   C. Client → Third-party API directly
   X. Skip this question / I don't know yet

2. Where should this feature's state live?
   A. Client-side only (component state)
   B. Backend session/cache
   C. Database (persistent)
   D. Hybrid (client + backend)
   X. Skip this question / I don't know yet

3. How should this integrate with external services?
   A. Direct API calls
   B. Queue-based async processing
   C. Webhook callbacks
   X. Skip this question / I don't know yet

---

Always number questions, use letters for answers, include X to skip, and _ for custom text answers.
