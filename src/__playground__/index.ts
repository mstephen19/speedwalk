import { walk } from '../index.js';

await walk('./', (path, dirent) => {
    console.log(path);
    if (dirent.name === 'node_modules') return true;
});
