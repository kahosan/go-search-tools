# Search Tools

ChatGPT 教我写的，README 都是它生成的，听我说 谢谢你

This is a command-line utility written in Go that allows you to search for text or files in a given directory.

## Usage

To use the program, you need to have Go installed on your system. Once you have Go, you can download and build the program by running the following commands:

```terminal
go get github.com/user/tools
cd $GOPATH/src/github.com/user/tools
go build
```

You can then run the `tools` binary to search for text or files in a directory:

```terminal
./tools -d <directory> -s <search-text> [-f]
```

The `-d` flag specifies the directory to search in, and the `-s` flag specifies the text to search for. If the `-f` flag is provided, the program will search for files instead of text.

Here is an example of how to search for the text "hello" in the current directory:

```terminal
./tools -d . -s hello
```

And here is an example of how to search for files with the ".txt" extension in the `/tmp` directory:

```terminal
./tools -d /tmp -s ".txt" -f
```

The program will print the paths of any files or text that match the search criteria.

## Implementation

The program uses the `flag` package to parse command-line arguments and the `app` package to search for text or files. The `app` package contains two structs, `SearchText` and `SearchFile`, that implement the `SearchTools` interface. The `main` function uses these structs to search for text or files based on the `-f` flag provided by the user.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
