txt-to-html
============================================================================

This simple utility converts notes written in txt/Markdown to HTML.

It takes one argument: a directory containing `*.txt` or `*.md` files.
HTML output files will be written to the same directory.
For example, an input file named `coolnotes.txt` will generate an output file named `coolnotes.txt.html`.

If the files `HEADER.html` and `FOOTER.html` exist in the directory, they will be prepended and appended to the output documents.

Supplying the `-i` flag generates an `index.html` file in the output directory.

Supplying the `-c` flag will substitute `.html` for the file name extension instead of appending, and clobber any existing HTML file with that name. Without the `-c` flag, input `file.txt` will output `file.txt.html`. With the flag, input `file.txt` will output to `file.html`.


License (2-clause BSD)
----------------------------------------------------------------------------

Copyright 2018 Paul Gorman

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
