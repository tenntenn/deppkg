# deppkg

## What is this?

`deppkg` analyzes coverprofile files of go test and outputs package list which are depended from given gofiles.
`deppkg` assumes that coverprofile files are put on each packages direcotry.

## How to use

### Basic

```
$ deppkg -d $GOPATH/src/myproject file1.go file2.go...
```

### With test

`deppkg` can be used for tests with diffrences without whole tests.
Because `deppkg` detects dependent packages with modified go files from coverprofile files.

```
for $PKG in `git diff --name-only HEAD | xargs deppkg -d $GOPATH/src/myproject`
do
    go test -coverprofile=$GOPATH/src/$PKG/coverprofile $PKG
done
```

If you have not created coverprofile files yet, you can run this commands.

```
for $PKG in `find $GOPATH/src/myproject -type d`
do
    go test -coverprofile=$PKG/coverprofile $PKG
done
```

When you use `deppkg` for tests, you should commit coverprofile files to your git (or other VCS) repository.
