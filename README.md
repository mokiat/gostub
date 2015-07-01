# GoStub

GoStub is a stubbing tool for the Go programming language.

You can use the tool to create stub implementations of your own Go interfaces. Those stubs can then be used to fake the implementation of a method of the interface or to verify the number of times a given method was called and the arguments that were passed.

Some of the ideas used in this tool are heavily based on the  **[counterfeiter](https://github.com/maxbrunsfeld/counterfeiter)** tool, where `gostub` seeks to improve on them and to solve some issues in the aforementioned tool.

## User's Guide

You can get the tool using the following command. (You should have your Go development environment already set up)

```bash
go get github.com/momchil-atanasov/gostub
```

You can then navigate to the folder where your interface is located and execute the following command.

```bash
gostub <interface_name>
```

This will generate a stub called `<interface_name>Stub` in the `<folder_name>_stubs/<interface_name>_stub.go` file.

**Note:** The directory needs to be part of the `src` sub-tree of your `$GOPATH`.

It's unlikely that you will want to write that statement each time you desire your stub be recreated. Instead, you can use Go's generate functionality. Your interface file might look something like this.

```go
package example

//go:generate gostub Person

type Person interface {
	// ...
}
```

All you need to do is run the `go generate ./...` command from the current or parent directory and all your stubs will be regenerated.

If you want to run the `gostub` from a different location, you can use the `-s` or `--source` flags to specify the location of the interface.

Example:

```bash
gostub -s $GOPATH/src/github.com/momchil-atanasov/example Person
```

If you wish to change to location where the generated stub will be saved, you can use the `-o` or `--output` flags to specify the file path of the generated stub.

Example:

```bash
gostub -o my_stubs/this_is_a_person_stub.go Person
```

If you wish to change the name of the generated stub, you can use the `-n` or `--name` flags to specify a new one. This will also affect the target file name, unless you use the `-o` or `--output` flags.

Example:

```bash
gostub -n StubbedPerson Person
```

## Developer's Guide

This project uses the [Ginkgo](https://github.com/onsi/ginkgo) for the tests, so you will need to download it.
