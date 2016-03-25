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
Listening on port 2525
---
from: john@doe.com
to: jane@doe.com
Reply-To:
subject: =?utf-8?Q?Hi?=
Message-ID: <1458919281.382535.501.70189175994080@doe.com>
Date: Fri, 25 Mar 2016 10:21:21 -0500

Hi!
```

License
-------

mat is released under the [MIT License](http://www.opensource.org/licenses/MIT).
