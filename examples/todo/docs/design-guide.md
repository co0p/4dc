# Visual Design Guide

## Purpose

Define the durable visual rules used by the todo example so future increments keep a consistent brutalism-like direction.

## Direction

- Intense contrast between surfaces and text.
- Sharp edges and strong border lines.
- Muted, grounded palette instead of highly saturated accents.
- Minimal decorative motion.

## Core Tokens

From `styles.css`:

- Background surface: `--bg`
- Primary text: `--ink`
- Border line: `--border`
- Accent text: `--accent`
- Focus ring: `--focus`
- Error colors: `--error-bg`, `--error-ink`

## Component Rules

- Panels and rows use explicit 2px borders.
- Containers may use offset hard shadows for depth.
- State/meta text should remain high contrast.
- Error containers use high-contrast border and fill.

## Form Input Patterns

Form elements follow the same brutalist direction:

- Text inputs: 2px solid border, white background, high-contrast text
- Submit buttons: accent background, white text, explicit hover/active states
- Focus states: 3px outline with offset for keyboard visibility
- Error display: high-contrast error colors with explicit border
- Visually hidden labels: Use SR-only pattern for minimalist forms while maintaining accessibility

Form layout should use flex with gap for responsive stacking on narrow viewports.

## Typography and Hierarchy

- Use strong weight shifts for headings and labels.
- Keep labels and state text clear at mobile widths.
- Prefer explicit section separation over subtle cards.

## Accessibility Baseline

- Keep visible keyboard focus using `:focus-visible` outline.
- Preserve readable contrast for body, metadata, and error text.
- Keep row spacing and touch targets usable on mobile.

## Responsive Rule

- Mobile-first layout.
- At wider viewports, constrain content width for readability rather than stretching full width.
