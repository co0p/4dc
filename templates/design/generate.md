
---
name: generate-design-prompt
argument-hint: none
---


# Prompt: Generate Self-Contained design.prompt.md

You are tasked with generating a fully self-contained `design.prompt.md` file for the 4DC project. Use the contents of all files in the `/templates/design` folder (including persona.md, process.md, interaction.md, output.md, prompt.md) as source material.

- Merge all relevant details from these files into a single prompt file.
- Do not reference or link to any external files or folders in the output.
- The resulting prompt must include the persona, process, interaction style, and output format directly in the file.
- Ensure clarity, completeness, and maintainability.
- The output should be ready to use as a standalone prompt for generating a technical design for an increment.

# Output
A single, self-contained `design.prompt.md` file containing all necessary instructions, context, and templates for design generation. The file should be created at the top level of the repository (e.g., `/design.prompt.md`).