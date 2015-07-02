package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/momchil-atanasov/gostub/util"
)

func RunGoStub(c *cli.Context) {
	input, err := parseInput(c)
	exitOnErr(err)
	fmt.Printf("Running `gostub` with input: %#v\n", input)
}

type goStubInput struct {
	InterfaceName   string
	SourceDirectory string
	StubName        string
	OutputFilePath  string
}

func parseInput(c *cli.Context) (goStubInput, error) {
	if len(c.Args()) == 0 {
		return goStubInput{}, errors.New("Interface name not specified! Run `gostub --help` for more information.")
	}
	interfaceName := c.Args().First()

	sourceDir := c.String("source")
	if sourceDir == "" {
		sourceDir = filepath.Dir(c.App.Name)
	}
	sourceDir, err := filepath.Abs(sourceDir)
	if err != nil {
		return goStubInput{}, err
	}

	stubName := c.String("name")
	if stubName == "" {
		stubName = interfaceName + "Stub"
	}

	outputFileName := c.String("output")
	if outputFileName == "" {
		outputFolder := path.Join(sourceDir, path.Base(sourceDir)+"_stubs")
		outputFile := util.SnakeCase(interfaceName) + "_stub.go"
		outputFileName = path.Join(outputFolder, outputFile)
	}
	if path.Ext(outputFileName) != ".go" {
		return goStubInput{}, errors.New("The output file needs to have the `go` extension! Run `gostub --help` for more information.")
	}
	outputFileName, err = filepath.Abs(outputFileName)
	if err != nil {
		return goStubInput{}, err
	}

	return goStubInput{
		InterfaceName:   interfaceName,
		SourceDirectory: sourceDir,
		StubName:        stubName,
		OutputFilePath:  outputFileName,
	}, nil
}

func exitOnErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
