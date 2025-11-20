To set up a GitHub repo you do two things: create the repository on GitHub, then connect/push your local code to it using Git.[1][2]

## 1. Create repo on GitHub

1. Log in to https://github.com and click the “New” or “New repository” button from the Repositories page or top-right menu.[2][1]
2. Enter a repository name, optional description, and choose visibility (Public or Private).[3][1]
3. Optionally tick “Add a README” (recommended) and then click “Create repository”.[1][2]

Now you have an empty (or README-only) repo on GitHub.[1]

## 2. Initialize Git in your project

In your terminal, inside your project folder:

```bash
cd /path/to/your/project
git init
```

This creates a hidden `.git` directory and turns the folder into a Git repository.[4][5]

Then stage and commit your files:

```bash
git add .
git commit -m "Initial commit"
```

This records your first snapshot of the project history.[6][5]

## 3. Connect local repo to GitHub

From your new GitHub repo page, copy the HTTPS URL from the green “Code” button.[7][8]
Then in your project folder:

```bash
git remote add origin https://github.com/<your-username>/<repo-name>.git
git branch -M main        # only if your default branch isn't main yet
git push -u origin main
```

This ties your local repo to GitHub and uploads your code to the `main` branch.[8][7]

## 4. Create and push future changes

For everyday work after setup:

```bash
# after editing files
git add .
git commit -m "Describe what you changed"
git push
```

This updates the same GitHub repo with new commits on the current branch.
