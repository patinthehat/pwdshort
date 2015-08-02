## pwdshort ##
`pwdshort` is a utility for linux written in Go used to shorten the $PWD for display in the terminal, such as in a bash PS1 prompt.

---
### Compiling ###
---

To compile, clone the repository and run `$ go build pwdshort.go`.

---
### Configuration ###
---
Configuration for `pwdshort` is stored in `pwdshort.json`.  It must reside in the same directory as the pwdshort binary.
Here you can modify how many path parts to display at the front and end of the abbreviated path.
The part of the config file that holds these settings is as follows by default:
```
    "MaxDisplayedParts":{
      "Prefix":1,
      "Suffix":2
    },
```

You can also specify pathname replacements, i.e. replace "projects" to "PRJ": `/home/user/projects` => `/home/user/PRJ`
Pathname replacements are *only* to make the CWD more compact, it does NOT rename the pathname.
Default settings in the configuration file:
```
    "Replacements": [
      { "Pathname": "Development",  "Replacement":"Dev" },
      { "Pathname": "projects",     "Replacement":"prj" },
      { "Pathname": "images",       "Replacement":"img" },
    ]
```

---
### Bash Usage ###
---

Open `~/.bashrc` in an editor, and update your PS1 variable definition:
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

Paths with up to four parts are displayed as-is (assuming the configuration file has not been modified).

---
### License ###
---
`pwdshort` is open source software and is available under the <a href="LICENSE">MIT License</a>.
