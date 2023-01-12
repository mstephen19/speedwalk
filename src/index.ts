import { opendir } from 'fs/promises';
import { join } from 'path';

import type { WalkDirCallback } from './types.js';

/**
 * Traverse an entire directory's files and all subsequent subdirectories.
 *
 * @param path The directory to start walking in.
 * @param callback A callback function to run for each file found.
 * @returns
 */
export async function walk(root: string, callback: WalkDirCallback): Promise<void> {
    // Open the directory
    const dir = await opendir(root);

    // An array of promises where recursive calls can
    // be stored and awaited later on.
    const promises: Promise<void>[] = [];

    // Asynchronously iterate through the directory
    for await (const file of dir) {
        // Grab the current path value for the file.
        const current = join(root, file.name);

        promises.push(
            // Run the callback handling, but without awaiting
            // it within the loop so there is no blocking for
            // the next file.
            (async () => {
                const ignore = await callback(current, file);
                if (ignore || !file.isDirectory()) return;
                // If "true" wasn't returned from the callback and
                // the file is a directory, recurse into it.
                promises.push(walk(current, callback));
            })()
        );
    }

    return void Promise.all(promises);
}
