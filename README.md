mat
===

A little mail tester daemon written in Go.

Ported from [mt](https://github.com/soveran/mt).

Description
-----------

`mat` fakes a SMTP server and prints incoming emails to stdout.

Installation
------------

```
go get github.com/frodsan/mat
```

Usage
-----

By default, it uses port 2525.

```
$ mat
from: john@doe.com
to: jane@doe.com
Reply-To:
subject: =?utf-8?Q?Hi?=
# ...
```

License
-------

mat is released under the [MIT License](http://www.opensource.org/licenses/MIT).
