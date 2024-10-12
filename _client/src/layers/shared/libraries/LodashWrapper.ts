/**
 * This file mostly generated using AI
 */

import { merge as lodashMerge } from "lodash";

function merge<T>(...objects: (Partial<T> | null | undefined)[]): T {
  const validObjects = objects.filter((obj): obj is Partial<T> => obj != null);
  return lodashMerge({}, ...validObjects);
}

export { merge };
