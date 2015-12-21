# PGREP

Multi pattern grep and count. Implemented in Golang.

Same functionality exists in `grep` but this one is slower and simpler :-)

## Synopsis

   pgrep <flag> <pattern> [<flag> <pattern> ...] [file]

## Usage

`pgrep -p foo -c bar -f baz foobar.txt` - grep three patterns in foobar.txt

### File

The file name is optional. If not supplied, STDIN is used.

### Patterns

Patterns is specified as in [Google RE2](https://github.com/google/re2/wiki/Syntax)

### Flags
    * `-p` print matching lines
    * `-c` count matching lines
    * `-pc` print and count lines

## Wishlist

_Future improvements_

    * Make it run faster
    * Add some nice features


