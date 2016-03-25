mat
===

Prints mail to stdout.

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
from: frodsan@me.com
to: hello@frodsan.com
Reply-To:
subject: =?utf-8?Q?Hi?=
```

License
-------

mat is released under the [MIT License](http://www.opensource.org/licenses/MIT).
