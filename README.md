# URLED

A URL encoder/decoder command line utility.

## Installation

```bash
go install github.com/madsaune/urled
```

## Usage

### Encode

Using `stdin`.

```bash
~ $ urled -encode
hello world
<CTRL-D>
hello+world
```

Using pipe.

```bash
~ $ echo "hello world" | urled -encode
hello+world
```

Using `-input` flag.

```
~ $ urled -encode -input "hello world"
hello+world
```

### Decode

Using `stdin`.

```bash
~ $ urled -decode
hello+world
<CTRL-D>
hello world
```

Using pipe.

```bash
~ $ echo "hello+world" | urled -decode
hello world
```

Using `-input` flag.

```
~ $ urled -decode -input "hello+world"
hello world
```
