And/Or
=====================================================

> <span class="red">An</span>other <span class="red">d</span>igital / <span class="red">O</span>bject <span class="red">r</span>epository

[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg?style=flat-square)](https://choosealicense.com/licenses/bsd-3-clause)

<!-- [![Latest release](https://img.shields.io/badge/Latest_release-0.0.1-b44e88.svg?style=flat-square)](http://shields.io) -->



**And/Or** was a proof of concept for a simple object repository library
based on Caltech Library's [dataset](https://caltechlibrary.github.io/dataset)
tool.  It implements a asynchronous service for build a web-based multi-user repository using Python suitible for a small group of users curating JSON objects. __And/Or__ takes care of creating a efficient storage model
for your Python service, permissioning, access, control of those resources
are then implemented as a Python web service. This leaves the storage system
simple while allowing for easy construction of workflow oriented JSON
object curation.


Table of contents
-----------------

* [Introduction](#introduction)
* [Installation](install.html)
* [Documentation](docs/)
* [Known issues and limitations](#known-issues-and-limitations)
* [Getting help](#getting-help)
* [Contributing](contributing.html)
* [License](#license)
* [Authors and history](#authors-and-history)
* [Acknowledgments](#authors-and-acknowledgments)


Introduction
------------

**And/Or** is an asynchrous shared library based 
[dataset](https://github.com/caltechlibrary/dataset). __dataset__ 
has proven to be a useful tool for managing library metadata using 
a data science approach.  It is built for continious migration dataflows. 
It lacks multi-user curation support or the convienence of having 
web browser based edit forms. **And/Or** is a prototype of extended
libdataset.go to allow easy creation of Python based web services that
read/write to _dataset_ collections.  **And/Or** leverages Python as its
running environments for building multi-user extremely light weight JSON  bject repositories. 


Installation
------------

See [INSTALL.md](install.html). This software is experimental
and pre-compiled binaries are NOT provided.  This software is written in 
[Go](https://golang.org) programming language and needs a Go compiler
to be compiled. It is intended to be used from Python 3.7 or better so
that needs to be available too.

A "Makefile" has been provided for your convience if you wish to
manually compile andor library.

```bash
    git clone https://github.com/caltechlibrary/andor 
    cd andor
    go build -o bin/andor cmd/andor/andor.go
```



Known issues and limitations
----------------------------

This is a proof-of-concept project. It SHOULD NOT be used
in any production setting.  It is ONLY suitable for demonstrating
an approach to building light weight object repositories.

Getting help
------------

You can contact us via GitHub [issue tracker](https://github.com/caltechlibrary/andor/issues).

Contributing
------------

See [CONTRIBUTING.md](contributing.html)


License
-------

Software produced by the Caltech Library is Copyright (C) 2019, Caltech.  
This software is freely distributed under a BSD/MIT type license.  
Please see the [LICENSE](LICENSE) file for more information.


Authors and history
---------------------------

[Robert](https://rsdoiel.github.io) is the culprit responsible 
for this proof of concept


Acknowledgments
---------------

This work was funded by the California Institute of Technology Library.

<div align="center">
  <br>
  <a href="https://www.caltech.edu">
    <img width="100" height="100" src="assets/caltech-round.svg">
  </a>
</div>

