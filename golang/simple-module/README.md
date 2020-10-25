## Simple Module Example

A [tutorial example](https://golang.org/doc/tutorial/create-module) to build a simple Go module.

### Compile, Install and Run the Application 
```bash
cd hello/
# discover installation paths
go list -f '{{.Target}}'

# add the path to "bin/" discovered from the previous step to system shell path
export PATH=$PATH:/path/to/your/install/directory

# compile, install and run the package
go build && ./hello
```

## Test
To run test,
```bash
cd greetings/ && go test
```

### Lint and Format Correction
```bash
gofmt
```