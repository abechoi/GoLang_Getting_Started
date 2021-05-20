<h1 align="center">
<p>Getting started with</p>
  Go Programming Language
</h1>

### Table of Contents

1. [Commands](#commands)
2. [Data Structures](#data-structures)
3. [Compile & Install](#compile-and-install)

### 🚀 1. Commands

```
# Finds main.go and executes it
go run .
```

```
# Creates a module
go mod init example.com/hello
```

```
# Edit and replaces location of module
go mod edit -replace=example.com/greetings=../greetings
```

```
# Downloads and imports all external packages
go mod tidy
```

### 🚀 2. Data Structures

```
# array
var a [4]int
```

```
# slice
s := []int{0, 1, 2, 3}
```

### 🚀 3. Compile and Install

```
# go to hello directory and build the app
go build

# discover install path
go list -f '{{.Target}}'

# go to ~ and export path to directory, exclude app name
export PATH=$PATH:/Users/me/go/bin

# install
go install

# execute app
hello
```
