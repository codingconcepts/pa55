# pa55
Generates a cryptographically secure password on the command line.

## Installation

``` bash
$ go get -u github.com/codingconcepts/pa55
```

## Usage

``` bash
$ pa55 --out base32 --len 50
```

## Arguments

``` bash
$ pa55
Usage of pa55:
  -len int
        specify the output length (default 20)
  -out string
        specify the output encoding ([ascii, hex, base32, base64]) (default "hex")
```

Available output types:

* hex
``` bash
$ pa55 --out hex --len 20
0a06f90380e443389d77c5d4809a749be2f28af9
```

* base32
``` bash
$ pa55 --out base32 --len 20
5EBBYLOC2HFXFAFWCBPMYXDCPJLONTK6
```

* base64
``` bash
$ pa55 --out base64 --len 20
tJG/ih/TNNnr4+7bSaGC+CqkBqo=
```

* ascii
``` bash
$ pa55 --out ascii --len 20
!Qv2ACRj0c1BjsQP8x*D
```
