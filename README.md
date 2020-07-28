#### Overview

This is an example of calling a C shared library from Go and vice-versa. The J
interpreter shared library connects to a front-end through user supplied
callbacks. This Go program calls the J shared library functions to initialize
the interpreter, then starts a repl. The J interpreter passes the output to the
Go function `GoOutput` which prints it to the console. The J shared library and
headers files are read from the present directory.

#### Running
```shell script
go run c.go main.go
```

#### Example Sesion
go run c.go main.go
```shell script
> %. ? 5 5 $ 0
_0.391485   1.53743   0.310594  0.62827    _1.181
_0.721259 0.0468615    1.53972 _3.30281   3.05491
   1.7293  _1.65349   _1.46677 _1.14991   1.91428
 0.872482 _0.536556 _0.0481831  2.71585   _2.9795
  _1.6354   1.44829   0.631947  1.65261 _0.922678
```
