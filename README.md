# SpeedWalk

[![Install size](https://packagephobia.com/badge?p=speedwalk@latest)](https://packagephobia.com/result?p=speedwalk@latest)

Walk an entire directory. Fast, simple, and asynchronous.

## Features

- Zero dependencies. Small bundle size.
- Written in Typescript.
- Fast and asynchronous. Suitable for large directories. Fastest tree-traversal solution on NPM.
- Dead simple. Modeled after Golang's [`filepath.Walk`](https://pkg.go.dev/path/filepath#Walk).
- Modern [ESModules](https://hacks.mozilla.org/2018/03/es-modules-a-cartoon-deep-dive/) support.

## Comparisons

Some comparisons between **speedwalk** and other Node.js libraries that provide tree-traversal solutions.

> All tests were conducted on the same large directory. Each result is the average of each library's time to complete the walk with no custom functions provided.

|Package|Result|
|-|-|
|`speedwalk`|6.95ms|
|`walk`|13826.33ms|
|`@root/walk`|276.96ms|
|`walker`|296.05ms|

## Usage

Import the `walk` function from the package. To call it, pass in a root directory and a callback function accepting a `path` and a `dirent`.

```TypeScript
import { walk } from 'speedwalk';

await walk('./', (path, dirent) => {
    console.log('Path:', path);
    console.log('Is a file:', dirent.isFile());
});
```

The function will asynchronously walk through the root directory and all subsequent directories until the entire tree has been traversed.

You may want to ignore a directory. Tell `walk` to skip the traversal of a certain directory by returning `true` from the callback.

```TypeScript
import { walk } from 'speedwalk';

await walk('./', (path, dirent) => {
    console.log('Path:', path);
    console.log('Is a file:', dirent.isFile());

    if (dirent.isDirectory() && dirent.name === 'node_modules') {
        return true;
    }
});
```
