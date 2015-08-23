Israeli Supermarket Price Project
=================================

This is a collection of tools and scripts to download and manage price data
from Israeli vendors.

1. [License](#license)
1. [What's in the Box](#whats-in-the-box)
1. [Requirements](#requirements)
1. [How to Build](#how-to-build)
1. [TODOs](#todos)

License
-------

Copyright (c) 2015 Amit Lavon

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

What's in the Box
-----------------

The following binaries can be created using the code:

* `prices` - Downloads XMLs from all known vendors.
* `items` - Parses XMLs and generates SQL statements.

The following SQL scripts can be used:

* `pre.sql` - Generates the table schema. Should be executed prior to using
  the SQL generated by `items`.
* `post.sql` - Executes some counting and cleanup. Should be executed after
  building the database with `items`.

Requirements
------------

1. [Go](http://golang.org/) compiler added to your system path, to compile the
   go code.
2. [SQLite](http://sqlite.org/), to run the SQL generated by the Go programs.

How to Build
------------

1. Update GOPATH:
   * **Linux** - `export GOPATH=/path/to/project`
   * **Windows** - `set GOPATH=\path\to\project`
2. Compile everything. `cd` into project/src folder and execute:
   `go install ./...`
3. A `bin` folder will be created in the project's folder, containing all
   generated binaries.

TODOs
-----

Things to do in no particular order (in addition to TODOs in code).

* Create a table for manufacturers.
* Change binary names from `prices` and `items` to something more informative.
* Embed SQL scripts in 'items' binary?

