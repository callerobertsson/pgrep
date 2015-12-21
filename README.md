# PGREP

Multi pattern grep and count. Implemented in Golang.

Same functionality exists in `grep` but this one is slower and simpler :-)

## Synopsis

    pgrep <flag> <pattern> [<flag> <pattern> ...] [file]

## Usage

Grep three patterns in foobar.txt

    pgrep -p foo -c bar -pc baz foobar.txt

Lines matching foo will be printed, number of lines matching bar will be printed,
and the lines matching baz will be printed and the number of lines will be printed.

### File

The file name is optional. If not supplied, STDIN is used.

### Patterns

Patterns are regular expressions as specified in [Google RE2](https://github.com/google/re2/wiki/Syntax).

### Flags

* `-p` print matching lines
* `-c` count matching lines
* `-pc` print and count lines

## Future improvements

* Break out pgrep funcs in package
* Make it run faster
* Add some nice features


