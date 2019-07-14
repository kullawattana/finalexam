# GO WORKSHOP

# Run Go
## $ go run main.go

- Postman (latest vesion https://www.getpostman.com/downloads/)
- git (latest version https://git-scm.com/downloads)
- vscode (latest version https://code.visualstudio.com/download)
- vscode xtension (https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)
- GOROOT ใน enviroment variable (https://golang.org/dl/)

# Gin Go Command
- go mod init
- go env
- go mod init https://github.com/kullawattana/school
- go mod tidy
- echo $GOPATH

# Set port
- PORT=1234 go run main.go

# Run
- go run main.go

# Gin Go Install

## $ go mod init
go: cannot determine module path for source directory /Users/toppee/Desktop/school (outside GOPATH, no import comments)

## $ go run main.go
main.go:3:8: cannot find package "github.com/gin-gonic/gin" in any of:
        /usr/local/go/src/github.com/gin-gonic/gin (from $GOROOT)
        /Users/toppee/go/src/github.com/gin-gonic/gin (from $GOPATH)

## $ go env
GOARCH="amd64" \
GOBIN="" \
GOCACHE="/Users/toppee/Library/Caches/go-build" \
GOEXE="" \
GOFLAGS="" \
GOHOSTARCH="amd64" \
GOHOSTOS="darwin" \
GOOS="darwin" \
GOPATH="/Users/toppee/go" \
GOPROXY="" \
GORACE="" \
GOROOT="/usr/local/go" \
GOTMPDIR="" \
GOTOOLDIR="/usr/local/go/pkg/tool/darwin_amd64" \
GCCGO="gccgo" \
CC="clang" \
CXX="clang++" \
CGO_ENABLED="1" \
GOMOD="" \
CGO_CFLAGS="-g -O2" \
CGO_CPPFLAGS="" \
CGO_CXXFLAGS="-g -O2" \
CGO_FFLAGS="-g -O2" \
CGO_LDFLAGS="-g -O2" \
PKG_CONFIG="pkg-config" \
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/bj/0q111_8s7fjbh8mfjfvyhpdh0000gn/T/go-build613639050=/tmp/go-build -gno-record-gcc-switches -fno-common" 

## $ go mod init https://github.com/kullawattana/school
go: creating new go.mod: module https://github.com/kullawattana/school

## $ go mod tidy
go: finding github.com/gin-gonic/gin v1.4.0 \
go: downloading github.com/gin-gonic/gin v1.4.0 \
go: extracting github.com/gin-gonic/gin v1.4.0 \
go: finding github.com/modern-go/reflect2 v1.0.1 \
go: finding github.com/stretchr/testify v1.3.0 \
go: finding github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd \
go: finding github.com/golang/protobuf v1.3.1 \
go: finding github.com/gin-contrib/sse v0.0.0-20190301062529-5545eab6dad3 \
go: finding github.com/json-iterator/go v1.1.6 \
go: finding github.com/ugorji/go v1.1.4 \
go: finding golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c \
go: finding gopkg.in/go-playground/validator.v8 v8.18.2 \
go: finding gopkg.in/go-playground/assert.v1 v1.2.1 \
go: finding gopkg.in/yaml.v2 v2.2.2 \
go: finding github.com/mattn/go-isatty v0.0.7 \
go: finding github.com/stretchr/objx v0.1.0 \
go: finding github.com/pmezard/go-difflib v1.0.0 \
go: finding github.com/davecgh/go-spew v1.1.0 \
go: finding golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2 \
go: finding golang.org/x/text v0.3.0 \
go: finding golang.org/x/sys v0.0.0-20190222072716-a9d3bda3a223 \
go: finding gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405 \
go: finding golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a \
go: golang.org/x/sys@v0.0.0-20190222072716-a9d3bda3a223: git fetch -f origin refs/heads/*:refs/heads/* refs/tags/*:refs/tags/* in /Users/toppee/go/pkg/mod/cache/vcs/76a8992ccba6d77c6bcf031ff2b6d821cf232e4ad8d1f2362404fbd0a798d846: exit status 128: \
        error: RPC failed; curl 56 LibreSSL SSL_read: SSL_ERROR_SYSCALL, errno 54 \
        fatal: the remote end hung up unexpectedly \
        fatal: early EOF \
        fatal: index-pack failed \
go: golang.org/x/sys@v0.0.0-20190215142949-d0b11bdaac8a: unknown revision d0b11bdaac8a \
go: error loading module requirements 

## $ go $GOPATH
Go is a tool for managing Go source code.

## Usage:

        go <command> [arguments]

## The commands are:

        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         download and install packages and dependencies
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages

## Use "go help <command>" for more information about a command.

## Additional help topics:

        buildmode   build modes
        c           calling between Go and C
        cache       build and test caching
        environment environment variables
        filetype    file types
        go.mod      the go.mod file
        gopath      GOPATH environment variable
        gopath-get  legacy GOPATH go get
        goproxy     module proxy protocol
        importpath  import path syntax
        modules     modules, module versions, and more
        module-get  module-aware go get
        packages    package lists and patterns
        testflag    testing flags
        testfunc    testing functions

## Use "go help <topic>" for more information about that topic.

## $ go %GOPATH
go %GOPATH: unknown command \
Run 'go help' for usage. 

# $ go mod tidy
go: finding golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a \
go: finding golang.org/x/sys v0.0.0-20190222072716-a9d3bda3a223 \
go: downloading github.com/gin-contrib/sse v0.0.0-20190301062529-5545eab6dad3 \
go: downloading github.com/mattn/go-isatty v0.0.7 \
go: downloading github.com/ugorji/go v1.1.4 \
go: downloading gopkg.in/go-playground/validator.v8 v8.18.2 \
go: downloading gopkg.in/yaml.v2 v2.2.2 \
go: extracting github.com/gin-contrib/sse v0.0.0-20190301062529-5545eab6dad3 \
go: extracting github.com/mattn/go-isatty v0.0.7 \
go: downloading github.com/json-iterator/go v1.1.6 \
go: downloading github.com/golang/protobuf v1.3.1 \
go: downloading golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c \
go: downloading golang.org/x/sys v0.0.0-20190222072716-a9d3bda3a223 \
go: downloading github.com/stretchr/testify v1.3.0 \
go: extracting github.com/json-iterator/go v1.1.6 \
go: downloading github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd \
go: downloading github.com/modern-go/reflect2 v1.0.1 \
go: extracting github.com/stretchr/testify v1.3.0 \
go: extracting github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd \
go: downloading github.com/davecgh/go-spew v1.1.0 \
go: downloading github.com/pmezard/go-difflib v1.0.0 \
go: extracting github.com/modern-go/reflect2 v1.0.1 \
go: extracting github.com/ugorji/go v1.1.4 \
go: extracting github.com/pmezard/go-difflib v1.0.0 \
go: extracting github.com/davecgh/go-spew v1.1.0 \
go: extracting github.com/golang/protobuf v1.3.1 \
go: extracting golang.org/x/sys v0.0.0-20190222072716-a9d3bda3a223 \
go: extracting golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c \
go: extracting gopkg.in/go-playground/validator.v8 v8.18.2 \
go: downloading gopkg.in/go-playground/assert.v1 v1.2.1 \
go: extracting gopkg.in/yaml.v2 v2.2.2 \
go: downloading gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405 \
go: extracting gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405 \
go: extracting gopkg.in/go-playground/assert.v1 v1.2.1 

# Postgres
- url := "postgres://fqwnvlfk:lv3nDmkzmBXgk6dup77dO6CbsjcJa2-L@satao.db.elephantsql.com:5432/fqwnvlfk"
# Database : https://github.com/lib/pq
## go mod init github.com/suttipong/postgres

## $ go mod init github.com/suttipong/postgres
go: creating new go.mod: module github.com/suttipong/postgres

## $ go mod tiny
go mod tiny: unknown command \
Run 'go help mod' for usage.

## $ go run postgres.go
go: finding github.com/lib/pq v1.1.1 \
go: downloading github.com/lib/pq v1.1.1 \
go: extracting github.com/lib/pq v1.1.1 \
OK

- SET DATABASE_URL = “postgres://dfpsncgm:ZvvFpVgiPUz5xo25lRmoQS1Rf41uzrQo@satao.db.elephantsql.com:5432/dfpsncgm” go run insert.go 

# REST SERVICE
- Command + shift + P (Setup Plugin)

```go
"go.alternateTools": {
        "go-langserver": []
},
"go.languageServerFlags": [],
"go.useLanguageServer": false
```

# Unit Test
_test \
*testing.t 

--------------------------------------------------
#Toppees-MacBook-Pro:school-service-master toppee$ go run main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached. \

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production. \
 - using env:   export GIN_MODE=release \
 - using code:  gin.SetMode(gin.ReleaseMode) \

[GIN-debug] GET    /api/todos                --> github.com/..../school-service/service.GetTodos (4 handlers)\
[GIN-debug] GET    /api/todos/:id            --> github.com/..../school-service/service.GetTodosById (4 handlers)\
[GIN-debug] POST   /api/todos/               --> github.com/..../school-service/service.PostTodos (4 handlers)\
[GIN-debug] DELETE /api/todos/:id            --> github.com/..../school-service/service.DeleteTodosById (4 handlers)\
[GIN-debug] PUT    /api/todos/:id            --> github.com/..../school-service/service.UpdateTodo (4 handlers)\

No Port In Heroku1234 \
[GIN-debug] Listening and serving HTTP on :1234 \
[GIN-debug] redirecting request 301: /api/todos/ --> /api/todos \
Hello from middleware \
token :Bearer token123 \
[{1 Homework active} {2 Homework Inactive} {3 Homework Inactive}] \
Ending \

[GIN] 2019/06/29 - 10:36:03 | 200 |  1.106725518s |             ::1 | GET      /api/todos 

# Commit

## $ ls
go.mod  go.sum  main.go

## $ git init
Initialized empty Git repository in /Users/toppee/Desktop/Github Workspace/KBTG-workshop-Go-Programming-2019/school/.git/

## $ git status
On branch master

No commits yet

Untracked files:
  (use "git add <file>..." to include in what will be committed)

        .vscode/
        go.mod
        go.sum
        main.go

nothing added to commit but untracked files present (use "git add" to track)
## $ git add main.go go.sum go.mod
## $ git status
On branch master
## $ git commit -m "first commit"
[master (root-commit) 854819c] first commit \
 3 files changed, 161 insertions(+) \
 create mode 100644 go.mod \
 create mode 100644 go.sum \
 create mode 100644 main.go
 
## Toppees-MacBook-Pro:school toppee$ 
nothing added to commit but untracked files present (use "git add" to track)

## $ git remote add origin https://github.com/kullawattana/school.git

## $ git push -u origin master 

Enumerating objects: 5, done. \
Counting objects: 100% (5/5), done. \
Delta compression using up to 8 threads \
Compressing objects: 100% (5/5), done. \
Writing objects: 100% (5/5), 3.13 KiB | 3.13 MiB/s, done. \
Total 5 (delta 0), reused 0 (delta 0) \
To https://github.com/kullawattana/school.git \
 * [new branch]      master -> master \
Branch 'master' set up to track remote branch 'master' from 'origin'. 

# Build 
## $ go build -o school .
## $ go build . 
## $ ./school-service

# RUN !!!
## $ go env
## $ GOOS=windows GOARCH=amd64 go build -o school.exe .

# Code Convention

- how to call function by writing big character name of method  (public)
- Static Language 
- Dynamic Language
- ฟังก์ชั่นที่ขึ้นต้นด้วย test เป็น Code Convension ที่ใช้กับ testcase
- จัด format go fmt use swap.go

```go
pingHandler := func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
}

ss := map[string]Student{
        "1111": Student{Name: "Suttipong", ID: "1111"},
}

response = map[string]interface{}{
        "message": "success",
}
```
