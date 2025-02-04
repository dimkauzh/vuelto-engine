# Contributing rules

We are happy and welcome if you want to contribute to Vuelto. But please consider a few details before continuing:

1. Branch: Please when working on your fork, work in the dev branch, because we only will accept commits to the dev branch. It will later be released with the next version of Vuelto.
2. Explain: Please explain why this should be considered and merged. That will make our life easier.
3. Test: Please test your code before even opening a new pull request.
4. Documentation: Please, if you’re adding something new, like a feature, please document everything.
5. Format: Please, run `make format` for formatting of the code.

Not following these rules
If we see a pull request that doesn't follow these rules, we will tell you that, and close the pull request.
We allow you to re-open a new pull request, but we expect you to have your code fixed.
So make sure that you followed [the rules](#contributing-rules)

Some technologies we are using

- GLFW: We use this for the desktop windowing
- “syscall/js”: We use this to interact with the JS runtime
- glow: Generate OpenGL Go bindings (only happened one)
- CGo: Interface to use C with Go

## How to setup the work environment

If you want to contribute, you have to set up the work environment, so you can develop vuelto the right way.

- Run `go mod tidy` to install all the packages.
- Fork [the repository](https://github.com/vuelto-org/vuelto).
- Clone your forked github repository:

```bash
git clone https://github.com/your_username/vuelto.git
```

- Change the branch to the dev branch to follow rule \#1:

```bash
git checkout dev
```

## Pull Request

When you're ready with your changes, make sure you run `make` to format your code before pull requesting:

```bash
make format
```

Then make sure your code works without erroring and you're following the [contribution rules](#contributing-rules).

After all of this, you can create a pull request and one of our main members will look at it (and hopefully merge it!).
