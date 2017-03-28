polyfill-size
==

Given a particular set of features, grab a polyfill for many many user agents to determine and plan for file size.

installing
--

```bash
$ go install github.com/jspc/polyfill-size
```

building/ developing
--

```bash
$ go build
```

running
---

```bash
$ polyfill-size -h
Usage of ./polyfil-size:
  -agents string
        Path to file containing useragents to test (default "./agents.txt")
  -concurrency int
        Number of downloads to run at once (default 50)
  -features string
        'Features' string to request a polyfil for/from (default "default,fetch,includes,HTMLPictureElement,Array.prototype.entries,Object.assign")
```

licence
--

MIT License

Copyright (c) 2017 jspc

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
