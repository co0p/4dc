const vscode = require('vscode');
const path = require('path');

function activate(context) {
  const root = vscode.workspace.workspaceFolders?.[0]?.uri.fsPath;
  if (!root) return;

  function openPrompt(relPath) {
    const full = path.join(root, '.prompt', relPath);
    vscode.workspace.openTextDocument(full).then(doc => vscode.window.showTextDocument(doc));
  }

  async function copyPromptToChat(relPath) {
    const full = path.join(root, '.prompt', relPath);
    try {
      const doc = await vscode.workspace.openTextDocument(full);
      const text = doc.getText();
      await vscode.env.clipboard.writeText(text);
      // Try to open the Chat view; this command exists in newer VS Code builds.
      try {
        await vscode.commands.executeCommand('workbench.action.openChat');
      } catch (e) {
        // best-effort: ignore if not available
      }
      vscode.window.showInformationMessage('Prompt copied to clipboard. Paste into Chat (Cmd+V) or open a new chat and paste.');
    } catch (err) {
      vscode.window.showErrorMessage('Could not copy prompt: ' + err.message);
    }
  }

  context.subscriptions.push(
    vscode.commands.registerCommand('4dc.open.createConstitution', () => openPrompt('create-constitution.prompt.md')),
    vscode.commands.registerCommand('4dc.open.designPrompt', () => openPrompt('design.prompt.md')),
    vscode.commands.registerCommand('4dc.open.implementPrompt', () => openPrompt('implement.prompt.md')),
    vscode.commands.registerCommand('4dc.open.incrementPrompt', () => openPrompt('increment.prompt.md')),
    vscode.commands.registerCommand('4dc.open.improvePrompt', () => openPrompt('improve.prompt.md')),

    vscode.commands.registerCommand('4dc.prompts.copyToChat.createConstitution', () => copyPromptToChat('create-constitution.prompt.md')),
    vscode.commands.registerCommand('4dc.prompts.copyToChat.implementPrompt', () => copyPromptToChat('implement.prompt.md'))
  );
}

function deactivate() {}

module.exports = { activate, deactivate };