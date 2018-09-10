# vendiff

This tool encapsulates a workflow for modifying dependencies within a project so it is possible to change and modify your dependencies to debug problems with them in your code base. It can be used with any project that uses Go modules that does not commit their vendor directory to source control.

    $ vendiff init
    $ vim vendor/github.com/spf13/cobra/command.go # modify this file
    $ vendiff # output the diff for the vendor directory
    $ vendiff -- github.com/spf13/cobra # output a diff for a specific folder
    $ vendiff clean -f # remove the vendor directory when you are done

This makes it easy to modify the vendored dependencies and then use the diff to generate a patch that can be applied to the original using `patch`. As an example for if you found a bug in `github.com/spf13/cobra` while working on your own project:

    $ vendiff -- github.com/spf13/cobra > ~/cobra-my-bugfix.patch
    $ cd ~/go/src/github.com/spf13/cobra
    $ patch -p4 < ~/cobra-my-bugfix.patch
    $ git commit -a
