package main

import (
	"bytes"
	"compress/zlib"
	"errors"
	"fmt"
	"io"
	"strconv"

	// Uncomment this block to pass the first stage!
	"os"

	"github.com/alecthomas/kong"
)

type Context struct {
}

type InitCmd struct {
}

type CatFileCmd struct {
	PrettyPrint bool   `short:"p" help:"Pretty print the content of blob."`
	BlobSha     string `arg:"" type:"string"`
}

var CLI struct {
	CatFile CatFileCmd `cmd:"cat-file"`
	Init    InitCmd    `cmd:"init"`
}

func assertEqual(a, b interface{}) error {
	if a != b {
		return errors.New("assertion failed: values are not equal")
	}
	return nil
}

func decompress(compressedData []byte) ([]byte, error) {
	rawReader := bytes.NewReader(compressedData)
	reader, err := zlib.NewReader(rawReader)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *InitCmd) Run(ctx *Context) error {
	for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating directory: %s\n", err)
		}
	}

	headFileContents := []byte("ref: refs/heads/main\n")
	if err := os.WriteFile(".git/HEAD", headFileContents, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %s\n", err)
	}

	fmt.Println("Initialized git directory")
	return nil
}

func (r *CatFileCmd) Run(ctx *Context) error {
	fn := fmt.Sprintf(".git/objects/%s/%s", r.BlobSha[:2], r.BlobSha[2:])

	compressedContent, err := os.ReadFile(fn)
	if err != nil {
		return err
	}

	data, err := decompress(compressedContent)

	if err != nil {
		return err
	}

	if len(data) == 0 {
		return nil
	}

	assertEqual(data[0], byte('b'))
	assertEqual(data[1], byte('l'))
	assertEqual(data[2], byte('o'))
	assertEqual(data[3], byte('b'))

	var i = 4
	i += 1

	var buffer bytes.Buffer
	for ; data[i] != '\x00'; i++ {
		buffer.WriteByte(data[i])
	}

	length, err := strconv.Atoi(buffer.String())

	if err != nil {
		return err
	}

	i += 1

	fmt.Print(string(data[i:(i + length)]))
	return nil
}

func main() {
	ctx := kong.Parse(&CLI)
	err := ctx.Run(&Context{})
	ctx.FatalIfErrorf(err)
}
