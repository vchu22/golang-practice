# Golang Practice

Practice files for learning [Go](https://golang.org/)

### Folder Structure

```
- bin
- src
    - program1
    - program2
```

### Execute a source code

Navigate into the program's folder using `cd folder-path/` with `folder-path/` being the path to your file and execute `go run filename.go` with `filename.go` being whatever name you named your Go file with.

### Compile

Executing `go build` inside the program's folder will create a binary with a filename based on your current working directory.
Executing `go build path-to-file/filename.go` will create a binary with filename based on the name of the Go file.
