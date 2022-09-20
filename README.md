# filep

[![GitHub license](https://img.shields.io/github/license/onozaty/filep)](https://github.com/onozaty/filep/blob/main/LICENSE)
[![Test](https://github.com/onozaty/filep/actions/workflows/test.yaml/badge.svg)](https://github.com/onozaty/filep/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/onozaty/filep/branch/main/graph/badge.svg?token=VSU64LAK8P)](https://codecov.io/gh/onozaty/filep)

`filep` is a file processing tool.

`filep` has the following subcommands

* [replace](#replace) Replace file contents.
* [truncate](#truncate) Truncate file contents.


## replace

Replace the contents of the input file to create an output file.  

### Usage

```
filep replace -i INPUT -o OUTPUT [-r REGEX | -s STRING] -t REPLACEMENT [--escape] [--recursive] [-encoding ENCODING]
```

```
Usage:
  filep replace [flags]

Flags:
  -i, --input string         Input file/dir path.
  -o, --output string        Output file/dir path.
  -r, --regex string         Target regex.
  -s, --string string        Target string.
  -t, --replacement string   Replacement.
      --escape               Enable escape sequence.
      --recursive            Recursively traverse the input dir.
      --encoding string      Encoding. (default "UTF-8")
  -h, --help                 help for replace
```

#### Replacement method

The target can be a regular expression or a string.

The regular expression is specified by `-r`.

```
$ filep replace -i input.txt -o output.txt -r "[0-9]" -t ""
```

You can also use the capture group as `-t`.  
The following is a method for specifying "N" at the beginning of a number.

```
$ filep replace -i input.txt -o output.txt -r "([0-9]+)" -t "N$1"
```

Please refer to the following for the syntax of regular expressions.

* https://pkg.go.dev/regexp/syntax

String is specified by `-s`.

```
$ filep replace -i input.txt -o output.txt -s a -t z
```

To treat backslash as an escape sequence, specify `--escape`.

```
$ filep replace -i input.txt -o output.txt -s "\u3000" -t "" --escape
```

#### Input / Output

If specified with `-i`, only the specified file will be processed.

```
$ filep replace -i input.txt -o output.txt -s a -t z
```

If a directory is specified, files under the directory are processed.  

```
$ filep replace -i in_dir -o out_dir -s a -t z
```

The default is to process only files directly under the specified directory.  
If `--recursive` is specified, subdirectories are processed recursively.

```
$ filep replace -i in_dir -o out_dir -s a -t z --recursive
```

#### Encoding

When processing non UTF-8 files, specify the encoding with `--encoding`.

```
$ filep replace -i input.txt -o output.txt -s a -t z --encoding sjis
```

Charset must be one that can be specified in `htmlindex.Get`.

* https://pkg.go.dev/golang.org/x/text/encoding/htmlindex#Get

A special charset is `binary`.  
If `binary` is specified, it can be treated as a hexadecimal character.  
A hexadecimal character represents a byte with three characters prefixed by `x`, such as `x00` or `xFF`.

To remove two consecutive bytes, such as 0x00 0x01, specify as follows

```
$ filep replace -i input.txt -o output.txt -s x00x01 -t "" --encoding binary
```

## truncate

Truncate the contents of the input file to create an output file.  

### Usage

```
filep truncate -i INPUT -o OUTPUT [-b BYTES | -c CHARS | -l LINES] [--recursive] [-encoding ENCODING]
```

```
Usage:
  filep truncate [flags]

Flags:
  -i, --input string      Input file/dir path.
  -o, --output string     Output file/dir path.
  -b, --byte int          Number of bytes.
  -c, --char int          Number of characters.
  -l, --line int          Number of lines.
      --recursive         Recursively traverse the input dir.
      --encoding string   Encoding. (default "UTF-8")
  -h, --help              help for truncate
```

#### Truncate method

The truncation size is specified in bytes, characters, or lines.

The number of bytes is specified by `-b`.

```
$ filep truncate -i input.txt -o output.txt -b 100
```

The number of characters is specified by `-c`.

```
$ filep truncate -i input.txt -o output.txt -c 100
```

The number of lines is specified by `-l`.

```
$ filep truncate -i input.txt -o output.txt -l 100
```

#### Input / Output

If specified with `-i`, only the specified file will be processed.

```
$ filep truncate -i input.txt -o output.txt -b 100
```

If a directory is specified, files under the directory are processed.  

```
$ filep truncate -i in_dir -o out_dir -b 100
```

The default is to process only files directly under the specified directory.  
If `--recursive` is specified, subdirectories are processed recursively.

```
$ filep truncate -i in_dir -o out_dir -b 100 --recursive
```

#### Encoding

When processing non UTF-8 files, specify the encoding with `--encoding`.

```
$ filep truncate -i input.txt -o output.txt -c 100 --encoding sjis
```

Charset must be one that can be specified in `htmlindex.Get`.

* https://pkg.go.dev/golang.org/x/text/encoding/htmlindex#Get

## Install

`filep` is implemented in golang and runs on all major platforms such as Windows, Mac OS, and Linux.  
You can download the binaries for each OS from the links below.

You can download the binary from the following.

* https://github.com/onozaty/filep/releases/latest

## License

MIT

## Author

[onozaty](https://github.com/onozaty)
