## pwdshort ##
`pwdshort` is a utility for linux written in Go used to shorten the $PWD for display in the terminal, such as in a bash PS1 prompt.

---
### Compiling ###
---

To compile, clone the repository and run `$ go build pwdshort.go`.

---
### Bash Usage ###
---

Open ~/.bashrc in an editor, and update your PS1 variable definition:
```bash
    pathAbbrev='/path/to/pwdshort-binary'
    PS1='\[\033[01;32m\]\u@\h\[\033[00m\]:\n\e[01;36;19m`${pathAbbrev}`\e[0m\$ '
```

---
### Sample Output ###
---

Assuming the current directory is:
`/home/user/Development/projects/go/src/pwdshort`;
the path displayed in bash would be:
`~/Dev/.../src/pwdshort`

Paths with up to four parts are displayed as-is.

---
### License ###
---
`pwdshort` is available for free under the <a href="LICENSE">MIT License</a>.
