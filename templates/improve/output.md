# Improve Output Format

The improve output must be concise and parsable by the /increment prompt.

## Output Schema: Improve.md

### 1. Assessment
- **Constitution Alignment:** [Brief evaluation]
- **Design Alignment:** [Brief evaluation]
- **Quality:** [Brief evaluation]
- **Risks:** [List]

### 2. Lessons
- **Worked Well:** [List]
- **To Improve:** [List]
- **Emerging Patterns:** [List]

### 3. ADRs Created
When the user approves an ADR candidate during the improve process, create the ADR file immediately using the ADR Output Template. List only the created ADR files here:
- `design/adr-[number]-[name].md` - [Brief description]

### 4. Improvements
Each improvement is a separate section with explicit file references. Format for parsability by /increment:

#### Improvement 1: [Title]
- **Lens:** [Naming/Modularity/Architecture/Testing/Duplication/Documentation]
- **Priority:** [H/M/L]
- **Effort:** [X min]
- **Files:** `path/to/file.ext`
- **Change:** [Specific change description]

#### Improvement 2: [Title]
- **Lens:** [...]
- **Priority:** [...]
- **Effort:** [...]
- **Files:** `path/to/file.ext`
- **Change:** [...]

---

## Example

```markdown
# Improve: Add Todo Item

## 1. Assessment
- **Constitution Alignment:** Adheres to browser-native principles
- **Design Alignment:** Follows event-driven architecture
- **Quality:** Good readability, some duplication
- **Risks:** No error handling for localStorage failures

## 2. Lessons
- **Worked Well:** Event delegation pattern, localStorage abstraction
- **To Improve:** Validation logic scattered across components
- **Emerging Patterns:** Read-Modify-Save-Render cycle for state

## 3. ADRs Created
- `design/adr-001-state-mutation-pattern.md` - Standardizes state mutation cycle

## 4. Improvements

#### Improvement 1: Extract validation helper
- **Lens:** Modularity & Separation
- **Priority:** H
- **Effort:** 15 min
- **Files:** `src/utils/validation.js`, `src/components/TodoForm.js`
- **Change:** Move duplicate validation logic to shared helper

#### Improvement 2: Add localStorage error handling
- **Lens:** Testing & Reliability
- **Priority:** M
- **Effort:** 20 min
- **Files:** `src/storage.js`
- **Change:** Wrap localStorage calls in try-catch with fallback
```
