import type { Dirent } from 'fs';

type Awaitable<T> = T | Promise<T>;

/**
 * A callback function to run for each file found.
 */
export type WalkDirCallback = (path: string, dir: Dirent) => Awaitable<void | boolean>;
