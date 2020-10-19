# rinko
rinko is a trivial source code checker (especially for Android application development).

## Install

```
$ go get -u github.com/kwQt/rinko
```

## Commands

### comment
Check whether link comments are written above the definition of particular classes.
```
OPTIONS:
   --name value                    specify suffix of file name (default: "Fragment")
   --extension value, --ext value  specify file extension (default: "kt")
   --all, -a                       display all results (default: false)
```

### extfunc
List all Kotlin files including extension function.
```
OPTIONS:
   --name value  specify receiver type (default: "ALL")
   --help, -h    show help (default: false)
```
   
