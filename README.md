Israeli Supermarket Price Project
=================================

This is a collection of tools and scripts to download and manage price data
from Israeli vendors.

1. [License](#license)
3. [Requirements](#requirements)
4. [How to Build](#how-to-build)
2. [How to Use](#how-to-use)
5. [TODOs](#todos)

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

A link to the Github project (https://github.com/fluhus/prices) shall be
included in all copies or substantial portions of the Software. Any product
based on the Software must present the link either in its welcome, about or
help section, explicitly declaring that it is based on the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

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

How to Use
----------

### Downloading Price Data

1. Build the executables.
2. An executable called `prices` should be in the the `bin` folder. The program
   downloads all files from all known vendors.
3. Run `prices` without parameters for usage instructions.

### Creating the Database

1. Build the executables.
2. An executable called `items` should be in the the `bin` folder. The program
   parses XMLs and outputs SQL statements. It can handle XML, GZIP and ZIP
   files.
3. Prepare the database schema using `pre.sql`:
   ```sqlite3 my_database.db < pre.sql```
4. Use `items` to parse the desired files and pipe the SQL into SQLite. Run
   without parameters for usage instructions. Example:
   ```items my_prices1.xml my_prices2.zip | sqlite3 my_database.db```
5. Clean up the database using `post.sql`:
   ```sqlite3 my_database.db < post.sql```


