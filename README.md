# filep

[![GitHub license](https://img.shields.io/github/license/onozaty/filep)](https://github.com/onozaty/filep/blob/main/LICENSE)
[![Test](https://github.com/onozaty/filep/actions/workflows/test.yaml/badge.svg)](https://github.com/onozaty/filep/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/onozaty/filep/branch/main/graph/badge.svg?token=VSU64LAK8P)](https://codecov.io/gh/onozaty/filep)

`filep` is a powerful command-line file processing tool designed to efficiently handle text files with various operations.

With `filep`, you can easily perform the following file operations:

* **[replace](#replace)** - Replace specific content in files using strings or regular expressions
* **[truncate](#truncate)** - Truncate files to a specified size (by bytes, characters, or lines)
* **[extract](#extract)** - Extract specific portions of files based on position ranges

## Common

The following is a description of the parts that are common to each subcommand.

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

Encoding must be one that can be specified in `htmlindex.Get`.

* https://pkg.go.dev/golang.org/x/text/encoding/htmlindex#Get

## replace

The `replace` command allows you to find and replace text in files using either string matching or powerful regular expressions. This is perfect for batch text replacements, data transformation, or content cleaning across multiple files.  

### Usage

```
filep replace -i INPUT -o OUTPUT [-r REGEX | -s STRING] -t REPLACEMENT [--escape] [--recursive] [--encoding ENCODING]
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

#### Binary encoding

The special `binary` encoding allows you to work directly with binary data using hexadecimal notation.  
Each byte is represented as a three-character sequence starting with `x`, such as `x00` for null byte or `xFF` for 255.

For example, to remove a specific byte sequence like 0x00 0x01:

```
$ filep replace -i input.txt -o output.txt -s x00x01 -t "" --encoding binary
```

#### Note

* See [Common / Input Output](#input--output) for input/output.
* See [Common / Encoding](#encoding) for file encoding.

## truncate

The `truncate` command helps you cut files to a specific size, whether you need to limit by bytes, characters, or line count. This is useful for creating file previews, removing excess content, or processing large log files.  

### Usage

```
filep truncate -i INPUT -o OUTPUT [-b BYTES | -c CHARS | -l LINES] [--recursive] [--encoding ENCODING]
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

#### Note

* See [Common / Input Output](#input--output) for input/output.
* See [Common / Encoding](#encoding) for file encoding.

## extract

The `extract` command enables precise extraction of specific portions from files by defining start and end positions. You can extract content by bytes, characters, or lines, making it ideal for data parsing and content isolation.  

### Usage

```
filep extract -i INPUT -o OUTPUT [-s START] [-e END] [-b | -c | -l] [--recursive] [--encoding ENCODING]
```

```
Usage:
  filep extract [flags]

Flags:
  -i, --input string      Input file/dir path.
  -o, --output string     Output file/dir path.
  -s, --start int         Start position.
  -e, --end int           End position.
  -b, --byte              Handle by bytes.
  -c, --char              Handle by characters.
  -l, --line              Handle by lines.
      --recursive         Recursively traverse the input dir.
      --encoding string   Encoding. (default "UTF-8")
  -h, --help              help for extract
```

#### Extract method

Extracts a portion of a file by specifying a start and end position.  
The position can be specified in bytes(`-b`), characters(`-c`), or lines(`-l`).

For example, if you want to extract the 3rd to 5th bytes, specify as follows

```
$ filep extract -i input.txt -o output.txt -s 3 -e 5 -b
```

Start(`-s`) and end(`-e`) can be omitted.

If start(`-s`) is omitted, the value is `1`.  
For example, if you want to extract only the first line, you can specify as follows.

```
$ filep extract -i input.txt -o output.txt -e 1 -l
```

If end(`-e`) is omitted, it will be at the end of the file.  
For example, if you want to erase only the first character, you can specify as follows.

```
$ filep extract -i input.txt -o output.txt -s 2 -c
```

#### Note

* See [Common / Input Output](#input--output) for input/output.
* See [Common / Encoding](#encoding) for file encoding.

## Install

### Homebrew (macOS/Linux)

```
brew install onozaty/tap/filep
```

### Scoop (Windows)

```
scoop bucket add onozaty https://github.com/onozaty/scoop-bucket
scoop install filep
```

### Binary Download

Download the latest binary from [GitHub Releases](https://github.com/onozaty/filep/releases/latest).

## License

MIT

## Author

[onozaty](https://github.com/onozaty)
