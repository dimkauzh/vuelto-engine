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
- glow: Generate OpenGL Go bindings  
- CGo: Interface to use C with Go

## How to setup the work environment

If you want to contribute, you have to set up the work environment, so you can develop vuelto the right way.

1. Install  the packages

```bash
go mod tidy  
```

2. Fork [the repository](https://github.com/vuelto-org/vuelto).

3. Clone your forked github repository:

```bash
git clone https://github.com/your_username_/vuelto.git
cd vuelto  
```

4. Change the branch to the dev branch to follow rule \#1:

```bash
git checkout dev  
```

## Pull Request

If you're ready with your changes, then you must follow a few steps before pull requesting.

1. Run `make` to format your code:

```bash
Make format  
```

Then make sure your pull request code works without erroring and you followed the [contribution rules](#contributing-rules)

After all of this, you can create a pull request and one of our main organization members will look at it.
