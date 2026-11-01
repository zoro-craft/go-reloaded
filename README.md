# Go Reloaded

A small command line tool that reads a text file, applies a fixed set of
formatting rules, and writes the result to a second file. Built in Go as part of
the Zone01 / 42 curriculum.

## Usage

```bash
go run . sample.txt output.txt
```

Or build a binary first:

```bash
go build -o go-reloaded .
./go-reloaded input.txt output.txt
```

Both arguments must end in `.txt`, the input must not be larger than **100 KiB**,
and exactly two file names are accepted — anything else prints a message and
exits without touching the output file.

## Rules

Every rule is applied to the whole file, line by line, in this order:

| Marker | Effect | Example |
|---|---|---|
| `(hex)` | Converts the previous word from hexadecimal to decimal | `1a (hex)` → `26` |
| `(bin)` | Converts the previous word from binary to decimal | `101 (bin)` → `5` |
| `(up)` | Uppercases the previous word | `stop (up)` → `STOP` |
| `(low)` | Lowercases the previous word | `BREAKFAST (low)` → `breakfast` |
| `(cap)` | Capitalises the previous word, lowercases the rest | `harold (cap)` → `Harold` |
| `(up, N)` | Uppercases the previous N words | `in my house (up, 2)` → `in MY HOUSE` |
| `(low, N)` | Lowercases the previous N words | `BREAKFAST IN BED (low, 3)` → `breakfast in bed` |
| `(cap, N)` | Capitalises the previous N words | `it was the (cap, 2)` → `It Was the` |
| punctuation | Removes the space before `. , ! ? : ;` and tidies the space after | `boring ,what` → `boring, what` |
| quotes | Attaches single quotes to the text they wrap | `wilson : ' I am '` → `Wilson: 'I am'` |
| articles | Turns `a` / `A` into `an` / `An` before a vowel | `a optimist` → `an optimist` |

The numbered markers arrive as two tokens (`(up,` and `2)`), because the input is
split on whitespace before processing.

## Tests

```bash
go test ./...
```

`main_test.go` checks that tokenizing and rebuilding `sample.txt` is lossless, and
that the full pipeline turns `sample.txt` into exactly `expectedResult.txt`.

## Structure

| File | Role |
|---|---|
| `main.go` | Argument and extension checks, file reading, pipeline calls, writing |
| `operations.go` | Argument validation, tokenizing, the stage pipeline, reconstruction |
| `textprocessor/` | One file per rule: hex/bin, case, numbered case, punctuation, quotes, vowels |
| `textprocessor/utils.go` | Shared helpers: word detection, quote scanning, token removal |
| `sample.txt` / `expectedResult.txt` | Input and expected output used by the tests |
