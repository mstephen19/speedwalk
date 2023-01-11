# DirWalk

[![TypeScript](https://badgen.net/badge/-/TypeScript/blue?icon=typescript&label)](https://www.typescriptlang.org/) [![Install size](https://packagephobia.com/badge?p=dirwalk@latest)](https://packagephobia.com/result?p=dirwalk@latest)

Walk an entire directory. Fast, simple, and asynchronous.

## Features

- Zero dependencies. Small bundle size.
- Written in Typescript.
- Fast and asynchronous. Suitable for large directories. Fastest tree-traversal solution on NPM.
- Dead simple. Modeled after Golang's [`filepath.Walk`](https://pkg.go.dev/path/filepath#Walk).
- Modern [ESModules](https://hacks.mozilla.org/2018/03/es-modules-a-cartoon-deep-dive/)-only support.

## Usage

Import the `walk` function from the package. To call it, pass in a root directory and a callback function accepting a `path` and a `dirent`.

```TypeScript
import { walk } from 'dirwalk';

await walk('./', (path, dirent) => {
    console.log('Path:', path);
    console.log('Is a file:', dirent.isFile());
});
```

The function will asynchronously traverse the directory and any subsequent directories until the entire tree has been traversed.

You may want to ignore a directory. Tell `walk` to skip the traversal of a certain directory by returning `true` from the callback.

```TypeScript
import { walk } from '../index.js';

await walk('./', (path, dirent) => {
    console.log('Path:', path);
    console.log('Is a file:', dirent.isFile());

    if (dirent.isDirectory() && dirent.name === 'node_modules') {
        return true;
    }
});
```

## License

The MIT License (MIT)

Copyright (c) 2023 Matthias Stephens

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
