# AI Dev Tasks

A structured documentation system that separates project concerns across different roles and perspectives, enabling AI-assisted development with clear boundaries between strategy, requirements, architecture, and implementation.

## The System

This framework creates four types of documents that work together to guide development:

### 1. Constitution (WHY) - CTO/Architect View
**File:** `CONSTITUTION.md`

**Purpose:** Defines the foundational technical decisions, principles, and strategies that guide all development.

**Contains:**
- Core development principles (e.g., Speed of Delivery, Test Critical Paths Only)
- Technology stack decisions (languages, frameworks, tools)
- Architectural guidelines and constraints
- Development philosophy (testing, deployment, error handling)

**Role:** Sets the guardrails. All other documents must align with constitutional principles.

### 2. Features (WHAT) - Product Owner View
**Files:** `[feature-name]/feature.md`

**Purpose:** Describes individual features/capabilities from the user's perspective without technical implementation details.

**Contains:**
- User goals and context
- Assumptions being tested
- Success scenarios and alternative paths
- Acceptance criteria
- Success metrics and failure signals
- What's explicitly out of scope

**Role:** Defines what needs to be built and why it matters to users. No "how" — that's for the ADR.

### 3. ADR/Design (HOW) - Senior Developer View
**Files:** `[feature-name]/adr.md`

**Purpose:** Documents architectural decisions, trade-offs, and technical approach without prescribing exact implementation code.

**Contains:**
- Key architectural decisions with rationale
- Trade-offs and alternatives considered
- Constitutional alignment for each decision
- System architecture overview
- API contracts, data models, component responsibilities
- Open questions and risks

**Role:** Explains which technical approaches were chosen and why. Guides implementation without dictating code.

### 4. Tasks (HOW - Detailed) - Junior Developer View
**Files:** `[feature-name]/tasks.md`

**Purpose:** Breaks down the implementation into actionable, step-by-step tasks.

**Contains:**
- Parent tasks and detailed sub-tasks
- Relevant files to create/modify
- Specific commands to run
- Checkboxes for progress tracking

**Role:** Provides a clear roadmap for implementation. Can be followed by developers or AI agents.

## Document Hierarchy

```
Constitution (WHY - Strategic)
    ↓
Feature (WHAT - Requirements)
    ↓
ADR (HOW - Architecture)
    ↓
Tasks (HOW - Implementation)
```

## Workflow

1. **Create Constitution** - Define your project's technical foundation and principles (the broad foundation)
2. **Write Feature** - Describe the feature from the user's perspective
3. **Generate ADR** - Document architectural decisions that comply with the constitution and implement the feature
4. **Generate Tasks** - Break down implementation steps that follow the ADR within all given constraints
5. **Implement** - Follow the tasks, checking them off as you go

## Templates

Each document type has a template to guide AI generation:

- `create-constitution.md` - Guide for creating project constitutions
- `create-feature.md` - Guide for generating features
- `create-adr.md` - Guide for generating architecture decision records
- `create-tasks.md` - Guide for generating task lists

## Example

See the `example/` directory for a complete example:

- `example/CONSTITUTION.md` - EpicSum CSV time aggregation tool
- `example/sum-time-by-ticket/feature.md` - Time aggregation feature
- `example/sum-time-by-ticket/adr.md` - Architectural decisions for the feature
- `example/sum-time-by-ticket/tasks.md` - Implementation task breakdown

## Key Principles

**Separation of Concerns:**
- Features capture WHAT (user goals, acceptance criteria)
- ADRs capture HOW (technical decisions, architecture)
- Constitution captures WHY (principles, strategic choices)
- Tasks capture implementation steps (actionable work)

**Constitutional Alignment:**
- Constitution establishes the broad technical foundation
- Features describe user needs while respecting constitutional constraints
- ADRs make technical decisions that comply with the constitution and implement the feature
- Tasks execute the feature implementation following ADR guidance within all constraints (constitution, feature, ADR)