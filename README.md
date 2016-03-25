mat
===

Prints mail to stdout. Ported from [mt](https://github.com/soveran/mt).

Installation
------------

```
go get github.com/frodsan/mat
```

Usage
-----

`mat` fakes a SMTP server and prints incoming emails to stdout.
By default, it binds to port 2525.

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
