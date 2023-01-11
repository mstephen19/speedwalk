import { walk } from '../index.js';

await walk('./', (path, dirent) => {
    console.log('Path:', path);
    console.log('Is a file:', dirent.isFile());

    if (dirent.isDirectory() && dirent.name === 'node_modules') {
        return true;
    }
});
