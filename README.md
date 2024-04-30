# pa55
Generates a cryptographically secure password on the command line and copies it to the clipboard.

## Installation

``` bash
$ go get -u github.com/codingconcepts/pa55
```

## Usage

``` bash
$ pa55
**************************************************
```

## Arguments

``` bash
$ pa55
Usage of pa55:
  -len int
        specify the output length (default 50)
  -out string
        specify the output encoding ([ascii, hex, base32, base64, unicode]) (default "ascii")
  -print
        print the password rather than copying it
  -set string
        character set to use for passwords (default " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~")
```

Available output types:

* hex
``` bash
$ pa55 --out hex --len 20
****************************************

0a06f90380e443389d77c5d4809a749be2f28af9 // -> To clipboard
```

* base32
``` bash
$ pa55 --out base32 --len 20
********************************

5EBBYLOC2HFXFAFWCBPMYXDCPJLONTK6 // -> To clipboard
```

* base64
``` bash
$ pa55 --out base64 --len 20
****************************

tJG/ih/TNNnr4+7bSaGC+CqkBqo= // -> To clipboard
```

* ascii
``` bash
$ pa55 --out ascii --len 20
********************

!Qv2ACRj0c1BjsQP8x*D // -> To clipboard
```

* unicode
``` bash
$ pass --out unicode --len 20
********************************************************

៧Ἀจያا्༙॰क໅྾Ҍཙ෰᠛ǃΣႍឫႶ // -> To clipboard
```