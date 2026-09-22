# AGENTS.md

## Rules

These rules are mandatory. Follow them unless the user explicitly gives different instructions.

## 1. Dependencies and Versions

- Always prefer the `latest` version when adding or updating dependencies.
- Do not intentionally pin dependencies to old versions.
- For GitHub Actions, use `@main` or `@master`.
- Do not use commit SHA pins for GitHub Actions unless explicitly requested.
- Do not add unnecessary dependencies.
- Do not upgrade unrelated dependencies.
- Do not modify dependency versions unless required by the task.

## 2. Code Style

- Keep the implementation simple and concise.
- Keep variable names short, clear, and meaningful.
- Keep function names short, clear, and meaningful.
- Do not use unnecessarily verbose names.
- Do not add comments unless they are genuinely necessary to explain non-obvious behavior.
- Do not add comments that merely restate what the code does.
- Avoid unnecessary abstractions.
- Avoid unnecessary helper functions.
- Avoid over-engineering.
- Follow the existing project conventions when they are more specific than these rules.

## 3. Scope of Changes

- Make the smallest change that fully solves the task.
- Do not modify unrelated files.
- Do not refactor unrelated code.
- Do not change formatting in unrelated files.
- Do not rename files, variables, functions, or modules unless required by the task.
- Do not change public APIs unless required by the task.
- Do not introduce new architecture unless required by the task.
- Do not silently change behavior outside the requested scope.

## 4. pnpm Projects

If the project uses `pnpm`, use `pnpm` commands.

### Development

Use:

```bash
pnpm dev
```

Do not use:

```bash
pnpm build
```

for normal local development verification.

`pnpm build` must only be run when the user explicitly asks for a production build or the task specifically requires build verification.

### Checks

After making code changes, run:

```bash
pnpm check
```

### Formatting

After making code changes, run:

```bash
pnpm format
```

Formatting should happen before the final review of the changes.

### Required pnpm workflow

For a normal code change, follow this order:

```bash
pnpm dev
pnpm check
pnpm format
```

If a command fails, investigate and fix the issue before considering the task complete.

Do not skip `pnpm check` or `pnpm format` unless the user explicitly asks you to.

## 5. Local Verification

- Always inspect the result locally after making changes.
- Do not assume the code works based only on static inspection.
- For `pnpm` projects, use `pnpm dev` for local verification.
- Check the actual result of the change before finishing.
- If verification cannot be completed, clearly state what could not be verified.

## 6. Git

### No Automatic Commits

- Never create a Git commit automatically.
- Never run `git commit` unless the user explicitly requests a commit.
- Never run `git push` unless the user explicitly requests it.
- Leave changes in the working tree after completing the task.
- The user must be able to review the changes before committing.

### Before Commit

Only when the user explicitly requests a commit:

1. Review the complete diff.
2. Run the required checks.
3. Run formatting.
4. Update the `version` field in `package.json`.
5. Increment the version by exactly `0.0.1`.
6. Review the diff again.
7. Create the commit only after the above steps are complete.

Example:

```json
{
  "version": "1.2.3"
}
```

must become:

```json
{
  "version": "1.2.4"
}
```

Do not change the version by `major`, `minor`, or any other increment.

### Commit Scope

- Do not include unrelated changes in a commit.
- Do not revert user changes.
- Do not reset the working tree unless explicitly requested.
- Do not use destructive Git commands unless explicitly requested.
- Preserve existing user modifications.

## 7. Existing User Changes

Before modifying files:

- Check the current working tree state.
- Treat existing uncommitted changes as user-owned.
- Do not overwrite, discard, reset, or revert existing user changes.
- Keep task-related changes separate from unrelated existing changes whenever possible.
- Do not assume existing modifications were created by you.

Never use destructive commands such as:

```bash
git reset --hard
git checkout -- .
git restore .
```

unless the user explicitly requests the operation.

## 8. Package Manager

- Detect the package manager from the repository.
- If the project uses `pnpm`, use `pnpm`.
- Do not switch package managers.
- Do not generate a different lockfile.
- Do not run `npm install`, `yarn install`, or equivalent commands in a pnpm project.

## 9. Lockfiles

- Preserve the existing lockfile format.
- Do not delete lockfiles.
- Do not regenerate lockfiles unnecessarily.
- Only modify the lockfile when a dependency change actually requires it.

## 10. Final Review

Before finishing a task:

- Review the changed files.
- Review the Git diff.
- Confirm there are no unrelated modifications.
- Confirm the requested behavior works locally.
- Confirm `pnpm check` passes for pnpm projects.
- Confirm `pnpm format` has been run for pnpm projects.
- Confirm no commit was created unless explicitly requested.

## 11. Priority

When instructions conflict, use this priority:

1. Explicit instructions from the user for the current task.
2. This `AGENTS.md`.
3. Existing project conventions.
4. General coding conventions.

Never invent requirements that are not necessary for the task.

## 12. Default Behavior

Unless the user explicitly says otherwise:

- Use latest versions.
- Use `@main` or `@master` for GitHub Actions.
- Keep code concise.
- Minimize comments.
- Keep names short and clear.
- Make minimal changes.
- Verify locally.
- Run `pnpm check`.
- Run `pnpm format`.
- Do not build with `pnpm build` for normal verification.
- Do not commit.
- Do not push.
- Only bump `package.json` version by `0.0.1` when a commit has been explicitly requested.
